package ui

// Digits contains ASCII-art representation of digits 0-9 and colon
// Each character is 5 lines tall
var Digits = map[rune][]string{
	'0': {
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	},
	'1': {
		" █ ",
		"██ ",
		" █ ",
		" █ ",
		"███",
	},
	'2': {
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	},
	'3': {
		"███",
		"  █",
		"███",
		"  █",
		"███",
	},
	'4': {
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	},
	'5': {
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	},
	'6': {
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	},
	'7': {
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	},
	'8': {
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	},
	'9': {
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	},
	':': {
		"   ",
		" █ ",
		"   ",
		" █ ",
		"   ",
	},
	' ': {
		"   ",
		"   ",
		"   ",
		"   ",
		"   ",
	},
}

// RenderBigText renders a string with large ASCII characters
func RenderBigText(s string) string {
	if len(s) == 0 {
		return ""
	}

	// Build lines for each level
	lines := make([]string, 5)

	for _, ch := range s {
		digit, ok := Digits[ch]
		if !ok {
			// Unknown character - skip or use space
			for i := range lines {
				lines[i] += "   "
			}
			continue
		}
		for i, row := range digit {
			lines[i] += row + " "
		}
	}

	// Join lines
	result := ""
	for i, line := range lines {
		result += line
		if i < len(lines)-1 {
			result += "\n"
		}
	}

	return result
}
