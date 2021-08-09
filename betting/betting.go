// Package betting provides functions to compute bet types and sizes in order to achieve a free bet
// or greenbook.
package betting

import (
	"fmt"

	"github.com/gustavooferreira/bfutils"
	"github.com/shopspring/decimal"
)

// FreeBetDecimal returns percentage of P&L (or multiplier factor).
// Example: back:4 lay:2 the multiplier factor is 2, which means if you back at odd 4 with £10
// and lay at odd 2 with £10 you secure a free bet of 2 * £10 = £20
func FreeBetDecimal(oddBack decimal.Decimal, oddLay decimal.Decimal) decimal.Decimal {
	return oddBack.Sub(oddLay)
}

// FreeBetPL returns the profit in case selection wins.
// Note that 'stake' is the backer's stake not the layer's liability
func FreeBetPL(oddBack decimal.Decimal, oddLay decimal.Decimal, stake decimal.Decimal) decimal.Decimal {
	return FreeBetDecimal(oddBack, oddLay).Mul(stake)
}

// GreenBookOpenBackDecimal returns percentage of P&L.
func GreenBookOpenBackDecimal(oddBack decimal.Decimal, oddLay decimal.Decimal) (decimal.Decimal, error) {
	if oddLay.LessThan(bfutils.Odds[0]) {
		return decimal.Zero, fmt.Errorf("oddLay cannot be outside of trading range")
	}
	return oddBack.DivRound(oddLay, 2).Sub(decimal.RequireFromString("1")), nil
}

// GreenBookOpenBackAmount returns lay stake to greenbook.
func GreenBookOpenBackAmount(oddBack decimal.Decimal, stakeBack decimal.Decimal, oddLay decimal.Decimal) (decimal.Decimal, error) {
	if oddLay.LessThan(bfutils.Odds[0]) {
		return decimal.Zero, fmt.Errorf("oddLay cannot be outside of trading range")
	}
	return oddBack.DivRound(oddLay, 2).Mul(stakeBack), nil
}

// GreenBookOpenBackAmountByPerc returns oddLay for a given perc P&L.
// Note that when Backing, you cannot lose more than 100% of your stake
// therefore feeding perc with a number less or equal to -1 is an error!
// perc is a representation in decimal, meaning if you want to know at what LAY odd you should
// place a bet at in order to get 100% profit, then perc is == 1
func GreenBookOpenBackAmountByPerc(oddBack decimal.Decimal, perc decimal.Decimal) (decimal.Decimal, error) {
	if perc.LessThanOrEqual(decimal.RequireFromString("-1")) {
		return decimal.Zero, fmt.Errorf("cannot lose more than 100%% of stake when backing")
	}
	temp := perc.Add(decimal.RequireFromString("1"))
	return oddBack.DivRound(temp, 2), nil
}

// GreenBookOpenLayDecimal returns percentage of P&L.
func GreenBookOpenLayDecimal(oddLay decimal.Decimal, oddBack decimal.Decimal) (decimal.Decimal, error) {
	if oddLay.LessThan(bfutils.Odds[0]) {
		return decimal.Zero, fmt.Errorf("oddBack cannot be outside of trading range")
	}
	temp := oddLay.DivRound(oddBack, 2)
	return decimal.RequireFromString("1").Sub(temp), nil
}

// GreenBookOpenLayAmount returns back stake to greenbook.
func GreenBookOpenLayAmount(oddLay decimal.Decimal, stakeLay decimal.Decimal, oddBack decimal.Decimal) (decimal.Decimal, error) {
	if oddLay.LessThan(bfutils.Odds[0]) {
		return decimal.Zero, fmt.Errorf("oddBack cannot be outside of trading range")
	}
	temp := stakeLay.Mul(oddLay)
	return temp.DivRound(oddBack, 2), nil
}

// GreenBookOpenLayAmountByPerc returns oddBack for a given perc P&L.
// Note that when Laying, you cannot win more than 100% of your stake
// therefore feeding perc with a number greater or equal to 1 is an error!
func GreenBookOpenLayAmountByPerc(oddLay decimal.Decimal, perc decimal.Decimal) (decimal.Decimal, error) {
	if perc.GreaterThanOrEqual(decimal.RequireFromString("1")) {
		return decimal.Zero, fmt.Errorf("cannot win more than 100%% of stake when laying")
	}
	temp := decimal.RequireFromString("1").Sub(perc)
	return oddLay.DivRound(temp, 2), nil
}

