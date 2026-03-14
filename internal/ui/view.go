package ui

import (
	"fmt"
	"numbergoapp/internal/game"
)

func progressBar(progress float64) string {
	width := 30
	filled := int(progress * float64(width))

	bar := ""

	for i := 0; i < width; i++ {
		if i < filled {
			bar += "█"
		} else {
			bar += "░"
		}
	}

	return bar
}

func upgradeState(m Model, up *game.Upgrade) string {
	cost := up.Cost()

	switch {

	case m.Game.Levels.Level < up.RequiredLevel:
		return lockedStyle.Render("LOCKED")

	case m.Game.Money < cost:
		return expensiveStyle.Render("Too Expensive")

	default:
		return affordableStyle.Render("Available")
	}
}

func (m Model) View() string {
	g := m.Game

	progress := g.Levels.Progress(g.Money)
	next := g.Levels.NextMilestone()

	bar := progressBar(progress)

	upgradeText := ""

	for i, key := range game.UpgradeList {

		up := m.Game.Upgrades[key]

		state := upgradeState(m, up)

		line := fmt.Sprintf(
			"[%d] %-18s $%.0f - %d  %s\n",
			i+1,
			up.Name,
			up.Cost(),
			up.Purchased,
			state,
		)

		upgradeText += line
	}

	return fmt.Sprintf(
		`
NumberGoApp 🚀

Users:        %.0f
Traffic Cap:  %.0f
Money:        $%.2f

Revenue/tick: $%.2f

Level: %d
[%s]

Next milestone: $%.0f

Upgrades
%s

Status: %s

Press q to quit
`,
		g.Users,
		g.TrafficCap,
		g.Money,
		g.Users*g.RevenuePerUser,
		g.Levels.Level,
		bar,
		next,
		upgradeText,
		m.StatusMessage,
	)
}
