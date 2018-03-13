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
    
    duration, _ := humanizer.GetDuration("2 days, 5 hours and 40 seconds")
    fmt.Println(duration)
    // Prints: 53h0m40s

    firstDate := time.Date(2017, 3, 21, 0, 0, 0, 0, time.UTC)
    secondDate := time.Date(2017, 6, 21, 0, 0, 0, 0, time.UTC)

    fmt.Println(humanizer.TimeDiff(firstDate, secondDate))
    // Prints: in 3 months

    fmt.Println(humanizer.TimeDiff(secondDate, firstDate))
    // Prints: 3 months ago
}
```