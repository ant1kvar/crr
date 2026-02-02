package ui

import "github.com/mattn/go-runewidth"

// Truncate truncates text to maxWidth columns with ellipsis
func Truncate(s string, maxWidth int) string {
	w := runewidth.StringWidth(s)
	if w <= maxWidth {
		return s
	}
	return runewidth.Truncate(s, maxWidth, "...")
}

// Marquee returns a scrolling text "window"
// offset is current scroll position, width is window width in columns
func Marquee(s string, offset, width int) string {
	textWidth := runewidth.StringWidth(s)

	// If text fits - return as is
	if textWidth <= width {
		return s
	}

	// Add separator for smooth cycling
	separator := "   "
	loopText := s + separator
	loopRunes := []rune(loopText)
	loopLen := len(loopRunes)

	// Normalize offset for cycling
	offset = offset % loopLen

	// Build visible window accounting for character widths
	var result []rune
	currentWidth := 0
	for i := 0; currentWidth < width; i++ {
		idx := (offset + i) % loopLen
		r := loopRunes[idx]
		rw := runewidth.RuneWidth(r)
		if currentWidth+rw > width {
			// Add space if wide character doesn't fit
			result = append(result, ' ')
			break
		}
		result = append(result, r)
		currentWidth += rw
	}

	return string(result)
}
