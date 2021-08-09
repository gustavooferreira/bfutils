package betting_test

import (
	"fmt"
	"testing"

	"github.com/gustavooferreira/bfutils/betting"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCalcFreeBetDecimal(t *testing.T) {
	tests := map[string]struct {
		oddBack  decimal.Decimal
		oddLay   decimal.Decimal
		expected decimal.Decimal
	}{
		"calculate free bet [B:0,L:0]": {
			oddBack:  decimal.RequireFromString("0"),
			oddLay:   decimal.RequireFromString("0"),
			expected: decimal.RequireFromString("0.0")},
		"calculate free bet [B:2,L:2]": {
			oddBack:  decimal.RequireFromString("2"),
			oddLay:   decimal.RequireFromString("2"),
			expected: decimal.RequireFromString("0.0")},
		"calculate free bet [B:20,L:10]": {
			oddBack:  decimal.RequireFromString("20"),
			oddLay:   decimal.RequireFromString("10"),
			expected: decimal.RequireFromString("10")},
		"calculate free bet [B:18,L:7]": {
			oddBack:  decimal.RequireFromString("18"),
			oddLay:   decimal.RequireFromString("7"),
			expected: decimal.RequireFromString("11")},
		"calculate free bet [B:100,L:250]": {
			oddBack:  decimal.RequireFromString("100"),
			oddLay:   decimal.RequireFromString("250"),
			expected: decimal.RequireFromString("-150")},
		"calculate free bet [B:50,L:5]": {
			oddBack:  decimal.RequireFromString("50"),
			oddLay:   decimal.RequireFromString("5"),
			expected: decimal.RequireFromString("45")},
		"calculate free bet [B:1.35,L:2.02]": {
			oddBack:  decimal.RequireFromString("1.35"),
			oddLay:   decimal.RequireFromString("2.02"),
			expected: decimal.RequireFromString("-0.67")},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			value := betting.FreeBetDecimal(test.oddBack, test.oddLay)
			assert.True(t, test.expected.Equal(value))
		})
	}
}

func TestCalcFreeBetPL(t *testing.T) {
	tests := map[string]struct {
		oddBack  decimal.Decimal
		oddLay   decimal.Decimal
		stake    decimal.Decimal
		expected decimal.Decimal
	}{
		"calculate free bet [B:0,L:0;S:0]": {
			oddBack:  decimal.RequireFromString("0"),
			oddLay:   decimal.RequireFromString("0"),
			stake:    decimal.RequireFromString("0"),
			expected: decimal.RequireFromString("0.0")},
		"calculate free bet [B:2,L:2;S:10]": {
			oddBack:  decimal.RequireFromString("2"),
			oddLay:   decimal.RequireFromString("2"),
			stake:    decimal.RequireFromString("10"),
			expected: decimal.RequireFromString("0.0")},
		"calculate free bet [B:20,L:10;S:4]": {
			oddBack:  decimal.RequireFromString("20"),
			oddLay:   decimal.RequireFromString("10"),
			stake:    decimal.RequireFromString("4"),
			expected: decimal.RequireFromString("40")},
		"calculate free bet [B:18,L:7;S:0]": {
			oddBack:  decimal.RequireFromString("18"),
			oddLay:   decimal.RequireFromString("7"),
			stake:    decimal.RequireFromString("0"),
			expected: decimal.RequireFromString("0")},
		"calculate free bet [B:100,L:250;S:2]": {
			oddBack:  decimal.RequireFromString("100"),
			oddLay:   decimal.RequireFromString("250"),
			stake:    decimal.RequireFromString("2"),
			expected: decimal.RequireFromString("-300")},
		"calculate free bet [B:50,L:5;S:5]": {
			oddBack:  decimal.RequireFromString("50"),
			oddLay:   decimal.RequireFromString("5"),
			stake:    decimal.RequireFromString("5"),
			expected: decimal.RequireFromString("225")},
		"calculate free bet [B:1.35,L:2.02;S:2]": {
			oddBack:  decimal.RequireFromString("1.35"),
			oddLay:   decimal.RequireFromString("2.02"),
			stake:    decimal.RequireFromString("2"),
			expected: decimal.RequireFromString("-1.34")},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			value := betting.FreeBetPL(test.oddBack, test.oddLay, test.stake)
			assert.True(t, test.expected.Equal(value))
		})
	}
}

