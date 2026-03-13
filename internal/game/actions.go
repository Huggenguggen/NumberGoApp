package game

func (g *GameState) BuyUpgrade(key string) string {
	up, ok := g.Upgrades[key]
	if !ok {
		return "Unknown Upgrade"
	}

	if g.Levels.Level < up.RequiredLevel {
		return "Upgade locked by level"
	}

	cost := up.Cost()

	if g.Money < cost {
		return "Not enough money"
	}

	g.Money -= cost
	up.Apply(g)
	up.Purchased += 1
	return up.Name + " purchased!"
}
