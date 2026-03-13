// Package engine implements the ticks
package engine

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type TickMsg struct{}

func Tick() tea.Cmd {
	return func() tea.Msg {
		time.Sleep(time.Second)
		return TickMsg{}
	}
}
