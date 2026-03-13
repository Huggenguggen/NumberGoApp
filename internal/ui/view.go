package ui

import "fmt"

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

func upgradeText(m Model, key string) string {
	up := m.Game.Upgrades[key]
	cost := up.Cost()

	label := ""

	switch {

	case m.Game.Levels.Level < up.RequiredLevel:
		label = lockedStyle.Render("LOCKED")

	case m.Game.Money < cost:
		label = expensiveStyle.Render("Too Expensive")

	default:
		label = affordableStyle.Render("Available")
	}

	return label
}

func (m Model) View() string {
	g := m.Game

	progress := g.Levels.Progress(g.Money)
	next := g.Levels.NextMilestone()

	bar := progressBar(progress)

	userUp := g.Upgrades["user"]
	trafficUp := g.Upgrades["traffic"]

	userCost := userUp.Cost()
	trafficCost := trafficUp.Cost()

	userState := upgradeText(m, "user")
	trafficState := upgradeText(m, "traffic")

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
[1] %s		$%.0f - %d 	%s
[2] %s		$%.0f - %d 	%s

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
		userUp.Name,
		userCost,
		userUp.Purchased,
		userState,
		trafficUp.Name,
		trafficCost,
		trafficUp.Purchased,
		trafficState,
		m.StatusMessage,
	)
}