// -------------

// SelectionIsEdged returns true if selection is already been edged or if there are no bets in this selection.
// This might not give an accurate result in the sense that the selection might not be edged perfectly,
// because it might not be possible to edge it "even" across all outcomes at the current odds.
func SelectionIsEdged(bets []Bet) (bool, error) {
	if len(bets) == 0 {
		return true, nil
	}

	layAvgOdd := decimal.Decimal{}
	layAmount := decimal.Decimal{}
	backAvgOdd := decimal.Decimal{}
	backAmount := decimal.Decimal{}

	for _, bet := range bets {
		if bet.Amount.Equal(decimal.Zero) {
			continue
		}

		// Check Odd is valid
		match, _, err := bfutils.FindOdd(bet.Odd)
		if err != nil {
			return false, err
		}
		if !match {
			return false, fmt.Errorf("odd provided [%s] does not exist in the ladder", bet.Odd.String())
		}

		if bet.Type == BetType_Back {
			numerator := backAvgOdd.Mul(backAmount).Add(bet.Odd.Mul(bet.Amount))
			denominator := backAmount.Add(bet.Amount)
			backAvgOdd = numerator.DivRound(denominator, 2)
			backAmount = backAmount.Add(bet.Amount)
		} else if bet.Type == BetType_Lay {
			numerator := layAvgOdd.Mul(layAmount).Add(bet.Odd.Mul(bet.Amount))
			denominator := layAmount.Add(bet.Amount)
			layAvgOdd = numerator.DivRound(denominator, 2)
			layAmount = layAmount.Add(bet.Amount)
		} else {
			return false, fmt.Errorf("unknown bet type")
		}
	}

	backTotal := backAvgOdd.Mul(backAmount)
	layTotal := layAvgOdd.Mul(layAmount)
	sum := backTotal.Sub(layTotal)

	if sum.IsZero() {
		return true, nil
	}
	return false, nil
}

// GreenBookSelection computes what bet to make in order to greenbook a selection.
func GreenBookSelection(selection Selection) (bet Bet, err error) {
	one := decimal.RequireFromString("1")

	layAvgOdd := decimal.Decimal{}
	layAmount := decimal.Decimal{}
	backAvgOdd := decimal.Decimal{}
	backAmount := decimal.Decimal{}

	bets := selection.Bets
	currentBackOdd := selection.CurrentBackOdd
	currentLayOdd := selection.CurrentLayOdd

	if len(bets) == 0 {
		return bet, fmt.Errorf("no bets in this selection")
	}

	// Check current back Odd is valid
	match, _, err := bfutils.FindOdd(currentBackOdd)
	if err != nil {
		return bet, err
	}
	if !match {
		return bet, fmt.Errorf("odd provided [%s] does not exist in the ladder", currentBackOdd.String())
	}

	// Check current lay Odd is valid
	match, _, err = bfutils.FindOdd(currentLayOdd)
	if err != nil {
		return bet, err
	}
	if !match {
		return bet, fmt.Errorf("odd provided [%s] does not exist in the ladder", currentLayOdd.String())
	}

	for _, b := range bets {
		if b.Amount.Equal(decimal.Zero) {
			continue
		}

		// Check Odd is valid
		match, _, err := bfutils.FindOdd(b.Odd)
		if err != nil {
			return b, err
		}
		if !match {
			return b, fmt.Errorf("odd provided [%s] does not exist in the ladder", b.Odd.String())
		}

		if b.Type == BetType_Back {
			numerator := backAvgOdd.Mul(backAmount).Add(b.Odd.Mul(b.Amount))
			denominator := backAmount.Add(b.Amount)
			backAvgOdd = numerator.DivRound(denominator, 2)
			backAmount = backAmount.Add(b.Amount)
		} else if b.Type == BetType_Lay {
			numerator := layAvgOdd.Mul(layAmount).Add(b.Odd.Mul(b.Amount))
			denominator := layAmount.Add(b.Amount)
			layAvgOdd = numerator.DivRound(denominator, 2)
			layAmount = layAmount.Add(b.Amount)
		} else {
			return bet, fmt.Errorf("unknown bet type")
		}
	}

	// Compute bet
	// Decide whether it's a BACK or LAY bet
	numerator := layAvgOdd.Mul(layAmount).Sub(backAvgOdd.Mul(backAmount))
	backBetAmount := numerator.DivRound(currentBackOdd, 2)

	numerator = backAvgOdd.Mul(backAmount).Sub(layAvgOdd.Mul(layAmount))
	layBetAmount := numerator.DivRound(currentLayOdd, 2)

	if backBetAmount.IsZero() && layBetAmount.IsZero() {
		return bet, &AlreadyEdgedError{}
	} else if backBetAmount.GreaterThan(decimal.Zero) {
		bet.Type = BetType_Back
		bet.Odd = currentBackOdd
		bet.Amount = backBetAmount
		bet.WinPL = backAmount.Mul(backAvgOdd.Sub(one)).
			Sub(layAmount.Mul(layAvgOdd.Sub(one))).
			Add(backBetAmount.Mul(currentBackOdd.Sub(one)))
		bet.LosePL = layAmount.Sub(backAmount).Sub(backBetAmount)
	} else if layBetAmount.GreaterThan(decimal.Zero) {
		bet.Type = BetType_Lay
		bet.Odd = currentLayOdd
		bet.Amount = layBetAmount
		bet.WinPL = backAmount.Mul(backAvgOdd.Sub(one)).
			Sub(layAmount.Mul(layAvgOdd.Sub(one))).
			Sub(layBetAmount.Mul(currentLayOdd.Sub(one)))
		bet.LosePL = layAmount.Sub(backAmount).Add(layBetAmount)
	}

	return bet, nil
}

