package internal

import "math"

// EqualWithTolerance is a helper function to help estimate whether a given odd matches or not (within a threshold).
func EqualWithTolerance(a float64, b float64) bool {
	const float64EqualityThreshold = 1e-9
	return math.Abs(a-b) <= float64EqualityThreshold
}
