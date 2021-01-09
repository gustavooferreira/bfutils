package horserace

// BETFAIR NOTE: In the case of Beginners Handicap Hurdle, Novices Handicap Chase etc,
// disregard Novice, Beginners etc as using the three abbreviations will not fit.
// Therefore Hcap Hrd, or Hcap Chs.

// ClassToAbbrev is a map from Classification to Abbreviation
var ClassToAbbrev = map[string]string{
	"Listed Race":               "Listed",
	"Group 1":                   "Grp1",
	"Group 2":                   "Grp2",
	"Group 3":                   "Grp3",
	"Maiden Stakes":             "Mdn Stks",
	"Selling Stakes":            "Sell Stks",
	"Claiming Stakes":           "Claim Stks",
	"Rated Stakes":              "Hcap",
	"Classified Stakes":         "Class Stks",
	"Nursery Handicap":          "Nursery",
	"Handicap":                  "Hcap",
	"Showcase Handicap":         "Hcap",
	"Maiden":                    "Mdn",
	"Stakes":                    "Stks",
	"Hurdle":                    "Hrd",
	"Chase":                     "Chs",
	"Hunters":                   "Hunt",
	"Novice":                    "Nov",
	"Selling":                   "Sell",
	"Claiming":                  "Claim",
	"Classified":                "Class",
	"Beginners":                 "Beg",
	"National Hunt Flat":        "NHF",
	"Irish NH Flat":             "INHF",
	"Flat":                      "Flat",
	"Beginners Handicap Hurdle": "Hcap Hrd",
	"Novices Handicap Chase":    "Hcap Chs",
	"(Non Of the Above)":        "Stks",
}

// AbbrevToClass is a map from Abbreviation to Classification
var AbbrevToClass = map[string][]string{
	"Listed":     {"Listed Race"},
	"Grp1":       {"Group 1"},
	"Grp2":       {"Group 2"},
	"Grp3":       {"Group 3"},
	"Mdn Stks":   {"Maiden Stakes"},
	"Sell Stks":  {"Selling Stakes"},
	"Claim Stks": {"Claiming Stakes"},
	"Hcap":       {"Rated Stakes", "Handicap", "Showcase Handicap"},
	"Class Stks": {"Classified Stakes"},
	"Nursery":    {"Nursery Handicap"},
	"Mdn":        {"Maiden"},
	"Hrd":        {"Hurdle"},
	"Chs":        {"Chase"},
	"Hunt":       {"Hunters"},
	"Nov":        {"Novice"},
	"Sell":       {"Selling"},
	"Claim":      {"Claiming"},
	"Class":      {"Classified"},
	"Beg":        {"Beginners"},
	"NHF":        {"National Hunt Flat"},
	"INHF":       {"Irish NH Flat"},
	"Flat":       {"Flat"},
	"Hcap Hrd":   {"Beginners Handicap Hurdle"},
	"Hcap Chs":   {"Novices Handicap Chase"},
	"Stks":       {"Stakes", "(Non Of the Above)"},
}
