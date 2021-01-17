// Package betting provides functions to compute bet types and sizes in order to achieve a free bet
// or greenbook.
package betting

import (
	"fmt"

	"github.com/gustavooferreira/bfutils"
	"github.com/gustavooferreira/bfutils/internal"
)

// FreeBetDecimal returns the P&L multiplier factor.
// Example: back:4 lay:2 the multiplier factor is 2, which means if you back at odd 4 with £10
// and lay at odd 2 with £10 you secure a free bet of 2 * £10 = £20
func FreeBetDecimal(oddBack float64, oddLay float64) float64 {
	return oddBack - oddLay
}

// FreeBetPL returns the profit in case selection wins.
// Note that 'stake' is the backer's stake not the layer's liability
func FreeBetPL(oddBack float64, oddLay float64, stake float64) float64 {
	return FreeBetDecimal(oddBack, oddLay) * stake
}

// GreenBookOpenBackDecimal returns percentage of P&L.
func GreenBookOpenBackDecimal(oddBack float64, oddLay float64) (float64, error) {
	if oddLay < 1.01 {
		return 0, fmt.Errorf("oddLay cannot be outside of trading range")
	}
	return oddBack/oddLay - 1, nil
}

// GreenBookOpenBackAmount returns lay stake to greenbook.
func GreenBookOpenBackAmount(oddBack float64, stakeBack float64, oddLay float64) (float64, error) {
	if oddLay < 1.01 {
		return 0, fmt.Errorf("oddLay cannot be outside of trading range")
	}
	return (stakeBack * oddBack) / oddLay, nil
}

// GreenBookOpenBackAmountByPerc returns oddLay for a given perc P&L
// Note that when Backing, you cannot lose more than 100% of your stake
// therefore feeding perc with a number less or equal to -1 is an error!
// perc is a representation in decimal, meaning if you want to know at what LAY odd you should
// place a bet at in order to get 100% profit, then perc is == 1
func GreenBookOpenBackAmountByPerc(oddBack float64, perc float64) (float64, error) {
	if perc <= -1 {
		return 0, fmt.Errorf("cannot lose more than 100%% of stake when backing")
	}
	return oddBack / (perc + 1), nil
}

// GreenBookOpenLayDecimal returns percentage of P&L
func GreenBookOpenLayDecimal(oddLay float64, oddBack float64) (float64, error) {
	if oddBack < 1.01 {
		return 0, fmt.Errorf("oddBack cannot be outside of trading range")
	}
	return 1 - oddLay/oddBack, nil
}

// GreenBookOpenLayAmount returns back stake to greenbook.
func GreenBookOpenLayAmount(oddLay float64, stakeLay float64, oddBack float64) (float64, error) {
	if oddBack < 1.01 {
		return 0, fmt.Errorf("oddBack cannot be outside of trading range")
	}
	return (stakeLay * oddLay) / oddBack, nil
}

// GreenBookOpenLayAmountByPerc returns oddBack for a given perc P&L
// Note that when Laying, you cannot win more than 100% of your stake
// therefore feeding perc with a number greater or equal to 1 is an error!
func GreenBookOpenLayAmountByPerc(oddLay float64, perc float64) (float64, error) {
	if perc >= 1 {
		return 0, fmt.Errorf("cannot win more than 100%% of stake when laying")
	}
	return oddLay / (1 - perc), nil
}

// -------------

// SelectionIsEdged returns true if selection is already been edged or if there are no bets in this selection.
// This might not give an accurate result in the sense that the selection might not be edged perfectly,
// because it might not be possible to edge it "even" across all outcomes at the current odds.
func SelectionIsEdged(bets []Bet) (bool, error) {
	layAvgOdd := 0.0
	layAmount := 0.0
	backAvgOdd := 0.0
	backAmount := 0.0

	for _, bet := range bets {
		if bet.Amount == 0 {
			continue
		}

		// Check Odd is valid
		match, _, err := bfutils.FindOdd(bet.Odd)
		if err != nil {
			return false, err
		}
		if !match {
			return false, fmt.Errorf("odd provided [%f] does not exist in the ladder", bet.Odd)
		}

		if bet.Type == BetType_Back {
			backAvgOdd = (backAvgOdd*backAmount + bet.Odd*bet.Amount) / (backAmount + bet.Amount)
			backAmount += bet.Amount
		} else if bet.Type == BetType_Lay {
			layAvgOdd = (layAvgOdd*layAmount + bet.Odd*bet.Amount) / (layAmount + bet.Amount)
			layAmount += bet.Amount
		} else {
			return false, fmt.Errorf("unknown bet type")
		}
	}

	if internal.EqualWithTolerance(0.0, backAvgOdd*backAmount-layAvgOdd*layAmount) {
		return true, nil
	}
	return false, nil
}

