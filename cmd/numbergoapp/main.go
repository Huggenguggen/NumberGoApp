package main

import (
	"numbergoapp/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(ui.NewModel())

	if err := p.Start(); err != nil {
		panic(err)
	}
}
