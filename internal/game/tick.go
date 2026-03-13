package game

func (g *GameState) Tick() {
	g.TrafficCap += g.TrafficGrowthRate
	g.Users += g.UserGrowthRate

	if g.Users > g.TrafficCap {
		g.Users = g.TrafficCap
	}

	g.Money += g.Users * g.RevenuePerUser

	g.Levels.UpdateLevel(g.Money)
}
