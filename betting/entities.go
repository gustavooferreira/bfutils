package betting

type Bet struct {
	Type   BetType
	Odd    float64
	Amount float64
	// This value is meant to be treated as read-only
	WinPL float64
	// This value is meant to be treated as read-only
	LosePL float64
}

type Selection struct {
	Bets           []Bet
	CurrentBackOdd float64
	CurrentLayOdd  float64
}

type LadderStep struct {
	Odd         float64
	GreenBookPL float64
	VolMatched  float64
}
