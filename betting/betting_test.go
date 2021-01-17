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

func TestCalcFreeBetPL(t *testing.T) {
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
			value := betting.FreeBetPL(test.oddBack, test.oddLay, test.stake)
			assert.InDelta(t, test.expected, value, float64EqualityThreshold)
		})
	}
}

func TestCalcGreenBookOpenBackDecimal(t *testing.T) {
	tests := map[string]struct {
		oddBack      float64
		oddLay       float64
		expectedPerc float64
		expectedErr  bool
	}{
		"calculate green book [B:0,L:0]":       {oddBack: 0, oddLay: 0, expectedErr: true},
		"calculate green book [B:2,L:2]":       {oddBack: 2, oddLay: 2, expectedPerc: 0.0},
		"calculate green book [B:20,L:10]":     {oddBack: 20, oddLay: 10, expectedPerc: 1},
		"calculate green book [B:18,L:7]":      {oddBack: 18, oddLay: 7, expectedPerc: 1.57142},
		"calculate green book [B:100,L:250]":   {oddBack: 100, oddLay: 250, expectedPerc: -0.6},
		"calculate green book [B:50,L:5]":      {oddBack: 50, oddLay: 5, expectedPerc: 9},
		"calculate green book [B:1.35,L:2.02]": {oddBack: 1.35, oddLay: 2.02, expectedPerc: -0.33168},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var errBool bool
			var errMsg string
			value, err := betting.GreenBookOpenBackDecimal(test.oddBack, test.oddLay)
			if err != nil {
				errBool = true
				errMsg = fmt.Sprintf(" - err: %s", err.Error())
			}
			require.Equal(t, test.expectedErr, errBool, "error field"+errMsg)

			assert.InDelta(t, test.expectedPerc, value, float64EqualityThreshold)
		})
	}
}

func TestCalcGreenBookOpenBackAmount(t *testing.T) {
	tests := map[string]struct {
		oddBack          float64
		oddLay           float64
		stake            float64
		expectedLayStake float64
		expectedErr      bool
	}{
		"calculate green book amount [B:0,L:0;S:0]":       {oddBack: 0, oddLay: 0, stake: 0, expectedErr: true},
		"calculate green book amount [B:2,L:2;S:10]":      {oddBack: 2, oddLay: 2, stake: 10, expectedLayStake: 10},
		"calculate green book amount [B:20,L:10;S:4]":     {oddBack: 20, oddLay: 10, stake: 4, expectedLayStake: 8},
		"calculate green book amount [B:18,L:7;S:0]":      {oddBack: 18, oddLay: 7, stake: 0, expectedLayStake: 0},
		"calculate green book amount [B:100,L:250;S:2]":   {oddBack: 100, oddLay: 250, stake: 2, expectedLayStake: 0.8},
		"calculate green book amount [B:50,L:5;S:5]":      {oddBack: 50, oddLay: 5, stake: 5, expectedLayStake: 50},
		"calculate green book amount [B:1.35,L:2.02;S:2]": {oddBack: 1.35, oddLay: 2.02, stake: 2, expectedLayStake: 1.33663366},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var errBool bool
			var errMsg string
			value, err := betting.GreenBookOpenBackAmount(test.oddBack, test.stake, test.oddLay)
			if err != nil {
				errBool = true
				errMsg = fmt.Sprintf(" - err: %s", err.Error())
			}
			require.Equal(t, test.expectedErr, errBool, "error field"+errMsg)

			assert.InDelta(t, test.expectedLayStake, value, float64EqualityThreshold)
		})
	}
}

