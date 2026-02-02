package model

import (
	"context"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"crr/internal/client"
	"crr/internal/data"
	"crr/internal/player"
)

// TickMsg is a timer message for marquee animation
type TickMsg time.Time

// TickInterval is the marquee update interval
const TickInterval = 300 * time.Millisecond

// DoTick creates a timer command
func DoTick() tea.Cmd {
	return tea.Tick(TickInterval, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

// ClockTickMsg is a message for clock updates
type ClockTickMsg time.Time

// ClockTickInterval is the clock update interval (1 second)
const ClockTickInterval = 1 * time.Second

// DoClockTick creates a clock timer command
func DoClockTick() tea.Cmd {
	return tea.Tick(ClockTickInterval, func(t time.Time) tea.Msg {
		return ClockTickMsg(t)
	})
}

// FetchDebounceInterval is the delay before station request
const FetchDebounceInterval = 3 * time.Second

// FetchDebounceMsg is a debounce timer message
type FetchDebounceMsg struct {
	ID int // ID for validity check
}

// DoFetchDebounce creates a debounce timer command
func DoFetchDebounce(id int) tea.Cmd {
	return tea.Tick(FetchDebounceInterval, func(t time.Time) tea.Msg {
		return FetchDebounceMsg{ID: id}
	})
}

// FetchStationsMsg contains station loading result
type FetchStationsMsg struct {
	Stations []data.Station
	Err      error
}

// DoFetchStations creates a station loading command
func DoFetchStations(country, genre string) tea.Cmd {
	return func() tea.Msg {
		stations, err := client.GetStations(context.Background(), country, genre)
		return FetchStationsMsg{Stations: stations, Err: err}
	}
}

// PlayChunkMsg signals chunk playback completion
type PlayChunkMsg struct {
	Err error
}

// DoPlayChunk creates a command to play random chunk
func DoPlayChunk(p *player.Player) tea.Cmd {
	return func() tea.Msg {
		err := p.PlayChunk()
		return PlayChunkMsg{Err: err}
	}
}

// PlayStreamMsg signals stream playback start
type PlayStreamMsg struct {
	URL string
	Err error
}

// DoPlayStream creates a command to play audio stream
func DoPlayStream(p *player.Player, url string) tea.Cmd {
	return func() tea.Msg {
		err := p.PlayStream(url)
		return PlayStreamMsg{URL: url, Err: err}
	}
}

// SwitchStationMsg signals station switch
type SwitchStationMsg struct {
	URL string
	Err error
}

// DoSwitchStation creates instant chunk + stream connection command
func DoSwitchStation(p *player.Player, url string) tea.Cmd {
	return func() tea.Msg {
		err := p.SwitchStation(url)
		return SwitchStationMsg{URL: url, Err: err}
	}
}

// MetadataTickMsg is a timer message for metadata updates
type MetadataTickMsg time.Time

// MetadataTickInterval is the metadata polling interval (2 seconds)
const MetadataTickInterval = 2 * time.Second

// DoMetadataTick creates a metadata timer command
func DoMetadataTick() tea.Cmd {
	return tea.Tick(MetadataTickInterval, func(t time.Time) tea.Msg {
		return MetadataTickMsg(t)
	})
}

// MetadataMsg contains metadata fetch result
type MetadataMsg struct {
	Title  string
	Artist string
	Err    error
}

// DoFetchMetadata creates a command to fetch stream metadata
func DoFetchMetadata(url string) tea.Cmd {
	return func() tea.Msg {
		if url == "" {
			return MetadataMsg{Err: nil}
		}
		info, err := player.GetStreamMetadata(url)
		if err != nil {
			return MetadataMsg{Err: err}
		}
		return MetadataMsg{Title: info.Title, Artist: info.Artist}
	}
}
