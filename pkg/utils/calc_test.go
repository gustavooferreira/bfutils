package utils

import (
	"fmt"
	"math"
	"testing"
)

// Helper function to compare float numbers
func floatEqual(x float64, y float64, tolerance float64) bool {
	diff := math.Abs(x - y)
	mean := math.Abs(x+y) / 2.0
	if math.IsNaN(diff / mean) {
		return true
	}
	return (diff / mean) < tolerance
}

func TestCalcFreeBetDecimal(t *testing.T) {
	paramsTC := []struct {
		inOddBack float64
		inOddLay  float64
		out       float64
	}{
		{2, 2, 0.0},
		{20, 10, 10},
		{18, 7, 11},
		{100, 250, -150},
		{50, 5, 45},
		{1.35, 2.02, -0.67},
	}

	for i, tc := range paramsTC {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {

			result := FreeBetDecimal(tc.inOddBack, tc.inOddLay)

			if !floatEqual(result, tc.out, 0.001) {
				t.Fatalf("Got '%.2f', wanted '%.2f'", result, tc.out)
			}
		})
	}
}

func TestCalcFreeBetAmount(t *testing.T) {
	paramsTC := []struct {
		inOddBack float64
		inOddLay  float64
		inStake   float64
		out       float64
	}{
		{2, 2, 4, 0.0},
		{20, 10, 10, 100},
		{18, 7, 5, 55},
		{100, 250, 2, -300},
		{50, 5, 5, 225},
		{1.35, 2.02, 2, -1.34},
	}

	for i, tc := range paramsTC {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {

			result := FreeBetAmount(tc.inOddBack, tc.inOddLay, tc.inStake)

			if !floatEqual(result, tc.out, 0.001) {
				t.Fatalf("Got '%.2f', wanted '%.2f'", result, tc.out)
			}
		})
	}
}

func TestCalcGreenBookOpenBackDecimal(t *testing.T) {
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

			result := GreenBookOpenBackDecimal(tc.inOddBack, tc.inOddLay)

			if !floatEqual(result, tc.out, 0.001) {
				t.Fatalf("Got '%.2f', wanted '%.2f'", result, tc.out)
			}
		})
	}
}

func TestCalcGreenBookOpenBackAmountByPerc(t *testing.T) {
	paramsTC := []struct {
		inOddBack float64
		inPerc    float64
		out       float64
	}{
		{20, 1, 10.0},
		{18, 0.7, 10.59},
		{100, 1.5, 40.0},
		{50, 2, 16.67},
		{10, -0.5, 20.0},
		{20, -0.8, 100.0},
	}

	for i, tc := range paramsTC {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {

			result := GreenBookOpenBackAmountByPerc(tc.inOddBack, tc.inPerc)

			if !floatEqual(result, tc.out, 0.001) {
				t.Fatalf("Got '%.2f', wanted '%.2f'", result, tc.out)
			}
		})
	}
}

func TestCalcGreenBookOpenLayDecimal(t *testing.T) {
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

			result := GreenBookOpenLayDecimal(tc.inOddLay, tc.inOddBack)

			if !floatEqual(result, tc.out, 0.001) {
				t.Fatalf("Got '%.2f', wanted '%.2f'", result, tc.out)
			}
		})
	}
}
