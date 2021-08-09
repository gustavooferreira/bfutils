package bfutils_test

import (
	"testing"

	"github.com/gustavooferreira/bfutils"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOddsSize(t *testing.T) {
	OddsLen := len(bfutils.Odds)
	OddsStrLen := len(bfutils.OddsStr)
	OddsDecimalLen := len(bfutils.Odds)

	t.Run("Odds array length equal to OddsCount", func(t *testing.T) {
		assert.Equal(t, bfutils.OddsCount, OddsLen)
	})

	t.Run("OddsStr array length equal to OddsCount", func(t *testing.T) {
		assert.Equal(t, bfutils.OddsCount, OddsStrLen)
	})

	t.Run("OddsDecimal array length equal to OddsCount", func(t *testing.T) {
		assert.Equal(t, bfutils.OddsCount, OddsDecimalLen)
	})
}

func TestOddsDecimal(t *testing.T) {
	tests := map[string]struct {
		index    int
		expected decimal.Decimal
	}{
		"odd[1.01] exists": {index: 0, expected: decimal.RequireFromString("1.01")},
		"odd[1.1] exists":  {index: 9, expected: decimal.RequireFromString("1.1")},
		"odd[2] exists":    {index: 99, expected: decimal.RequireFromString("2")},
		"odd[510] exists":  {index: 300, expected: decimal.RequireFromString("510")},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.True(t, test.expected.Equal(bfutils.Odds[test.index]))
		})
	}
}

func TestOddRange(t *testing.T) {
	tests := map[string]struct {
		index         int
		expectedBegin string
		expectedEnd   string
	}{
		"odd range [0] boundaries": {index: 0, expectedBegin: "1.01", expectedEnd: "2"},
		"odd range [3] boundaries": {index: 3, expectedBegin: "4.1", expectedEnd: "6"},
		"odd range [9] boundaries": {index: 9, expectedBegin: "110", expectedEnd: "1000"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expectedBegin, bfutils.OddsRange[test.index]["begin"])
			assert.Equal(t, test.expectedEnd, bfutils.OddsRange[test.index]["end"])
		})
	}
}

func TestOddFloor(t *testing.T) {
	tests := map[string]struct {
		odd           decimal.Decimal
		expectedIndex int
		expectedOdd   decimal.Decimal
		expectedErr   bool
	}{
		// Boundary tests
		"odd[-1] match":    {odd: decimal.RequireFromString("-1"), expectedErr: true},
		"odd[0] match":     {odd: decimal.RequireFromString("0"), expectedErr: true},
		"odd[1] match":     {odd: decimal.RequireFromString("1"), expectedErr: true},
		"odd[99999] match": {odd: decimal.RequireFromString("99999"), expectedErr: true},

		"odd[1.01] match":    {odd: decimal.RequireFromString("1.01"), expectedIndex: 0, expectedOdd: decimal.RequireFromString("1.01")},
		"odd[1000] match":    {odd: decimal.RequireFromString("1000"), expectedIndex: 349, expectedOdd: decimal.RequireFromString("1000")},
		"odd[1.0100] match":  {odd: decimal.RequireFromString("1.0100"), expectedIndex: 0, expectedOdd: decimal.RequireFromString("1.01")},
		"odd[1000.00] match": {odd: decimal.RequireFromString("1000.00"), expectedIndex: 349, expectedOdd: decimal.RequireFromString("1000")},

		"odd[3.2552321] match":      {odd: decimal.RequireFromString("3.2552321"), expectedIndex: 154, expectedOdd: decimal.RequireFromString("3.25")},
		"odd[3.299999999999] match": {odd: decimal.RequireFromString("3.299999999999"), expectedIndex: 154, expectedOdd: decimal.RequireFromString("3.25")},

		"odd[2.0] match":  {odd: decimal.RequireFromString("2.0"), expectedIndex: 99, expectedOdd: decimal.RequireFromString("2.0")},
		"odd[4.05] match": {odd: decimal.RequireFromString("4.05"), expectedIndex: 169, expectedOdd: decimal.RequireFromString("4.0")},
		"odd[33] match":   {odd: decimal.RequireFromString("33"), expectedIndex: 240, expectedOdd: decimal.RequireFromString("32")},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var errBool bool
			index, odd, err := bfutils.OddFloor(test.odd)
			if err != nil {
				errBool = true
			}

			require.Equal(t, test.expectedErr, errBool)
			assert.Equal(t, test.expectedIndex, index)
			assert.True(t, test.expectedOdd.Equal(odd))
		})
	}
}