// GreenBookAcrossSelections computes odd and amount in order to greenbook all selections.
// func GreenBookAcrossSelections(selections []Selection) (bets []Bet, err error) {
// 	return
// }

// // GreenBookAtAllOdds returns the ladder with P&L and volumed matched by bets.
// func GreenBookAtAllOdds(bets []Bet) (Ladder, error) {
// 	layAvgOdd := 0.0
// 	layAmount := 0.0
// 	backAvgOdd := 0.0
// 	backAmount := 0.0

// 	oddsMatched := map[float64]float64{}

// 	for _, bet := range bets {
// 		if bet.Amount == 0 {
// 			continue
// 		}

// 		// Check Odd is valid
// 		match, _, err := bfutils.FindOdd(bet.Odd)
// 		if err != nil {
// 			return nil, err
// 		}
// 		if !match {
// 			return nil, fmt.Errorf("odd provided [%f] does not exist in the ladder", bet.Odd)
// 		}

// 		oddsMatched[bet.Odd] += bet.Amount

// 		if bet.Type == BetType_Back {
// 			backAvgOdd = (backAvgOdd*backAmount + bet.Odd*bet.Amount) / (backAmount + bet.Amount)
// 			backAmount += bet.Amount
// 		} else if bet.Type == BetType_Lay {
// 			layAvgOdd = (layAvgOdd*layAmount + bet.Odd*bet.Amount) / (layAmount + bet.Amount)
// 			layAmount += bet.Amount
// 		}
// 	}

// 	ladder := make(Ladder, bfutils.OddsCount)

// 	for i, odd := range bfutils.Odds {
// 		ls := LadderStep{Odd: odd}

// 		// Compute bet
// 		// Decide whether it's a BACK or LAY bet
// 		backBetAmount := (layAvgOdd*layAmount - backAvgOdd*backAmount) / odd
// 		layBetAmount := (backAvgOdd*backAmount - layAvgOdd*layAmount) / odd

// 		if internal.EqualWithTolerance(0.0, backBetAmount) && internal.EqualWithTolerance(0.0, layBetAmount) {
// 			ls.GreenBookPL = 0.0
// 		} else if backBetAmount > 0 {
// 			ls.GreenBookPL = backAmount*(backAvgOdd-1) - layAmount*(layAvgOdd-1) + backBetAmount*(odd-1)
// 			ls.VolMatched = oddsMatched[odd] + backBetAmount
// 		} else if layBetAmount > 0 {
// 			ls.GreenBookPL = backAmount*(backAvgOdd-1) - layAmount*(layAvgOdd-1) - layBetAmount*(odd-1)
// 			ls.VolMatched = oddsMatched[odd] + layBetAmount
// 		}

// 		ladder[i] = ls
// 	}

// 	return ladder, nil
// }

// AlreadyEdgedError is the error used in case a selection is already edged.
type AlreadyEdgedError struct {
}

func (e *AlreadyEdgedError) Error() string {
	return "selection is already edged"
}
