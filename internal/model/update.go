package model

import (
	tea "github.com/charmbracelet/bubbletea"

	"crr/internal/logger"
)

// Update handles events (required by tea.Model interface)
// Called on each event (key press, window resize, etc.)
func (d Drums) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		// Save terminal dimensions
		d.Width = msg.Width
		d.Height = msg.Height
		return d, nil

	case TickMsg:
		// Increment offset for marquee animation
		d.ScrollOffset++
		return d, DoTick() // Continue ticking

	case ClockTickMsg:
		// Update clock
		d.Clock.Update()
		return d, DoClockTick() // Continue clock updates

	case MetadataTickMsg:
		// Request current stream metadata
		return d, tea.Batch(
			DoMetadataTick(), // Continue ticking
			DoFetchMetadata(d.CurrentStreamURL),
		)

	case MetadataMsg:
		// Update track info
		if msg.Err == nil && msg.Title != "" {
			d.Track.SetTrack(msg.Title, msg.Artist)
		}
		return d, nil

	case FetchDebounceMsg:
		// Check debounce validity
		if msg.ID != d.DebounceID {
			return d, nil // Outdated debounce, ignore
		}
		// Start station loading (chunk is already playing since key press)
		d.Loading = true
		return d, DoFetchStations(d.CurrentCountryCode(), d.CurrentGenre())

	case FetchStationsMsg:
		logger.Log.Printf("FetchStationsMsg received: %d stations, err=%v", len(msg.Stations), msg.Err)
		d.Loading = false
		if msg.Err != nil {
			// Load error - keep current list
			logger.Log.Printf("Error loading stations: %v", msg.Err)
			return d, nil
		}
		// Update station list
		d.Stations = msg.Stations
		// Update third drum with station names
		names := make([]string, len(msg.Stations))
		for i, s := range msg.Stations {
			names[i] = s.Name
		}
		logger.Log.Printf("Updating drum with %d names", len(names))
		if len(names) > 0 {
			d.List[2].Items = names
			d.List[2].Active = 0
			// Auto-play first station
			d.CurrentStreamURL = d.Stations[0].Link
			d.Track.SetTrack(d.Stations[0].Name, "") // Show station name for now
			return d, tea.Batch(
				DoPlayStream(d.Player, d.Stations[0].Link),
				DoFetchMetadata(d.Stations[0].Link), // Request metadata immediately
			)
		}
		return d, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			// Stop playback on exit
			if d.Player != nil {
				d.Player.Cleanup()
			}
			return d, tea.Quit

		// Drum navigation
		case "up", "k":
			oldCountry := d.CurrentCountry()
			oldGenre := d.CurrentGenre()
			d.List[d.Active].MoveUp()
			d.ScrollOffset = 0
			// If in Station column - instant chunk + switch
			if d.Active == 2 && len(d.Stations) > 0 {
				idx := d.List[2].Active
				d.CurrentStreamURL = d.Stations[idx].Link
				d.Track.SetTrack(d.Stations[idx].Name, "") // Station name for now
				return d, tea.Batch(
					DoSwitchStation(d.Player, d.Stations[idx].Link),
					DoFetchMetadata(d.Stations[idx].Link),
				)
			}
			// Instant chunk + debounce on country/genre change
			return d, d.checkFetchDebounceWithChunk(oldCountry, oldGenre)

		case "down", "j":
			oldCountry := d.CurrentCountry()
			oldGenre := d.CurrentGenre()
			d.List[d.Active].MoveDown()
			d.ScrollOffset = 0
			// If in Station column - instant chunk + switch
			if d.Active == 2 && len(d.Stations) > 0 {
				idx := d.List[2].Active
				d.CurrentStreamURL = d.Stations[idx].Link
				d.Track.SetTrack(d.Stations[idx].Name, "") // Station name for now
				return d, tea.Batch(
					DoSwitchStation(d.Player, d.Stations[idx].Link),
					DoFetchMetadata(d.Stations[idx].Link),
				)
			}
			// Instant chunk + debounce on country/genre change
			return d, d.checkFetchDebounceWithChunk(oldCountry, oldGenre)

		case "left", "h":
			d.MoveLeft()
			d.ScrollOffset = 0

		case "right", "l":
			d.MoveRight()
			d.ScrollOffset = 0

		// Volume control
		case "+", "=":
			d.Volume.Up()

		case "-", "_":
			d.Volume.Down()

		case "m":
			d.Volume.ToggleMute()
		}
	}

	return d, nil
}

// checkFetchDebounce checks if country/genre changed and starts debounce
func (d *Drums) checkFetchDebounce(oldCountry, oldGenre string) tea.Cmd {
	newCountry := d.CurrentCountry()
	newGenre := d.CurrentGenre()

	// If country or genre changed - start debounce
	if newCountry != oldCountry || newGenre != oldGenre {
		d.DebounceID++
		d.List[2].Items = []string{"Scanning..."}
		d.List[2].Active = 0
		d.Track.SetTrack("", "") // Reset track to scanning state
		return DoFetchDebounce(d.DebounceID)
	}
	return nil
}

// checkFetchDebounceWithChunk same as above + instant chunk on change
func (d *Drums) checkFetchDebounceWithChunk(oldCountry, oldGenre string) tea.Cmd {
	newCountry := d.CurrentCountry()
	newGenre := d.CurrentGenre()

	// If country or genre changed - instant chunk + debounce
	if newCountry != oldCountry || newGenre != oldGenre {
		d.DebounceID++
		d.List[2].Items = []string{"Scanning..."}
		d.List[2].Active = 0
		d.Track.SetTrack("", "") // Reset track to scanning state
		// Stop current stream and play chunk immediately
		d.Player.Stop()
		d.Player.PlayChunkImmediately()
		return DoFetchDebounce(d.DebounceID)
	}
	return nil
}
