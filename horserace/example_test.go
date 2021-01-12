package horserace_test

import (
	"fmt"

	"github.com/gustavooferreira/bfutils/horserace"
)

// This example parses a horse race's betfair market name and returns the distance and then
// the race classification.
func Example_a() {
	raceBetfairName := "2m3f Hcap"

	name, distance := horserace.GetClassAndDistance(raceBetfairName)
	class, err := horserace.GetClassificationFromAbbrev(name)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Race classification: %s\n", class)
	fmt.Printf("Race distance: %s\n", distance)

	// Output:
	// Race classification: Handicap
	// Race distance: 2m3f
}