func TestCalcGreenBookOpenBackAmountByPerc(t *testing.T) {
	tests := map[string]struct {
		oddBack          float64
		perc             float64
		expectedLayStake float64
		expectedErr      bool
	}{
		"calculate green book perc 1": {oddBack: 5, perc: -2, expectedErr: true},
		"calculate green book perc 2": {oddBack: 20, perc: 1, expectedLayStake: 10},
		"calculate green book perc 3": {oddBack: 18, perc: 0.7, expectedLayStake: 10.5882352},
		"calculate green book perc 4": {oddBack: 100, perc: 1.5, expectedLayStake: 40},
		"calculate green book perc 5": {oddBack: 50, perc: 2, expectedLayStake: 16.666666},
		"calculate green book perc 6": {oddBack: 10, perc: -0.5, expectedLayStake: 20},
		"calculate green book perc 7": {oddBack: 20, perc: -0.8, expectedLayStake: 100},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var errBool bool
			var errMsg string
			value, err := betting.GreenBookOpenBackAmountByPerc(test.oddBack, test.perc)
			if err != nil {
				errBool = true
				errMsg = fmt.Sprintf(" - err: %s", err.Error())
			}
			require.Equal(t, test.expectedErr, errBool, "error field"+errMsg)

			assert.InDelta(t, test.expectedLayStake, value, float64EqualityThreshold)
		})
	}
}

func TestCalcGreenBookOpenLayDecimal(t *testing.T) {
	tests := map[string]struct {
		oddLay       float64
		oddBack      float64
		expectedPerc float64
		expectedErr  bool
	}{
		"calculate green book 1": {oddLay: 0, oddBack: 0, expectedErr: true},
		"calculate green book 2": {oddLay: 20, oddBack: 10, expectedPerc: -1},
		"calculate green book 3": {oddLay: 18, oddBack: 7, expectedPerc: -1.5714285714},
		"calculate green book 4": {oddLay: 100, oddBack: 250, expectedPerc: 0.6},
		"calculate green book 5": {oddLay: 50, oddBack: 5, expectedPerc: -9},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var errBool bool
			var errMsg string
			value, err := betting.GreenBookOpenLayDecimal(test.oddLay, test.oddBack)
			if err != nil {
				errBool = true
				errMsg = fmt.Sprintf(" - err: %s", err.Error())
			}
			require.Equal(t, test.expectedErr, errBool, "error field"+errMsg)

			assert.InDelta(t, test.expectedPerc, value, float64EqualityThreshold)
		})
	}
}

func TestCalcGreenBookOpenLayAmount(t *testing.T) {
	tests := map[string]struct {
		oddBack          float64
		oddLay           float64
		stake            float64
		expectedLayStake float64
		expectedErr      bool
	}{
		"calculate green book amount 1": {oddBack: 0, oddLay: 0, stake: 0, expectedErr: true},
		"calculate green book amount 2": {oddBack: 2, oddLay: 2, stake: 10, expectedLayStake: 10},
		"calculate green book amount 3": {oddBack: 20, oddLay: 10, stake: 4, expectedLayStake: 2},
		"calculate green book amount 4": {oddBack: 18, oddLay: 7, stake: 0, expectedLayStake: 0},
		"calculate green book amount 5": {oddBack: 100, oddLay: 250, stake: 2, expectedLayStake: 5},
		"calculate green book amount 6": {oddBack: 50, oddLay: 5, stake: 5, expectedLayStake: 0.5},
		"calculate green book amount 7": {oddBack: 1.35, oddLay: 2.02, stake: 2, expectedLayStake: 2.992592592},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var errBool bool
			var errMsg string
			value, err := betting.GreenBookOpenLayAmount(test.oddLay, test.stake, test.oddBack)
			if err != nil {
				errBool = true
				errMsg = fmt.Sprintf(" - err: %s", err.Error())
			}
			require.Equal(t, test.expectedErr, errBool, "error field"+errMsg)

			assert.InDelta(t, test.expectedLayStake, value, float64EqualityThreshold)
		})
	}
}

func TestCalcGreenBookOpenLayAmountByPerc(t *testing.T) {
	tests := map[string]struct {
		oddLay           float64
		perc             float64
		expectedLayStake float64
		expectedErr      bool
	}{
		"calculate green book perc 1": {oddLay: 5, perc: 2, expectedErr: true},
		"calculate green book perc 2": {oddLay: 20, perc: -1, expectedLayStake: 10},
		"calculate green book perc 3": {oddLay: 18, perc: 0.75, expectedLayStake: 72},
		"calculate green book perc 4": {oddLay: 100, perc: -1.5, expectedLayStake: 40},
		"calculate green book perc 5": {oddLay: 50, perc: -2, expectedLayStake: 16.666666},
		"calculate green book perc 6": {oddLay: 10, perc: 0.5, expectedLayStake: 20},
		"calculate green book perc 7": {oddLay: 20, perc: 0.8, expectedLayStake: 100},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var errBool bool
			var errMsg string
			value, err := betting.GreenBookOpenLayAmountByPerc(test.oddLay, test.perc)
			if err != nil {
				errBool = true
				errMsg = fmt.Sprintf(" - err: %s", err.Error())
			}
			require.Equal(t, test.expectedErr, errBool, "error field"+errMsg)

			assert.InDelta(t, test.expectedLayStake, value, float64EqualityThreshold)
		})
	}
}

