// Package model contains data models for the TUI application
package model

// Drum represents a single column (drum) with selectable items
type Drum struct {
	Items  []string // List of items in the column (countries, genres, or stations)
	Active int      // Index of the currently selected item
	Title  string   // Column title
}

// Len returns the number of items in the drum
func (d *Drum) Len() int {
	return len(d.Items)
}

// MoveUp moves selection up with wrap-around
func (d *Drum) MoveUp() {
	d.Active--
	if d.Active < 0 {
		d.Active = d.Len() - 1 // Wrap to end
	}
}

// MoveDown moves selection down with wrap-around
func (d *Drum) MoveDown() {
	d.Active++
	if d.Active >= d.Len() {
		d.Active = 0 // Wrap to beginning
	}
}

// GetItem returns item at index with wrap-around support
func (d *Drum) GetItem(index int) string {
	n := d.Len()
	idx := ((index % n) + n) % n // Normalize index for wrap-around
	return d.Items[idx]
}
