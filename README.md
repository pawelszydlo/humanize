# humanize [![GoDoc](https://godoc.org/github.com/pawelszydlo/humanize?status.svg)](https://godoc.org/github.com/pawelszydlo/humanize) [![Build Status](https://travis-ci.org/pawelszydlo/humanize.svg?branch=master)](https://travis-ci.org/pawelszydlo/humanize) [![codecov](https://codecov.io/gh/pawelszydlo/humanize/branch/master/graph/badge.svg)](https://codecov.io/gh/pawelszydlo/humanize)
Human readable formatting and input parsing for Go with l18n support.

### Features
* Based partially on golang.org/x/text/message, hence it's locale aware
* Supports varying numerals if language calls for it, e.g.:
```
2 godziny
5 godzin
```
* Easily extensible with new languages.

### Supported languages
* English
* Polish

### Supported operations

Init with:
```golang
humanizer, err := humanize.New("en")
```
#### Time

Decoding duration from human input:
```golang
duration, _ := humanizer.ParseDuration("2 days, 5 hours and 40 seconds")
fmt.Println(duration) 
// Prints: 53h0m40s
```
Humanized time difference:
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
#### Numbers

Number humanization (locale aware):
```golang
fmt.Println(humanizer.HumanizeNumber(1234.567, 2))
// Prints: 1,234.57
```
#### Metric prefixes

Decoding value from human input:
```golang
value, _ := humanizer.ParsePrefix("1.5k")
fmt.Println(value)
// Prints: 1500
```
Number conversion:
```golang
// Quick usage.
fmt.Println(humanizer.PrefixFast(174512))
// Prints: 174.5k

// Controlled usage.
fmt.Println(humanizer.Prefix(1440000, 2, 1000, false))
// Prints: 1.44 mega
```