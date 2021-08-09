package betting

import "github.com/shopspring/decimal"

// Bet represents a bet in the market.
type Bet struct {
	// Bet type: Back or Lay.
	Type BetType
	// Odd in the market.
	Odd decimal.Decimal
	// Amount represents how much to bet or how much has been matched (backer's stake, or layer's payout)
	Amount decimal.Decimal
	// WinPL represents how much is the profit or loss in case this selection wins.
	// This value is meant to be treated as read-only.
	WinPL decimal.Decimal
	// LosePL represents how much is the profit or loss in case this selection loses.
	// This value is meant to be treated as read-only.
	LosePL decimal.Decimal
}

// Selection represents a selection in a market.
type Selection struct {
	// Bets matched in this specific selection.
	Bets []Bet
	// Current back odd being offered.
	CurrentBackOdd decimal.Decimal
	// Current lay odd being offered.
	CurrentLayOdd decimal.Decimal
}

// LadderStep represents a step in the trading ladder.
type LadderStep struct {
	// Odd in the market.
	Odd decimal.Decimal
	// Potential profit or loss in this selection in case of a greenbook operation.
	GreenBookPL decimal.Decimal
	// Volume matched by bets placed.
	VolMatched decimal.Decimal
}

// Ladder represents the trading ladder.
type Ladder []LadderStep
