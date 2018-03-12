package humanize

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	Minute   = 60
	Hour     = 60 * Minute
	Day      = 24 * Hour
	Week     = 7 * Day
	Month    = 30 * Day
	Year     = 12 * Month
	LongTime = 35 * Year
)

type timeRange struct {
	maxTime int64
	format  string
	divBy   int64
}

// TimeDiffNow is a convenience method returning humanized time from now till date.
func (humanizer *Humanizer) TimeDiffNow(date time.Time) string {
	return humanizer.TimeDiff(time.Now(), date)
}

// TimeDiff will return the humanized time difference between the given dates.
func (humanizer *Humanizer) TimeDiff(startDate, endDate time.Time) string {
	diff := endDate.Unix() - startDate.Unix()

	// Don't bother with Math.Abs
	absDiff := diff
	if absDiff < 0 {
		absDiff = -absDiff
	}

	// Find the range closest but bigger then diff.
	n := sort.Search(len(humanizer.provider.timeRanges), func(i int) bool {
		return humanizer.provider.timeRanges[i].maxTime > absDiff
	})

	timeRange := humanizer.provider.timeRanges[n]

	humanized := timeRange.format
	// If range has a placeholder for a number, insert it.
	if strings.Contains(timeRange.format, "%d") {
		humanized = fmt.Sprintf(timeRange.format, absDiff/timeRange.divBy)
	}

	// Past or future?
	if diff > 0 {
		return fmt.Sprintf(humanizer.provider.timeFuture, humanized)
	} else {
		return fmt.Sprintf(humanizer.provider.timePast, humanized)
	}
}

// GetDuration will return time duration as parsed from input string.
func (humanizer *Humanizer) GetDuration(input string) (time.Duration, error) {
	matched := humanizer.timeInputRe.FindStringSubmatch(input)
	if len(matched) != 4 {
		return time.Duration(0), errors.New(fmt.Sprintf("Cannot parse '%s'.", input))
	}
	// 0 - full match, 1 - number, 2 - decimal, 3 - unit
	if matched[2] == "" { // Decimal component is empty.
		matched[2] = "0"
	}
	// Parse first two groups into a float.
	number, err := strconv.ParseFloat(matched[1]+"."+matched[2], 64)
	if err != nil {
		return time.Duration(0), err
	}
	// Get the value of the unit in seconds.
	seconds, _ := humanizer.provider.timeUnits[matched[3]]

	return time.Duration(int64(number * float64(seconds) * float64(time.Second))), nil
}
