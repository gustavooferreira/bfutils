package betting

import (
	"fmt"
	"math"

	"github.com/gustavooferreira/bfutils"
)

// FreeBetDecimal returns the P&L multiplier factor
// Example: back:4 lay:2 the multiplier factor is 2, which means if you back at odd 4 with £10
// and lay at odd 2 with £10 you secure a free bet of 2 * £10 = £20
func FreeBetDecimal(oddBack float64, oddLay float64) float64 {
	return oddBack - oddLay
}

// FreeBetAmount returns the profit in case selection wins.
// Note that 'stake' is the backer's stake not the layer's liability
func FreeBetAmount(oddBack float64, oddLay float64, stake float64) float64 {
	return FreeBetDecimal(oddBack, oddLay) * stake
}

// GreenBookOpenBackAmount returns lay stake to greenbook.
func GreenBookOpenBackAmount(oddBack float64, stakeBack float64, oddLay float64) float64 {
	return (stakeBack * oddBack) / oddLay
}

// GreenBookOpenBackDecimal returns percentage of P&L
func GreenBookOpenBackDecimal(oddBack float64, oddLay float64) float64 {
	return oddBack/oddLay - 1
}

// GreenBookOpenBackAmountByPerc returns oddLay for a given perc P&L
// Note that when Backing, you cannot lose more than 100% of your stake
// therefore feeding perc with a number less or equal to -1 is an error!
// perc is a representation in decimal, meaning if you want to know at what LAY odd you should
// place a bet at in order to get 100% profit, then perc is == 1
func GreenBookOpenBackAmountByPerc(oddBack float64, perc float64) float64 {
	return oddBack / (perc + 1)
}

// GreenBookOpenBackAmount returns back stake to greenbook.
func GreenBookOpenLayAmount(oddLay float64, stakeLay float64, oddBack float64) float64 {
	return (stakeLay * oddLay) / oddBack
}

// GreenBookOpenLayDecimal returns percentage of P&L
func GreenBookOpenLayDecimal(oddLay float64, oddBack float64) float64 {
	return 1 - oddLay/oddBack
}

// GreenBookOpenLayAmountByPerc returns oddBack for a given perc P&L
// Note that when Laying, you cannot win more than 100% of your stake
// therefore feeding perc with a number less or equal to -1 is an error!
func GreenBookOpenLayAmountByPerc(oddLay float64, perc float64) float64 {
	return oddLay / (1 - perc)
}

// GreenBookSelection computes odd and amount in order to greenbook a selection
func GreenBookSelection(bets []Bet, currentBackOdd float64, currentLayOdd float64) (b Bet, pl float64, err error) {
	layAvgOdd := 0.0
	layAmount := 0.0
	backAvgOdd := 0.0
	backAmount := 0.0

	if len(bets) == 0 {
		return b, 0, &NoBet{"no bets in this selection"}
	}

	// Check Odd is valid
	match, _, err := bfutils.FindOdd(currentBackOdd)
	if err != nil {
		return b, 0, err
	}
	if !match {
		return b, 0, &NoBet{fmt.Sprintf("odd provided [%f] does not exist in the ladder", currentBackOdd)}
	}

	match, _, err = bfutils.FindOdd(currentLayOdd)
	if err != nil {
		return b, 0, err
	}
	if !match {
		return b, 0, &NoBet{fmt.Sprintf("odd provided [%f] does not exist in the ladder", currentLayOdd)}
	}

	for _, bet := range bets {
		if bet.Amount == 0 {
			continue
		}

		// Check Odd is valid
		match, _, err := bfutils.FindOdd(bet.Odd)
		if err != nil {
			return b, 0, err
		}
		if !match {
			return b, 0, &NoBet{fmt.Sprintf("odd provided [%f] does not exist in the ladder", bet.Odd)}
		}

		if bet.Type == BetType_Back {
			backAvgOdd = (backAvgOdd*backAmount + bet.Odd*bet.Amount) / (backAmount + bet.Amount)
			backAmount += bet.Amount
		} else if bet.Type == BetType_Lay {
			layAvgOdd = (layAvgOdd*layAmount + bet.Odd*bet.Amount) / (layAmount + bet.Amount)
			layAmount += bet.Amount
		}
	}

	// Compute bet
	// Decide whether it's a BACK or LAY bet
	backBetAmount := (layAvgOdd*layAmount - backAvgOdd*backAmount) / currentBackOdd
	layBetAmount := (backAvgOdd*backAmount - layAvgOdd*layAmount) / currentLayOdd

	if equalWithTolerance(0.0, backBetAmount) && equalWithTolerance(0.0, layBetAmount) {
		return b, 0, &NoBet{"selection is already edged"}
	} else if backBetAmount > 0 {
		b.Type = BetType_Back
		b.Odd = currentBackOdd
		b.Amount = backBetAmount
		pl = backAmount*(backAvgOdd-1) - layAmount*(layAvgOdd-1) + backBetAmount*(currentBackOdd-1)
	} else if layBetAmount > 0 {
		b.Type = BetType_Lay
		b.Odd = currentLayOdd
		b.Amount = layBetAmount
		pl = backAmount*(backAvgOdd-1) - layAmount*(layAvgOdd-1) - layBetAmount*(currentLayOdd-1)
	} else {
		return b, 0, &NoBet{"unknown error"}
	}

	return b, pl, nil
}

