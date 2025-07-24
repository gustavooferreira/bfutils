package betting

// Bet represents a bet in the market.
type Bet struct {
	// Bet type: Back or Lay.
	Type BetType
	// Odd in the market.
	Odd float64
	// Amount represents how much to bet or how much has been matched (backer's stake, or layer's payout)
	Amount float64
	// WinPL represents how much is the profit or loss in case this selection wins.
	// This value is meant to be treated as read-only.
	WinPL float64
	// LosePL represents how much is the profit or loss in case this selection loses.
	// This value is meant to be treated as read-only.
	LosePL float64
}

// Selection represents a selection in a market.
type Selection struct {
	// Bets matched in this specific selection.
	Bets []Bet
	// Current back odd being offered.
	CurrentBackOdd float64
	// Current lay odd being offered.
	CurrentLayOdd float64
}

// LadderStep represents a trading ladder.
type LadderStep struct {
	// Odd in the market.
	Odd float64
	// Potential profit or loss in this selection in case of a greenbook operation.
	GreenBookPL float64
	// Volume matched by bets placed.
	VolMatched float64
}

// BetType represents the type of bet to make in the market.
type BetType uint

const (
	// BetType_Back represents a back bet.
	BetType_Back = iota + 1
	// BetType_Lay represents a lay bet.
	BetType_Lay
)

// String returns the string representation of BetType.
func (bt BetType) String() string {
	return [...]string{"", "Back", "Lay"}[bt]
}
