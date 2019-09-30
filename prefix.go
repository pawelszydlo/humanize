package humanize

import (
	"errors"
	"fmt"
	"math/big"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// SI prefixing functions.

type Prefix struct {
	multiplier float64
	short      string
	long       string
}

// preparePrefixes will build a regular expression to match all possible prefix inputs.
func (humanizer *Humanizer) preparePrefixes() {
	// Save all possible prefixes.
	humanizer.allPrefixes = append(humanizer.allPrefixes, humanizer.provider.siPrefixes...)
	humanizer.allPrefixes = append(humanizer.allPrefixes, humanizer.provider.bitPrefixes...)
	// List of all prefixes as strings.
	prefixes := make([]string, 0, len(humanizer.allPrefixes))
	// Append prefixes.
	for _, prefix := range humanizer.allPrefixes {
		prefixes = append(prefixes, prefix.long)
		prefixes = append(prefixes, prefix.short)
	}
	// Regexp will match: number, optional coma or dot, optional second number, optional space, optional suffix.
	humanizer.prefixInputRe = regexp.MustCompile(
		`([0-9]+)[.,]?([0-9]*?) ?(` + strings.Join(prefixes, "|") + `)?$`)
}

// Hack to get rid of trailing zeroes (while keeping the precision if necessary)
func (humanizer *Humanizer) trimZeroes(value string) string {
	if strings.ContainsRune(value, '.') {
		value = strings.TrimRight(value, "0")
		value = strings.TrimRight(value, ".")
	}
	return value
}

// Performs the actual prefixing.
func (humanizer *Humanizer) prefix(value float64, decimals int, threshold int64, short bool, bit bool) string {
	prefixes := humanizer.provider.siPrefixes
	if bit {
		prefixes = humanizer.provider.bitPrefixes
	}
	if threshold < 10 {
		threshold = 10
	}
	// If value falls within ignored range then just format it.
	if value <= float64(threshold) && value >= 10.0/float64(threshold) {
		return humanizer.trimZeroes(strconv.FormatFloat(value, 'f', decimals, 64))
	}
	// Find most appropriate prefix.
	i := sort.Search(len(prefixes), func(i int) bool {
		return prefixes[i].multiplier < value
	})

	convertedValue := humanizer.trimZeroes(
		strconv.FormatFloat(value/prefixes[i].multiplier, 'f', decimals, 64))

	if short {
		return convertedValue + prefixes[i].short
	} else {
		return convertedValue + " " + prefixes[i].long
	}
}

// BitPrefixFast is a convenience wrapper over BitPrefix. See help for PrefixFast.
func (humanizer *Humanizer) BitPrefixFast(value float64) string {
	return humanizer.BitPrefix(value, 1, 1000, true)
}

// PrefixFast is a convenience function for easy prefixing with a SI prefix.
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
// TODO: precision is an issue with big, uniform numbers. Figure it out.
func (humanizer *Humanizer) ParsePrefix(input string) (float64, error) {
	matched := humanizer.prefixInputRe.FindStringSubmatch(strings.TrimSpace(input))
	if len(matched) != 4 {
		return 0, errors.New(fmt.Sprintf("Cannot parse '%s'.", input))
	}

	// 0 - full match, 1 - number, 2 - decimal, 3 - suffix
	if matched[2] == "" { // Decimal component is empty.
		matched[2] = "0"
	}
	// Parse first two groups a float.
	number, err := strconv.ParseFloat(matched[1]+"."+matched[2], 64)
	if err != nil { // This can only fail if the regexp is wrong and allows non numbers.
		return 0, err
	}
	// No suffix, no multiplication.
	if matched[3] == "" {
		return number, nil
	}
	// Get the multiplier for the prefix.
	for _, prefix := range humanizer.allPrefixes {
		if prefix.short == matched[3] || prefix.long == matched[3] {
			result, _ := new(big.Float).Mul(
				new(big.Float).SetFloat64(number),
				new(big.Float).SetFloat64(prefix.multiplier)).Float64()
			return result, nil
		}
	}

	// No prefix was found. This should never happen as the regexp covers all units.
	return 0, errors.New(fmt.Sprintf("Can't match prefix for '%s'.", matched[3]))
}
