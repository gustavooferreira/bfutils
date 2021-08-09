package conversion_test

import (
	"testing"

	"github.com/gustavooferreira/bfutils/horserace/conversion"
	"github.com/stretchr/testify/assert"
)

const float64EqualityThreshold = 1e-5

func TestMileToFromMeter(t *testing.T) {
	tests := map[string]struct {
		miles  float64
		meters float64
	}{
		"convert 0 miles":   {miles: 0, meters: 0},
		"convert 0.5 miles": {miles: 0.5, meters: 804.672},
		"convert 1 miles":   {miles: 1, meters: 1609.344},
		"convert 2 miles":   {miles: 2, meters: 3218.688},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			value := conversion.MileToMeter(test.miles)
			assert.InDelta(t, test.meters, value, float64EqualityThreshold)

			value = conversion.MeterToMile(test.meters)
			assert.InDelta(t, test.miles, value, float64EqualityThreshold)
		})
	}
}

func TestFurlongToFromMeter(t *testing.T) {
	tests := map[string]struct {
		furlongs float64
		meters   float64
	}{
		"convert 0 furlongs":   {furlongs: 0, meters: 0},
		"convert 0.5 furlongs": {furlongs: 0.5, meters: 100.584},
		"convert 1 furlongs":   {furlongs: 1, meters: 201.168},
		"convert 2 furlongs":   {furlongs: 2, meters: 402.336},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			value := conversion.FurlongToMeter(test.furlongs)
			assert.InDelta(t, test.meters, value, float64EqualityThreshold)

			value = conversion.MeterToFurlong(test.meters)
			assert.InDelta(t, test.furlongs, value, float64EqualityThreshold)
		})
	}
}

func TestYardToFromMeter(t *testing.T) {
	tests := map[string]struct {
		yards  float64
		meters float64
	}{
		"convert 0 yards":   {yards: 0, meters: 0},
		"convert 0.5 yards": {yards: 0.5, meters: 0.4572},
		"convert 1 yards":   {yards: 1, meters: 0.9144},
		"convert 2 yards":   {yards: 2, meters: 1.8288},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			value := conversion.YardToMeter(test.yards)
			assert.InDelta(t, test.meters, value, float64EqualityThreshold)

			value = conversion.MeterToYard(test.meters)
			assert.InDelta(t, test.yards, value, float64EqualityThreshold)
		})
	}
}

func TestFurlongToFromMile(t *testing.T) {
	tests := map[string]struct {
		furlongs float64
		miles    float64
	}{
		"convert 0 furlongs":  {furlongs: 0, miles: 0},
		"convert 7 furlongs":  {furlongs: 7, miles: 0.875},
		"convert 8 furlongs":  {furlongs: 8, miles: 1},
		"convert 16 furlongs": {furlongs: 16, miles: 2},
		"convert 20 furlongs": {furlongs: 20, miles: 2.5},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			value := conversion.FurlongToMile(test.furlongs)
			assert.InDelta(t, test.miles, value, float64EqualityThreshold)

			value = conversion.MileToFurlong(test.miles)
			assert.InDelta(t, test.furlongs, value, float64EqualityThreshold)
		})
	}
}

func TestMileToFromYard(t *testing.T) {
	tests := map[string]struct {
		miles float64
		yards float64
	}{
		"convert 0 yards":    {yards: 0, miles: 0},
		"convert 100 yards":  {yards: 100, miles: 0.056818182},
		"convert 220 yards":  {yards: 220, miles: 0.125},
		"convert 1760 yards": {yards: 1760, miles: 1},
		"convert 3520 yards": {yards: 3520, miles: 2},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			value := conversion.MileToYard(test.miles)
			assert.InDelta(t, test.yards, value, float64EqualityThreshold)

			value = conversion.YardToMile(test.yards)
			assert.InDelta(t, test.miles, value, float64EqualityThreshold)
		})
	}
}

