// Package main is the entry point for the application
package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"crr/internal/model"
)

func main() {
	p := tea.NewProgram(model.NewDrums(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
