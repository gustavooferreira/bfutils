//go:build ignore
// +build ignore

package main

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
)

type Mapping struct {
	ClassToAbbrev   map[string]string
	AbbrevToClasses map[string][]string
}

func main() {
	root := filepath.Join("..")

	csvFile := filepath.Join(root, "assets", "data", "horserace_classifications.csv")
	f, err := os.Open(csvFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	// Skip header row
	records = records[1:]

	classToAbbrev := make(map[string]string)
	abbrevToClasses := make(map[string][]string)

	for _, row := range records {
		if len(row) != 2 {
			continue
		}
		class := strings.TrimSpace(row[0])
		abbrev := strings.TrimSpace(row[1])
		classToAbbrev[class] = abbrev
		abbrevToClasses[abbrev] = append(abbrevToClasses[abbrev], class)
	}

	// Sort the slices to make output deterministic
	for _, list := range abbrevToClasses {
		sort.Strings(list)
	}

	mapping := Mapping{
		ClassToAbbrev:   classToAbbrev,
		AbbrevToClasses: abbrevToClasses,
	}

	// Load the template
	tmplPath := filepath.Join(root, "assets", "templates", "horserace_classifications.go.tmpl")
	tmpl := template.Must(template.ParseFiles(tmplPath))

	// Create the output file
	outputPath := filepath.Join(root, "horserace", "classifications.go")
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
