package model

// Volume represents volume level control
type Volume struct {
	Level    int  // Level from 0 to 100
	Muted    bool // Whether sound is muted
	MaxLevel int  // Maximum level
}

// NewVolume creates a new volume controller
func NewVolume() *Volume {
	return &Volume{
		Level:    50,
		Muted:    false,
		MaxLevel: 100,
	}
}

// Up increases volume by 5%
func (v *Volume) Up() {
	v.Level += 5
	if v.Level > v.MaxLevel {
		v.Level = v.MaxLevel
	}
}

// Down decreases volume by 5%
func (v *Volume) Down() {
	v.Level -= 5
	if v.Level < 0 {
		v.Level = 0
	}
}

// ToggleMute toggles mute state
func (v *Volume) ToggleMute() {
	v.Muted = !v.Muted
}

// SetLevel sets volume level
func (v *Volume) SetLevel(level int) {
	if level < 0 {
		level = 0
	}
	if level > v.MaxLevel {
		level = v.MaxLevel
	}
	v.Level = level
}

// DisplayBar returns visual volume bar representation
// width is the bar width in characters
func (v *Volume) DisplayBar(width int) string {
	if v.Muted {
		return "MUTE"
	}

	filled := (v.Level * width) / v.MaxLevel
	empty := width - filled

	bar := ""
	for i := 0; i < filled; i++ {
		bar += "#"
	}
	for i := 0; i < empty; i++ {
		bar += "-"
	}

	return "VOL " + bar
}