func TestFurlongToFromYard(t *testing.T) {
	tests := map[string]struct {
		furlongs float64
		yards    float64
	}{
		"convert 0 yards":    {yards: 0, furlongs: 0},
		"convert 100 yards":  {yards: 100, furlongs: 0.454545455},
		"convert 220 yards":  {yards: 220, furlongs: 1},
		"convert 1760 yards": {yards: 1760, furlongs: 8},
		"convert 3520 yards": {yards: 3520, furlongs: 16},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			value := conversion.FurlongToYard(test.furlongs)
			assert.InDelta(t, test.yards, value, float64EqualityThreshold)

			value = conversion.YardToFurlong(test.yards)
			assert.InDelta(t, test.furlongs, value, float64EqualityThreshold)
		})
	}
}

func TestParsingDistance(t *testing.T) {
	tests := map[string]struct {
		distance         string
		expectedMiles    uint
		expectedFurlongs uint
		expectedYards    uint
		expectedErr      bool
	}{

		"distance[1f2m] check error":   {distance: "1f2m", expectedErr: true},
		"distance[qwerty] check error": {distance: "qwerty", expectedErr: true},
		"distance[100F] check error":   {distance: "100F", expectedErr: true},

		"distance[0f] check":        {distance: "0f", expectedMiles: 0, expectedFurlongs: 0, expectedYards: 0},
		"distance[7f] check":        {distance: "7f", expectedMiles: 0, expectedFurlongs: 7, expectedYards: 0},
		"distance[16f] check":       {distance: "16f", expectedMiles: 2, expectedFurlongs: 0, expectedYards: 0},
		"distance[1m] check":        {distance: "1m", expectedMiles: 1, expectedFurlongs: 0, expectedYards: 0},
		"distance[1m1f] check":      {distance: "1m1f", expectedMiles: 1, expectedFurlongs: 1, expectedYards: 0},
		"distance[1m2f] check":      {distance: "1m2f", expectedMiles: 1, expectedFurlongs: 2, expectedYards: 0},
		"distance[2m2f] check":      {distance: "2m2f", expectedMiles: 2, expectedFurlongs: 2, expectedYards: 0},
		"distance[2m] check":        {distance: "2m", expectedMiles: 2, expectedFurlongs: 0, expectedYards: 0},
		"distance[2m100y] check":    {distance: "2m100y", expectedMiles: 2, expectedFurlongs: 0, expectedYards: 100},
		"distance[2m15f660y] check": {distance: "2m15f660y", expectedMiles: 4, expectedFurlongs: 2, expectedYards: 0},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var errBool bool
			dist, err := conversion.ParseDistance(test.distance)
			if err != nil {
				errBool = true
			}

			assert.Equal(t, test.expectedErr, errBool)
			assert.Equal(t, test.expectedMiles, dist.Miles)
			assert.Equal(t, test.expectedFurlongs, dist.Furlongs)
			assert.Equal(t, test.expectedYards, dist.Yards)
		})
	}
}

func TestDistanceToMeters(t *testing.T) {
	tests := map[string]struct {
		distance       []uint
		expectedMeters float64
	}{

		"distance[0,0,0] check": {distance: []uint{0, 0, 0}, expectedMeters: 0},
		"distance[1,0,0] check": {distance: []uint{1, 0, 0}, expectedMeters: 1609.344},
		"distance[0,1,0] check": {distance: []uint{0, 1, 0}, expectedMeters: 201.168},
		"distance[0,0,1] check": {distance: []uint{0, 0, 1}, expectedMeters: 0.9144},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			dist := conversion.NewDistance(test.distance[0], test.distance[1], test.distance[2])
			value := dist.ToMeters()
			assert.InDelta(t, test.expectedMeters, value, float64EqualityThreshold)
		})
	}
}

func TestDistanceToString(t *testing.T) {
	tests := map[string]struct {
		distance    []uint
		expectedStr string
	}{

		"distance[0,0,0] check":    {distance: []uint{0, 0, 0}, expectedStr: ""},
		"distance[1,0,0] check":    {distance: []uint{1, 0, 0}, expectedStr: "1m"},
		"distance[0,1,0] check":    {distance: []uint{0, 1, 0}, expectedStr: "1f"},
		"distance[0,0,1] check":    {distance: []uint{0, 0, 1}, expectedStr: "1y"},
		"distance[0,16,150] check": {distance: []uint{0, 16, 150}, expectedStr: "2m150y"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			dist := conversion.NewDistance(test.distance[0], test.distance[1], test.distance[2])
			value := dist.String()
			assert.Equal(t, test.expectedStr, value)
		})
	}
}
