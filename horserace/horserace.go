// Package horserace provides helper functions that facilitate operations related specifically
// with the horse racing markets.
package horserace

import (
	"fmt"
	"strings"
)

// GetClassAndDistance returns the classification of the race and the distance, given the name of the race (as returned by the betfair API).
func GetClassAndDistance(marketName string) (name string, distance string) {
	words := strings.Fields(marketName)
	if len(words) < 2 {
		return "", ""
	}

	distance = words[0]
	name = strings.Join(words[1:], " ")
	return name, distance
}

// GetTrackNameFromAbbrev returns the track name, given the the betfair track abbreviation and country.
func GetTrackNameFromAbbrev(country Country, abbrev string) (track string, err error) {
	var ok bool

	switch country {
	case Country_UK:
		track, ok = ukAbbrevToTrack[abbrev]
	case Country_IRE:
		track, ok = ireAbbrevToTrack[abbrev]
	default:
		return "", fmt.Errorf("country [%s] not supported", country)
	}

	if ok {
		return track, nil
	}
	return "", fmt.Errorf("couldn't find track with abbreviation [%s] in country [%s]", abbrev, country)
}

// GetAbbrevFromTrackName returns the betfair track abbreviation, given the country and track name.
func GetAbbrevFromTrackName(country Country, track string) (abbrev string, err error) {
	var ok bool

	switch country {
	case Country_UK:
		track, ok = ukTrackToAbbrev[track]
	case Country_IRE:
		track, ok = ireTrackToAbbrev[track]
	default:
		return "", fmt.Errorf("country [%s] not supported", country)
	}

	if ok {
		return track, nil
	}
	return "", fmt.Errorf("couldn't find track [%s] in country [%s]", abbrev, country)
}

// GetClassificationFromAbbrev returns the race classification, given the betfair classification abbreviation.
func GetClassificationFromAbbrev(abbrev string) (class string, err error) {
	if classes, ok := abbrevToClass[abbrev]; ok {
		return classes[0], nil
	}
	return "", fmt.Errorf("couldn't find match")
}

// GetAbbrevFromClassification returns the betfair classification abbreviation, given the race classification.
func GetAbbrevFromClassification(class string) (abbrev string, err error) {
	if abbrev, ok := classToAbbrev[class]; ok {
		return abbrev, nil
	}
	return "", fmt.Errorf("couldn't find match")
}

// Country represents a country.
type Country uint

const (
	// Country_UK represents UK country
	Country_UK = iota + 1
	// Country_IRE represents IRE country
	Country_IRE
)

// String returns the string representation of Country.
func (c Country) String() string {
	return [...]string{"", "UK", "IRE"}[c]
}
