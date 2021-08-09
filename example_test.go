package bfutils_test

import (
	"fmt"

	"github.com/gustavooferreira/bfutils"
	"github.com/shopspring/decimal"
)

// This example, computes the ticks difference between two odds.
func Example_a() {
	randomOdd, _ := decimal.NewFromString("4.051")

	index1, odd1, err := bfutils.OddShift(bfutils.RoundType_Floor, randomOdd, 10)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Odd1: %s - position in the ladder: %d\n", odd1.StringFixed(2), index1+1)

	index2, odd2, err := bfutils.OddShift(bfutils.RoundType_Floor, randomOdd, -10)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Odd2: %s - position in the ladder: %d\n", odd2.StringFixed(2), index2+1)

	ticksDiff, err := bfutils.OddsTicksDiff(bfutils.RoundType_Floor, odd1, odd2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Ticks difference between both odds: %d\n", ticksDiff)

	// Output:
	// Odd1: 5.00 - position in the ladder: 180
	// Odd2: 3.50 - position in the ladder: 160
	// Ticks difference between both odds: 20
}
