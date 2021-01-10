package betting_test

import (
	"fmt"
	"testing"

	"github.com/gustavooferreira/bfutils/betting"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const float64EqualityThreshold = 1e-5
const amountEqualityThreshold = 0.01

func TestCalcFreeBetDecimal(t *testing.T) {
	tests := map[string]struct {
		oddBack  float64
		oddLay   float64
		expected float64
	}{
		"calculate free bet [B:0,L:0]":       {oddBack: 0, oddLay: 0, expected: 0.0},
		"calculate free bet [B:2,L:2]":       {oddBack: 2, oddLay: 2, expected: 0.0},
		"calculate free bet [B:20,L:10]":     {oddBack: 20, oddLay: 10, expected: 10},
		"calculate free bet [B:18,L:7]":      {oddBack: 18, oddLay: 7, expected: 11},
		"calculate free bet [B:100,L:250]":   {oddBack: 100, oddLay: 250, expected: -150},
		"calculate free bet [B:50,L:5]":      {oddBack: 50, oddLay: 5, expected: 45},
		"calculate free bet [B:1.35,L:2.02]": {oddBack: 1.35, oddLay: 2.02, expected: -0.67},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			value := betting.FreeBetDecimal(test.oddBack, test.oddLay)
			assert.InDelta(t, test.expected, value, float64EqualityThreshold)
		})
	}
}

func TestCalcFreeBetAmount(t *testing.T) {
	tests := map[string]struct {
		oddBack  float64
		oddLay   float64
		stake    float64
		expected float64
	}{
		"calculate free bet [B:0,L:0;S:0]":       {oddBack: 0, oddLay: 0, stake: 0, expected: 0.0},
		"calculate free bet [B:2,L:2;S:10]":      {oddBack: 2, oddLay: 2, stake: 10, expected: 0.0},
		"calculate free bet [B:20,L:10;S:4]":     {oddBack: 20, oddLay: 10, stake: 4, expected: 40},
		"calculate free bet [B:18,L:7;S:0]":      {oddBack: 18, oddLay: 7, stake: 0, expected: 0},
		"calculate free bet [B:100,L:250;S:2]":   {oddBack: 100, oddLay: 250, stake: 2, expected: -300},
		"calculate free bet [B:50,L:5;S:5]":      {oddBack: 50, oddLay: 5, stake: 5, expected: 225},
		"calculate free bet [B:1.35,L:2.02;S:2]": {oddBack: 1.35, oddLay: 2.02, stake: 2, expected: -1.34},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			value := betting.FreeBetAmount(test.oddBack, test.oddLay, test.stake)
			assert.InDelta(t, test.expected, value, float64EqualityThreshold)
		})
	}
}

// func TestCalcGreenBookOpenBackDecimal(t *testing.T) {
// 	paramsTC := []struct {
// 		inOddBack float64
// 		inOddLay  float64
// 		out       float64
// 	}{
// 		{20, 10, 1},
// 		{18, 7, 1.57},
// 		{100, 250, -0.6},
// 		{50, 5, 9.0},
// 	}

// 	for i, tc := range paramsTC {
// 		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {

// 			result := GreenBookOpenBackDecimal(tc.inOddBack, tc.inOddLay)

// 			if !floatEqual(result, tc.out, 0.001) {
// 				t.Fatalf("Got '%.2f', wanted '%.2f'", result, tc.out)
// 			}
// 		})
// 	}
// }

// func TestCalcGreenBookOpenBackAmountByPerc(t *testing.T) {
// 	paramsTC := []struct {
// 		inOddBack float64
// 		inPerc    float64
// 		out       float64
// 	}{
// 		{20, 1, 10.0},
// 		{18, 0.7, 10.59},
// 		{100, 1.5, 40.0},
// 		{50, 2, 16.67},
// 		{10, -0.5, 20.0},
// 		{20, -0.8, 100.0},
// 	}

// 	for i, tc := range paramsTC {
// 		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {

// 			result := GreenBookOpenBackAmountByPerc(tc.inOddBack, tc.inPerc)

// 			if !floatEqual(result, tc.out, 0.001) {
// 				t.Fatalf("Got '%.2f', wanted '%.2f'", result, tc.out)
// 			}
// 		})
// 	}
// }

// func TestCalcGreenBookOpenLayDecimal(t *testing.T) {
// 	paramsTC := []struct {
// 		inOddLay  float64
// 		inOddBack float64
// 		out       float64
// 	}{
// 		{20, 10, -1},
// 		{18, 7, -1.57},
// 		{100, 250, 0.6},
// 		{50, 5, -9.0},
// 	}

// 	for i, tc := range paramsTC {
// 		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {

// 			result := GreenBookOpenLayDecimal(tc.inOddLay, tc.inOddBack)

// 			if !floatEqual(result, tc.out, 0.001) {
// 				t.Fatalf("Got '%.2f', wanted '%.2f'", result, tc.out)
// 			}
// 		})
// 	}
// }

