package internal

import "math"

const float64EqualityThreshold float64 = 1e-9

// EqualWithTolerance is a helper function and constant to help estimate whether odd matches or not
func EqualWithTolerance(a float64, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}
