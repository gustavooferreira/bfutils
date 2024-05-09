package bfutils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/gustavooferreira/bfutils"
)

func TestOddsSize(t *testing.T) {
	t.Run("Odds array length equal to OddsCount", func(t *testing.T) {
		OddsLen := len(bfutils.Odds)
		assert.Equal(t, bfutils.OddsCount, OddsLen)
	})

	t.Run("OddsStr array length equal to OddsCount", func(t *testing.T) {
		OddsStrLen := len(bfutils.OddsStr)
		assert.Equal(t, bfutils.OddsCount, OddsStrLen)
	})
}

func TestOddExists(t *testing.T) {
	testCases := []struct {
		name     string
		index    int
		expected float64
	}{
		{name: "odd[1.01] exists", index: 0, expected: 1.01},
		{name: "odd[1.1] exists", index: 9, expected: 1.1},
		{name: "odd[2] exists", index: 99, expected: 2},
		{name: "odd[510] exists", index: 300, expected: 510},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, bfutils.Odds[tc.index])
		})
	}
}

func TestOddRange(t *testing.T) {
	testCases := []struct {
		name              string
		index             int
		expectedBegin     float64
		expectedEnd       float64
		expectedIncrement float64
		expectedTicks     int
	}{
		{
			name:              "odd range [idx: 0] boundaries",
			index:             0,
			expectedBegin:     1.01,
			expectedEnd:       2.0,
			expectedIncrement: 0.01,
			expectedTicks:     100,
		},
		{
			name:              "odd range [idx: 3] boundaries",
			index:             3,
			expectedBegin:     4.1,
			expectedEnd:       6.0,
			expectedIncrement: 0.1,
			expectedTicks:     20,
		},
		{
			name:              "odd range [idx: 9] boundaries",
			index:             9,
			expectedBegin:     110,
			expectedEnd:       1000,
			expectedIncrement: 10,
			expectedTicks:     90,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedBegin, bfutils.OddsRangeTable[tc.index].Begin)
			assert.Equal(t, tc.expectedEnd, bfutils.OddsRangeTable[tc.index].End)
			assert.Equal(t, tc.expectedIncrement, bfutils.OddsRangeTable[tc.index].Increment)
			assert.Equal(t, tc.expectedTicks, bfutils.OddsRangeTable[tc.index].Ticks)
		})
	}
}

func TestOddFloor(t *testing.T) {
	testCases := []struct {
		name          string
		odd           float64
		expectedIndex int
		expectedOdd   float64
		expectedErr   bool
	}{
		// Boundary tests
		{name: "odd[-1] match", odd: -1, expectedErr: true},
		{name: "odd[0] match", odd: 0, expectedErr: true},
		{name: "odd[1] match", odd: 1, expectedErr: true},
		{name: "odd[99999] match", odd: 99999, expectedErr: true},

		{name: "odd[1.01] match", odd: 1.01, expectedIndex: 0, expectedOdd: 1.01},
		{name: "odd[1000] match", odd: 1000, expectedIndex: 349, expectedOdd: 1000},

		{name: "odd[3.2552321] match", odd: 3.2552321, expectedIndex: 154, expectedOdd: 3.25},
		{name: "odd[3.299999999999] match", odd: 3.299999999999, expectedIndex: 155, expectedOdd: 3.3},

		{name: "odd[3.29] match", odd: 3.29, expectedIndex: 154, expectedOdd: 3.25},
		{name: "odd[2.0] match", odd: 2.0, expectedIndex: 99, expectedOdd: 2.0},
		{name: "odd[4.05] match", odd: 4.05, expectedIndex: 169, expectedOdd: 4.0},
		{name: "odd[33] match", odd: 33, expectedIndex: 240, expectedOdd: 32},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			index, odd, err := bfutils.OddFloor(tc.odd)
			if tc.expectedErr {
				require.Error(t, err)
				return // end of test
			}

			require.NoError(t, err)
			assert.Equal(t, tc.expectedIndex, index)
			assert.Equal(t, tc.expectedOdd, odd)
		})
	}
}