type NoBet struct {
	msg string
}

func (e *NoBet) Error() string {
	return fmt.Sprintf("no bet: %s", e.msg)
}

// Helper function and constant to help estimate whether odd matches or not
func equalWithTolerance(a float64, b float64) bool {
	const float64EqualityThreshold = 1e-9
	return math.Abs(a-b) <= float64EqualityThreshold
}

// TODO: Provide multiple selections with array of lays and backs and work out what needs to be done.

// TODO: IsEdged func

// GreenBookAcrossSelection returns the ladder with P&L and Volumed matched by bets
func GreenBookAcrossSelection(bets []Bet) ([]LadderStep, error) {
	layAvgOdd := 0.0
	layAmount := 0.0
	backAvgOdd := 0.0
	backAmount := 0.0

	oddsMatched := map[float64]float64{}

	for _, bet := range bets {
		if bet.Amount == 0 {
			continue
		}

		// Check Odd is valid
		match, _, err := bfutils.FindOdd(bet.Odd)
		if err != nil {
			return nil, err
		}
		if !match {
			return nil, &NoBet{fmt.Sprintf("odd provided [%f] does not exist in the ladder", bet.Odd)}
		}

		oddsMatched[bet.Odd] += bet.Amount

		if bet.Type == BetType_Back {
			backAvgOdd = (backAvgOdd*backAmount + bet.Odd*bet.Amount) / (backAmount + bet.Amount)
			backAmount += bet.Amount
		} else if bet.Type == BetType_Lay {
			layAvgOdd = (layAvgOdd*layAmount + bet.Odd*bet.Amount) / (layAmount + bet.Amount)
			layAmount += bet.Amount
		}
	}

	ladder := make([]LadderStep, bfutils.OddsCount)

	for i, odd := range bfutils.Odds {
		ls := LadderStep{Odd: odd}

		// Compute bet
		// Decide whether it's a BACK or LAY bet
		backBetAmount := (layAvgOdd*layAmount - backAvgOdd*backAmount) / odd
		layBetAmount := (backAvgOdd*backAmount - layAvgOdd*layAmount) / odd

		if equalWithTolerance(0.0, backBetAmount) && equalWithTolerance(0.0, layBetAmount) {
			ls.GreenBookPL = 0.0
		} else if backBetAmount > 0 {
			ls.GreenBookPL = backAmount*(backAvgOdd-1) - layAmount*(layAvgOdd-1) + backBetAmount*(odd-1)
			ls.VolMatched = oddsMatched[odd] + backBetAmount
		} else if layBetAmount > 0 {
			ls.GreenBookPL = backAmount*(backAvgOdd-1) - layAmount*(layAvgOdd-1) - layBetAmount*(odd-1)
			ls.VolMatched = oddsMatched[odd] + layBetAmount
		} else {
			continue
		}

		ladder[i] = ls
	}

	return ladder, nil
}
