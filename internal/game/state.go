// Package game implements game logic such as
// State, Actions, Upgrades etc
package game

type GameState struct {
	Users      float64
	TrafficCap float64
	Money      float64

	UserGrowthRate    float64
	TrafficGrowthRate float64

	RevenuePerUser float64

	Levels   *LevelSystem
	Upgrades map[string]*Upgrade
}

func NewGame() *GameState {

	g := &GameState{
		Users:             100,
		TrafficCap:        100,
		Money:             0,
		UserGrowthRate:    1,
		TrafficGrowthRate: 1,
		RevenuePerUser:    0.1,
		Levels:            NewLevelSystem(),
		Upgrades:          map[string]*Upgrade{},
	}

	for _, key := range UpgradeList {

		def := UpgradeMap[key]

		copy := *def
		copy.Purchased = 0

		g.Upgrades[key] = &copy
	}

	return g
}