func TestCalcGreenBookOpenBackDecimal(t *testing.T) {
	tests := map[string]struct {
		oddBack      decimal.Decimal
		oddLay       decimal.Decimal
		expectedPerc decimal.Decimal
		expectedErr  bool
	}{
		"calculate green book [B:0,L:0]": {
			oddBack:     decimal.RequireFromString("0"),
			oddLay:      decimal.RequireFromString("0"),
			expectedErr: true},

		"calculate green book [B:2,L:2]": {
			oddBack:      decimal.RequireFromString("2"),
			oddLay:       decimal.RequireFromString("2"),
			expectedPerc: decimal.RequireFromString("0.0")},
		"calculate green book [B:20,L:10]": {
			oddBack:      decimal.RequireFromString("20"),
			oddLay:       decimal.RequireFromString("10"),
			expectedPerc: decimal.RequireFromString("1")},
		"calculate green book [B:18,L:7]": {
			oddBack:      decimal.RequireFromString("18"),
			oddLay:       decimal.RequireFromString("7"),
			expectedPerc: decimal.RequireFromString("1.57")},
		"calculate green book [B:100,L:250]": {
			oddBack:      decimal.RequireFromString("100"),
			oddLay:       decimal.RequireFromString("250"),
			expectedPerc: decimal.RequireFromString("-0.6")},
		"calculate green book [B:50,L:5]": {
			oddBack:      decimal.RequireFromString("50"),
			oddLay:       decimal.RequireFromString("5"),
			expectedPerc: decimal.RequireFromString("9")},
		"calculate green book [B:1.35,L:2.02]": {
			oddBack:      decimal.RequireFromString("1.35"),
			oddLay:       decimal.RequireFromString("2.02"),
			expectedPerc: decimal.RequireFromString("-0.33")},
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

			assert.True(t, test.expectedPerc.Equal(value), "expected: %s, got: %s",
				test.expectedPerc.StringFixed(3),
				value.StringFixed(3))
		})
	}
}

func TestCalcGreenBookOpenBackAmount(t *testing.T) {
	tests := map[string]struct {
		oddBack          decimal.Decimal
		oddLay           decimal.Decimal
		stake            decimal.Decimal
		expectedStakeLay decimal.Decimal
		expectedErr      bool
	}{
		"calculate green book amount [B:0,L:0;S:0]": {
			oddBack:     decimal.RequireFromString("0"),
			oddLay:      decimal.RequireFromString("0"),
			stake:       decimal.RequireFromString("0"),
			expectedErr: true},
		"calculate green book amount [B:2,L:2;S:10]": {
			oddBack:          decimal.RequireFromString("2"),
			oddLay:           decimal.RequireFromString("2"),
			stake:            decimal.RequireFromString("10"),
			expectedStakeLay: decimal.RequireFromString("10")},
		"calculate green book amount [B:20,L:10;S:4]": {
			oddBack:          decimal.RequireFromString("20"),
			oddLay:           decimal.RequireFromString("10"),
			stake:            decimal.RequireFromString("4"),
			expectedStakeLay: decimal.RequireFromString("8")},
		"calculate green book amount [B:18,L:7;S:0]": {
			oddBack:          decimal.RequireFromString("18"),
			oddLay:           decimal.RequireFromString("7"),
			stake:            decimal.RequireFromString("0"),
			expectedStakeLay: decimal.RequireFromString("0")},
		"calculate green book amount [B:100,L:250;S:2]": {
			oddBack:          decimal.RequireFromString("100"),
			oddLay:           decimal.RequireFromString("250"),
			stake:            decimal.RequireFromString("2"),
			expectedStakeLay: decimal.RequireFromString("0.8")},
		"calculate green book amount [B:50,L:5;S:5]": {
			oddBack:          decimal.RequireFromString("50"),
			oddLay:           decimal.RequireFromString("5"),
			stake:            decimal.RequireFromString("5"),
			expectedStakeLay: decimal.RequireFromString("50")},
		"calculate green book amount [B:1.35,L:2.02;S:2]": {
			oddBack:          decimal.RequireFromString("1.35"),
			oddLay:           decimal.RequireFromString("2.02"),
			stake:            decimal.RequireFromString("2"),
			expectedStakeLay: decimal.RequireFromString("1.34")},
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

			assert.True(t, test.expectedStakeLay.Equal(value), "expected: %s, got: %s",
				test.expectedStakeLay.StringFixed(3),
				value.StringFixed(3))
		})
	}
}

