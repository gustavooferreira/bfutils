package horserace

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetNameAndDistance(t *testing.T) {

	marketname := "1m3f Mdn Stks"
	name, dist := GetNameAndDistance(marketname)

	if name != "Mdn Stks" || dist != "1m3f" {
		t.Errorf("input: %q - output: %q, %q", marketname, name, dist)
	}
}

func TestDistanceFromFurlongs(t *testing.T) {
	horseRaceDistanceFromFurlongsTests := []struct {
		in  uint
		out string
	}{
		{18, "2m2f"},
		{10, "1m2f"},
		{9, "1m1f"},
		{8, "1m"},
		{7, "7f"},
		{0, "0f"},
	}

	for i, tc := range horseRaceDistanceFromFurlongsTests {
		t.Run(tc.out, func(t *testing.T) {

			result := DistanceFromFurlongs(tc.in)

			diff := cmp.Diff(tc.out, result)
			if diff != "" {
				t.Errorf("test %d: input %d, got %q, want %q\n%s", i, tc.in, result, tc.out, diff)
			}
		})
	}
}
