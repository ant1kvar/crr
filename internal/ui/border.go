package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/mattn/go-runewidth"
)

// Rounded border characters
const (
	borderTopLeft     = "╭"
	borderTopRight    = "╮"
	borderBottomLeft  = "╰"
	borderBottomRight = "╯"
	borderHorizontal  = "─"
	borderVertical    = "│"
)

// RenderBoxWithTitle renders content in a box with title in top border
// Example: ╭─ Title ─────╮
func RenderBoxWithTitle(content, title string, width int, borderColor lipgloss.Color, titleColor lipgloss.Color) string {
	// Inner width (without border)
	innerWidth := width - 2

	// Border style
	borderStyle := lipgloss.NewStyle().Foreground(borderColor)
	titleStyle := lipgloss.NewStyle().Foreground(titleColor).Bold(true)

	// Build top border with title
	titleText := " " + title + " "
	titleLen := runewidth.StringWidth(titleText)

	// Number of ─ characters left and right of title
	remaining := innerWidth - titleLen
	leftDash := 1 // Minimum one ─ on left
	rightDash := remaining - leftDash
	if rightDash < 0 {
		rightDash = 0
	}

	topLine := borderStyle.Render(borderTopLeft) +
		borderStyle.Render(strings.Repeat(borderHorizontal, leftDash)) +
		titleStyle.Render(titleText) +
		borderStyle.Render(strings.Repeat(borderHorizontal, rightDash)) +
		borderStyle.Render(borderTopRight)

	// Build bottom border
	bottomLine := borderStyle.Render(borderBottomLeft) +
		borderStyle.Render(strings.Repeat(borderHorizontal, innerWidth)) +
		borderStyle.Render(borderBottomRight)

	// Build content lines with side borders
	contentLines := strings.Split(content, "\n")
	var middleLines []string
	for _, line := range contentLines {
		// Pad line to required width accounting for character widths
		lineWidth := runewidth.StringWidth(line)
		padding := innerWidth - lineWidth
		if padding < 0 {
			padding = 0
		}
		paddedLine := line + strings.Repeat(" ", padding)

		middleLine := borderStyle.Render(borderVertical) +
			paddedLine +
			borderStyle.Render(borderVertical)
		middleLines = append(middleLines, middleLine)
	}

	// Assemble everything
	result := topLine + "\n" +
		strings.Join(middleLines, "\n") + "\n" +
		bottomLine

	return result
}