func TestCalcGreenBookOpenBackAmountByPerc(t *testing.T) {
	tests := map[string]struct {
		oddBack        decimal.Decimal
		perc           decimal.Decimal
		expectedOddLay decimal.Decimal
		expectedErr    bool
	}{
		"calculate green book perc 1": {
			oddBack:     decimal.RequireFromString("5"),
			perc:        decimal.RequireFromString("-2"),
			expectedErr: true},

		"calculate green book perc 2": {
			oddBack:        decimal.RequireFromString("20"),
			perc:           decimal.RequireFromString("1"),
			expectedOddLay: decimal.RequireFromString("10")},
		"calculate green book perc 3": {
			oddBack:        decimal.RequireFromString("18"),
			perc:           decimal.RequireFromString("0.7"),
			expectedOddLay: decimal.RequireFromString("10.59")},
		"calculate green book perc 4": {
			oddBack:        decimal.RequireFromString("100"),
			perc:           decimal.RequireFromString("1.5"),
			expectedOddLay: decimal.RequireFromString("40")},
		"calculate green book perc 5": {
			oddBack:        decimal.RequireFromString("50"),
			perc:           decimal.RequireFromString("2"),
			expectedOddLay: decimal.RequireFromString("16.67")},
		"calculate green book perc 6": {
			oddBack:        decimal.RequireFromString("10"),
			perc:           decimal.RequireFromString("-0.5"),
			expectedOddLay: decimal.RequireFromString("20")},
		"calculate green book perc 7": {
			oddBack:        decimal.RequireFromString("20"),
			perc:           decimal.RequireFromString("-0.8"),
			expectedOddLay: decimal.RequireFromString("100")},
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

			assert.True(t, test.expectedOddLay.Equal(value), "expected: %s, got: %s",
				test.expectedOddLay.StringFixed(3),
				value.StringFixed(3))
		})
	}
}

func TestCalcGreenBookOpenLayDecimal(t *testing.T) {
	tests := map[string]struct {
		oddLay       decimal.Decimal
		oddBack      decimal.Decimal
		expectedPerc decimal.Decimal
		expectedErr  bool
	}{
		"calculate green book 1": {
			oddLay:      decimal.RequireFromString("0"),
			oddBack:     decimal.RequireFromString("0"),
			expectedErr: true},
		"calculate green book 2": {
			oddLay:       decimal.RequireFromString("20"),
			oddBack:      decimal.RequireFromString("10"),
			expectedPerc: decimal.RequireFromString("-1")},
		"calculate green book 3": {
			oddLay:       decimal.RequireFromString("18"),
			oddBack:      decimal.RequireFromString("7"),
			expectedPerc: decimal.RequireFromString("-1.57")},
		"calculate green book 4": {
			oddLay:       decimal.RequireFromString("100"),
			oddBack:      decimal.RequireFromString("250"),
			expectedPerc: decimal.RequireFromString("0.6")},
		"calculate green book 5": {
			oddLay:       decimal.RequireFromString("50"),
			oddBack:      decimal.RequireFromString("5"),
			expectedPerc: decimal.RequireFromString("-9")},
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

			assert.True(t, test.expectedPerc.Equal(value), "expected: %s, got: %s",
				test.expectedPerc.StringFixed(3),
				value.StringFixed(3))
		})
	}
}

