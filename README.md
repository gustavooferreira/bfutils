# Betfair Utils

[![Build Status](https://travis-ci.com/gustavooferreira/bfutils.svg?branch=master)](https://travis-ci.com/gustavooferreira/bfutils)
[![codecov](https://codecov.io/gh/gustavooferreira/bfutils/branch/master/graph/badge.svg)](https://codecov.io/gh/gustavooferreira/bfutils)
[![Go Report Card](https://goreportcard.com/badge/github.com/gustavooferreira/bfutils)](https://goreportcard.com/report/github.com/gustavooferreira/bfutils)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/gustavooferreira/bfutils)](https://pkg.go.dev/github.com/gustavooferreira/bfutils)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fgustavooferreira%2Fbfutils.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fgustavooferreira%2Fbfutils?ref=badge_shield)

Go code (golang) set of packages that provide many utility functions that help with day to day automation operations in the Betfair exchange.

Features include:

- [Odds operations](#bfutils-package)
- [Betting calculations](#betting-package)
- [Horse Races helper methods](#horserace-package)
- [Distance conversions](#conversion-package)

## [`bfutils`](https://pkg.go.dev/github.com/gustavooferreira/bfutils "API documentation") package

The `bfutils` package provides some helpful methods that allow you to do some calculations with odds.

- Find if a given arbitrary float belongs to the set of tradeable odds
- Round, Floor and Ceiling rounding operations when the float doesn't match one of the tradeable values allowed
- Find how many ticks away two odds are from each other
- Shift an odd by X ticks.

See it in action:

```go
package main

import (
  "fmt"
  "github.com/gustavooferreira/bfutils"
)

func main() {
    randomOdd := 4.051

    index1, odd1, err := bfutils.OddShift(bfutils.RoundType_Floor, randomOdd, 10)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Odd1: %.2f - position in the ladder: %d\n", odd1, index1+1)

    index2, odd2, err := bfutils.OddShift(bfutils.RoundType_Floor, randomOdd, -10)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Odd2: %.2f - position in the ladder: %d\n", odd2, index2+1)

    ticksDiff, err := bfutils.OddsTicksDiff(bfutils.RoundType_Floor, odd1, odd2)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Ticks difference between both odds: %d", ticksDiff)
}

```

## [`betting`](https://pkg.go.dev/github.com/gustavooferreira/bfutils/betting "API documentation") package

The `betting` package provides some helpful methods that allow you to ...

- ...

See it in action:

```go
package main

import (
  "fmt"
  "github.com/gustavooferreira/bfutils/betting"
)

func main() {
}

```

## [`horserace`](https://pkg.go.dev/github.com/gustavooferreira/bfutils/horserace "API documentation") package

The `horserace` package provides some helpful methods that allow you to ...

- ...

See it in action:

```go
package main

import (
  "fmt"
  "github.com/gustavooferreira/bfutils/horserace"
)

func main() {
}

```

## [`conversion`](https://pkg.go.dev/github.com/gustavooferreira/bfutils/conversion "API documentation") package

The `conversion` package provides some helpful methods that allow you to ...

- ...

See it in action:

```go
package main

import (
  "fmt"
  "github.com/gustavooferreira/bfutils/conversion"
)

func main() {
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
    github.com/gustavooferreira/bfutils/conversion

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

---

# Contributing

Please feel free to submit issues, fork the repository and send pull requests!

When submitting an issue, we ask that you please include a complete test function that demonstrates the issue.

---

# License

This project is licensed under the terms of the MIT license.
