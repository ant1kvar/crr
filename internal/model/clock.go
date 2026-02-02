package model

import (
	"time"

	"crr/internal/ui"
)

// Clock displays date and time
type Clock struct {
	currentTime time.Time
}

// NewClock creates a new clock instance
func NewClock() *Clock {
	return &Clock{
		currentTime: time.Now(),
	}
}

// Update refreshes current time
func (c *Clock) Update() {
	c.currentTime = time.Now()
}

// DateString returns date in "Jan 2 2006" format
func (c *Clock) DateString() string {
	return c.currentTime.Format("Jan 2 2006")
}

// TimeString returns time in "15:04" format
func (c *Clock) TimeString() string {
	return c.currentTime.Format("15:04")
}

// DisplayLines returns date and time for display (two lines)
func (c *Clock) DisplayLines() (string, string) {
	return c.DateString(), c.TimeString()
}

// DisplayBig returns large time display with blinking colon
func (c *Clock) DisplayBig() string {
	hour := c.currentTime.Format("15")
	minute := c.currentTime.Format("04")

	// Colon blinks every second
	separator := ":"
	if c.currentTime.Second()%2 == 1 {
		separator = " "
	}

	return ui.RenderBigText(hour + separator + minute)
}