func TestCalcGreenBookOpenLayAmount(t *testing.T) {
	tests := map[string]struct {
		oddBack          decimal.Decimal
		oddLay           decimal.Decimal
		stake            decimal.Decimal
		expectedStakeLay decimal.Decimal
		expectedErr      bool
	}{
		"calculate green book amount 1": {
			oddBack:     decimal.RequireFromString("0"),
			oddLay:      decimal.RequireFromString("0"),
			stake:       decimal.RequireFromString("0"),
			expectedErr: true},
		"calculate green book amount 2": {
			oddBack:          decimal.RequireFromString("2"),
			oddLay:           decimal.RequireFromString("2"),
			stake:            decimal.RequireFromString("10"),
			expectedStakeLay: decimal.RequireFromString("10")},
		"calculate green book amount 3": {
			oddBack:          decimal.RequireFromString("20"),
			oddLay:           decimal.RequireFromString("10"),
			stake:            decimal.RequireFromString("4"),
			expectedStakeLay: decimal.RequireFromString("2")},
		"calculate green book amount 4": {
			oddBack:          decimal.RequireFromString("18"),
			oddLay:           decimal.RequireFromString("7"),
			stake:            decimal.RequireFromString("0"),
			expectedStakeLay: decimal.RequireFromString("0")},
		"calculate green book amount 5": {
			oddBack:          decimal.RequireFromString("100"),
			oddLay:           decimal.RequireFromString("250"),
			stake:            decimal.RequireFromString("2"),
			expectedStakeLay: decimal.RequireFromString("5")},
		"calculate green book amount 6": {
			oddBack:          decimal.RequireFromString("50"),
			oddLay:           decimal.RequireFromString("5"),
			stake:            decimal.RequireFromString("5"),
			expectedStakeLay: decimal.RequireFromString("0.5")},
		"calculate green book amount 7": {
			oddBack:          decimal.RequireFromString("1.35"),
			oddLay:           decimal.RequireFromString("2.02"),
			stake:            decimal.RequireFromString("2"),
			expectedStakeLay: decimal.RequireFromString("2.99")},
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

			assert.True(t, test.expectedStakeLay.Equal(value), "expected: %s, got: %s",
				test.expectedStakeLay.StringFixed(3),
				value.StringFixed(3))
		})
	}
}

