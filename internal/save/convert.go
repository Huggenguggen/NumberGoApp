package save

import "numbergoapp/internal/game"

func FromGame(g *game.GameState) SaveData {
	return SaveData{
		Users:      g.Users,
		TrafficCap: g.TrafficCap,
		Money:      g.Money,

		UserGrowthRate:    g.UserGrowthRate,
		TrafficGrowthRate: g.TrafficGrowthRate,

		Level: g.Levels.Level,

		UserUpgradePurchased:    g.Upgrades["user"].Purchased,
		TrafficUpgradePurchased: g.Upgrades["traffic"].Purchased,
	}
}

func ToGame(data SaveData) *game.GameState {
	g := game.NewGame()

	g.Users = data.Users
	g.TrafficCap = data.TrafficCap
	g.Money = data.Money

	g.UserGrowthRate = data.UserGrowthRate
	g.TrafficGrowthRate = data.TrafficGrowthRate

	g.Levels.Level = data.Level

	g.Upgrades["user"].Purchased = data.UserUpgradePurchased
	g.Upgrades["traffic"].Purchased = data.TrafficUpgradePurchased

	return g
}