func TestOddCeil(t *testing.T) {
	testCases := []struct {
		name          string
		odd           float64
		expectedIndex int
		expectedOdd   float64
		expectedErr   bool
	}{
		// Boundary tests
		{name: "odd[-1] match", odd: -1, expectedErr: true},
		{name: "odd[0] match", odd: 0, expectedErr: true},
		{name: "odd[1] match", odd: 1, expectedErr: true},
		{name: "odd[99999] match", odd: 99999, expectedErr: true},

		{name: "odd[1.01] match", odd: 1.01, expectedIndex: 0, expectedOdd: 1.01},
		{name: "odd[1000] match", odd: 1000, expectedIndex: 349, expectedOdd: 1000},

		{name: "odd[3.2552321] match", odd: 3.2552321, expectedIndex: 155, expectedOdd: 3.3},
		{name: "odd[3.249999999999] match", odd: 3.249999999999, expectedIndex: 154, expectedOdd: 3.25},

		{name: "odd[2.0] match", odd: 2.0, expectedIndex: 99, expectedOdd: 2.0},
		{name: "odd[4.05] match", odd: 4.05, expectedIndex: 170, expectedOdd: 4.1},
		{name: "odd[33] match", odd: 33, expectedIndex: 241, expectedOdd: 34},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			index, odd, err := bfutils.OddCeil(tc.odd)
			if tc.expectedErr {
				require.Error(t, err)
				return // end of test
			}

			require.NoError(t, err)
			assert.Equal(t, tc.expectedIndex, index)
			assert.Equal(t, tc.expectedOdd, odd)
		})
	}
}

func TestOddRound(t *testing.T) {
	testCases := []struct {
		name          string
		odd           float64
		expectedIndex int
		expectedOdd   float64
		expectedErr   bool
	}{
		// Boundary tests
		{name: "odd[-1] match", odd: -1, expectedErr: true},
		{name: "odd[0] match", odd: 0, expectedErr: true},
		{name: "odd[1] match", odd: 1, expectedErr: true},
		{name: "odd[99999] match", odd: 99999, expectedErr: true},

		{name: "odd[1.01] match", odd: 1.01, expectedIndex: 0, expectedOdd: 1.01},
		{name: "odd[1000] match", odd: 1000, expectedIndex: 349, expectedOdd: 1000},

		{name: "odd[3.2552321] match", odd: 3.2552321, expectedIndex: 154, expectedOdd: 3.25},
		{name: "odd[3.249999999999] match", odd: 3.249999999999, expectedIndex: 154, expectedOdd: 3.25},

		{name: "odd[2.0] match", odd: 2.0, expectedIndex: 99, expectedOdd: 2.0},
		{name: "odd[4.077] match", odd: 4.077, expectedIndex: 170, expectedOdd: 4.1},
		{name: "odd[32.9] match", odd: 32.9, expectedIndex: 240, expectedOdd: 32},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			index, odd, err := bfutils.OddRound(tc.odd)
			if tc.expectedErr {
				require.Error(t, err)
				return // end of test
			}

			require.NoError(t, err)
			assert.Equal(t, tc.expectedIndex, index)
			assert.Equal(t, tc.expectedOdd, odd)
		})
	}
}

func TestOddShift(t *testing.T) {
	testCases := []struct {
		name          string
		roundType     bfutils.RoundType
		odd           float64
		shift         int
		expectedIndex int
		expectedOdd   float64
		expectedErr   bool
	}{
		// Boundary tests
		{name: "odd[-1] shift", odd: -1, expectedErr: true},
		{name: "odd[0] shift", odd: 0, expectedErr: true},
		{name: "odd[1] shift", odd: 1, expectedErr: true},
		{name: "odd[99999] shift", odd: 99999, expectedErr: true},
		{name: "odd[10, 1000] shift", roundType: bfutils.RoundType_Ceil, odd: 10, shift: 1000, expectedErr: true},

		{
			name:          "odd[1.01, 3] shift",
			roundType:     bfutils.RoundType_Round,
			odd:           1.01,
			shift:         3,
			expectedIndex: 3,
			expectedOdd:   1.04,
		},
		{
			name:          "odd[4, -10] shift",
			roundType:     bfutils.RoundType_Ceil,
			odd:           4,
			shift:         -10,
			expectedIndex: 159,
			expectedOdd:   3.5,
		},
		{
			name:          "odd[10, 5] shift",
			roundType:     bfutils.RoundType_Floor,
			odd:           10,
			shift:         5,
			expectedIndex: 214,
			expectedOdd:   12.5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			index, odd, err := bfutils.OddShift(tc.odd, tc.shift, tc.roundType)
			if tc.expectedErr {
				require.Error(t, err)
				return // end of test
			}

			require.NoError(t, err)
			assert.Equal(t, tc.expectedIndex, index)
			assert.Equal(t, tc.expectedOdd, odd)
		})
	}
}

