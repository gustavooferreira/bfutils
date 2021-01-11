// Package bfutils provides a set of utility functions that help with day to day automation in the Betfair exchange
package bfutils

import (
	"fmt"
	"math"
)

// OddsCount is a constant defining the number of available odds in the ladder.
const OddsCount = 350

// Odds is an array of float64 holding the value of all tradable odds.
var Odds = [...]float64{
	1.01, 1.02, 1.03, 1.04, 1.05, 1.06, 1.07, 1.08, 1.09, 1.1,
	1.11, 1.12, 1.13, 1.14, 1.15, 1.16, 1.17, 1.18, 1.19, 1.2,
	1.21, 1.22, 1.23, 1.24, 1.25, 1.26, 1.27, 1.28, 1.29, 1.3,
	1.31, 1.32, 1.33, 1.34, 1.35, 1.36, 1.37, 1.38, 1.39, 1.4,
	1.41, 1.42, 1.43, 1.44, 1.45, 1.46, 1.47, 1.48, 1.49, 1.5,
	1.51, 1.52, 1.53, 1.54, 1.55, 1.56, 1.57, 1.58, 1.59, 1.6,
	1.61, 1.62, 1.63, 1.64, 1.65, 1.66, 1.67, 1.68, 1.69, 1.7,
	1.71, 1.72, 1.73, 1.74, 1.75, 1.76, 1.77, 1.78, 1.79, 1.8,
	1.81, 1.82, 1.83, 1.84, 1.85, 1.86, 1.87, 1.88, 1.89, 1.9,
	1.91, 1.92, 1.93, 1.94, 1.95, 1.96, 1.97, 1.98, 1.99, 2,
	2.02, 2.04, 2.06, 2.08, 2.1, 2.12, 2.14, 2.16, 2.18, 2.2,
	2.22, 2.24, 2.26, 2.28, 2.3, 2.32, 2.34, 2.36, 2.38, 2.4,
	2.42, 2.44, 2.46, 2.48, 2.5, 2.52, 2.54, 2.56, 2.58, 2.6,
	2.62, 2.64, 2.66, 2.68, 2.7, 2.72, 2.74, 2.76, 2.78, 2.8,
	2.82, 2.84, 2.86, 2.88, 2.9, 2.92, 2.94, 2.96, 2.98, 3,
	3.05, 3.1, 3.15, 3.2, 3.25, 3.3, 3.35, 3.4, 3.45, 3.5,
	3.55, 3.6, 3.65, 3.7, 3.75, 3.8, 3.85, 3.9, 3.95, 4,
	4.1, 4.2, 4.3, 4.4, 4.5, 4.6, 4.7, 4.8, 4.9, 5, 5.1, 5.2,
	5.3, 5.4, 5.5, 5.6, 5.7, 5.8, 5.9, 6, 6.2, 6.4, 6.6, 6.8,
	7, 7.2, 7.4, 7.6, 7.8, 8, 8.2, 8.4, 8.6, 8.8, 9, 9.2,
	9.4, 9.6, 9.8, 10, 10.5, 11, 11.5, 12, 12.5, 13, 13.5, 14,
	14.5, 15, 15.5, 16, 16.5, 17, 17.5, 18, 18.5, 19, 19.5, 20,
	21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 32, 34, 36, 38, 40,
	42, 44, 46, 48, 50, 55, 60, 65, 70, 75, 80, 85, 90, 95, 100,
	110, 120, 130, 140, 150, 160, 170, 180, 190, 200, 210, 220,
	230, 240, 250, 260, 270, 280, 290, 300, 310, 320, 330, 340,
	350, 360, 370, 380, 390, 400, 410, 420, 430, 440, 450, 460,
	470, 480, 490, 500, 510, 520, 530, 540, 550, 560, 570, 580,
	590, 600, 610, 620, 630, 640, 650, 660, 670, 680, 690, 700,
	710, 720, 730, 740, 750, 760, 770, 780, 790, 800, 810, 820,
	830, 840, 850, 860, 870, 880, 890, 900, 910, 920, 930, 940,
	950, 960, 970, 980, 990, 1000}

