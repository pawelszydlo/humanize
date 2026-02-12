package humanize

import (
	"fmt"
	"math"
	"math/big"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// Prefixing functions.

// Single prefix definition.
type prefixDef struct {
	value       *big.Float
	approxValue float64 // For faster comparisons. Is it needed though?
	short       string
	long        string
}

var siPrefixes = []prefixDef{
	{bigPow(10, 24), math.Pow10(24), "Y", "yotta"},
	{bigPow(10, 21), math.Pow10(21), "Z", "zetta"},
	{bigPow(10, 18), math.Pow10(18), "E", "exa"},
	{bigPow(10, 15), math.Pow10(15), "P", "peta"},
	{bigPow(10, 12), math.Pow10(12), "T", "tera"},
	{bigPow(10, 9), math.Pow10(9), "G", "giga"},
	{bigPow(10, 6), math.Pow10(6), "M", "mega"},
	{bigPow(10, 3), math.Pow10(3), "k", "kilo"},
	{bigPow(10, 2), math.Pow10(2), "h", "hecto"},
	{bigPow(10, 1), 10, "da", "deca"},
	{bigPow(10, -1), math.Pow10(-1), "d", "deci"},
	{bigPow(10, -2), math.Pow10(-2), "c", "centi"},
	{bigPow(10, -3), math.Pow10(-3), "m", "milli"},
	{bigPow(10, -6), math.Pow10(-6), "µ", "micro"},
	{bigPow(10, -9), math.Pow10(-9), "n", "nano"},
	{bigPow(10, -12), math.Pow10(-12), "p", "pico"},
	{bigPow(10, -15), math.Pow10(-15), "f", "femto"},
	{bigPow(10, -18), math.Pow10(-18), "a", "atto"},
	{bigPow(10, -21), math.Pow10(-21), "z", "zepto"},
	{bigPow(10, -24), math.Pow10(-24), "y", "yocto"},
}

var bitPrefixes = []prefixDef{
	{bigPow(2, 80), math.Pow(2, 80), "Yi", "yobi"},
	{bigPow(2, 70), math.Pow(2, 70), "Zi", "zebi"},
	{bigPow(2, 60), math.Pow(2, 60), "Ei", "exbi"},
	{bigPow(2, 50), math.Pow(2, 50), "Pi", "pebi"},
	{bigPow(2, 40), math.Pow(2, 40), "Ti", "tebi"},
	{bigPow(2, 30), math.Pow(2, 30), "Gi", "gibi"},
	{bigPow(2, 20), math.Pow(2, 20), "Mi", "mebi"},
	{bigPow(2, 10), math.Pow(2, 10), "Ki", "kibi"},
}

// preparePrefixes will build a regular expression to match all possible prefix inputs.
func (humanizer *Humanizer) preparePrefixes() {
	// Save all prefixes into one slice - for convenience.
	humanizer.allPrefixes = append(humanizer.allPrefixes, siPrefixes...)
	humanizer.allPrefixes = append(humanizer.allPrefixes, bitPrefixes...)
	// List of all prefixes as strings.
	prefixes := make([]string, 0, len(humanizer.allPrefixes))
	// Append prefixes.
	for i := range humanizer.allPrefixes {
		// Use this loop to also translate the long versions.
		humanizer.allPrefixes[i].long = humanizer.provider.prefixes[humanizer.allPrefixes[i].short]
		prefixes = append(prefixes, humanizer.allPrefixes[i].long)
		prefixes = append(prefixes, humanizer.allPrefixes[i].short)
	}
	// Regexp will match: number, optional coma or dot, optional second number, optional space, optional suffix.
	humanizer.prefixInputRe = regexp.MustCompile(
		`([0-9]+)[.,]?([0-9]*?) ?(` + strings.Join(prefixes, "|") + `)?$`)
}

// Performs the actual prefixing.
func (humanizer *Humanizer) prefix(value float64, decimals int, threshold int64, short bool, bit bool) string {
	prefixes := siPrefixes
	if bit {
		prefixes = bitPrefixes
	}
	if threshold < 10 {
		threshold = 10
	}
	// If value falls within ignored range then just format it.
	if value <= float64(threshold) && value >= 10.0/float64(threshold) {
		return trimZeroes(strconv.FormatFloat(value, 'f', decimals, 64))
	}
	// Find most appropriate prefix.
	i := sort.Search(len(prefixes), func(i int) bool {
		return prefixes[i].approxValue < value
	})
	if i == len(prefixes) { // prefixDef not found.
		return trimZeroes(strconv.FormatFloat(value, 'f', decimals, 64))
	}

	// For prefixing the approximate value should be enough.
	convertedValue := trimZeroes(
		strconv.FormatFloat(value/prefixes[i].approxValue, 'f', decimals, 64))

	if short {
		return convertedValue + prefixes[i].short
	}
	return convertedValue + " " + prefixes[i].long
}

// BitPrefixFast is a convenience wrapper over BitPrefix.
// Precision is 2 decimal place. Will not prefix values smaller than 1024 and will append only the short prefix.
func (humanizer *Humanizer) BitPrefixFast(value float64) string {
	return humanizer.BitPrefix(value, 2, 1024, true)
}

// SiPrefixFast is a convenience function for easy prefixing with a SI prefix.
// Precision is 1 decimal place. Will not prefix values in range 0.01 - 1000 and will append only the short prefix.
func (humanizer *Humanizer) SiPrefixFast(value float64) string {
	return humanizer.SiPrefix(value, 1, 1000, true)
}

// SiPrefix appends a SI prefix to the value and converts it accordingly.
// Arguments:
//  value - value to be converted.
//  decimals - decimal precision for the converted value.
//  threshold - upper bound of the value range to be ignored. Lower bound is 1/threshold.
//  short - whether to use short or long prefix.
func (humanizer *Humanizer) SiPrefix(value float64, decimals int, threshold int64, short bool) string {
	return humanizer.prefix(value, decimals, threshold, short, false)
}

// BitPrefix appends a bit prefix to the value and converts it accordingly.
// Arguments:
//  value - value to be converted (should be in bytes).
//  decimals - decimal precision for the converted value.
//  threshold - upper bound of the value range to be ignored. Lower bound is 1/threshold.
//  short - whether to use short or long prefix.
func (humanizer *Humanizer) BitPrefix(value float64, decimals int, threshold int64, short bool) string {
	return humanizer.prefix(value, decimals, threshold, short, true)
}

// ParsePrefix will return a number as parsed from input string.
func (humanizer *Humanizer) ParsePrefix(input string) (*big.Float, error) {
	matched := humanizer.prefixInputRe.FindStringSubmatch(strings.TrimSpace(input))
	// 0 - full match, 1 - number, 2 - decimal, 3 - suffix
	if len(matched) != 4 {
		return new(big.Float), fmt.Errorf("cannot parse %q", input)
	}

	// Parse first two groups as a float.
	// This can only fail if the regexp is wrong and allows non numbers.
	number, _ := new(big.Float).SetString(matched[1] + "." + matched[2])

	// No suffix, no multiplication.
	if matched[3] == "" {
		return number, nil
	}
	// Get the multiplier for the prefix.
	for _, prefix := range humanizer.allPrefixes {
		if prefix.short == matched[3] || prefix.long == matched[3] {
			result := new(big.Float).Mul(number, prefix.value)
			return result, nil
		}
	}

	// No prefix was found. This should never happen as the regexp covers all units.
	return new(big.Float), fmt.Errorf("can't match prefix for %q", matched[3])
}