func TestOddsTicksDiff(t *testing.T) {
	testCases := []struct {
		name         string
		roundType    bfutils.RoundType
		odd1         float64
		odd2         float64
		expectedDiff int
		expectedErr  bool
	}{
		// Boundary tests
		{name: "odds[-1, 10] tick diff", odd1: -1, odd2: 10, expectedErr: true},
		{name: "odds[0, 5] tick diff", odd1: 0, odd2: 5, expectedErr: true},
		{name: "odds[50, 1] tick diff", odd1: 50, odd2: 1, expectedErr: true},
		{name: "odds[10, 99999] tick diff", odd1: 10, odd2: 99999, expectedErr: true},

		{name: "odds[1.01, 3] tick diff", roundType: bfutils.RoundType_Round, odd1: 1.01, odd2: 3, expectedDiff: 149},
		{name: "odds[4, 4.5] tick diff", roundType: bfutils.RoundType_Ceil, odd1: 4, odd2: 4.5, expectedDiff: 5},
		{name: "odds[10, 5] tick diff", roundType: bfutils.RoundType_Floor, odd1: 10, odd2: 5, expectedDiff: 30},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var errBool bool
			diff, err := bfutils.OddsTicksDiff(tc.odd1, tc.odd2, tc.roundType)
			if err != nil {
				errBool = true
			}

			require.Equal(t, tc.expectedErr, errBool)
			assert.Equal(t, tc.expectedDiff, diff)
		})
	}
}

func TestOddWithinBoundaries(t *testing.T) {
	testCases := []struct {
		name           string
		odd            float64
		expectedResult bool
	}{
		// Boundary tests
		{name: "odds[-1] boundary check", odd: -1, expectedResult: false},
		{name: "odds[0] boundary check", odd: 0, expectedResult: false},
		{name: "odds[1] boundary check", odd: 1, expectedResult: false},
		{name: "odds[1.01] boundary check", odd: 1.01, expectedResult: true},
		{name: "odds[1001] boundary check", odd: 1001, expectedResult: false},

		{name: "odds[2.3] boundary check", odd: 2.3, expectedResult: true},
		{name: "odds[50] boundary check", odd: 50, expectedResult: true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			value := bfutils.IsOddWithinBoundaries(tc.odd)
			assert.Equal(t, tc.expectedResult, value)
		})
	}
}

func TestFindOdd(t *testing.T) {
	testCases := []struct {
		name          string
		odd           float64
		expectedMatch bool
		expectedIndex int
		expectedErr   bool
	}{
		{name: "find odd [-1]", odd: -1, expectedErr: true},
		{name: "find odd [0]", odd: 0, expectedErr: true},
		{name: "find odd [1]", odd: 1, expectedErr: true},

		{name: "find odd [1.01]", odd: 1.01, expectedErr: false, expectedIndex: 0, expectedMatch: true},
		{name: "find odd [1000]", odd: 1000, expectedErr: false, expectedIndex: 349, expectedMatch: true},

		{name: "find odd [2]", odd: 2, expectedErr: false, expectedIndex: 99, expectedMatch: true},

		{name: "find odd [6.1]", odd: 6.1, expectedErr: false, expectedIndex: 189, expectedMatch: false},
		{name: "find odd [1.0123]", odd: 1.0123, expectedErr: false, expectedIndex: 0, expectedMatch: false},
		{name: "find odd [999.1]", odd: 999.1, expectedErr: false, expectedIndex: 348, expectedMatch: false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			match, index, err := bfutils.FindOdd(tc.odd)
			if tc.expectedErr {
				require.Error(t, err)
				return // end of test
			}

			require.NoError(t, err)
			require.Equal(t, tc.expectedMatch, match)
			assert.Equal(t, tc.expectedIndex, index)

			t.Logf("Odd at index [%d]: %.2f", index, bfutils.Odds[index])
		})
	}
}

func TestRoundTypeEnum(t *testing.T) {
	var enum1 bfutils.RoundType = bfutils.RoundType_Ceil
	assert.Equal(t, "Ceil", enum1.String())

	var enum2 bfutils.RoundType = bfutils.RoundType_Round
	assert.Equal(t, "Round", enum2.String())

	var enum3 bfutils.RoundType = bfutils.RoundType_Floor
	assert.Equal(t, "Floor", enum3.String())
}
