package bfutils_test

import (
	"testing"

	"github.com/gustavooferreira/bfutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOddsSize(t *testing.T) {
	OddsLen := len(bfutils.Odds)
	OddsStrLen := len(bfutils.OddsStr)

	t.Run("Odds array length equal to OddsCount", func(t *testing.T) {
		assert.Equal(t, bfutils.OddsCount, OddsLen)
	})

	t.Run("OddsStr array length equal to OddsCount", func(t *testing.T) {
		assert.Equal(t, bfutils.OddsCount, OddsStrLen)
	})
}

func TestOddExists(t *testing.T) {
	tests := map[string]struct {
		index    int
		expected float64
	}{
		"odd[1.01] exists": {index: 0, expected: 1.01},
		"odd[1.1] exists":  {index: 9, expected: 1.1},
		"odd[2] exists":    {index: 99, expected: 2},
		"odd[510] exists":  {index: 300, expected: 510},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, bfutils.Odds[test.index])
		})
	}
}

func TestOddRange(t *testing.T) {
	tests := map[string]struct {
		index         int
		expectedBegin float64
		expectedEnd   float64
	}{
		"odd range [0] boundaries": {index: 0, expectedBegin: 1.01, expectedEnd: 2.0},
		"odd range [3] boundaries": {index: 3, expectedBegin: 4.1, expectedEnd: 6.0},
		"odd range [9] boundaries": {index: 9, expectedBegin: 110, expectedEnd: 1000},
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
		odd           float64
		expectedIndex int
		expectedOdd   float64
		expectedErr   bool
	}{
		// Boundary tests
		"odd[-1] match":    {odd: -1, expectedErr: true},
		"odd[0] match":     {odd: 0, expectedErr: true},
		"odd[1] match":     {odd: 1, expectedErr: true},
		"odd[99999] match": {odd: 99999, expectedErr: true},

		"odd[1.01] match": {odd: 1.01, expectedIndex: 0, expectedOdd: 1.01},
		"odd[1000] match": {odd: 1000, expectedIndex: 349, expectedOdd: 1000},

		"odd[3.2552321] match":      {odd: 3.2552321, expectedIndex: 154, expectedOdd: 3.25},
		"odd[3.299999999999] match": {odd: 3.299999999999, expectedIndex: 155, expectedOdd: 3.3},

		"odd[2.0] match":  {odd: 2.0, expectedIndex: 99, expectedOdd: 2.0},
		"odd[4.05] match": {odd: 4.05, expectedIndex: 169, expectedOdd: 4.0},
		"odd[33] match":   {odd: 33, expectedIndex: 240, expectedOdd: 32},
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
			assert.Equal(t, test.expectedOdd, odd)
		})
	}
}

func TestOddCeil(t *testing.T) {
	tests := map[string]struct {
		odd           float64
		expectedIndex int
		expectedOdd   float64
		expectedErr   bool
	}{
		// Boundary tests
		"odd[-1] match":    {odd: -1, expectedErr: true},
		"odd[0] match":     {odd: 0, expectedErr: true},
		"odd[1] match":     {odd: 1, expectedErr: true},
		"odd[99999] match": {odd: 99999, expectedErr: true},

		"odd[1.01] match": {odd: 1.01, expectedIndex: 0, expectedOdd: 1.01},
		"odd[1000] match": {odd: 1000, expectedIndex: 349, expectedOdd: 1000},

		"odd[3.2552321] match":      {odd: 3.2552321, expectedIndex: 155, expectedOdd: 3.3},
		"odd[3.249999999999] match": {odd: 3.249999999999, expectedIndex: 154, expectedOdd: 3.25},

		"odd[2.0] match":  {odd: 2.0, expectedIndex: 99, expectedOdd: 2.0},
		"odd[4.05] match": {odd: 4.05, expectedIndex: 170, expectedOdd: 4.1},
		"odd[33] match":   {odd: 33, expectedIndex: 241, expectedOdd: 34},
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
			assert.Equal(t, test.expectedOdd, odd)
		})
	}
}

