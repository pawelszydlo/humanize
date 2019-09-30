# humanize [![GoDoc](https://godoc.org/github.com/pawelszydlo/humanize?status.svg)](https://godoc.org/github.com/pawelszydlo/humanize) [![Build Status](https://travis-ci.org/pawelszydlo/humanize.svg?branch=master)](https://travis-ci.org/pawelszydlo/humanize) [![codecov](https://codecov.io/gh/pawelszydlo/humanize/branch/master/graph/badge.svg)](https://codecov.io/gh/pawelszydlo/humanize)
Human readable formatting and input parsing for Go, with i18n support.
Easily extendable with new languages.

### Supported languages
* English
* Polish

----

## Features
### Time related operations

Decoding duration from human input:
```golang
duration, _ := humanizer.ParseDuration("2 days, 5 hours and 40 seconds")
fmt.Println(duration) 
// Prints: 53h0m40s
```
Humanized time difference:*
```golang
firstDate := time.Date(2017, 3, 21, 12, 30, 15, 0, time.UTC)
secondDate := time.Date(2017, 6, 21, 0, 0, 0, 0, time.UTC)

// Approximate mode.
fmt.Println(humanizer.TimeDiff(firstDate, secondDate, false))
// Prints: in 3 months

// Precise mode.
fmt.Println(humanizer.TimeDiff(secondDate, firstDate, true))
// Prints: 3 months, 1 day, 11 hours, 29 minutes and 45 seconds ago
```
Pretty print timestamps (seconds only):
```golang
fmt.Println(humanizer.SecondsToTimeString(67))
// Prints: 01:07
```
### Number related operations

Number humanization (locale aware):
```golang
fmt.Println(humanizer.HumanizeNumber(1234.567, 2))
// Prints: 1,234.57
```
### Prefixes (metric and bit)

Decoding value from human input:
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
Convert big number into something readable:
```golang
// Quick usage.
fmt.Println(humanizer.SiPrefixFast(174512))
// Prints: 174.5k

// Controlled usage.
fmt.Println(humanizer.SiPrefix(1440000, 2, 1000, false))
// Prints: 1.44 mega
```
Same with bit prefixes:
```golang
// Quick usage.
fmt.Println(humanizer.BitPrefixFast(1509949))
// Prints: 1.44Mi
```
----

## TODO
* Float precision issues when parsing extreme prefixes.
* Smarter imprecise mode for time durations.
* More features?