func TestCalcGreenBookOpenLayAmountByPerc(t *testing.T) {
	tests := map[string]struct {
		oddLay           decimal.Decimal
		perc             decimal.Decimal
		expectedStakeLay decimal.Decimal
		expectedErr      bool
	}{
		"calculate green book perc 1": {
			oddLay:      decimal.RequireFromString("5"),
			perc:        decimal.RequireFromString("2"),
			expectedErr: true},
		"calculate green book perc 2": {
			oddLay:           decimal.RequireFromString("20"),
			perc:             decimal.RequireFromString("-1"),
			expectedStakeLay: decimal.RequireFromString("10")},
		"calculate green book perc 3": {
			oddLay:           decimal.RequireFromString("18"),
			perc:             decimal.RequireFromString("0.75"),
			expectedStakeLay: decimal.RequireFromString("72")},
		"calculate green book perc 4": {
			oddLay:           decimal.RequireFromString("100"),
			perc:             decimal.RequireFromString("-1.5"),
			expectedStakeLay: decimal.RequireFromString("40")},
		"calculate green book perc 5": {
			oddLay:           decimal.RequireFromString("50"),
			perc:             decimal.RequireFromString("-2"),
			expectedStakeLay: decimal.RequireFromString("16.67")},
		"calculate green book perc 6": {
			oddLay:           decimal.RequireFromString("10"),
			perc:             decimal.RequireFromString("0.5"),
			expectedStakeLay: decimal.RequireFromString("20")},
		"calculate green book perc 7": {
			oddLay:           decimal.RequireFromString("20"),
			perc:             decimal.RequireFromString("0.8"),
			expectedStakeLay: decimal.RequireFromString("100")},
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

			assert.True(t, test.expectedStakeLay.Equal(value), "expected: %s, got: %s", test.expectedStakeLay.StringFixed(3), value.StringFixed(3))
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
			bets: []betting.Bet{
				{
					Type: 0, Odd: decimal.RequireFromString("2"),
					Amount: decimal.RequireFromString("2"),
				},
			},
			expectedErr: true,
		},
		"calculate greenbook 2": {
			bets: []betting.Bet{
				{
					Type:   betting.BetType_Back,
					Odd:    decimal.RequireFromString("2.111"),
					Amount: decimal.RequireFromString("2"),
				},
			},
			expectedErr: true,
		},
		"calculate greenbook 3": {
			bets: []betting.Bet{
				{
					Type:   betting.BetType_Back,
					Odd:    decimal.RequireFromString("2"),
					Amount: decimal.RequireFromString("0"),
				},
				{
					Type:   betting.BetType_Lay,
					Odd:    decimal.RequireFromString("0.5"),
					Amount: decimal.RequireFromString("2"),
				},
			},
			expectedErr: true,
		},
		"calculate greenbook 4": {
			bets: []betting.Bet{
				{
					Type:   betting.BetType_Back,
					Odd:    decimal.RequireFromString("2.0"),
					Amount: decimal.RequireFromString("10"),
				},
				{
					Type:   betting.BetType_Lay,
					Odd:    decimal.RequireFromString("1.5"),
					Amount: decimal.RequireFromString("5"),
				},
			},
			expectedEdged: false,
		},
		"calculate greenbook 5": {
			bets: []betting.Bet{
				{
					Type:   betting.BetType_Back,
					Odd:    decimal.RequireFromString("4.0"),
					Amount: decimal.RequireFromString("10")},
				{
					Type:   betting.BetType_Lay,
					Odd:    decimal.RequireFromString("3.4"),
					Amount: decimal.RequireFromString("10"),
				},
				{
					Type:   betting.BetType_Lay,
					Odd:    decimal.RequireFromString("2.8"),
					Amount: decimal.RequireFromString("2"),
				},
			},
			expectedEdged: false,
		},
		"calculate greenbook 6": {
			bets: []betting.Bet{
				{
					Type:   betting.BetType_Back,
					Odd:    decimal.RequireFromString("16.5"),
					Amount: decimal.RequireFromString("50"),
				},
				{
					Type:   betting.BetType_Lay,
					Odd:    decimal.RequireFromString("16"),
					Amount: decimal.RequireFromString("51.56"),
				},
			},
			expectedEdged: false,
		},
		"calculate greenbook 7": {
			bets: []betting.Bet{
				{
					Type:   betting.BetType_Back,
					Odd:    decimal.RequireFromString("8"),
					Amount: decimal.RequireFromString("10"),
				},
				{
					Type:   betting.BetType_Lay,
					Odd:    decimal.RequireFromString("4"),
					Amount: decimal.RequireFromString("20"),
				},
			},
			expectedEdged: true,
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
		expectedWinPL  decimal.Decimal
		expectedLosePL decimal.Decimal
		expectedErr    bool
	}{
		"calculate greenbook 1": {expectedErr: true},
		"calculate greenbook 2": {
			selection: betting.Selection{
				Bets: []betting.Bet{},
			},
			expectedErr: true},
		"calculate greenbook 3": {
			selection: betting.Selection{
				Bets: []betting.Bet{
					{
						Type:   betting.BetType_Back,
						Odd:    decimal.RequireFromString("1.5"),
						Amount: decimal.RequireFromString("10"),
					},
				},
			},
			expectedErr: true},
		"calculate greenbook 4": {
			selection: betting.Selection{
				Bets: []betting.Bet{
					{
						Type:   betting.BetType_Back,
						Odd:    decimal.RequireFromString("1.5"),
						Amount: decimal.RequireFromString("10"),
					},
				},
				CurrentBackOdd: decimal.RequireFromString("1.525"),
			},
			expectedErr: true},
		"calculate greenbook 5": {
			selection: betting.Selection{
				Bets: []betting.Bet{
					{
						Type:   betting.BetType_Back,
						Odd:    decimal.RequireFromString("1.5"),
						Amount: decimal.RequireFromString("10"),
					},
				},
				CurrentBackOdd: decimal.RequireFromString("1.5"),
			},
			expectedErr: true},
		"calculate greenbook 6": {
			selection: betting.Selection{
				Bets: []betting.Bet{
					{
						Type:   betting.BetType_Back,
						Odd:    decimal.RequireFromString("1.5"),
						Amount: decimal.RequireFromString("10"),
					},
				},
				CurrentBackOdd: decimal.RequireFromString("1.5"),
				CurrentLayOdd:  decimal.RequireFromString("1.525"),
			},
			expectedErr: true},
		"calculate greenbook 7": {
			selection: betting.Selection{
				Bets: []betting.Bet{
					{
						Type:   betting.BetType_Back,
						Odd:    decimal.RequireFromString("4"),
						Amount: decimal.RequireFromString("0"),
					},
					{
						Type:   betting.BetType_Lay,
						Odd:    decimal.RequireFromString("1"),
						Amount: decimal.RequireFromString("5"),
					},
				},
				CurrentBackOdd: decimal.RequireFromString("2"),
				CurrentLayOdd:  decimal.RequireFromString("2.1"),
			},
			expectedErr: true,
		},
		"calculate greenbook 8": {
			selection: betting.Selection{
				Bets: []betting.Bet{
					{
						Type:   betting.BetType_Lay,
						Odd:    decimal.RequireFromString("1.525"),
						Amount: decimal.RequireFromString("5"),
					},
				},
				CurrentBackOdd: decimal.RequireFromString("2"),
				CurrentLayOdd:  decimal.RequireFromString("2.1"),
			},
			expectedErr: true,
		},
		"calculate greenbook 9": {
			selection: betting.Selection{
				Bets: []betting.Bet{
					{
						Type:   betting.BetType_Back,
						Odd:    decimal.RequireFromString("8"),
						Amount: decimal.RequireFromString("10"),
					},
					{
						Type:   betting.BetType_Lay,
						Odd:    decimal.RequireFromString("4"),
						Amount: decimal.RequireFromString("20"),
					},
				},
				CurrentBackOdd: decimal.RequireFromString("2"),
				CurrentLayOdd:  decimal.RequireFromString("2.1"),
			},
			expectedErr: true,
		},
		"calculate greenbook 10": {
			selection: betting.Selection{
				Bets: []betting.Bet{
					{
						Type:   betting.BetType_Back,
						Odd:    decimal.RequireFromString("4"),
						Amount: decimal.RequireFromString("10"),
					},
					{
						Type:   betting.BetType_Lay,
						Odd:    decimal.RequireFromString("2"),
						Amount: decimal.RequireFromString("5"),
					},
				},
				CurrentBackOdd: decimal.RequireFromString("2"),
				CurrentLayOdd:  decimal.RequireFromString("2.1"),
			},
			expectedBet: betting.Bet{
				Type:   betting.BetType_Lay,
				Odd:    decimal.RequireFromString("2.1"),
				Amount: decimal.RequireFromString("14.29"),
			},
			expectedWinPL:  decimal.RequireFromString("9.29"),
			expectedLosePL: decimal.RequireFromString("9.29"),
		},
		// "calculate greenbook 11": {
		// 	selection: betting.Selection{
		// 		Bets: []betting.Bet{
		// 			{Type: betting.BetType_Back, Odd: 3, Amount: 5},
		// 		},
		// 		CurrentBackOdd: 2,
		// 		CurrentLayOdd:  2.1,
		// 	},
		// 	expectedBet:    betting.Bet{Type: betting.BetType_Lay, Odd: 2.1, Amount: 7.14},
		// 	expectedWinPL:  2.15,
		// 	expectedLosePL: 2.15,
		// },
		// "calculate greenbook 12": {
		// 	selection: betting.Selection{
		// 		Bets: []betting.Bet{
		// 			{Type: betting.BetType_Lay, Odd: 3, Amount: 5},
		// 		},
		// 		CurrentBackOdd: 4,
		// 		CurrentLayOdd:  4.2,
		// 	},
		// 	expectedBet:    betting.Bet{Type: betting.BetType_Back, Odd: 4, Amount: 3.75},
		// 	expectedWinPL:  1.25,
		// 	expectedLosePL: 1.25,
		// },
		// "calculate greenbook 13": {
		// 	selection: betting.Selection{
		// 		Bets: []betting.Bet{
		// 			{Type: betting.BetType_Back, Odd: 4, Amount: 5},
		// 			{Type: betting.BetType_Lay, Odd: 3, Amount: 5},
		// 		},
		// 		CurrentBackOdd: 3.1,
		// 		CurrentLayOdd:  3,
		// 	},
		// 	expectedBet:    betting.Bet{Type: betting.BetType_Lay, Odd: 3, Amount: 1.67},
		// 	expectedWinPL:  1.67,
		// 	expectedLosePL: 1.67,
		// },
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

			if !errBool {
				assert.Equal(t, test.expectedBet.Type, bet.Type, "bet type field")
				assert.Equal(t, test.expectedBet.Odd, bet.Odd, "bet odd field")
				assert.True(t, test.expectedBet.Amount.Equal(bet.Amount), "Bet Amount - expected: %s, got: %s",
					test.expectedBet.Amount.StringFixed(3),
					bet.Amount.StringFixed(3))

				assert.True(t, test.expectedWinPL.Equal(bet.WinPL), "WinP&L - expected: %s, got: %s",
					test.expectedWinPL.StringFixed(3),
					bet.WinPL.StringFixed(3))
				assert.True(t, test.expectedLosePL.Equal(bet.LosePL), "LoseP&L - expected: %s, got: %s",
					test.expectedLosePL.StringFixed(3),
					bet.LosePL.StringFixed(3))
			}
		})
	}
}

