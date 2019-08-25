package utils

import (
	"fmt"
	"math"
	"testing"
)

func floatEqual(x float64, y float64, tolerance float64) bool {
	diff := math.Abs(x - y)
	mean := math.Abs(x+y) / 2.0
	if math.IsNaN(diff / mean) {
		return false
	}
	return (diff / mean) < tolerance
}

func TestCalcPercOpenBack(t *testing.T) {
	paramsTC := []struct {
		inOddBack float64
		inOddLay  float64
		out       float64
	}{
		{20, 10, 1},
		{18, 7, 1.57},
		{100, 250, -0.6},
		{50, 5, 9.0},
	}

	for i, tc := range paramsTC {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {

			result := PercPLOpenBack(tc.inOddBack, tc.inOddLay)

			if !floatEqual(result, tc.out, 0.001) {
				t.Fatalf("Got '%.2f', wanted '%.2f'", result, tc.out)
			}
		})
	}
}

func TestCalcPercOpenLay(t *testing.T) {
	paramsTC := []struct {
		inOddLay  float64
		inOddBack float64
		out       float64
	}{
		{20, 10, -1},
		{18, 7, -1.57},
		{100, 250, 0.6},
		{50, 5, -9.0},
	}

	for i, tc := range paramsTC {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {

			result := PercPLOpenLay(tc.inOddLay, tc.inOddBack)

			if !floatEqual(result, tc.out, 0.001) {
				t.Fatalf("Got '%.2f', wanted '%.2f'", result, tc.out)
			}
		})
	}
}
