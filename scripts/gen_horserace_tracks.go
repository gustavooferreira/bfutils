//go:build ignore
// +build ignore

package main

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Mapping struct {
	UKTrackToAbbrev  map[string]string
	UKAbbrevToTrack  map[string]string
	IRETrackToAbbrev map[string]string
	IREAbbrevToTrack map[string]string
}

func main() {
	root := filepath.Join("..")

	ukCSV := filepath.Join(root, "assets", "data", "horserace_tracks_uk.csv")
	ireCSV := filepath.Join(root, "assets", "data", "horserace_tracks_ire.csv")

	ukTrackToAbbrev, ukAbbrevToTrack := loadMapping(ukCSV)
	ireTrackToAbbrev, ireAbbrevToTrack := loadMapping(ireCSV)

	mapping := Mapping{
		UKTrackToAbbrev:  ukTrackToAbbrev,
		UKAbbrevToTrack:  ukAbbrevToTrack,
		IRETrackToAbbrev: ireTrackToAbbrev,
		IREAbbrevToTrack: ireAbbrevToTrack,
	}

	tmplPath := filepath.Join(root, "assets", "templates", "horserace_tracks.go.tmpl")
	tmpl := template.Must(template.ParseFiles(tmplPath))

	outputPath := filepath.Join(root, "horserace", "racecourses.go")
	out, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	err = tmpl.Execute(out, mapping)
	if err != nil {
		panic(err)
	}
}

func loadMapping(csvPath string) (map[string]string, map[string]string) {
	f, err := os.Open(csvPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	// skip header
	records = records[1:]

	trackToAbbrev := make(map[string]string)
	abbrevToTrack := make(map[string]string)

	for _, row := range records {
		if len(row) < 2 {
			continue
		}
		track := strings.TrimSpace(row[0])
		abbrev := strings.TrimSpace(row[1])
		trackToAbbrev[track] = abbrev
		abbrevToTrack[abbrev] = track
	}

	return trackToAbbrev, abbrevToTrack
}