func TestOddCeil(t *testing.T) {
	tests := map[string]struct {
		odd           decimal.Decimal
		expectedIndex int
		expectedOdd   decimal.Decimal
		expectedErr   bool
	}{
		// Boundary tests
		"odd[-1] match":    {odd: decimal.RequireFromString("-1"), expectedErr: true},
		"odd[0] match":     {odd: decimal.RequireFromString("0"), expectedErr: true},
		"odd[1] match":     {odd: decimal.RequireFromString("1"), expectedErr: true},
		"odd[99999] match": {odd: decimal.RequireFromString("99999"), expectedErr: true},

		"odd[1.01] match":    {odd: decimal.RequireFromString("1.01"), expectedIndex: 0, expectedOdd: decimal.RequireFromString("1.01")},
		"odd[1000] match":    {odd: decimal.RequireFromString("1000"), expectedIndex: 349, expectedOdd: decimal.RequireFromString("1000")},
		"odd[1.0100] match":  {odd: decimal.RequireFromString("1.0100"), expectedIndex: 0, expectedOdd: decimal.RequireFromString("1.01")},
		"odd[1000.00] match": {odd: decimal.RequireFromString("1000.00"), expectedIndex: 349, expectedOdd: decimal.RequireFromString("1000")},

		"odd[3.2552321] match":      {odd: decimal.RequireFromString("3.2552321"), expectedIndex: 155, expectedOdd: decimal.RequireFromString("3.3")},
		"odd[3.249999999999] match": {odd: decimal.RequireFromString("3.249999999999"), expectedIndex: 154, expectedOdd: decimal.RequireFromString("3.25")},

		"odd[2.0] match":  {odd: decimal.RequireFromString("2.0"), expectedIndex: 99, expectedOdd: decimal.RequireFromString("2.0")},
		"odd[4.05] match": {odd: decimal.RequireFromString("4.05"), expectedIndex: 170, expectedOdd: decimal.RequireFromString("4.1")},
		"odd[33] match":   {odd: decimal.RequireFromString("33"), expectedIndex: 241, expectedOdd: decimal.RequireFromString("34")},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var errBool bool
			index, odd, err := bfutils.OddCeil(test.odd)
			if err != nil {
				errBool = true
			}

			require.Equal(t, test.expectedErr, errBool)
			assert.Equal(t, test.expectedIndex, index)
			assert.True(t, test.expectedOdd.Equal(odd))
		})
	}
}

func TestOddRound(t *testing.T) {
	tests := map[string]struct {
		odd           decimal.Decimal
		expectedIndex int
		expectedOdd   decimal.Decimal
		expectedErr   bool
	}{
		// Boundary tests
		"odd[-1] match":    {odd: decimal.RequireFromString("-1"), expectedErr: true},
		"odd[0] match":     {odd: decimal.RequireFromString("0"), expectedErr: true},
		"odd[1] match":     {odd: decimal.RequireFromString("1"), expectedErr: true},
		"odd[99999] match": {odd: decimal.RequireFromString("99999"), expectedErr: true},

		"odd[1.01] match":    {odd: decimal.RequireFromString("1.01"), expectedIndex: 0, expectedOdd: decimal.RequireFromString("1.01")},
		"odd[1000] match":    {odd: decimal.RequireFromString("1000"), expectedIndex: 349, expectedOdd: decimal.RequireFromString("1000")},
		"odd[1.0100] match":  {odd: decimal.RequireFromString("1.0100"), expectedIndex: 0, expectedOdd: decimal.RequireFromString("1.01")},
		"odd[1000.00] match": {odd: decimal.RequireFromString("1000.00"), expectedIndex: 349, expectedOdd: decimal.RequireFromString("1000")},

		"odd[3.2552321] match":      {odd: decimal.RequireFromString("3.2552321"), expectedIndex: 154, expectedOdd: decimal.RequireFromString("3.25")},
		"odd[3.249999999999] match": {odd: decimal.RequireFromString("3.249999999999"), expectedIndex: 154, expectedOdd: decimal.RequireFromString("3.25")},

		"odd[21.05] match": {odd: decimal.RequireFromString("21.05"), expectedIndex: 230, expectedOdd: decimal.RequireFromString("21")},
		"odd[2.0] match":   {odd: decimal.RequireFromString("2.0"), expectedIndex: 99, expectedOdd: decimal.RequireFromString("2.0")},
		"odd[4.077] match": {odd: decimal.RequireFromString("4.077"), expectedIndex: 170, expectedOdd: decimal.RequireFromString("4.1")},
		"odd[32.9] match":  {odd: decimal.RequireFromString("32.9"), expectedIndex: 240, expectedOdd: decimal.RequireFromString("32")},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var errBool bool
			index, odd, err := bfutils.OddRound(test.odd)
			if err != nil {
				errBool = true
			}

			require.Equal(t, test.expectedErr, errBool)
			assert.Equal(t, test.expectedIndex, index)
			assert.True(t, test.expectedOdd.Equal(odd))
		})
	}
}

