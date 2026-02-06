package player

import (
	"embed"
	"io/fs"
	"os"
	"path/filepath"
)

//go:embed chunks/*.mp3
var embeddedChunks embed.FS

var extractedChunksDir string

// ExtractChunks extracts embedded chunks to temp directory
// Returns path to extracted chunks directory
func ExtractChunks() (string, error) {
	if extractedChunksDir != "" {
		return extractedChunksDir, nil
	}

	// Create temp directory
	tmpDir, err := os.MkdirTemp("", "crr-chunks-*")
	if err != nil {
		return "", err
	}

	// Extract all mp3 files
	entries, err := fs.ReadDir(embeddedChunks, "chunks")
	if err != nil {
		return "", err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		data, err := embeddedChunks.ReadFile("chunks/" + entry.Name())
		if err != nil {
			continue
		}

		outPath := filepath.Join(tmpDir, entry.Name())
		if err := os.WriteFile(outPath, data, 0644); err != nil {
			continue
		}
	}

	extractedChunksDir = tmpDir
	return tmpDir, nil
}

// CleanupChunks removes extracted chunks directory
func CleanupChunks() {
	if extractedChunksDir != "" {
		os.RemoveAll(extractedChunksDir)
		extractedChunksDir = ""
	}
}
