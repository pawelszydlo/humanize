# humanize [![GoDoc](https://godoc.org/github.com/pawelszydlo/humanize?status.svg)](https://godoc.org/github.com/pawelszydlo/humanize)
Human readable formatting and input parsing for Go. 

#### Supported values
* Time

#### Supported languages
* English
* Polish

#### Example usage

```golang
package main

import (
	"fmt"
	"github.com/pawelszydlo/humanize"
	"time"
)

func main() {
	humanizer, _ := humanize.New("en")

	// Decode human duration input.
	duration, _ := humanizer.ParseDuration("2 days, 5 hours and 40 seconds")
	fmt.Println(duration) 
	// Prints: 53h0m40s

	firstDate := time.Date(2017, 3, 21, 12, 30, 15, 0, time.UTC)
	secondDate := time.Date(2017, 6, 21, 0, 0, 0, 0, time.UTC)

	// Approximate mode.
	fmt.Println(humanizer.TimeDiff(firstDate, secondDate, false))
    // Prints: in 3 months

	// Precise mode.
	fmt.Println(humanizer.TimeDiff(secondDate, firstDate, true))
    // Prints: 3 months, 1 day, 11 hours, 29 minutes and 45 seconds ago
}
```