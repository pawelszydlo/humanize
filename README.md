# humanize [![GoDoc](https://godoc.org/github.com/pawelszydlo/humanize?status.svg)](https://godoc.org/github.com/pawelszydlo/humanize) [![Go Report Card](https://goreportcard.com/badge/github.com/pawelszydlo/humanize)](https://goreportcard.com/report/github.com/pawelszydlo/humanize) [![Build Status](https://travis-ci.org/pawelszydlo/humanize.svg?branch=master)](https://travis-ci.org/pawelszydlo/humanize) [![codecov](https://codecov.io/gh/pawelszydlo/humanize/branch/master/graph/badge.svg)](https://codecov.io/gh/pawelszydlo/humanize)
Human readable formatting and input parsing for Go, with i18n support.
Easily extendable with new languages.

### Supported languages
* English
* Polish

### Table of contents

 - [Features](#features)
    - [Decode duration from human input](#decode-duration-from-human-input)
    - [Humanize date difference](#humanize-date-difference)
    - [Pretty print timestamps](#pretty-print-timestamps)
    - [Add decimal separators to numbers](#add-decimal-separators-to-numbers)
    - [Decode value from human input with a prefix](#decode-value-from-human-input-with-a-prefix)
    - [Humanize big numbers with prefixes](#humanize-big-numbers-with-prefixes)
    - [Humanize parts of one](#humanize-parts-of-one)
  - [TODO](#todo)

----

## Features

### Decode duration from human input
```golang
duration, _ := humanizer.ParseDuration("2 days, 5 hours and 40 seconds")
fmt.Println(duration) 
// Prints: 53h0m40s
```
### Humanize date difference
```golang
firstDate := time.Date(2017, 3, 21, 12, 30, 15, 0, time.UTC)
secondDate := time.Date(2017, 6, 21, 0, 0, 0, 0, time.UTC)
```
Approximate mode:
```golang
fmt.Println(humanizer.TimeDiff(firstDate, secondDate, false))
// Prints: in 3 months
```
Precise mode:
```golang
fmt.Println(humanizer.TimeDiff(secondDate, firstDate, true))
// Prints: 3 months, 1 day, 11 hours, 29 minutes and 45 seconds ago
```
### Pretty print timestamps
```golang
fmt.Println(humanizer.SecondsToTimeString(67))
// Prints: 01:07
```

### Add decimal separators to numbers
Uses x/text/number and is locale aware.
```golang
fmt.Println(humanizer.HumanizeNumber(1234.567, 2))
// Prints: 1,234.57
```

### Decode value from human input with a prefix
```golang
value, _ := humanizer.ParsePrefix("1.5k")
fmt.Println(value)
// Prints: 1500
```
Bit prefixes are recognized as well:
```golang
value, _ := humanizer.ParsePrefix("1.5Ki")
fmt.Println(value)
// Prints: 1536
```
NOTE: ParsePrefix will return a precise value (big.Float), so you might get fractions
where you wouldn't expect them (e.g. bytes). It's up to you to handle that.

### Humanize big numbers with prefixes
Quick usage:
```golang
fmt.Println(humanizer.SiPrefixFast(174512))
// Prints: 174.5k
```
Controlled usage:
```golang
fmt.Println(humanizer.SiPrefix(1440000, 2, 1000, false))
// Prints: 1.44 mega
```
Using bit prefixes instead of metric:
```golang
fmt.Println(humanizer.BitPrefixFast(1509949))
// Prints: 1.44Mi
```

### Humanize parts of one
Avoid leading zeroes:
```golang
fmt.Println(humanizer.HumanizeParts(0.25, 0))
// Prints: 25%
fmt.Println(humanizer.HumanizeParts(0.0023, 0))
// Prints: 2‰
```
Allow leading zeroes to avoid jumping to a smaller unit:
```golang
fmt.Println(humanizer.HumanizeParts(0.0023, 1))
// Prints: 0.23%
```

----

## TODO
* Smarter imprecise mode for time durations.
* More features?