func TestOddRound(t *testing.T) {
	tests := map[string]struct {
		odd           float64
		expectedIndex int
		expectedOdd   float64
		expectedErr   bool
	}{
		// Boundary tests
		"odd[-1] match":    {odd: -1, expectedErr: true},
		"odd[0] match":     {odd: 0, expectedErr: true},
		"odd[1] match":     {odd: 1, expectedErr: true},
		"odd[99999] match": {odd: 99999, expectedErr: true},

		"odd[1.01] match": {odd: 1.01, expectedIndex: 0, expectedOdd: 1.01},
		"odd[1000] match": {odd: 1000, expectedIndex: 349, expectedOdd: 1000},

		"odd[3.2552321] match":      {odd: 3.2552321, expectedIndex: 154, expectedOdd: 3.25},
		"odd[3.249999999999] match": {odd: 3.249999999999, expectedIndex: 154, expectedOdd: 3.25},

		"odd[2.0] match":   {odd: 2.0, expectedIndex: 99, expectedOdd: 2.0},
		"odd[4.077] match": {odd: 4.077, expectedIndex: 170, expectedOdd: 4.1},
		"odd[32.9] match":  {odd: 32.9, expectedIndex: 240, expectedOdd: 32},
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
			assert.Equal(t, test.expectedOdd, odd)
		})
	}
}

func TestOddShift(t *testing.T) {
	tests := map[string]struct {
		roundType     bfutils.RoundType
		odd           float64
		shift         int
		expectedIndex int
		expectedOdd   float64
		expectedErr   bool
	}{
		// Boundary tests
		"odd[-1] shift":       {odd: -1, expectedErr: true},
		"odd[0] shift":        {odd: 0, expectedErr: true},
		"odd[1] shift":        {odd: 1, expectedErr: true},
		"odd[99999] shift":    {odd: 99999, expectedErr: true},
		"odd[10, 1000] shift": {roundType: bfutils.RoundType_Ceil, odd: 10, shift: 1000, expectedErr: true},

		"odd[1.01, 3] shift": {roundType: bfutils.RoundType_Round, odd: 1.01, shift: 3, expectedIndex: 3, expectedOdd: 1.04},
		"odd[4, -10] shift":  {roundType: bfutils.RoundType_Ceil, odd: 4, shift: -10, expectedIndex: 159, expectedOdd: 3.5},
		"odd[10, 5] shift":   {roundType: bfutils.RoundType_Floor, odd: 10, shift: 5, expectedIndex: 214, expectedOdd: 12.5},
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
			assert.Equal(t, test.expectedOdd, odd)
		})
	}
}

func TestOddsTicksDiff(t *testing.T) {
	tests := map[string]struct {
		roundType    bfutils.RoundType
		odd1         float64
		odd2         float64
		expectedDiff int
		expectedErr  bool
	}{
		// Boundary tests
		"odds[-1, 10] tick diff":    {odd1: -1, odd2: 10, expectedErr: true},
		"odds[0, 5] tick diff":      {odd1: 0, odd2: 5, expectedErr: true},
		"odds[50, 1] tick diff":     {odd1: 50, odd2: 1, expectedErr: true},
		"odds[10, 99999] tick diff": {odd1: 10, odd2: 99999, expectedErr: true},

		"odds[1.01, 3] tick diff": {roundType: bfutils.RoundType_Round, odd1: 1.01, odd2: 3, expectedDiff: 149},
		"odds[4, 4.5] tick diff":  {roundType: bfutils.RoundType_Ceil, odd1: 4, odd2: 4.5, expectedDiff: 5},
		"odds[10, 5] tick diff":   {roundType: bfutils.RoundType_Floor, odd1: 10, odd2: 5, expectedDiff: 30},
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
		odd      float64
		expected bool
	}{
		// Boundary tests
		"odds[-1] boundary check":   {odd: -1, expected: false},
		"odds[0] boundary check":    {odd: 0, expected: false},
		"odds[1] boundary check":    {odd: 1, expected: false},
		"odds[1001] boundary check": {odd: 1001, expected: false},

		"odds[2.3] boundary check": {odd: 2.3, expected: true},
		"odds[50] boundary check":  {odd: 50, expected: true},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			value := bfutils.OddWithinBoundaries(test.odd)
			assert.Equal(t, test.expected, value)
		})
	}
}

func TestRoundTypeEnum(t *testing.T) {
	var enum bfutils.RoundType = bfutils.RoundType_Floor
	assert.Equal(t, "Floor", enum.String())
}
