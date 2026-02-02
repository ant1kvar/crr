package player

import (
	"fmt"
	"math/rand"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

// Volume settings (in dB)
const (
	ChunkVolumeDB  = "3" // Chunk volume (+3dB, louder)
	StreamVolumeDB = "0" // Stream volume (0dB, normal)
)

// Player manages audio playback
type Player struct {
	ffmpeg    *exec.Cmd  // ffmpeg process (crossfade)
	ffplay    *exec.Cmd  // ffplay process (playback)
	chunksDir string     // path to chunks folder
	mu        sync.Mutex // race condition protection
}

// New creates a new Player
func New(chunksDir string) *Player {
	return &Player{
		chunksDir: chunksDir,
	}
}

// getRandomChunk returns path to a random chunk file
func (p *Player) getRandomChunk() (string, error) {
	files, err := filepath.Glob(filepath.Join(p.chunksDir, "*.mp3"))
	if err != nil || len(files) == 0 {
		return "", fmt.Errorf("no chunks found in %s", p.chunksDir)
	}
	return files[rand.Intn(len(files))], nil
}

// PlayStream plays stream with crossfade from chunk
// ffmpeg -i chunk.mp3 -i stream_url -filter_complex "acrossfade=d=0" -f matroska - | ffplay -
func (p *Player) PlayStream(url string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	// Stop current playback
	p.stopLocked()

	// Get random chunk
	chunk, err := p.getRandomChunk()
	if err != nil {
		// If no chunks - just play stream directly
		return p.playDirectLocked(url)
	}

	// ffmpeg: crossfade chunk → stream
	p.ffmpeg = exec.Command("ffmpeg",
		"-i", chunk,
		"-i", url,
		"-filter_complex", "[0:a]apad[a0];[a0][1:a]acrossfade=d=1:c1=tri:c2=tri",
		"-f", "matroska",
		"-loglevel", "quiet",
		"-",
	)

	// ffplay: playback from pipe
	p.ffplay = exec.Command("ffplay",
		"-nodisp",
		"-loglevel", "quiet",
		"-i", "-",
	)

	// Connect ffmpeg stdout → ffplay stdin
	pipe, err := p.ffmpeg.StdoutPipe()
	if err != nil {
		return err
	}
	p.ffplay.Stdin = pipe

	// Start both processes
	if err := p.ffmpeg.Start(); err != nil {
		return err
	}
	if err := p.ffplay.Start(); err != nil {
		p.ffmpeg.Process.Kill()
		return err
	}

	return nil
}

// PlayChunkThenStream plays chunk immediately, then connects to stream
func (p *Player) PlayChunkThenStream(url string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	// Stop current playback
	p.stopLocked()

	// Get random chunk
	chunk, err := p.getRandomChunk()
	if err != nil {
		return p.playDirectLocked(url)
	}

	// Start ffmpeg with concat: chunk first, then stream
	// Using concat demuxer via pipe
	p.ffmpeg = exec.Command("ffmpeg",
		"-i", chunk, // chunk as first input
		"-i", url, // stream as second input
		"-filter_complex",
		"[0:a]apad=pad_dur=1.2[a0];[a0][1:a]acrossfade=d=1:c1=tri:c2=tri[out]",
		"-map", "[out]",
		"-f", "matroska",
		"-loglevel", "quiet",
		"-",
	)

	p.ffplay = exec.Command("ffplay",
		"-nodisp",
		"-loglevel", "quiet",
		"-i", "-",
	)

	pipe, err := p.ffmpeg.StdoutPipe()
	if err != nil {
		return err
	}
	p.ffplay.Stdin = pipe

	if err := p.ffmpeg.Start(); err != nil {
		return err
	}
	if err := p.ffplay.Start(); err != nil {
		p.ffmpeg.Process.Kill()
		return err
	}

	return nil
}

// PlayChunkImmediately instantly starts chunk playback (separate process, louder)
func (p *Player) PlayChunkImmediately() error {
	chunk, err := p.getRandomChunk()
	if err != nil {
		return err
	}
	// Start separate ffplay for chunk with increased volume
	cmd := exec.Command("ffplay",
		"-nodisp", "-autoexit", "-loglevel", "quiet",
		"-af", "volume="+ChunkVolumeDB+"dB",
		chunk,
	)
	return cmd.Start()
}

// SwitchStation plays instant chunk + connects to stream
func (p *Player) SwitchStation(url string) error {
	// 1. Stop current stream immediately
	p.Stop()

	// 2. Instantly start chunk (separate process)
	p.PlayChunkImmediately()

	// 3. Start connecting to new stream (in background)
	go func() {
		p.mu.Lock()
		defer p.mu.Unlock()
		p.playDirectLocked(url)
	}()

	return nil
}

// playDirectLocked plays stream directly without crossfade (fallback)
func (p *Player) playDirectLocked(url string) error {
	p.ffplay = exec.Command("ffplay",
		"-nodisp", "-loglevel", "quiet",
		"-af", "volume="+StreamVolumeDB+"dB",
		url,
	)
	return p.ffplay.Start()
}

// PlayChunk plays only chunk (without stream, with increased volume)
func (p *Player) PlayChunk() error {
	chunk, err := p.getRandomChunk()
	if err != nil {
		return err
	}

	cmd := exec.Command("ffplay",
		"-nodisp", "-autoexit", "-loglevel", "quiet",
		"-af", "volume="+ChunkVolumeDB+"dB",
		chunk,
	)
	return cmd.Start()
}

// Stop stops current playback
func (p *Player) Stop() error {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.stopLocked()
}

// stopLocked is internal stop method (call with mutex held)
func (p *Player) stopLocked() error {
	if p.ffplay != nil && p.ffplay.Process != nil {
		p.ffplay.Process.Kill()
		p.ffplay = nil
	}
	if p.ffmpeg != nil && p.ffmpeg.Process != nil {
		p.ffmpeg.Process.Kill()
		p.ffmpeg = nil
	}
	return nil
}

// IsPlaying checks if stream is playing
func (p *Player) IsPlaying() bool {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.ffplay != nil && p.ffplay.Process != nil
}

// Cleanup terminates all processes on exit
func (p *Player) Cleanup() {
	p.Stop()
	// Kill all remaining processes
	exec.Command("pkill", "-f", "ffplay").Run()
	exec.Command("pkill", "-f", "ffmpeg").Run()
}

// TrackInfo contains track information from stream
type TrackInfo struct {
	Title  string
	Artist string
}

// GetStreamMetadata fetches metadata from stream via ffprobe
func GetStreamMetadata(url string) (*TrackInfo, error) {
	// ffprobe -v quiet -show_entries format_tags=StreamTitle,icy-title,title,artist
	// -of default=nw=1:nk=1 "url"
	ctx, cancel := exec.Command("ffprobe",
		"-v", "quiet",
		"-show_entries", "format_tags=StreamTitle,icy-title,title,artist",
		"-of", "json",
		"-i", url,
	).Output()
	cancel = nil
	_ = cancel

	if len(ctx) == 0 {
		return nil, fmt.Errorf("no metadata")
	}

	// Parse JSON response
	return parseMetadata(ctx)
}

// parseMetadata parses JSON from ffprobe
func parseMetadata(data []byte) (*TrackInfo, error) {
	// Simple parsing - look for StreamTitle or title
	str := string(data)
	info := &TrackInfo{}

	// Look for StreamTitle (ICY metadata)
	if idx := strings.Index(str, "StreamTitle"); idx != -1 {
		info.Title = extractJSONValue(str[idx:])
	} else if idx := strings.Index(str, "icy-title"); idx != -1 {
		info.Title = extractJSONValue(str[idx:])
	} else if idx := strings.Index(str, "title"); idx != -1 {
		info.Title = extractJSONValue(str[idx:])
	}

	// Look for artist
	if idx := strings.Index(str, "artist"); idx != -1 {
		info.Artist = extractJSONValue(str[idx:])
	}

	// If "Artist - Title" format in StreamTitle
	if info.Title != "" && info.Artist == "" {
		if parts := strings.SplitN(info.Title, " - ", 2); len(parts) == 2 {
			info.Artist = parts[0]
			info.Title = parts[1]
		}
	}

	if info.Title == "" {
		return nil, fmt.Errorf("no title found")
	}

	return info, nil
}

// extractJSONValue extracts value from JSON string
func extractJSONValue(s string) string {
	// Look for ": "value"
	start := strings.Index(s, "\":")
	if start == -1 {
		return ""
	}
	s = s[start+2:]

	// Skip whitespace
	s = strings.TrimLeft(s, " ")

	// Find value start
	if len(s) == 0 || s[0] != '"' {
		return ""
	}
	s = s[1:]

	// Find value end
	end := strings.Index(s, "\"")
	if end == -1 {
		return ""
	}

	return s[:end]
}
