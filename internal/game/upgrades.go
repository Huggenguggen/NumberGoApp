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

var UpgradeList = []string{
	"user01",
	"traffic01",
	"user02",
}

var UpgradeMap = map[string]*Upgrade{
	"user01": {
		Name:           "User Expansion",
		BaseCost:       100,
		CostMultiplier: 1.15,
		RequiredLevel:  1,
		Purchased:      0,
		Apply: func(g *GameState) {
			g.UserGrowthRate += 1
		},
	},

	"traffic01": {
		Name:           "Traffic Expansion",
		BaseCost:       100,
		CostMultiplier: 1.15,
		RequiredLevel:  1,
		Purchased:      0,
		Apply: func(g *GameState) {
			g.TrafficGrowthRate += 1
		},
	},
	"user02": {
		Name:           "User Expansion 2",
		BaseCost:       500,
		CostMultiplier: 1.25,
		RequiredLevel:  2,
		Purchased:      0,
		Apply: func(g *GameState) {
			g.UserGrowthRate += 5
		},
	},
}
