package horserace_test

import (
	"fmt"
	"testing"

	"github.com/gustavooferreira/bfutils/horserace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetClassificationAndDistance(t *testing.T) {
	tests := map[string]struct {
		marketName       string
		expectedDistance string
		expectedClass    string
	}{
		"race: <Empty>":                            {marketName: "<Empty>", expectedDistance: "", expectedClass: ""},
		"race: 1m3f Mdn Stks":                      {marketName: "1m3f Mdn Stks", expectedDistance: "1m3f", expectedClass: "Mdn Stks"},
		"race: 2m Hcap":                            {marketName: "2m Hcap", expectedDistance: "2m", expectedClass: "Hcap"},
		"race: 2m <with spaces around> Hcap":       {marketName: "   2m    Hcap    ", expectedDistance: "2m", expectedClass: "Hcap"},
		"race: 1m3f Mdn <with spaces around> Stks": {marketName: "1m3f     Mdn     Stks    ", expectedDistance: "1m3f", expectedClass: "Mdn Stks"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			name, dist := horserace.GetClassAndDistance(test.marketName)

			assert.Equal(t, test.expectedDistance, dist)
			assert.Equal(t, test.expectedClass, name)
		})
	}
}

func TestGetTrackNameFromAbbrev(t *testing.T) {
	tests := map[string]struct {
		country       horserace.Country
		abbrev        string
		expectedTrack string
		expectedErr   bool
	}{
		"race: no country and no abbrev": {expectedErr: true},
		"race: UK, no abbrev":            {country: horserace.Country_UK, abbrev: "NoTrack", expectedErr: true},
		"race: UK, Leicester":            {country: horserace.Country_UK, abbrev: "Leic", expectedTrack: "Leicester"},
		"race: IRE, Clonmel":             {country: horserace.Country_IRE, abbrev: "Clon", expectedTrack: "Clonmel"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var errBool bool
			var errMsg string
			track, err := horserace.GetTrackNameFromAbbrev(test.country, test.abbrev)
			if err != nil {
				errBool = true
				errMsg = fmt.Sprintf(" - err: %s", err.Error())
			}

			require.Equal(t, test.expectedErr, errBool, "error field"+errMsg)

			assert.Equal(t, test.expectedTrack, track)
		})
	}
}

func TestGetAbbrevFromTrackName(t *testing.T) {
	tests := map[string]struct {
		country        horserace.Country
		track          string
		expectedAbbrev string
		expectedErr    bool
	}{
		"race: no country and no track": {expectedErr: true},
		"race: UK, no track":            {country: horserace.Country_UK, track: "NoTrack", expectedErr: true},
		"race: UK, Leicester":           {country: horserace.Country_UK, track: "Leicester", expectedAbbrev: "Leic"},
		"race: IRE, Clonmel":            {country: horserace.Country_IRE, track: "Clonmel", expectedAbbrev: "Clon"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var errBool bool
			var errMsg string
			track, err := horserace.GetAbbrevFromTrackName(test.country, test.track)
			if err != nil {
				errBool = true
				errMsg = fmt.Sprintf(" - err: %s", err.Error())
			}

			require.Equal(t, test.expectedErr, errBool, "error field"+errMsg)

			assert.Equal(t, test.expectedAbbrev, track)
		})
	}
}

func TestGetClassificationFromAbbrev(t *testing.T) {
	tests := map[string]struct {
		abbrev        string
		expectedClass string
		expectedErr   bool
	}{
		"race: no abbrev": {expectedErr: true},
		"race: Grp1":      {abbrev: "Grp1", expectedClass: "Group 1"},
		"race: Hcap":      {abbrev: "Hcap", expectedClass: "Handicap"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var errBool bool
			var errMsg string
			class, err := horserace.GetClassificationFromAbbrev(test.abbrev)
			if err != nil {
				errBool = true
				errMsg = fmt.Sprintf(" - err: %s", err.Error())
			}

			require.Equal(t, test.expectedErr, errBool, "error field"+errMsg)

			assert.Equal(t, test.expectedClass, class)
		})
	}
}

func TestGetAbbrevFromClassification(t *testing.T) {
	tests := map[string]struct {
		class          string
		expectedAbbrev string
		expectedErr    bool
	}{
		"race: no class": {expectedErr: true},
		"race: Group 1":  {class: "Group 1", expectedAbbrev: "Grp1"},
		"race: Handicap": {class: "Handicap", expectedAbbrev: "Hcap"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var errBool bool
			var errMsg string
			abbrev, err := horserace.GetAbbrevFromClassification(test.class)
			if err != nil {
				errBool = true
				errMsg = fmt.Sprintf(" - err: %s", err.Error())
			}

			require.Equal(t, test.expectedErr, errBool, "error field"+errMsg)

			assert.Equal(t, test.expectedAbbrev, abbrev)
		})
	}
}
