package save

type SaveData struct {
	Users      float64
	TrafficCap float64
	Money      float64

	UserGrowthRate    float64
	TrafficGrowthRate float64

	Level int

	UpgradePurchased map[string]int
}
