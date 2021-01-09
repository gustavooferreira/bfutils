package horserace

import (
	"strings"
)

// GetClassificationAndDistance returns the classification of the race and the distance, given the name of the race (as returned by the betfair API).
func GetClassificationAndDistance(marketName string) (name string, distance string) {
	words := strings.Fields(marketName)

	// check length as well
	distance = words[0]

	name = strings.Join(words[1:], " ")

	return name, distance
}
