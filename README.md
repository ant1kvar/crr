# Cool Retro Radio

A vintage-style terminal radio station selector with a three-drum interface. Browse stations by country and genre, enjoy smooth transitions with audio chunks, and experience the nostalgic feel of old-school radio tuning.

### Interface Layout
<img width="800" height="600" alt="image" src="https://github.com/user-attachments/assets/f9206f5b-b377-42c9-ae25-8976b262f98f" />

[![Bubble Tea](https://img.shields.io/badge/TUI-Bubble%20Tea-ff69b4)](https://github.com/charmbracelet/bubbletea)
[![Cool Retro Term](https://img.shields.io/badge/Best%20with-Cool%20Retro%20Term-orange)](https://github.com/Swordfish90/cool-retro-term)
[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-green)](LICENSE)

## Features

- **Three-Drum Selector** - Infinite scroll columns for Country, Genre, and Station
- **Live Radio Streaming** - Powered by Radio Browser API with thousands of stations worldwide
- **Smooth Transitions** - Audio chunks play during station switching for seamless experience
- **Track Info Display** - Real-time metadata extraction (artist & title)
- **Big Digital Clock** - Retro ASCII-art clock with blinking colon
- **Marquee Animation** - Long text scrolls smoothly in active selections

## Best Experience

For the ultimate retro experience, use **Cool Retro Radio** with [Cool Retro Term](https://github.com/Swordfish90/cool-retro-term) - a terminal emulator that mimics old cathode displays with scan lines, screen curvature, and phosphor glow effects.

```
Cool Retro Term + Cool Retro Radio = Pure Nostalgia
```

## Installation

### Prerequisites

- Go 1.21 or higher
- `ffmpeg` and `ffplay` (for audio playback)

```bash
# macOS
brew install ffmpeg

# Ubuntu/Debian
sudo apt install ffmpeg

# Arch
sudo pacman -S ffmpeg
```

### Build from Source

```bash
git clone https://github.com/ant1kvar/crr.git
cd crr
go build
./crr
```

## Usage

### Keyboard Controls

| Key | Action |
|-----|--------|
| `↑` / `k` | Scroll up |
| `↓` / `j` | Scroll down |
| `←` / `h` | Previous column |
| `→` / `l` | Next column |
| `m` | Toggle mute |
| `q` | Quit |

## How It Works

1. **Station Discovery** - Fetches stations from Radio Browser API based on selected country and genre
2. **Debounced Loading** - 3-second delay before fetching to avoid excessive API calls during navigation
3. **Audio Chunks** - Short audio clips play immediately when switching stations for instant feedback
4. **Crossfade** - Smooth audio transition from chunk to live stream using ffmpeg filters
5. **Metadata Polling** - Periodically fetches ICY metadata from the stream for track info

## Project Structure

```
crr/
├── main.go                 # Entry point
├── chunks/                 # Audio chunks for transitions (*.mp3)
└── internal/
    ├── model/              # Bubble Tea model (MVC pattern)
    │   ├── drums.go        # Main model with three columns
    │   ├── drum.go         # Single column with infinite scroll
    │   ├── update.go       # Event handling (keyboard, timers)
    │   ├── view.go         # UI rendering
    │   ├── track.go        # Track info component
    │   ├── volume.go       # Volume control
    │   ├── clock.go        # Digital clock
    │   └── tick.go         # Timer commands and async operations
    ├── ui/                 # UI utilities
    │   ├── styles.go       # Colors and constants
    │   ├── text.go         # Marquee and truncation
    │   ├── border.go       # Rounded box rendering
    │   └── digits.go       # ASCII-art digits
    ├── data/               # Static data
    │   ├── items.go        # Countries and genres
    │   └── station.go      # Station type
    ├── client/             # Radio Browser API client
    ├── cache/              # File-based station cache
    ├── player/             # ffmpeg/ffplay audio player
    └── logger/             # Debug logging
```

## Dependencies

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - TUI framework
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - Terminal styling
- [radiobrowser-go](https://github.com/randomtoy/radiobrowser-go) - Radio Browser API client
- [go-runewidth](https://github.com/mattn/go-runewidth) - Unicode character width

## Audio Chunks

Place short MP3 files (1-2 seconds) in the `chunks/` directory. These play instantly when switching stations, providing immediate audio feedback while the new stream buffers.

Example chunks: vinyl scratches, radio static, tuning sounds, click sounds.

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## License

MIT License - see [LICENSE](LICENSE) for details.

---

*Tune in. Turn on. Drop out.*
