# Betfair Utils

[![Build Status](https://travis-ci.com/gustavooferreira/bfutils.svg?branch=master)](https://travis-ci.com/gustavooferreira/bfutils)
[![codecov](https://codecov.io/gh/gustavooferreira/bfutils/branch/master/graph/badge.svg)](https://codecov.io/gh/gustavooferreira/bfutils)
[![Go Report Card](https://goreportcard.com/badge/github.com/gustavooferreira/bfutils)](https://goreportcard.com/report/github.com/gustavooferreira/bfutils)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fgustavooferreira%2Fbfutils.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fgustavooferreira%2Fbfutils?ref=badge_shield)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/gustavooferreira/bfutils)](https://pkg.go.dev/github.com/gustavooferreira/bfutils)

Go code (golang) set of packages that provide many utility functions that help with day to day automation operations in the Betfair exchange.

Features include:

- [Betfair Utils](#betfair-utils)
	- [`bfutils` package](#bfutils-package)
	- [`betting` package](#betting-package)
	- [`horserace` package](#horserace-package)
	- [`conversion` package](#conversion-package)
- [Installation](#installation)
- [Staying up to date](#staying-up-to-date)
- [Tests](#tests)
- [Contributing](#contributing)
- [License](#license)

## [`bfutils`](https://pkg.go.dev/github.com/gustavooferreira/bfutils "API documentation") package

The `bfutils` package provides some helpful methods that allow you to do some calculations with odds.

- Find if a given arbitrary float belongs to the set of tradeable odds
- Round, Floor and Ceiling rounding operations when the float doesn't match one of the tradeable values allowed
- Find how many ticks away two odds are from each other
- Shift an odd by X ticks

See it in action:

```go
package main

import (
    "fmt"
    "github.com/gustavooferreira/bfutils"
)

func main() {
	randomOdd := decimal.RequireFromString("4.051")

	index1, odd1, err := bfutils.OddShift(bfutils.RoundType_Floor, randomOdd, 10)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Odd1: %s - position in the ladder: %d\n", odd1.StringFixed(2), index1+1)

	index2, odd2, err := bfutils.OddShift(bfutils.RoundType_Floor, randomOdd, -10)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Odd2: %s - position in the ladder: %d\n", odd2.StringFixed(2), index2+1)

	ticksDiff, err := bfutils.OddsTicksDiff(bfutils.RoundType_Floor, odd1, odd2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Ticks difference between both odds: %d\n", ticksDiff)
}
```

## [`betting`](https://pkg.go.dev/github.com/gustavooferreira/bfutils/betting "API documentation") package

The `betting` package provides some helpful methods that allow you to calculate Free bets, GreenBooks
and various other related operations like computing green books at all odds, to be displayed on a
ladder.

- Compute free bets
- Compute green books
- Compute P&L on all odds in the ladder

See it in action:

```go
package main

import (
    "fmt"
    "github.com/gustavooferreira/bfutils/betting"
)

func main() {
	selection := betting.Selection{
		Bets: []betting.Bet{
			{Type: betting.BetType_Back, Odd: 4, Amount: 5},
			{Type: betting.BetType_Lay, Odd: 3, Amount: 5},
			{Type: betting.BetType_Back, Odd: 3.5, Amount: 10},
			{Type: betting.BetType_Lay, Odd: 3.2, Amount: 10},
		},
		CurrentBackOdd: 2.4,
		CurrentLayOdd:  2.42,
	}

	bet, err := betting.GreenBookSelection(selection)
	if err != nil {
		panic(err)
	}

	fmt.Printf("In order to green book this selection, put a {%s} bet at {%.2f} for £%.2f.\n",
		bet.Type, bet.Odd, bet.Amount)

	fmt.Printf("P&L\n")
	fmt.Printf("---\n")
	fmt.Printf("If this selection wins:  £%.2f\n", bet.WinPL)
	fmt.Printf("If this selection loses: £%.2f\n", bet.LosePL)
}
```

## [`horserace`](https://pkg.go.dev/github.com/gustavooferreira/bfutils/horserace "API documentation") package

The `horserace` package provides helper functions that facilitate operations specifically with the horse racing markets.

- Get race classification and distance from betfair market name
- Get race track name and classification from betfair abbreviations and vice-versa

See it in action:

```go
package main

import (
    "fmt"
    "github.com/gustavooferreira/bfutils/horserace"
)

func main() {
	raceBetfairName := "2m3f Hcap"

	name, distance := horserace.GetClassAndDistance(raceBetfairName)
	class, err := horserace.GetClassificationFromAbbrev(name)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Race classification: %s\n", class)
	fmt.Printf("Race distance: %s\n", distance)
}
```

## [`conversion`](https://pkg.go.dev/github.com/gustavooferreira/bfutils/conversion "API documentation") package

The `conversion` package provides helper functions to convert distance from/to various different distance units.

- Convert from/to meters
- Parse betfair race distances

See it in action:

```go
package main

import (
    "fmt"
    "github.com/gustavooferreira/bfutils/horserace/conversion"
)

func main() {
	raceDistance := "2m5f100y"

	d, err := conversion.ParseDistance(raceDistance)
	if err != nil {
		panic(err)
	}

	fmt.Printf("This race distance [%s] has %d miles, %d furlongs and %d yards.\n",
		d, d.Miles, d.Furlongs, d.Yards)
	fmt.Printf("This race distance in meters is: %.2f\n", d.ToMeters())
}
```

---

# Installation

To install bfutils, use `go get`:

    go get github.com/gustavooferreira/bfutils

This will then make the following packages available to you:

    github.com/gustavooferreira/bfutils
    github.com/gustavooferreira/bfutils/betting
    github.com/gustavooferreira/bfutils/horserace
    github.com/gustavooferreira/bfutils/horserace/conversion

Import the `bfutils/betting` package into your code using this template:

```go
package main

import (
    "fmt"
    "github.com/gustavooferreira/bfutils/betting"
)

func main() {
}
```

---

# Staying up to date

To update bfutils to the latest version, use `go get -u github.com/gustavooferreira/bfutils`.

---

# Tests

To run tests:

```
make test
```

To get coverage:

```
make coverage
```

---

# Contributing

Please feel free to submit issues, fork the repository and send pull requests!

When submitting an issue, we ask that you please include a complete test function that demonstrates the issue.

---

# License

This project is licensed under the terms of the MIT license.