// OddsStr is an array of strings holding the value of all tradable odds.
var OddsStr = [...]string{
	"1.01", "1.02", "1.03", "1.04", "1.05", "1.06", "1.07", "1.08", "1.09", "1.1",
	"1.11", "1.12", "1.13", "1.14", "1.15", "1.16", "1.17", "1.18", "1.19", "1.2",
	"1.21", "1.22", "1.23", "1.24", "1.25", "1.26", "1.27", "1.28", "1.29", "1.3",
	"1.31", "1.32", "1.33", "1.34", "1.35", "1.36", "1.37", "1.38", "1.39", "1.4",
	"1.41", "1.42", "1.43", "1.44", "1.45", "1.46", "1.47", "1.48", "1.49", "1.5",
	"1.51", "1.52", "1.53", "1.54", "1.55", "1.56", "1.57", "1.58", "1.59", "1.6",
	"1.61", "1.62", "1.63", "1.64", "1.65", "1.66", "1.67", "1.68", "1.69", "1.7",
	"1.71", "1.72", "1.73", "1.74", "1.75", "1.76", "1.77", "1.78", "1.79", "1.8",
	"1.81", "1.82", "1.83", "1.84", "1.85", "1.86", "1.87", "1.88", "1.89", "1.9",
	"1.91", "1.92", "1.93", "1.94", "1.95", "1.96", "1.97", "1.98", "1.99", "2",
	"2.02", "2.04", "2.06", "2.08", "2.1", "2.12", "2.14", "2.16", "2.18", "2.2",
	"2.22", "2.24", "2.26", "2.28", "2.3", "2.32", "2.34", "2.36", "2.38", "2.4",
	"2.42", "2.44", "2.46", "2.48", "2.5", "2.52", "2.54", "2.56", "2.58", "2.6",
	"2.62", "2.64", "2.66", "2.68", "2.7", "2.72", "2.74", "2.76", "2.78", "2.8",
	"2.82", "2.84", "2.86", "2.88", "2.9", "2.92", "2.94", "2.96", "2.98", "3",
	"3.05", "3.1", "3.15", "3.2", "3.25", "3.3", "3.35", "3.4", "3.45", "3.5",
	"3.55", "3.6", "3.65", "3.7", "3.75", "3.8", "3.85", "3.9", "3.95", "4",
	"4.1", "4.2", "4.3", "4.4", "4.5", "4.6", "4.7", "4.8", "4.9", "5", "5.1", "5.2",
	"5.3", "5.4", "5.5", "5.6", "5.7", "5.8", "5.9", "6", "6.2", "6.4", "6.6", "6.8",
	"7", "7.2", "7.4", "7.6", "7.8", "8", "8.2", "8.4", "8.6", "8.8", "9", "9.2",
	"9.4", "9.6", "9.8", "10", "10.5", "11", "11.5", "12", "12.5", "13", "13.5", "14",
	"14.5", "15", "15.5", "16", "16.5", "17", "17.5", "18", "18.5", "19", "19.5", "20",
	"21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "32", "34", "36", "38", "40",
	"42", "44", "46", "48", "50", "55", "60", "65", "70", "75", "80", "85", "90", "95", "100",
	"110", "120", "130", "140", "150", "160", "170", "180", "190", "200", "210", "220",
	"230", "240", "250", "260", "270", "280", "290", "300", "310", "320", "330", "340",
	"350", "360", "370", "380", "390", "400", "410", "420", "430", "440", "450", "460",
	"470", "480", "490", "500", "510", "520", "530", "540", "550", "560", "570", "580",
	"590", "600", "610", "620", "630", "640", "650", "660", "670", "680", "690", "700",
	"710", "720", "730", "740", "750", "760", "770", "780", "790", "800", "810", "820",
	"830", "840", "850", "860", "870", "880", "890", "900", "910", "920", "930", "940",
	"950", "960", "970", "980", "990", "1000"}

// OddsRange is a map[string]float64 with all the odd ranges available.
var OddsRange = []map[string]float64{
	{"begin": 1.01, "end": 2, "var": 0.01, "ticks": 100},
	{"begin": 2.02, "end": 3, "var": 0.02, "ticks": 50},
	{"begin": 3.05, "end": 4, "var": 0.05, "ticks": 20},
	{"begin": 4.1, "end": 6, "var": 0.1, "ticks": 20},
	{"begin": 6.2, "end": 10, "var": 0.2, "ticks": 20},
	{"begin": 10.5, "end": 20, "var": 0.5, "ticks": 20},
	{"begin": 21, "end": 30, "var": 1, "ticks": 10},
	{"begin": 32, "end": 50, "var": 2, "ticks": 10},
	{"begin": 55, "end": 100, "var": 5, "ticks": 10},
	{"begin": 110, "end": 1000, "var": 10, "ticks": 90},
}

