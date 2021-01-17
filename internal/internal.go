package internal

import "math"

// EqualWithTolerance is a helper function and constant to help estimate whether odd matches or not
func EqualWithTolerance(a float64, b float64) bool {
	const float64EqualityThreshold = 1e-9
	return math.Abs(a-b) <= float64EqualityThreshold
}
