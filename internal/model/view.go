package model

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/mattn/go-runewidth"

	"crr/internal/ui"
)

// View renders the interface (required by tea.Model interface)
// Called after each Update to redraw the screen
func (d Drums) View() string {
	// If size not yet received, show loading
	if d.Width == 0 {
		return "Search..."
	}

	// Render header panel
	header := d.renderHeader()

	// Render drums
	var columns []string
	for i, drum := range d.List {
		column := d.renderDrum(&drum, i == d.Active)
		columns = append(columns, column)
	}
	drums := lipgloss.JoinHorizontal(lipgloss.Top, columns...)

	return header + "\n\n" + drums
}

// renderHeader renders the header panel (track on left, clock on right)
func (d *Drums) renderHeader() string {
	// Get clock display
	clockText := d.Clock.DisplayBig()
	clockLines := strings.Split(clockText, "\n")

	// Clock width
	clockWidth := 0
	for _, line := range clockLines {
		w := runewidth.StringWidth(line)
		if w > clockWidth {
			clockWidth = w
		}
	}

	// Track width
	trackWidth := d.Width - clockWidth - 2

	// Track (artist and title)
	artist, name := d.Track.DisplayLines()
	artist = ui.Truncate(artist, trackWidth-2)
	name = ui.Truncate(name, trackWidth-2)

	// Build lines: track on left, clock on right
	// Clock is 5 lines, track on lines 2 and 3 (with top padding)
	var lines []string
	for i, clockLine := range clockLines {
		trackPart := ""
		if i == 1 {
			trackPart = "  " + artist
		} else if i == 2 {
			trackPart = "  " + name
		}

		// Padding between track and clock
		trackPartWidth := runewidth.StringWidth(trackPart)
		padding := d.Width - trackPartWidth - runewidth.StringWidth(clockLine)
		if padding < 0 {
			padding = 0
		}

		lines = append(lines, trackPart+strings.Repeat(" ", padding)+clockLine)
	}

	return strings.Join(lines, "\n")
}

// renderDrum renders a single drum column
func (d *Drums) renderDrum(drum *Drum, isActiveColumn bool) string {
	var items []string

	// Calculate visible window center
	middle := ui.VisibleItems / 2

	// Dynamic sizes
	itemWidth := d.ItemWidth()
	maxTextWidth := d.MaxTextWidth()

	// Styles with dynamic width
	activeItemStyle := lipgloss.NewStyle().
		Foreground(ui.ActiveColor).
		Bold(true).
		Border(lipgloss.NormalBorder()).
		BorderForeground(ui.ActiveColor)

	itemStyle := lipgloss.NewStyle().
		Foreground(ui.InactiveColor)

	// Total line width (including active item border)
	totalLineWidth := itemWidth + 2

	// Display only VisibleItems items with active in center
	for j := 0; j < ui.VisibleItems; j++ {
		offset := j - middle
		item := drum.GetItem(drum.Active + offset)

		isCenter := j == middle

		if isCenter {
			displayText := ui.Marquee(item, d.ScrollOffset, maxTextWidth)
			// Center text inside active item
			displayText = centerText(displayText, itemWidth)
			items = append(items, activeItemStyle.Render(displayText))
		} else {
			displayText := ui.Truncate(item, maxTextWidth)
			// Center text and pad to total width
			displayText = centerText(displayText, totalLineWidth)
			items = append(items, itemStyle.Render(displayText))
		}
	}

	content := strings.Join(items, "\n")

	var borderColor lipgloss.Color
	if isActiveColumn {
		borderColor = ui.ActiveColor
	} else {
		borderColor = ui.InactiveColor
	}

	return ui.RenderBoxWithTitle(content, drum.Title, d.ColumnWidth(), borderColor, borderColor)
}

// centerText centers text within given width accounting for character widths
func centerText(s string, width int) string {
	textWidth := runewidth.StringWidth(s)
	if textWidth >= width {
		return s
	}
	padding := width - textWidth
	leftPad := padding / 2
	rightPad := padding - leftPad
	return strings.Repeat(" ", leftPad) + s + strings.Repeat(" ", rightPad)
}