func TestGreenBookSelection(t *testing.T) {
	tests := map[string]struct {
		selection      betting.Selection
		expectedBet    betting.Bet
		expectedWinPL  float64
		expectedLosePL float64
		expectedErr    bool
	}{
		"calculate greenbook 0": {expectedErr: true},
		"calculate greenbook 1": {selection: betting.Selection{
			Bets: []betting.Bet{},
		}, expectedErr: true},
		"calculate greenbook 2": {
			selection: betting.Selection{
				Bets: []betting.Bet{
					{Type: betting.BetType_Back, Odd: 4, Amount: 10},
					{Type: betting.BetType_Lay, Odd: 2, Amount: 5},
				},
				CurrentBackOdd: 2,
				CurrentLayOdd:  2.1,
			},
			expectedBet:    betting.Bet{Type: betting.BetType_Lay, Odd: 2.1, Amount: 14.29},
			expectedWinPL:  9.29,
			expectedLosePL: 9.29,
		},
		"calculate greenbook 3": {
			selection: betting.Selection{
				Bets: []betting.Bet{
					{Type: betting.BetType_Back, Odd: 3, Amount: 5},
				},
				CurrentBackOdd: 2,
				CurrentLayOdd:  2.1,
			},
			expectedBet:    betting.Bet{Type: betting.BetType_Lay, Odd: 2.1, Amount: 7.14},
			expectedWinPL:  2.15,
			expectedLosePL: 2.15,
		},
		"calculate greenbook 4": {
			selection: betting.Selection{
				Bets: []betting.Bet{
					{Type: betting.BetType_Lay, Odd: 3, Amount: 5},
				},
				CurrentBackOdd: 4,
				CurrentLayOdd:  4.2,
			},
			expectedBet:    betting.Bet{Type: betting.BetType_Back, Odd: 4, Amount: 3.75},
			expectedWinPL:  1.25,
			expectedLosePL: 1.25,
		},
		"calculate greenbook 5": {
			selection: betting.Selection{
				Bets: []betting.Bet{
					{Type: betting.BetType_Back, Odd: 4, Amount: 5},
					{Type: betting.BetType_Lay, Odd: 3, Amount: 5},
				},
				CurrentBackOdd: 3.1,
				CurrentLayOdd:  3,
			},
			expectedBet:    betting.Bet{Type: betting.BetType_Lay, Odd: 3, Amount: 1.67},
			expectedWinPL:  1.67,
			expectedLosePL: 1.67,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var errBool bool
			var errMsg string
			bet, err := betting.GreenBookSelection(test.selection)
			if err != nil {
				errBool = true
				errMsg = fmt.Sprintf(" - err: %s", err.Error())
			}

			require.Equal(t, test.expectedErr, errBool, "error field"+errMsg)

			assert.Equal(t, test.expectedBet.Type, bet.Type, "bet type field")
			assert.Equal(t, test.expectedBet.Odd, bet.Odd, "bet odd field")
			assert.InDelta(t, test.expectedBet.Amount, bet.Amount, amountEqualityThreshold, "bet amount field")

			assert.InDelta(t, test.expectedWinPL, bet.WinPL, amountEqualityThreshold, "Win P&L field")
			assert.InDelta(t, test.expectedLosePL, bet.LosePL, amountEqualityThreshold, "Lose P&L field")
		})
	}
}

func TestGreenBookAtAllOdds(t *testing.T) {
	tests := map[string]struct {
		bets               []betting.Bet
		index              int
		expectedLadderStep betting.LadderStep
		expectedErr        bool
	}{
		"calculate greenbook across selection 1": {
			bets: []betting.Bet{
				{Type: betting.BetType_Back, Odd: 4, Amount: 5},
				{Type: betting.BetType_Lay, Odd: 3, Amount: 5},
			},
			index: 159,
			expectedLadderStep: betting.LadderStep{
				Odd:         3.5,
				GreenBookPL: 1.43,
				VolMatched:  1.43,
			}},
		"calculate greenbook across selection 2": {
			bets: []betting.Bet{
				{Type: betting.BetType_Back, Odd: 2, Amount: 10},
				{Type: betting.BetType_Lay, Odd: 1.5, Amount: 5},
			},
			index: 49,
			expectedLadderStep: betting.LadderStep{
				Odd:         1.5,
				GreenBookPL: 3.34,
				VolMatched:  13.33,
			}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var errBool bool
			ladder, err := betting.GreenBookAtAllOdds(test.bets)
			if err != nil {
				errBool = true
			}

			require.Equal(t, test.expectedErr, errBool)

			value := ladder[test.index]

			assert.Equal(t, test.expectedLadderStep.Odd, value.Odd)
			assert.InDelta(t, test.expectedLadderStep.GreenBookPL, value.GreenBookPL, amountEqualityThreshold)
			assert.InDelta(t, test.expectedLadderStep.VolMatched, value.VolMatched, amountEqualityThreshold)
		})
	}
}
