package betting

type Bet struct {
	Type   BetType
	Odd    float64
	Amount float64
}

type LadderStep struct {
	Odd         float64
	GreenBookPL float64
	VolMatched  float64
}
