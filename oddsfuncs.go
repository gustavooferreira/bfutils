package bfutils

import (
	"fmt"
	"math"

	"github.com/gustavooferreira/bfutils/internal"
)

// FindOdd tries to find the odd in the ladder.
// If it finds it, then index is the odd index in the ladder.
// If it doesn't find it, it will return the index in the ladder that is the closest to the odd on the left side
// (towards 1.01).
func FindOdd(odd float64) (match bool, index int, err error) {
	if withinBoundary := IsOddWithinBoundaries(odd); !withinBoundary {
		return false, 0, fmt.Errorf("odd provided [%.2f] is outside of trading range", odd)
	}

	if internal.EqualWithTolerance(odd, 1000) {
		return true, OddsCount - 1, nil
	} else if internal.EqualWithTolerance(odd, 1.01) {
		return true, 0, nil
	}

	lo := 0
	hi := OddsCount

	for lo < hi {
		mid := (lo + hi) / 2

		if internal.EqualWithTolerance(odd, Odds[mid]) {
			return true, mid, nil
		} else if odd < Odds[mid] {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	return false, lo - 1, nil
}

// IsOddWithinBoundaries checks if odd is within trading range.
// I.e., odd is between 1.01 and 1000.
func IsOddWithinBoundaries(odd float64) bool {
	if internal.EqualWithTolerance(odd, 1000) {
		return true
	} else if odd > 1000 {
		return false
	}

	if internal.EqualWithTolerance(odd, 1.01) {
		return true
	} else if odd < 1.01 {
		return false
	}

	return true
}

// FindFixedOdd tries to find the fixed odd in the ladder.
// If it finds it, then index is the odd index in the ladder and ok is set to true.
// Otherwise, ok is set to false.
func FindFixedOdd(odd FixedOdd) (index int, ok bool) {
	// boundaries check
	if odd > OddsFixed[OddsCount-1] || odd < OddsFixed[0] {
		return 0, false
	}

	lo := 0
	hi := OddsCount

	for lo < hi {
		mid := (lo + hi) / 2

		if odd == OddsFixed[mid] {
			return mid, true
		} else if odd < OddsFixed[mid] {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	// didn't find the odd
	return 0, false
}

// FindStrOdd tries to find the odd string in the ladder.
// If it finds it, then index is the odd index in the ladder and ok is set to true.
// Otherwise, ok is set to false.
// NOTE: This function performs a linear search!
// NOTE: This function performs string exact matches. Therefore, "2.3" won't match "2.30".
func FindStrOdd(odd string) (index int, ok bool) {
	for i, oddTarger := range OddsStr {
		if odd == oddTarger {
			return i, true
		}
	}

	// didn't find the odd
	return 0, false
}

// OddFloor returns the same odd input rounded towards 1.01.
// If the odd supplied is one of the available odds in the ladder, than the same odd is returned.
// index returns the index of the odd in the ladder.
func OddFloor(odd float64) (index int, fixedOdd FixedOdd, err error) {
	_, index, err = FindOdd(odd)
	if err != nil {
		return 0, 0, err
	}

	return index, OddsFixed[index], nil
}

// OddCeil returns the same odd input rounded towards 1000.
// If the odd supplied is one of the available odds in the ladder, than the same odd is returned.
// index returns the index of the odd in the ladder.
func OddCeil(odd float64) (index int, fixedOdd FixedOdd, err error) {
	match, index, err := FindOdd(odd)
	if err != nil {
		return 0, 0, err
	}

	if match {
		return index, OddsFixed[index], nil
	}
	return index + 1, OddsFixed[index+1], nil
}

// OddRound returns the same odd input rounded to the nearest odd in the ladder.
// If the odd supplied is one of the available odds in the ladder, than the same odd is returned.
// index returns the index of the odd in the ladder.
func OddRound(odd float64) (index int, fixedOdd FixedOdd, err error) {
	match, index, err := FindOdd(odd)
	if err != nil {
		return 0, 0, err
	}

	if match {
		return index, OddsFixed[index], nil
	}

	// Compute deltas
	delta1 := math.Abs(odd - Odds[index])
	delta2 := math.Abs(odd - Odds[index+1])

	if delta1 <= delta2 {
		return index, OddsFixed[index], nil
	}
	return index + 1, OddsFixed[index+1], nil
}

// OddShift shifts the Odd up or down in the ladder.
// If shift is higher than zero, it shifts the odd towards 1000, if it's less than zero it shifts the odd towards 1.01.
// roundType is the round method to be used.
// shift represents the number of ticks to shift the odd.
func OddShift(fixedOdd FixedOdd, shift int) (index int, fixedOddOut FixedOdd, err error) {
	index, ok := FindFixedOdd(fixedOdd)
	if !ok {
		return 0, 0, fmt.Errorf("input odd not found in the ladder")
	}

	index += shift

	if (index >= OddsCount) || (index < 0) {
		return 0, 0, fmt.Errorf("computed odd outside of tradable range")
	}
	return index, OddsFixed[index], nil
}

// OddsTicksDiff computes the number of ticks between two odds.
func OddsTicksDiff(fixedOdd1 FixedOdd, fixedOdd2 FixedOdd) (ticksDiff int, err error) {
	index1, ok := FindFixedOdd(fixedOdd1)
	if !ok {
		return 0, fmt.Errorf("input odd 1 not found in the ladder")
	}

	index2, ok := FindFixedOdd(fixedOdd2)
	if !ok {
		return 0, fmt.Errorf("input odd 2 not found in the ladder")
	}

	return int(math.Abs(float64(index2 - index1))), nil
}

// RoundType is the round method to be used.
type RoundType uint

const (
	// RoundType_Ceil round towards 1000.
	RoundType_Ceil = iota
	// RoundType_Round round to the nearest odd in the ladder.
	RoundType_Round
	// RoundType_Floor round towards 1.01.
	RoundType_Floor
)

func (rt RoundType) String() string {
	return [...]string{"Ceil", "Round", "Floor"}[rt]
}
