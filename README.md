# humanize [![GoDoc](https://godoc.org/github.com/pawelszydlo/humanize?status.svg)](https://godoc.org/github.com/pawelszydlo/humanize)
Human readable formatting and input parsing for Go. 

#### Supported operations
* [Time humanization](#humanized-time-difference) and [parsing](#decoding-duration-from-human-input)
* [Metric prefixes](#metric-prefixes)

#### Supported languages
* English
* Polish

#### Example usage

Init with:
```golang
humanizer, err := humanize.New("en")
```
##### Decoding duration from human input:
```golang
duration, _ := humanizer.ParseDuration("2 days, 5 hours and 40 seconds")
fmt.Println(duration) 
// Prints: 53h0m40s
```
##### Humanized time difference:
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
##### Metric prefixes:
```golang
// Quick usage.
fmt.Println(humanizer.PrefixFast(174512))
// Prints: 174.5k

// Controlled usage.
fmt.Println(humanizer.Prefix(1440000, 2, 1000, true))
// Prints: 1.44M
```