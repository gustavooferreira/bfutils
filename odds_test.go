package bfutils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gustavooferreira/bfutils"
)

func TestOddsSize(t *testing.T) {
	OddsFixedLen := len(bfutils.OddsFixed)
	OddsLen := len(bfutils.Odds)
	OddsStrLen := len(bfutils.OddsStr)

	t.Run("OddsFixed array length equal to OddsCount", func(t *testing.T) {
		assert.Equal(t, bfutils.OddsCount, OddsFixedLen)
	})

	t.Run("Odds array length equal to OddsCount", func(t *testing.T) {
		assert.Equal(t, bfutils.OddsCount, OddsLen)
	})

	t.Run("OddsStr array length equal to OddsCount", func(t *testing.T) {
		assert.Equal(t, bfutils.OddsCount, OddsStrLen)
	})
}

func TestOddExists(t *testing.T) {
	tests := map[string]struct {
		index            int
		expectedOddFixed bfutils.FixedOdd
		expectedOdd      float64
		expectedOddStr   string
	}{
		"odd[1.01] exists": {index: 0, expectedOddFixed: 101, expectedOdd: 1.01, expectedOddStr: "1.01"},
		"odd[1.1] exists":  {index: 9, expectedOddFixed: 110, expectedOdd: 1.1, expectedOddStr: "1.1"},
		"odd[2] exists":    {index: 99, expectedOddFixed: 200, expectedOdd: 2, expectedOddStr: "2"},
		"odd[510] exists":  {index: 300, expectedOddFixed: 51000, expectedOdd: 510, expectedOddStr: "510"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expectedOddFixed, bfutils.OddsFixed[test.index])
			assert.Equal(t, test.expectedOdd, bfutils.Odds[test.index])
			assert.Equal(t, test.expectedOddStr, bfutils.OddsStr[test.index])
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
			assert.Equal(t, test.expectedBegin, bfutils.OddsRange[test.index].Begin.Float)
			assert.Equal(t, test.expectedEnd, bfutils.OddsRange[test.index].End.Float)
		})
	}
}
