package internal_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gustavooferreira/bfutils/internal"
)

func TestEqualWithTolerance(t *testing.T) {
	testCases := []struct {
		name           string
		a              float64
		b              float64
		expectedResult bool
	}{
		{name: "compare 1", a: 0, b: 0, expectedResult: true},
		{name: "compare 2", a: 10.0000000001, b: 10.0000000002, expectedResult: true},
		{name: "compare 3", a: 9.999999999999, b: 10, expectedResult: true},
		{name: "compare 4", a: 5, b: 10, expectedResult: false},
		{name: "compare 5", a: 9.555555, b: 9.5555556, expectedResult: false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			value := internal.EqualWithTolerance(tc.a, tc.b)
			assert.Equal(t, tc.expectedResult, value)
		})
	}
}
