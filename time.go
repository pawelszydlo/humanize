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
		return humanizer.provider.times.now
	}
	secondsLeft := seconds
	humanized := []string{}

	for secondsLeft > 0 {
		// Within all the ranges, find the one matching our time best (closest, but bigger).
		rangeIndex := sort.Search(len(humanizer.provider.times.ranges), func(i int) bool {
			// If we are in precise mode, and next range would be a fit but should be skipped, use this one.
			if precise && i < len(humanizer.provider.times.ranges)-1 &&
				humanizer.provider.times.ranges[i+1].upperLimit > secondsLeft &&
				humanizer.provider.times.ranges[i+1].skipWhenPrecise {
				return true
			}
			return humanizer.provider.times.ranges[i].upperLimit > secondsLeft
		})

		// Select this unit range and convert actualTime to it.
		unitRanges := humanizer.provider.times.ranges[rangeIndex]
		actualTime := secondsLeft / unitRanges.divideBy // Integer division!

		// Subtract the time span covered by this part.
		secondsLeft -= actualTime * unitRanges.divideBy
		// TODO: smarter imprecise mode.
		if !precise { // We don't care about the reminder.
			secondsLeft = 0
		}

		if actualTime == 1 { // Special case for singular unit.
			humanized = append(humanized, unitRanges.singular)
			continue
		}

		// Within the unit range, find the unit best fitting our actualTime (closest, but bigger).
		searchTime := actualTime
		if unitRanges.onlyLastDigitAfter != 0 && actualTime > unitRanges.onlyLastDigitAfter {
			searchTime = actualTime % 10
		}
		unitIndex := sort.Search(len(unitRanges.ranges), func(i int) bool {
			return unitRanges.ranges[i].upperLimit > searchTime
		})
		timeRange := unitRanges.ranges[unitIndex]

		humanized = append(humanized, fmt.Sprintf(timeRange.format, actualTime))
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
//   precise=true  -> "2 months and 10 days"
//
func (humanizer *Humanizer) TimeDiff(startDate, endDate time.Time, precise bool) string {
	diff := endDate.Unix() - startDate.Unix()

	// Don't bother with Math.Abs
	absDiff := diff
	if absDiff < 0 {
		absDiff = -absDiff
	}

	humanized := humanizer.humanizeDuration(absDiff, precise)

	// Past or future?
	if diff == 0 {
		return humanized
	} else if diff > 0 {
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
		if err != nil { // This can only fail if the regexp is wrong and allows non numbers.
			return time.Duration(0), err
		}
		// Get the value of the unit in seconds.
		seconds, _ := humanizer.provider.times.units[matched[3]]
		// Parser will simply sum up all the found durations.
		totalDuration += time.Duration(int64(number * float64(seconds) * float64(time.Second)))
	}

	return totalDuration, nil
}

// SecondsToTimeString converts the time in seconds into a human readable timestamp, eg.:
// 	 76 -> 01:16
// 	 3620 -> 1:00:20
func (humanizer *Humanizer) SecondsToTimeString(duration int64) string {
	h := duration / 3600
	duration -= h * 3600
	m := duration / 60
	duration -= m * 60
	s := duration
	if h > 0 {
		return fmt.Sprintf("%d:%02d:%02d", h, m, s)
	}
	return fmt.Sprintf("%02d:%02d", m, s)
}
