package model

// Track holds current track information
type Track struct {
	Name   string // Track title
	Artist string // Artist name
}

// NewTrack creates a new Track instance
func NewTrack() *Track {
	return &Track{
		Name:   "",
		Artist: "",
	}
}

// SetTrack sets track information
func (t *Track) SetTrack(name, artist string) {
	t.Name = name
	t.Artist = artist
}

// DisplayName returns formatted track display name
func (t *Track) DisplayName() string {
	if t.Artist == "" {
		return "SCANNING..."
	}
	if t.Name == "" {
		return t.Artist
	}
	return t.Artist + " - " + t.Name
}

// DisplayLines returns artist and title as two separate lines
func (t *Track) DisplayLines() (string, string) {
	// If artist is empty - metadata is being scanned
	if t.Artist == "" {
		return "SCANNING...", ""
	}
	return t.Artist, t.Name
}
