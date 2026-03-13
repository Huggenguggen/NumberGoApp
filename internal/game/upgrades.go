package game

import "math"

type Upgrade struct {
	Name           string
	BaseCost       float64
	CostMultiplier float64
	RequiredLevel  int
	Purchased      int
	Apply          func(*GameState)
}

func (u *Upgrade) Cost() float64 {
	return u.BaseCost * math.Pow(u.CostMultiplier, float64(u.Purchased))
}

func DefaultUpgrades() map[string]*Upgrade {
	return map[string]*Upgrade{
		"user": {
			Name:           "User Expansion",
			BaseCost:       100,
			CostMultiplier: 1.15,
			RequiredLevel:  1,
			Purchased:      0,
			Apply: func(g *GameState) {
				g.UserGrowthRate += 1
			},
		},

		"traffic": {
			Name:           "Traffic Expansion",
			BaseCost:       100,
			CostMultiplier: 1.15,
			RequiredLevel:  1,
			Purchased:      0,
			Apply: func(g *GameState) {
				g.TrafficGrowthRate += 1
			},
		},
	}
}