func TestSelectionIsEdged(t *testing.T) {
	tests := map[string]struct {
		bets          []betting.Bet
		expectedEdged bool
		expectedErr   bool
	}{
		"calculate greenbook 1": {
			bets:        []betting.Bet{{Type: 0, Odd: 2, Amount: 2}},
			expectedErr: true,
		},
		"calculate greenbook 2": {
			bets:        []betting.Bet{{Type: betting.BetType_Back, Odd: 2.111, Amount: 2}},
			expectedErr: true,
		},
		"calculate greenbook 3": {
			bets: []betting.Bet{
				{Type: betting.BetType_Back, Odd: 2, Amount: 0},
				{Type: betting.BetType_Lay, Odd: 0.5, Amount: 2},
			},
			expectedErr: true,
		},
		"calculate greenbook 4": {
			bets: []betting.Bet{
				{Type: betting.BetType_Back, Odd: 2.0, Amount: 10},
				{Type: betting.BetType_Lay, Odd: 1.5, Amount: 5}},
			expectedEdged: false,
		},
		"calculate greenbook 5": {
			bets: []betting.Bet{
				{Type: betting.BetType_Back, Odd: 4.0, Amount: 10},
				{Type: betting.BetType_Lay, Odd: 3.4, Amount: 10},
				{Type: betting.BetType_Lay, Odd: 3, Amount: 2}},
			expectedEdged: true,
		},
		"calculate greenbook 6": {
			bets: []betting.Bet{
				{Type: betting.BetType_Back, Odd: 16.5, Amount: 50},
				{Type: betting.BetType_Lay, Odd: 16, Amount: 51.56}},
			expectedEdged: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var errBool bool
			var errMsg string
			edgedFlag, err := betting.SelectionIsEdged(test.bets)
			if err != nil {
				errBool = true
				errMsg = fmt.Sprintf(" - err: %s", err.Error())
			}

			require.Equal(t, test.expectedErr, errBool, "error field"+errMsg)

			assert.Equal(t, test.expectedEdged, edgedFlag, "edged result")
		})
	}
}

