# humanize [![GoDoc](https://godoc.org/github.com/pawelszydlo/humanize?status.svg)](https://godoc.org/github.com/pawelszydlo/humanize)
Human readable formatting and input parsing for Go. 

Supports only a simple _\<duration> \<unit>_ format.

### Supported values
* Time

### Supported languages
* English
* Polish

### Example usage

```golang
package main
  
import (
    "fmt"
    "github.com/pawelszydlo/humanize"
    "time"
)

func main() {
    humanizer, _ := humanize.New("en")

    firstDate := time.Date(2017, 3, 21, 0, 0, 0, 0, time.UTC)
    secondDate := time.Date(2017, 6, 21, 0, 0, 0, 0, time.UTC)

    fmt.Println(humanizer.TimeDiff(firstDate, secondDate))
    // Prints: in 3 months

    fmt.Println(humanizer.TimeDiff(secondDate, firstDate))
    // Prints: 3 months ago

    duration, _ := humanizer.GetDuration("3.5 days")
    fmt.Println(duration)
    // Prints: 84h0m0s 
}
```