// func TestGreenBookAtAllOdds(t *testing.T) {
// 	tests := map[string]struct {
// 		bets               []betting.Bet
// 		index              int
// 		expectedLadderStep betting.LadderStep
// 		expectedErr        bool
// 	}{
// 		"calculate greenbook at all odds 1": {
// 			bets: []betting.Bet{
// 				{Type: betting.BetType_Back, Odd: 4, Amount: 0},
// 				{Type: betting.BetType_Lay, Odd: 1, Amount: 5},
// 			},
// 			expectedErr: true,
// 		},
// 		"calculate greenbook at all odds 2": {
// 			bets: []betting.Bet{
// 				{Type: betting.BetType_Lay, Odd: 1.525, Amount: 5},
// 			},
// 			expectedErr: true,
// 		},
// 		"calculate greenbook at all odds 3": {
// 			bets: []betting.Bet{
// 				{Type: betting.BetType_Back, Odd: 4.0, Amount: 10},
// 				{Type: betting.BetType_Lay, Odd: 3.4, Amount: 10},
// 				{Type: betting.BetType_Lay, Odd: 3, Amount: 2},
// 			},
// 			index: 0,
// 			expectedLadderStep: betting.LadderStep{
// 				Odd:         1.01,
// 				GreenBookPL: 0,
// 				VolMatched:  0,
// 			},
// 		},
// 		"calculate greenbook at all odds 4": {
// 			bets: []betting.Bet{
// 				{Type: betting.BetType_Back, Odd: 4, Amount: 5},
// 				{Type: betting.BetType_Lay, Odd: 3, Amount: 5},
// 			},
// 			index: 159,
// 			expectedLadderStep: betting.LadderStep{
// 				Odd:         3.5,
// 				GreenBookPL: 1.43,
// 				VolMatched:  1.43,
// 			},
// 		},
// 		"calculate greenbook at all odds 5": {
// 			bets: []betting.Bet{
// 				{Type: betting.BetType_Back, Odd: 2, Amount: 10},
// 				{Type: betting.BetType_Lay, Odd: 1.5, Amount: 5},
// 			},
// 			index: 49,
// 			expectedLadderStep: betting.LadderStep{
// 				Odd:         1.5,
// 				GreenBookPL: 3.34,
// 				VolMatched:  13.33,
// 			},
// 		},
// 		"calculate greenbook at all odds 6": {
// 			bets: []betting.Bet{
// 				{Type: betting.BetType_Lay, Odd: 1.5, Amount: 5},
// 			},
// 			index: 99,
// 			expectedLadderStep: betting.LadderStep{
// 				Odd:         2,
// 				GreenBookPL: 1.25,
// 				VolMatched:  3.75,
// 			},
// 		},
// 	}

// 	for name, test := range tests {
// 		t.Run(name, func(t *testing.T) {
// 			var errBool bool
// 			var errMsg string
// 			ladder, err := betting.GreenBookAtAllOdds(test.bets)
// 			if err != nil {
// 				errBool = true
// 				errMsg = fmt.Sprintf(" - err: %s", err.Error())
// 				require.Equal(t, test.expectedErr, errBool, "error field"+errMsg)
// 				return
// 			}

// 			value := ladder[test.index]

// 			assert.Equal(t, test.expectedLadderStep.Odd, value.Odd)
// 			assert.InDelta(t, test.expectedLadderStep.GreenBookPL, value.GreenBookPL, amountEqualityThreshold)
// 			assert.InDelta(t, test.expectedLadderStep.VolMatched, value.VolMatched, amountEqualityThreshold)
// 		})
// 	}
// }