// GreenBookSelection computes what bet to make in order to greenbook a selection.
func GreenBookSelection(selection Selection) (bet Bet, err error) {
	layAvgOdd := 0.0
	layAmount := 0.0
	backAvgOdd := 0.0
	backAmount := 0.0

	bets := selection.Bets
	currentBackOdd := selection.CurrentBackOdd
	currentLayOdd := selection.CurrentLayOdd

	if bets == nil || len(bets) == 0 {
		return bet, fmt.Errorf("no bets in this selection")
	}

	// Check current back Odd is valid
	match, _, err := bfutils.FindOdd(currentBackOdd)
	if err != nil {
		return bet, err
	}
	if !match {
		return bet, fmt.Errorf("odd provided [%f] does not exist in the ladder", currentBackOdd)
	}

	// Check current lay Odd is valid
	match, _, err = bfutils.FindOdd(currentLayOdd)
	if err != nil {
		return bet, err
	}
	if !match {
		return bet, fmt.Errorf("odd provided [%f] does not exist in the ladder", currentLayOdd)
	}

	for _, b := range bets {
		if b.Amount == 0 {
			continue
		}

		// Check Odd is valid
		match, _, err := bfutils.FindOdd(b.Odd)
		if err != nil {
			return bet, err
		}
		if !match {
			return bet, fmt.Errorf("odd provided [%f] does not exist in the ladder", b.Odd)
		}

		if b.Type == BetType_Back {
			backAvgOdd = (backAvgOdd*backAmount + b.Odd*b.Amount) / (backAmount + b.Amount)
			backAmount += b.Amount
		} else if b.Type == BetType_Lay {
			layAvgOdd = (layAvgOdd*layAmount + b.Odd*b.Amount) / (layAmount + b.Amount)
			layAmount += b.Amount
		}
	}

	// Compute bet
	// Decide whether it's a BACK or LAY bet
	backBetAmount := (layAvgOdd*layAmount - backAvgOdd*backAmount) / currentBackOdd
	layBetAmount := (backAvgOdd*backAmount - layAvgOdd*layAmount) / currentLayOdd

	if internal.EqualWithTolerance(0.0, backBetAmount) && internal.EqualWithTolerance(0.0, layBetAmount) {
		return bet, &AlreadyEdgedError{}
	} else if backBetAmount > 0 {
		bet.Type = BetType_Back
		bet.Odd = currentBackOdd
		bet.Amount = backBetAmount
		bet.WinPL = backAmount*(backAvgOdd-1) - layAmount*(layAvgOdd-1) + backBetAmount*(currentBackOdd-1)
		bet.LosePL = layAmount - backAmount - backBetAmount
	} else if layBetAmount > 0 {
		bet.Type = BetType_Lay
		bet.Odd = currentLayOdd
		bet.Amount = layBetAmount
		bet.WinPL = backAmount*(backAvgOdd-1) - layAmount*(layAvgOdd-1) - layBetAmount*(currentLayOdd-1)
		bet.LosePL = layAmount - backAmount + layBetAmount
	}

	return bet, nil
}

// GreenBookAcrossSelections computes odd and amount in order to greenbook all selections.
// func GreenBookAcrossSelections(selections []Selection) (bets []Bet, err error) {
// 	return
// }

// GreenBookAtAllOdds returns the ladder with P&L and Volumed matched by bets.
func GreenBookAtAllOdds(bets []Bet) ([]LadderStep, error) {
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
			return nil, fmt.Errorf("odd provided [%f] does not exist in the ladder", bet.Odd)
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

		if internal.EqualWithTolerance(0.0, backBetAmount) && internal.EqualWithTolerance(0.0, layBetAmount) {
			ls.GreenBookPL = 0.0
		} else if backBetAmount > 0 {
			ls.GreenBookPL = backAmount*(backAvgOdd-1) - layAmount*(layAvgOdd-1) + backBetAmount*(odd-1)
			ls.VolMatched = oddsMatched[odd] + backBetAmount
		} else if layBetAmount > 0 {
			ls.GreenBookPL = backAmount*(backAvgOdd-1) - layAmount*(layAvgOdd-1) - layBetAmount*(odd-1)
			ls.VolMatched = oddsMatched[odd] + layBetAmount
		}

		ladder[i] = ls
	}

	return ladder, nil
}

// AlreadyEdgedError is the error used in case a betting operation cannot be performed.
type AlreadyEdgedError struct {
}

func (e *AlreadyEdgedError) Error() string {
	return "selection is already edged"
}
