package humanize

// Time values humanization functions.

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Time constants.
const (
	Second   = 1
	Minute   = 60
	Hour     = 60 * Minute
	Day      = 24 * Hour
	Week     = 7 * Day
	Month    = 30 * Day
	Year     = 12 * Month
	LongTime = 35 * Year
)

// buildTimeInputRe will build a regular expression to match all possible time inputs.
func (humanizer *Humanizer) buildTimeInputRe() {
	// Get all possible time units.
	units := make([]string, 0, len(humanizer.provider.times.units))
	for unit := range humanizer.provider.times.units {
		units = append(units, unit)
	}
	// Regexp will match: number, optional coma or dot, optional second number, unit name
	humanizer.timeInputRe = regexp.MustCompile("([0-9]+)[.,]?([0-9]*?) (" + strings.Join(units, "|") + ")")
}

// humanizeDuration will return a humanized form of time duration.
func (humanizer *Humanizer) humanizeDuration(seconds int64, precise bool) string {
	if seconds < 1 {
		panic("Cannot humanize durations < 1 sec.")
	}
	secondsLeft := seconds
	humanized := []string{}

	for secondsLeft > 0 {
		// Find the ranges closest but bigger then diff.
		n := sort.Search(len(humanizer.provider.times.ranges), func(i int) bool {
			return humanizer.provider.times.ranges[i].upperLimit > secondsLeft
		})

		// Within the ranges find the one matching our time best.
		timeRanges := humanizer.provider.times.ranges[n]
		k := sort.Search(len(timeRanges.ranges), func(i int) bool {
			return timeRanges.ranges[i].upperLimit > secondsLeft
		})
		timeRange := timeRanges.ranges[k]
		actualTime := secondsLeft / timeRanges.divideBy // Integer division!

		// If range has a placeholder for a number, insert it.
		if strings.Contains(timeRange.format, "%d") {
			humanized = append(humanized, fmt.Sprintf(timeRange.format, actualTime))
		} else {
			humanized = append(humanized, timeRange.format)
		}

		// Subtract the time span covered by this part.
		secondsLeft -= actualTime * timeRanges.divideBy
		if !precise { // We don't care about the reminder.
			secondsLeft = 0
		}
	}

	if len(humanized) == 1 {
		return humanized[0]
	} else {
		return fmt.Sprintf(
			"%s %s %s",
			strings.Join(humanized[:len(humanized)-1], ", "),
			humanizer.provider.times.remainderSep,
			humanized[len(humanized)-1],
		)
	}
}

// TimeDiffNow is a convenience method returning humanized time from now till date.
func (humanizer *Humanizer) TimeDiffNow(date time.Time, precise bool) string {
	return humanizer.TimeDiff(time.Now(), date, precise)
}

// TimeDiff will return the humanized time difference between the given dates.
// Precise setting determines whether a rough approximation or exact description should be returned, e.g.:
//   precise=false -> "3 months"
//   precise=true  -> "2 months, 1 week and 3 days"
//
// TODO: in precise mode some ranges should be skipped, like weeks in the example above.
func (humanizer *Humanizer) TimeDiff(startDate, endDate time.Time, precise bool) string {
	diff := endDate.Unix() - startDate.Unix()

	if diff == 0 {
		return humanizer.provider.times.now
	}

	// Don't bother with Math.Abs
	absDiff := diff
	if absDiff < 0 {
		absDiff = -absDiff
	}

	humanized := humanizer.humanizeDuration(absDiff, precise)

	// Past or future?
	if diff > 0 {
		return fmt.Sprintf(humanizer.provider.times.future, humanized)
	} else {
		return fmt.Sprintf(humanizer.provider.times.past, humanized)
	}
}

// ParseDuration will return time duration as parsed from input string.
func (humanizer *Humanizer) ParseDuration(input string) (time.Duration, error) {
	allMatched := humanizer.timeInputRe.FindAllStringSubmatch(input, -1)
	if len(allMatched) == 0 {
		return time.Duration(0), errors.New(fmt.Sprintf("Cannot parse '%s'.", input))
	}

	totalDuration := time.Duration(0)
	for _, matched := range allMatched {
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
		seconds, _ := humanizer.provider.times.units[matched[3]]
		// Parser will simply sum up all the found durations.
		totalDuration += time.Duration(int64(number * float64(seconds) * float64(time.Second)))
	}

	return totalDuration, nil
}