// OddFloor returns the same odd input rounded towards 1.01.
// If the odd supplied is one of the available odds in the ladder, than the same odd is returned.
// index returns the index of the odd in the ladder.
func OddFloor(odd float64) (index int, oddRounded float64, err error) {
	_, index, err = FindOdd(odd)
	if err != nil {
		return 0, 0, err
	}

	return index, Odds[index], nil
}

// OddCeil returns the same odd input rounded towards 1000.
// If the odd supplied is one of the available odds in the ladder, than the same odd is returned.
// index returns the index of the odd in the ladder.
func OddCeil(odd float64) (index int, oddRounded float64, err error) {
	match, index, err := FindOdd(odd)
	if err != nil {
		return 0, 0, err
	}

	if match {
		return index, Odds[index], nil
	}
	return index + 1, Odds[index+1], nil
}

// OddRound returns the same odd input rounded to the nearest odd in the ladder.
// If the odd supplied is one of the available odds in the ladder, than the same odd is returned.
// index returns the index of the odd in the ladder.
func OddRound(odd float64) (index int, oddRounded float64, err error) {
	match, index, err := FindOdd(odd)
	if err != nil {
		return 0, 0, err
	}

	if match {
		return index, Odds[index], nil
	}

	// Compute deltas
	delta1 := math.Abs(odd - Odds[index])
	delta2 := math.Abs(odd - Odds[index+1])

	if delta1 <= delta2 {
		return index, Odds[index], nil
	}
	return index + 1, Odds[index+1], nil
}

// OddShift shifts the Odd up or down in the ladder.
// If shift is higher than zero, it shifts the odd towards 1000, if it's less than zero it shifts the odd towards 1.01.
// roundType is the round method to be used.
// shift represents the number of ticks to shift the odd.
func OddShift(roundType RoundType, odd float64, shift int) (index int, oddOut float64, err error) {
	switch roundType {
	case RoundType_Ceil:
		index, _, err = OddCeil(odd)
	case RoundType_Round:
		index, _, err = OddRound(odd)
	case RoundType_Floor:
		index, _, err = OddFloor(odd)
	}

	if err != nil {
		return 0, 0, err
	}

	index += shift

	if (index >= OddsCount) || (index < 0) {
		return 0, 0, fmt.Errorf("odd outside of tradable range")
	}
	return index, Odds[index], nil
}

// OddsTicksDiff computes the number of ticks between two odds.
// roundType is the round method to be used.
func OddsTicksDiff(roundType RoundType, odd1 float64, odd2 float64) (ticksDiff int, err error) {
	var index1 int
	var index2 int
	var err1 error
	var err2 error

	switch roundType {
	case RoundType_Ceil:
		index1, _, err1 = OddCeil(odd1)
		index2, _, err2 = OddCeil(odd2)
	case RoundType_Round:
		index1, _, err1 = OddRound(odd1)
		index2, _, err2 = OddRound(odd2)
	case RoundType_Floor:
		index1, _, err1 = OddFloor(odd1)
		index2, _, err2 = OddFloor(odd2)
	}

	if err1 != nil {
		return 0, err1
	}

	if err2 != nil {
		return 0, err2
	}

	return int(math.Abs(float64(index2 - index1))), nil
}

// IsOddWithinBoundaries checks if odd is within trading range
func IsOddWithinBoundaries(odd float64) bool {
	if equalWithTolerance(odd, 1000) {
		return true
	} else if odd > 1000 {
		return false
	}

	if equalWithTolerance(odd, 1.01) {
		return true
	} else if odd < 1.01 {
		return false
	}

	return true
}

// FindOdd tries to find the odd in the ladder.
// If it finds it, then index is the odd index in the ladder.
// If it doesn't find it, it will return the index in the ladder that is the closest to the odd on the left side.
func FindOdd(odd float64) (match bool, index int, err error) {
	// Boundaries
	if withinBoundary := IsOddWithinBoundaries(odd); !withinBoundary {
		return false, 0, fmt.Errorf("odd provided [%f] is outside of trading range", odd)
	}

	if equalWithTolerance(odd, 1000) {
		return true, OddsCount - 1, nil
	} else if equalWithTolerance(odd, 1.01) {
		return true, 0, nil
	}

	lo := 0
	hi := OddsCount

	for lo < hi {
		mid := (lo + hi) / 2

		if equalWithTolerance(odd, Odds[mid]) {
			return true, mid, nil
		} else if odd < Odds[mid] {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	return false, lo - 1, nil
}

// Helper function and constant to help estimate whether odd matches or not
func equalWithTolerance(a float64, b float64) bool {
	const float64EqualityThreshold = 1e-9

	return math.Abs(a-b) <= float64EqualityThreshold
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
