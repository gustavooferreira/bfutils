package horserace

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetClassificationAndDistance(t *testing.T) {
	tests := map[string]struct {
		marketName       string
		expectedDistance string
		expectedClass    string
	}{
		"race: 1m3f Mdn Stks":                      {marketName: "1m3f Mdn Stks", expectedDistance: "1m3f", expectedClass: "Mdn Stks"},
		"race: 2m Hcap":                            {marketName: "2m Hcap", expectedDistance: "2m", expectedClass: "Hcap"},
		"race: 2m <with spaces around> Hcap":       {marketName: "   2m    Hcap    ", expectedDistance: "2m", expectedClass: "Hcap"},
		"race: 1m3f Mdn <with spaces around> Stks": {marketName: "1m3f     Mdn     Stks    ", expectedDistance: "1m3f", expectedClass: "Mdn Stks"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			name, dist := GetClassificationAndDistance(test.marketName)

			assert.Equal(t, test.expectedDistance, dist)
			assert.Equal(t, test.expectedClass, name)
		})
	}
}
