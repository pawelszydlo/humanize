package humanize

import (
	"sort"
	"strconv"
	"strings"
)

// SI prefixing functions.

type siPrefix struct {
	multiplier float64
	short      string
	long       string
}

var siPrefixes = []siPrefix{
	{1000000000000000000000000, "Y", "yotta"},
	{1000000000000000000000, "Z", "zetta"},
	{1000000000000000000, "E", "exa"},
	{1000000000000000, "P", "peta"},
	{1000000000000, "T", "tera"},
	{1000000000, "G", "giga"},
	{1000000, "M", "mega"},
	{1000, "k", "kilo"},
	{100, "h", "hecto"},
	{10, "da", "deca"},
	{0.1, "d", "deci"},
	{0.01, "c", "centi"},
	{0.001, "m", "milli"},
	{0.000001, "µ", "micro"},
	{0.000000001, "n", "nano"},
	{0.000000000001, "p", "pico"},
	{0.000000000000001, "f", "femto"},
	{0.000000000000000001, "a", "atto"},
	{0.000000000000000000001, "z", "zepto"},
	{0.000000000000000000000001, "y", "yocto"},
}

// PrefixFastInt like PrefixFast but accepting integer values.
func (humanizer *Humanizer) PrefixFastInt(value int64) string {
	return humanizer.PrefixFast(float64(value))
}

// PrefixFast is a convenience function for easy prefixing.
// Precision is 1 decimal place. Will not prefix values in range 0.01 - 100 and will append only the short prefix.
func (humanizer *Humanizer) PrefixFast(value float64) string {
	return humanizer.Prefix(value, 1, 1000, true)
}

// Hack to get rid of trailing zeroes (while keeping the precision if necessary)
func (humanizer *Humanizer) trimZeroes(value string) string {
	if strings.ContainsRune(value, '.') {
		value = strings.TrimRight(value, "0")
		value = strings.TrimRight(value, ".")
	}
	return value
}

// Prefix appends a SI prefix to the value and converts it accordingly.
// Arguments:
//	value - value to be converted
//  decimals - decimal precision for the converted value
//	threshold - upper bound of the range to be ignored. Lower bound is 1/threshold.
//	short - whether to use short or long prefix.
func (humanizer *Humanizer) Prefix(value float64, decimals int, threshold int64, short bool) string {
	if threshold < 10 {
		threshold = 10
	}
	// If value falls within ignored range then just format it.
	if value <= float64(threshold) && value >= 10.0/float64(threshold) {
		return humanizer.trimZeroes(strconv.FormatFloat(value, 'f', decimals, 64))
	}
	// Find most appropriate prefix.
	i := sort.Search(len(siPrefixes), func(i int) bool {
		return siPrefixes[i].multiplier < value
	})

	convertedValue := humanizer.trimZeroes(
		strconv.FormatFloat(value/siPrefixes[i].multiplier, 'f', decimals, 64))

	if short {
		return  convertedValue + siPrefixes[i].short
	} else {
		return convertedValue + " " + siPrefixes[i].long
	}
}
