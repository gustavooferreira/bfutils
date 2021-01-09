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

// FurlongsInMiles msg
const FurlongsInMiles uint = 8

// YardsInFurlongs msg
const YardsInFurlongs uint = 220

// YardMeterConst msg
const YardMeterConst float64 = 0.9144

// FurlongMeterConst msg
const FurlongMeterConst float64 = 201.168

// MileMeterConst msg
const MileMeterConst float64 = 1609.344

// Meters conversions

// MileToMeter converts Miles to Meters.
func MileToMeter(miles float64) float64 {
	return miles * MileMeterConst
}

// MeterToMile converts Meters to Miles.
func MeterToMile(meters float64) float64 {
	return meters / MileMeterConst
}

// FurlongToMeter converts Furlongs (furrow long) to Meters.
func FurlongToMeter(furlong float64) float64 {
	return furlong * FurlongMeterConst
}

// MeterToFurlong converts Meters to Furlongs (furrow long).
func MeterToFurlong(meters float64) float64 {
	return meters / FurlongMeterConst
}

// YardToMeter converts Yards to Meters.
func YardToMeter(yards float64) float64 {
	return yards * YardMeterConst
}

// MeterToYard converts Meters to Yards.
func MeterToYard(meters float64) float64 {
	return meters / YardMeterConst
}

// ------------

// MileToFurlong converts Miles to Furlongs (furrow long).
func MileToFurlong(miles float64) float64 {
	return miles * float64(FurlongsInMiles)
}

// FurlongToMile converts Furlongs (furrow long) to Miles.
func FurlongToMile(furlongs float64) float64 {
	return furlongs / float64(FurlongsInMiles)
}

// MileToYard converts Miles to Yards.
func MileToYard(miles float64) float64 {
	return miles * float64(YardsInFurlongs*FurlongsInMiles)
}

// YardToMile converts Yards to Miles.
func YardToMile(yards float64) float64 {
	return yards / float64(YardsInFurlongs*FurlongsInMiles)
}

// FurlongToYard converts Furlongs (furrow long) to Yards.
func FurlongToYard(furlongs float64) float64 {
	return furlongs * float64(YardsInFurlongs)
}

// YardToFurlong converts Yards to Furlongs (furrow long).
func YardToFurlong(yards float64) float64 {
	return yards / float64(YardsInFurlongs)
}

// Horse races distance

type Distance struct {
	Miles    uint
	Furlongs uint
	Yards    uint
}

func NewDistance(miles uint, furlongs uint, yards uint) Distance {
	return standardizeMetrics(miles, furlongs, yards)
}

var rgx = regexp.MustCompile(`^((?P<Miles>\d+)m){0,1}((?P<Furlongs>\d+)f){0,1}((?P<Yards>\d+)y){0,1}$`)

// ParseDistance parses distance in the betfair format
// string representation XmYfZy
func ParseDistance(dist string) (Distance, error) {

	if match := rgx.MatchString(dist); !match {
		return Distance{}, errors.New("distance not properly formatted")
	}

	rs := rgx.FindStringSubmatch(dist)
	miles, _ := strconv.ParseUint(rs[2], 10, 32)
	furlongs, _ := strconv.ParseUint(rs[4], 10, 32)
	yards, _ := strconv.ParseUint(rs[6], 10, 32)

	return standardizeMetrics(uint(miles), uint(furlongs), uint(yards)), nil
}

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

func (d Distance) ToMeters() float64 {
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