func TestGreenBookSelection(t *testing.T) {
	tests := map[string]struct {
		selection      betting.Selection
		expectedBet    betting.Bet
		expectedWinPL  float64
		expectedLosePL float64
		expectedErr    bool
	}{
		"calculate greenbook 1": {expectedErr: true},
		"calculate greenbook 2": {selection: betting.Selection{
			Bets: []betting.Bet{},
		}, expectedErr: true},
		"calculate greenbook 3": {selection: betting.Selection{
			Bets: []betting.Bet{{Type: betting.BetType_Back, Odd: 1.5, Amount: 10}},
		}, expectedErr: true},
		"calculate greenbook 4": {selection: betting.Selection{
			Bets:           []betting.Bet{{Type: betting.BetType_Back, Odd: 1.5, Amount: 10}},
			CurrentBackOdd: 1.525,
		}, expectedErr: true},
		"calculate greenbook 5": {selection: betting.Selection{
			Bets:           []betting.Bet{{Type: betting.BetType_Back, Odd: 1.5, Amount: 10}},
			CurrentBackOdd: 1.5,
		}, expectedErr: true},
		"calculate greenbook 6": {selection: betting.Selection{
			Bets:           []betting.Bet{{Type: betting.BetType_Back, Odd: 1.5, Amount: 10}},
			CurrentBackOdd: 1.5,
			CurrentLayOdd:  1.525,
		}, expectedErr: true},
		"calculate greenbook 7": {
			selection: betting.Selection{
				Bets: []betting.Bet{
					{Type: betting.BetType_Back, Odd: 4, Amount: 0},
					{Type: betting.BetType_Lay, Odd: 1, Amount: 5},
				},
				CurrentBackOdd: 2,
				CurrentLayOdd:  2.1,
			}, expectedErr: true,
		},
		"calculate greenbook 8": {
			selection: betting.Selection{
				Bets: []betting.Bet{
					{Type: betting.BetType_Lay, Odd: 1.525, Amount: 5},
				},
				CurrentBackOdd: 2,
				CurrentLayOdd:  2.1,
			}, expectedErr: true,
		},
		"calculate greenbook 9": {
			selection: betting.Selection{
				Bets: []betting.Bet{
					{Type: betting.BetType_Back, Odd: 4.0, Amount: 10},
					{Type: betting.BetType_Lay, Odd: 3.4, Amount: 10},
					{Type: betting.BetType_Lay, Odd: 3, Amount: 2}},
				CurrentBackOdd: 2,
				CurrentLayOdd:  2.1,
			}, expectedErr: true,
		},
		"calculate greenbook 10": {
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
		"calculate greenbook 11": {
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
		"calculate greenbook 12": {
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
		"calculate greenbook 13": {
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
		"calculate greenbook at all odds 1": {
			bets: []betting.Bet{
				{Type: betting.BetType_Back, Odd: 4, Amount: 0},
				{Type: betting.BetType_Lay, Odd: 1, Amount: 5},
			},
			expectedErr: true,
		},
		"calculate greenbook at all odds 2": {
			bets: []betting.Bet{
				{Type: betting.BetType_Lay, Odd: 1.525, Amount: 5},
			},
			expectedErr: true,
		},
		"calculate greenbook at all odds 3": {
			bets: []betting.Bet{
				{Type: betting.BetType_Back, Odd: 4.0, Amount: 10},
				{Type: betting.BetType_Lay, Odd: 3.4, Amount: 10},
				{Type: betting.BetType_Lay, Odd: 3, Amount: 2},
			},
			index: 0,
			expectedLadderStep: betting.LadderStep{
				Odd:         1.01,
				GreenBookPL: 0,
				VolMatched:  0,
			},
		},
		"calculate greenbook at all odds 4": {
			bets: []betting.Bet{
				{Type: betting.BetType_Back, Odd: 4, Amount: 5},
				{Type: betting.BetType_Lay, Odd: 3, Amount: 5},
			},
			index: 159,
			expectedLadderStep: betting.LadderStep{
				Odd:         3.5,
				GreenBookPL: 1.43,
				VolMatched:  1.43,
			},
		},
		"calculate greenbook at all odds 5": {
			bets: []betting.Bet{
				{Type: betting.BetType_Back, Odd: 2, Amount: 10},
				{Type: betting.BetType_Lay, Odd: 1.5, Amount: 5},
			},
			index: 49,
			expectedLadderStep: betting.LadderStep{
				Odd:         1.5,
				GreenBookPL: 3.34,
				VolMatched:  13.33,
			},
		},
		"calculate greenbook at all odds 6": {
			bets: []betting.Bet{
				{Type: betting.BetType_Lay, Odd: 1.5, Amount: 5},
			},
			index: 99,
			expectedLadderStep: betting.LadderStep{
				Odd:         2,
				GreenBookPL: 1.25,
				VolMatched:  3.75,
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var errBool bool
			var errMsg string
			ladder, err := betting.GreenBookAtAllOdds(test.bets)
			if err != nil {
				errBool = true
				errMsg = fmt.Sprintf(" - err: %s", err.Error())
				require.Equal(t, test.expectedErr, errBool, "error field"+errMsg)
				return
			}

			value := ladder[test.index]

			assert.Equal(t, test.expectedLadderStep.Odd, value.Odd)
			assert.InDelta(t, test.expectedLadderStep.GreenBookPL, value.GreenBookPL, amountEqualityThreshold)
			assert.InDelta(t, test.expectedLadderStep.VolMatched, value.VolMatched, amountEqualityThreshold)
		})
	}
}