func TestOddShift(t *testing.T) {
	tests := map[string]struct {
		roundType     bfutils.RoundType
		odd           decimal.Decimal
		shift         int
		expectedIndex int
		expectedOdd   decimal.Decimal
		expectedErr   bool
	}{
		// Boundary tests
		"odd[-1] shift":       {odd: decimal.RequireFromString("-1"), expectedErr: true},
		"odd[0] shift":        {odd: decimal.RequireFromString("0"), expectedErr: true},
		"odd[1] shift":        {odd: decimal.RequireFromString("1"), expectedErr: true},
		"odd[99999] shift":    {odd: decimal.RequireFromString("99999"), expectedErr: true},
		"odd[10, 1000] shift": {roundType: bfutils.RoundType_Ceil, odd: decimal.RequireFromString("10"), shift: 1000, expectedErr: true},

		"odd[1.01, 3] shift": {roundType: bfutils.RoundType_Round, odd: decimal.RequireFromString("1.01"), shift: 3, expectedIndex: 3, expectedOdd: decimal.RequireFromString("1.04")},
		"odd[4, -10] shift":  {roundType: bfutils.RoundType_Ceil, odd: decimal.RequireFromString("4"), shift: -10, expectedIndex: 159, expectedOdd: decimal.RequireFromString("3.5")},
		"odd[10, 5] shift":   {roundType: bfutils.RoundType_Floor, odd: decimal.RequireFromString("10"), shift: 5, expectedIndex: 214, expectedOdd: decimal.RequireFromString("12.5")},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var errBool bool
			index, odd, err := bfutils.OddShift(test.roundType, test.odd, test.shift)
			if err != nil {
				errBool = true
			}

			require.Equal(t, test.expectedErr, errBool)
			assert.Equal(t, test.expectedIndex, index)
			assert.True(t, test.expectedOdd.Equal(odd))
		})
	}
}

func TestOddsTicksDiff(t *testing.T) {
	tests := map[string]struct {
		roundType    bfutils.RoundType
		odd1         decimal.Decimal
		odd2         decimal.Decimal
		expectedDiff int
		expectedErr  bool
	}{
		// Boundary tests
		"odds[-1, 10] tick diff":    {odd1: decimal.RequireFromString("-1"), odd2: decimal.RequireFromString("10"), expectedErr: true},
		"odds[0, 5] tick diff":      {odd1: decimal.RequireFromString("0"), odd2: decimal.RequireFromString("5"), expectedErr: true},
		"odds[50, 1] tick diff":     {odd1: decimal.RequireFromString("50"), odd2: decimal.RequireFromString("1"), expectedErr: true},
		"odds[10, 99999] tick diff": {odd1: decimal.RequireFromString("10"), odd2: decimal.RequireFromString("99999"), expectedErr: true},

		"odds[1.01, 3] tick diff": {roundType: bfutils.RoundType_Round, odd1: decimal.RequireFromString("1.01"), odd2: decimal.RequireFromString("3"), expectedDiff: 149},
		"odds[4, 4.5] tick diff":  {roundType: bfutils.RoundType_Ceil, odd1: decimal.RequireFromString("4"), odd2: decimal.RequireFromString("4.5"), expectedDiff: 5},
		"odds[10, 5] tick diff":   {roundType: bfutils.RoundType_Floor, odd1: decimal.RequireFromString("10"), odd2: decimal.RequireFromString("5"), expectedDiff: 30},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var errBool bool
			diff, err := bfutils.OddsTicksDiff(test.roundType, test.odd1, test.odd2)
			if err != nil {
				errBool = true
			}

			require.Equal(t, test.expectedErr, errBool)
			assert.Equal(t, test.expectedDiff, diff)
		})
	}
}

func TestOddWithinBoundaries(t *testing.T) {
	tests := map[string]struct {
		odd      decimal.Decimal
		expected bool
	}{
		// Boundary tests
		"odds[-1] boundary check":   {odd: decimal.RequireFromString("-1"), expected: false},
		"odds[0] boundary check":    {odd: decimal.RequireFromString("0"), expected: false},
		"odds[1] boundary check":    {odd: decimal.RequireFromString("1"), expected: false},
		"odds[1001] boundary check": {odd: decimal.RequireFromString("1001"), expected: false},

		"odds[2.3] boundary check": {odd: decimal.RequireFromString("2.3"), expected: true},
		"odds[50] boundary check":  {odd: decimal.RequireFromString("50"), expected: true},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			value := bfutils.IsOddWithinBoundaries(test.odd)
			assert.Equal(t, test.expected, value)
		})
	}
}

func TestRoundTypeEnum(t *testing.T) {
	var enum bfutils.RoundType = bfutils.RoundType_Floor
	assert.Equal(t, "Floor", enum.String())
}
