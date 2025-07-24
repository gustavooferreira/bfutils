package bfutils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/gustavooferreira/bfutils"
)

func TestFindOdd(t *testing.T) {
	tests := map[string]struct {
		odd           float64
		expectedMatch bool
		expectedIndex int
		expectedErr   bool
	}{
		// Boundary tests
		"odds[-1] boundary check":   {odd: -1, expectedErr: true},
		"odds[0] boundary check":    {odd: 0, expectedErr: true},
		"odds[1] boundary check":    {odd: 1, expectedErr: true},
		"odds[1001] boundary check": {odd: 1001, expectedErr: true},

		"odds[2.3] boundary check":  {odd: 2.3, expectedMatch: true, expectedIndex: 114},
		"odds[2.29] boundary check": {odd: 2.29, expectedMatch: false, expectedIndex: 113},
		"odds[50] boundary check":   {odd: 50, expectedMatch: true, expectedIndex: 249},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			match, index, err := bfutils.FindOdd(test.odd)

			if test.expectedErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expectedMatch, match)
				assert.Equal(t, test.expectedIndex, index)
				t.Logf("Odd matched: %.2f", bfutils.Odds[index])
			}
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
			value := bfutils.IsOddWithinBoundaries(test.odd)
			assert.Equal(t, test.expected, value)
		})
	}
}

func TestFindFixedOdd(t *testing.T) {
	tests := map[string]struct {
		odd           bfutils.FixedOdd
		expectedIndex int
		expectedOk    bool
	}{
		// Boundary tests
		"odds[-1] boundary check":     {odd: -1, expectedOk: false},
		"odds[0] boundary check":      {odd: 0, expectedOk: false},
		"odds[1] boundary check":      {odd: 1, expectedOk: false},
		"odds[100001] boundary check": {odd: 100001, expectedOk: false},

		"odds[2.3] boundary check": {odd: 230, expectedIndex: 114, expectedOk: true},
		"odds[50] boundary check":  {odd: 5000, expectedIndex: 249, expectedOk: true},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			index, ok := bfutils.FindFixedOdd(test.odd)

			require.Equal(t, test.expectedOk, ok)

			if ok {
				assert.Equal(t, test.expectedIndex, index)
				t.Logf("Odd matched: %.2f", bfutils.Odds[index])
			}
		})
	}
}

func TestFindStrOdd(t *testing.T) {
	tests := map[string]struct {
		odd           string
		expectedIndex int
		expectedOk    bool
	}{
		// Boundary tests
		"odds[-1] boundary check":     {odd: "-1", expectedOk: false},
		"odds[0] boundary check":      {odd: "0", expectedOk: false},
		"odds[1] boundary check":      {odd: "1", expectedOk: false},
		"odds[100001] boundary check": {odd: "1001", expectedOk: false},

		"odds[2.3] boundary check": {odd: "2.3", expectedIndex: 114, expectedOk: true},
		"odds[50] boundary check":  {odd: "50", expectedIndex: 249, expectedOk: true},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			index, ok := bfutils.FindStrOdd(test.odd)

			require.Equal(t, test.expectedOk, ok)

			if ok {
				assert.Equal(t, test.expectedIndex, index)
				t.Logf("Odd matched: %.2f", bfutils.Odds[index])
			}
		})
	}
}

func TestOddFloor(t *testing.T) {
	tests := map[string]struct {
		odd              float64
		expectedIndex    int
		expectedFixedOdd bfutils.FixedOdd
		expectedErr      bool
	}{
		// Boundary tests
		"odd[-1] match":    {odd: -1, expectedErr: true},
		"odd[0] match":     {odd: 0, expectedErr: true},
		"odd[1] match":     {odd: 1, expectedErr: true},
		"odd[99999] match": {odd: 99999, expectedErr: true},

		"odd[1.01] match": {odd: 1.01, expectedIndex: 0, expectedFixedOdd: 101},
		"odd[1000] match": {odd: 1000, expectedIndex: 349, expectedFixedOdd: 100000},

		"odd[3.2552321] match":      {odd: 3.2552321, expectedIndex: 154, expectedFixedOdd: 325},
		"odd[3.299999999999] match": {odd: 3.299999999999, expectedIndex: 155, expectedFixedOdd: 330},

		"odd[2.0] match":  {odd: 2.0, expectedIndex: 99, expectedFixedOdd: 200},
		"odd[4.05] match": {odd: 4.05, expectedIndex: 169, expectedFixedOdd: 400},
		"odd[33] match":   {odd: 33, expectedIndex: 240, expectedFixedOdd: 3200},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			index, odd, err := bfutils.OddFloor(test.odd)

			if test.expectedErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expectedIndex, index)
				assert.Equal(t, test.expectedFixedOdd, odd)
				t.Logf("Odd matched: %.2f", bfutils.Odds[index])
			}
		})
	}
}

