// Package ui deals with TUI things
// mostly leaning on bubbletea and has 2 files
// model.go and view.go
package ui

import (
	tea "github.com/charmbracelet/bubbletea"

	"numbergoapp/internal/engine"
	"numbergoapp/internal/game"
	"numbergoapp/internal/save"
)

type Model struct {
	Game          *game.GameState
	StatusMessage string
}

func NewModel() Model {
	var data save.SaveData

	err := save.LoadGame("save.json", &data)

	if err == nil {
		gameState := save.ToGame(data)

		return Model{
			Game: gameState,
		}
	}
	return Model{
		Game: game.NewGame(),
	}
}

func (m Model) Init() tea.Cmd {
	return engine.Tick()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case engine.TickMsg:

		m.Game.Tick()

		data := save.FromGame(m.Game)
		save.SaveGame("save.json", data)

		return m, engine.Tick()
	case tea.KeyMsg:
		switch msg.(tea.KeyMsg).String() {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "1":
			m.StatusMessage = m.Game.BuyUpgrade("user")
			return m, nil

		case "2":
			m.StatusMessage = m.Game.BuyUpgrade("traffic")
			return m, nil
		}
	}

	return m, nil
}
