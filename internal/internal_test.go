package internal_test

import (
	"testing"

	"github.com/gustavooferreira/bfutils/internal"
	"github.com/stretchr/testify/assert"
)

func TestEqualWithTolerance(t *testing.T) {
	tests := map[string]struct {
		a        float64
		b        float64
		expected bool
	}{
		"compare 1": {a: 0, b: 0, expected: true},
		"compare 2": {a: 10.0000000001, b: 10.0000000002, expected: true},
		"compare 3": {a: 9.999999999999, b: 10, expected: true},
		"compare 4": {a: 5, b: 10, expected: false},
		"compare 5": {a: 9.555555, b: 9.5555556, expected: false},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			value := internal.EqualWithTolerance(test.a, test.b)
			assert.Equal(t, test.expected, value, "expected boolean")
		})
	}
}