func TestOddCeil(t *testing.T) {
	tests := map[string]struct {
		odd              float64
		expectedIndex    int
		expectedFixedOdd bfutils.FixedOdd
		expectedErr      bool
	}{
		// Boundary tests
		"odd[-1] match":    {odd: -1, expectedErr: true},
		"odd[0] match":     {odd: 0, expectedErr: true},
		"odd[1] match":     {odd: 1, expectedErr: true},
		"odd[99999] match": {odd: 99999, expectedErr: true},

		"odd[1.01] match": {odd: 1.01, expectedIndex: 0, expectedFixedOdd: 101},
		"odd[1000] match": {odd: 1000, expectedIndex: 349, expectedFixedOdd: 100000},

		"odd[3.2552321] match":      {odd: 3.2552321, expectedIndex: 155, expectedFixedOdd: 330},
		"odd[3.249999999999] match": {odd: 3.249999999999, expectedIndex: 154, expectedFixedOdd: 325},

		"odd[2.0] match":  {odd: 2.0, expectedIndex: 99, expectedFixedOdd: 200},
		"odd[4.05] match": {odd: 4.05, expectedIndex: 170, expectedFixedOdd: 410},
		"odd[33] match":   {odd: 33, expectedIndex: 241, expectedFixedOdd: 3400},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			index, odd, err := bfutils.OddCeil(test.odd)

			if test.expectedErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expectedIndex, index)
				assert.Equal(t, test.expectedFixedOdd, odd)
				t.Logf("Odd matched: %.2f", bfutils.Odds[index])
			}
		})
	}
}

func TestOddRound(t *testing.T) {
	tests := map[string]struct {
		odd              float64
		expectedIndex    int
		expectedFixedOdd bfutils.FixedOdd
		expectedErr      bool
	}{
		// Boundary tests
		"odd[-1] match":    {odd: -1, expectedErr: true},
		"odd[0] match":     {odd: 0, expectedErr: true},
		"odd[1] match":     {odd: 1, expectedErr: true},
		"odd[99999] match": {odd: 99999, expectedErr: true},

		"odd[1.01] match": {odd: 1.01, expectedIndex: 0, expectedFixedOdd: 101},
		"odd[1000] match": {odd: 1000, expectedIndex: 349, expectedFixedOdd: 100000},

		"odd[3.2552321] match":      {odd: 3.2552321, expectedIndex: 154, expectedFixedOdd: 325},
		"odd[3.249999999999] match": {odd: 3.249999999999, expectedIndex: 154, expectedFixedOdd: 325},

		"odd[2.0] match":   {odd: 2.0, expectedIndex: 99, expectedFixedOdd: 200},
		"odd[4.077] match": {odd: 4.077, expectedIndex: 170, expectedFixedOdd: 410},
		"odd[32.9] match":  {odd: 32.9, expectedIndex: 240, expectedFixedOdd: 3200},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			index, odd, err := bfutils.OddRound(test.odd)

			if test.expectedErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expectedIndex, index)
				assert.Equal(t, test.expectedFixedOdd, odd)
				t.Logf("Odd matched: %.2f", bfutils.Odds[index])
			}
		})
	}
}

func TestOddShift(t *testing.T) {
	tests := map[string]struct {
		fixedOdd         bfutils.FixedOdd
		shift            int
		expectedIndex    int
		expectedFixedOdd bfutils.FixedOdd
		expectedErr      bool
	}{
		// Boundary tests
		"odd[-1] shift":       {fixedOdd: -1, expectedErr: true},
		"odd[0] shift":        {fixedOdd: 0, expectedErr: true},
		"odd[1] shift":        {fixedOdd: 1, expectedErr: true},
		"odd[99999] shift":    {fixedOdd: 9999900, expectedErr: true},
		"odd[10, 1000] shift": {fixedOdd: 1000, shift: 1000, expectedErr: true},

		"odd[1.01, 3] shift": {fixedOdd: 101, shift: 3, expectedIndex: 3, expectedFixedOdd: 104},
		"odd[4, -10] shift":  {fixedOdd: 400, shift: -10, expectedIndex: 159, expectedFixedOdd: 350},
		"odd[10, 5] shift":   {fixedOdd: 1000, shift: 5, expectedIndex: 214, expectedFixedOdd: 1250},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			index, odd, err := bfutils.OddShift(test.fixedOdd, test.shift)

			if test.expectedErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expectedIndex, index)
				assert.Equal(t, test.expectedFixedOdd, odd)
				t.Logf("Odd matched: %.2f", bfutils.Odds[index])
			}
		})
	}
}

func TestOddsTicksDiff(t *testing.T) {
	tests := map[string]struct {
		fixedOdd1    bfutils.FixedOdd
		fixedOdd2    bfutils.FixedOdd
		expectedDiff int
		expectedErr  bool
	}{
		// Boundary tests
		"odds[-1, 10] tick diff":    {fixedOdd1: -100, fixedOdd2: 1000, expectedErr: true},
		"odds[0, 5] tick diff":      {fixedOdd1: 0, fixedOdd2: 500, expectedErr: true},
		"odds[50, 1] tick diff":     {fixedOdd1: 5000, fixedOdd2: 100, expectedErr: true},
		"odds[10, 99999] tick diff": {fixedOdd1: 1000, fixedOdd2: 9999900, expectedErr: true},

		"odds[1.01, 3] tick diff": {fixedOdd1: 101, fixedOdd2: 300, expectedDiff: 149},
		"odds[4, 4.5] tick diff":  {fixedOdd1: 400, fixedOdd2: 450, expectedDiff: 5},
		"odds[10, 5] tick diff":   {fixedOdd1: 1000, fixedOdd2: 500, expectedDiff: 30},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			diff, err := bfutils.OddsTicksDiff(test.fixedOdd1, test.fixedOdd2)

			assert.Equal(t, test.expectedDiff, diff)

			if test.expectedErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expectedDiff, diff)
			}
		})
	}
}
