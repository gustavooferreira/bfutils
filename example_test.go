package bfutils_test

import (
	"fmt"

	"github.com/gustavooferreira/bfutils"
)

// In this example, we computing the ticks difference between two odds.
func Example_a() {
	randomOdd := 4.051

	index1, odd1, err := bfutils.OddShift(bfutils.RoundType_Floor, randomOdd, 10)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Odd1: %.2f - position in the ladder: %d\n", odd1, index1+1)

	index2, odd2, err := bfutils.OddShift(bfutils.RoundType_Floor, randomOdd, -10)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Odd2: %.2f - position in the ladder: %d\n", odd2, index2+1)

	ticksDiff, err := bfutils.OddsTicksDiff(bfutils.RoundType_Floor, odd1, odd2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Ticks difference between both odds: %d", ticksDiff)

	// Output:
	// Odd1: 5.00 - position in the ladder: 180
	// Odd2: 3.50 - position in the ladder: 160
	// Ticks difference between both odds: 20
}
