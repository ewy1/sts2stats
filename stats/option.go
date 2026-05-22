package stats

type ChoiceStat struct {
	Amount int
	Wins   int
}

type Choice struct {
	Taken   *ChoiceStat
	Ignored *ChoiceStat
}
