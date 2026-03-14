package save

import "numbergoapp/internal/game"

func FromGame(g *game.GameState) SaveData {

	data := SaveData{
		Users:             g.Users,
		TrafficCap:        g.TrafficCap,
		Money:             g.Money,
		UserGrowthRate:    g.UserGrowthRate,
		TrafficGrowthRate: g.TrafficGrowthRate,
		Level:             g.Levels.Level,
		UpgradePurchased:  map[string]int{},
	}

	for k, up := range g.Upgrades {
		data.UpgradePurchased[k] = up.Purchased
	}

	return data
}

func ToGame(data SaveData) *game.GameState {
	g := game.NewGame()

	g.Users = data.Users
	g.TrafficCap = data.TrafficCap
	g.Money = data.Money

	g.UserGrowthRate = data.UserGrowthRate
	g.TrafficGrowthRate = data.TrafficGrowthRate

	g.Levels.Level = data.Level

	for k, v := range data.UpgradePurchased {

		if up, ok := g.Upgrades[k]; ok {
			up.Purchased = v
		}
	}

	return g
}
