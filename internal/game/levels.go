package game

type LevelSystem struct {
	Level      int
	Milestones []float64
}

func NewLevelSystem() *LevelSystem {
	return &LevelSystem{
		Level: 1,
		Milestones: []float64{
			0,
			10000,
			500000,
			10000000,
		},
	}
}

func (l *LevelSystem) UpdateLevel(money float64) {
	for i := len(l.Milestones) - 1; i >= 0; i-- {
		if money >= l.Milestones[i] {
			l.Level = i + 1
			return
		}
	}
}

func (l *LevelSystem) Progress(money float64) float64 {
	if l.Level >= len(l.Milestones) {
		return 1
	}

	prev := l.Milestones[l.Level-1]
	next := l.Milestones[l.Level]

	return (money - prev) / (next - prev)
}

func (l *LevelSystem) NextMilestone() float64 {
	if l.Level >= len(l.Milestones) {
		return -1
	}

	return l.Milestones[l.Level]
}
