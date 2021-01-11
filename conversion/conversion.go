// Package conversion provides helper functions to convert distance from/to various different distance units.
package conversion

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

// Conversions:
//
// 1 mile <-> 1609.344 meters
// 1 mile <-> 8 furlong
// 1 mile <-> 1760 yards
//
// 1 furlong <-> 201.168 meters
// 1 furlong <-> 220 yards
//
// 1 yard <-> 0.9144 meters

// FurlongsInMiles provides a constant for how many furlongs fit in a mile.
const FurlongsInMiles uint = 8

// YardsInFurlongs provides a constant for how many yards fit in a furlong.
const YardsInFurlongs uint = 220

// YardMeterConst represents the Yard/Meter conversion ratio.
const YardMeterConst float64 = 0.9144

// FurlongMeterConst represents the Furlong/Meter conversion ratio.
const FurlongMeterConst float64 = 201.168

// MileMeterConst represents the Mile/Meter conversion ratio.
const MileMeterConst float64 = 1609.344

// Meters conversions

// MileToMeter converts Miles to Meters.
func MileToMeter(miles float64) (meters float64) {
	return miles * MileMeterConst
}

// MeterToMile converts Meters to Miles.
func MeterToMile(meters float64) (miles float64) {
	return meters / MileMeterConst
}

// FurlongToMeter converts Furlongs (furrow long) to Meters.
func FurlongToMeter(furlong float64) (meters float64) {
	return furlong * FurlongMeterConst
}

// MeterToFurlong converts Meters to Furlongs (furrow long).
func MeterToFurlong(meters float64) (furlongs float64) {
	return meters / FurlongMeterConst
}

// YardToMeter converts Yards to Meters.
func YardToMeter(yards float64) (meters float64) {
	return yards * YardMeterConst
}

// MeterToYard converts Meters to Yards.
func MeterToYard(meters float64) (yards float64) {
	return meters / YardMeterConst
}

// ------------

// MileToFurlong converts Miles to Furlongs (furrow long).
func MileToFurlong(miles float64) (furlongs float64) {
	return miles * float64(FurlongsInMiles)
}

// FurlongToMile converts Furlongs (furrow long) to Miles.
func FurlongToMile(furlongs float64) (miles float64) {
	return furlongs / float64(FurlongsInMiles)
}

// MileToYard converts Miles to Yards.
func MileToYard(miles float64) (yards float64) {
	return miles * float64(YardsInFurlongs*FurlongsInMiles)
}

// YardToMile converts Yards to Miles.
func YardToMile(yards float64) (miles float64) {
	return yards / float64(YardsInFurlongs*FurlongsInMiles)
}

// FurlongToYard converts Furlongs (furrow long) to Yards.
func FurlongToYard(furlongs float64) (yards float64) {
	return furlongs * float64(YardsInFurlongs)
}

// YardToFurlong converts Yards to Furlongs (furrow long).
func YardToFurlong(yards float64) (furlongs float64) {
	return yards / float64(YardsInFurlongs)
}

// Horse races distance

// Distance represents a distance split in 3 different units: miles, furlongs and yards.
type Distance struct {
	Miles    uint
	Furlongs uint
	Yards    uint
}

// NewDistance returns a new Distance.
// If the number of furlongs is higher than 7 or the number of yards higher than 220, this function
// will automatically compute "overflow" the excess into the next available distance unit.
// Example:
// Input: 2 miles, 16 furlongs, 300 yards
// Output: Distance{Miles: 4, Furlongs: 1, Yards: 80}
func NewDistance(miles uint, furlongs uint, yards uint) Distance {
	return standardizeMetrics(miles, furlongs, yards)
}

var rgx = regexp.MustCompile(`^((?P<Miles>\d+)m){0,1}((?P<Furlongs>\d+)f){0,1}((?P<Yards>\d+)y){0,1}$`)

// ParseDistance parses distance in the betfair format.
// Format: XmYfZy
func ParseDistance(distStr string) (dist Distance, err error) {

	if match := rgx.MatchString(distStr); !match {
		return dist, errors.New("distance not properly formatted")
	}

	rs := rgx.FindStringSubmatch(distStr)
	miles, _ := strconv.ParseUint(rs[2], 10, 32)
	furlongs, _ := strconv.ParseUint(rs[4], 10, 32)
	yards, _ := strconv.ParseUint(rs[6], 10, 32)

	return standardizeMetrics(uint(miles), uint(furlongs), uint(yards)), nil
}

// String is the string representation of Distance in the betfair format.
// Format: XmYfZy
func (d Distance) String() string {
	var result string

	if d.Miles != 0 {
		result += fmt.Sprintf("%dm", d.Miles)
	}
	if d.Furlongs != 0 {
		result += fmt.Sprintf("%df", d.Furlongs)
	}
	if d.Yards != 0 {
		result += fmt.Sprintf("%dy", d.Yards)
	}

	return result
}

// ToMeters returns distance in meters.
func (d Distance) ToMeters() (meters float64) {
	var result float64

	result += MileToMeter(float64(d.Miles))
	result += FurlongToMeter(float64(d.Furlongs))
	result += YardToMeter(float64(d.Yards))
	return result
}

func standardizeMetrics(miles uint, furlongs uint, yards uint) Distance {
	y1 := yards / YardsInFurlongs
	y2 := yards % YardsInFurlongs

	f1 := (furlongs + y1) / FurlongsInMiles
	f2 := (furlongs + y1) % FurlongsInMiles

	return Distance{Miles: miles + f1, Furlongs: f2, Yards: y2}
}
