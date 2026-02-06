package model

import (
	tea "github.com/charmbracelet/bubbletea"

	"crr/internal/data"
	"crr/internal/player"
)

// Drums is the main application model containing three columns
type Drums struct {
	List         [3]Drum // Array of three columns (Countries, Genre, Station)
	Active       int     // Index of currently active column (0, 1, or 2)
	ScrollOffset int     // Offset for marquee text animation
	Width        int     // Terminal width in characters
	Height       int     // Terminal height in characters

	// Header panel components
	Track  *Track  // Track info (left side)
	Volume *Volume // Volume control (center)
	Clock  *Clock  // Clock display (right side)

	// Station loading
	Stations   []data.Station // Loaded stations
	DebounceID int            // Current debounce timer ID
	Loading    bool           // Station loading flag

	// Playback
	Player           *player.Player // Audio player
	CurrentStreamURL string         // Current stream URL (for metadata)
}

// ColumnWidth returns the width of a single column (one third of terminal)
func (d *Drums) ColumnWidth() int {
	if d.Width < 30 {
		return 20 // Minimum width
	}
	return d.Width / 3
}

// InnerWidth returns inner column width (without column border)
func (d *Drums) InnerWidth() int {
	return d.ColumnWidth() - 2 // Minus left and right border
}

// ItemWidth returns item content width for lipgloss.Width()
// Item border will add 2 characters to total width
func (d *Drums) ItemWidth() int {
	return d.InnerWidth() - 2 // Item border is added automatically
}

// MaxTextWidth returns maximum text width inside an item
func (d *Drums) MaxTextWidth() int {
	return d.ItemWidth() - 4 // Minus inner borders and padding
}

// NewDrums creates a new Drums instance
func NewDrums() *Drums {
	countries := Drum{data.CountryNames(), 0, "Countries"}
	genre := Drum{data.Genre, 0, "Genre"}
	station := Drum{[]string{"Loading..."}, 0, "Station"}

	// Extract embedded chunks to temp directory
	chunksDir, err := player.ExtractChunks()
	if err != nil {
		chunksDir = "chunks" // Fallback to local directory
	}

	return &Drums{
		List:    [3]Drum{countries, genre, station},
		Active:  0,
		Track:   NewTrack(),
		Volume:  NewVolume(),
		Clock:   NewClock(),
		Loading: true,
		Player:  player.New(chunksDir),
	}
}

// Init initializes the model (required by tea.Model interface)
func (d Drums) Init() tea.Cmd {
	countryCode := d.CurrentCountryCode()
	genre := d.CurrentGenre()
	// Play chunk immediately on startup
	d.Player.PlayChunkImmediately()
	return tea.Batch(
		DoTick(),
		DoClockTick(),
		DoMetadataTick(),                    // Metadata update timer
		DoFetchStations(countryCode, genre), // Initial station load
	)
}

// CurrentCountry returns the name of currently selected country (for UI)
func (d *Drums) CurrentCountry() string {
	return d.List[0].GetItem(d.List[0].Active)
}

// CurrentCountryCode returns the code of currently selected country (for API)
func (d *Drums) CurrentCountryCode() string {
	name := d.CurrentCountry()
	return data.CountryCodeByName(name)
}

// CurrentGenre returns the currently selected genre
func (d *Drums) CurrentGenre() string {
	return d.List[1].GetItem(d.List[1].Active)
}

// ActiveDrum returns a pointer to the currently active column
func (d *Drums) ActiveDrum() *Drum {
	return &d.List[d.Active]
}

// MoveLeft switches to the column on the left
func (d *Drums) MoveLeft() {
	if d.Active > 0 {
		d.Active--
	}
}

// MoveRight switches to the column on the right
func (d *Drums) MoveRight() {
	if d.Active < 2 {
		d.Active++
	}
}
