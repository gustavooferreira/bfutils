package conversion_test

import (
	"fmt"

	"github.com/gustavooferreira/bfutils/horserace/conversion"
)

// This example parses a horse race's distance and then prints a human friendly
// message showing the distance. It also converts the distance into meters.
func Example_a() {
	raceDistance := "2m5f100y"

	d, err := conversion.ParseDistance(raceDistance)
	if err != nil {
		panic(err)
	}

	fmt.Printf("This race distance [%s] has %d miles, %d furlongs and %d yards.\n",
		d, d.Miles, d.Furlongs, d.Yards)
	fmt.Printf("This race distance in meters is: %.2f\n", d.ToMeters())

	// Output:
	// This race distance [2m5f100y] has 2 miles, 5 furlongs and 100 yards.
	// This race distance in meters is: 4315.97
}
