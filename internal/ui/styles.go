package ui

import "github.com/charmbracelet/lipgloss"

var (
	affordableStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("10")) // green

	lockedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("8")) // gray

	expensiveStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("9")) // red

	statusStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("14"))

	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("205"))
)
