// Package ui contains styles for interface rendering
package ui

import "github.com/charmbracelet/lipgloss"

// VisibleItems is the number of visible items in a drum
// Selected item is always in center (index 2 for 5 items)
const VisibleItems = 5

// UI Colors
var (
	ActiveColor   = lipgloss.Color("212") // Pink
	InactiveColor = lipgloss.Color("240") // Gray
)
