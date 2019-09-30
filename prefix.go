package humanize

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// Prefixing functions.

const PREC = 64

type Prefix struct {
	base        int
	power       int
	approxValue float64
	short       string
	long        string
}

var siPrefixes = []Prefix{
	{10, 24, math.Pow10(24), "Y", "yotta"},
	{10, 21, math.Pow10(21), "Z", "zetta"},
	{10, 18, math.Pow10(18), "E", "exa"},
	{10, 15, math.Pow10(15), "P", "peta"},
	{10, 12, math.Pow10(12), "T", "tera"},
	{10, 9, math.Pow10(9), "G", "giga"},
	{10, 6, math.Pow10(6), "M", "mega"},
	{10, 3, math.Pow10(3), "k", "kilo"},
	{10, 2, math.Pow10(2), "h", "hecto"},
	{10, 1, 10, "da", "deca"},
	{10, -1, math.Pow10(-1), "d", "deci"},
	{10, -2, math.Pow10(-2), "c", "centi"},
	{10, -3, math.Pow10(-3), "m", "milli"},
	{10, -6, math.Pow10(-6), "µ", "micro"},
	{10, -9, math.Pow10(-9), "n", "nano"},
	{10, -12, math.Pow10(-12), "p", "pico"},
	{10, -15, math.Pow10(-15), "f", "femto"},
	{10, -18, math.Pow10(-18), "a", "atto"},
	{10, -21, math.Pow10(-21), "z", "zepto"},
	{10, -24, math.Pow10(-24), "y", "yocto"},
}

var bitPrefixes = []Prefix{
	{2, 80, math.Pow(2, 80), "Yi", "yobi"},
	{2, 70, math.Pow(2, 70), "Zi", "zebi"},
	{2, 60, math.Pow(2, 60), "Ei", "exbi"},
	{2, 50, math.Pow(2, 50), "Pi", "pebi"},
	{2, 40, math.Pow(2, 40), "Ti", "tebi"},
	{2, 30, math.Pow(2, 30), "Gi", "gibi"},
	{2, 20, math.Pow(2, 20), "Mi", "mebi"},
	{2, 10, math.Pow(2, 10), "Ki", "kibi"},
}

func bigPow(x int, y int) *big.Float {
	one := new(big.Float).SetInt64(1)
	if y == 0 {
		return one
	}
	bigX := new(big.Float).SetInt64(int64(x))
	product := new(big.Float).Copy(bigX)
	// Always calculate positive power, inverse later.
	isNeg := false
	if y < 0 {
		isNeg = true
		y = -y
	}
	for i := int(0); i < y-1; i++ {
		product = new(big.Float).Mul(product, bigX)
	}
	if isNeg {
		return new(big.Float).Quo(one, product)
	}
	return product
}

// preparePrefixes will build a regular expression to match all possible prefix inputs.
func (humanizer *Humanizer) preparePrefixes() {
	// Save all prefixes into one slice - for convenience.
	humanizer.allPrefixes = append(humanizer.allPrefixes, siPrefixes...)
	humanizer.allPrefixes = append(humanizer.allPrefixes, bitPrefixes...)
	// List of all prefixes as strings.
	prefixes := make([]string, 0, len(humanizer.allPrefixes))
	// Append prefixes.
	for _, prefix := range humanizer.allPrefixes {
		// Use this loop to also translate the long versions.
		prefix.long = humanizer.provider.prefixes[prefix.short]
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
	prefixes := siPrefixes
	if bit {
		prefixes = bitPrefixes
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
		return prefixes[i].approxValue < value
	})

	// For prefixing the approximate value should be enough.
	convertedValue := humanizer.trimZeroes(
		strconv.FormatFloat(value/prefixes[i].approxValue, 'f', decimals, 64))

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
func (humanizer *Humanizer) ParsePrefix(input string) (*big.Float, error) {
	matched := humanizer.prefixInputRe.FindStringSubmatch(strings.TrimSpace(input))
	// 0 - full match, 1 - number, 2 - decimal, 3 - suffix
	if len(matched) != 4 {
		return new(big.Float), errors.New(fmt.Sprintf("Cannot parse '%s'.", input))
	}

	// Parse first two groups as a float.
	number, ok := new(big.Float).SetString(matched[1] + "." + matched[2])
	if !ok { // This can only fail if the regexp is wrong and allows non numbers.
		return new(big.Float), errors.New("Can't parse the number.")
	}
	// No suffix, no multiplication.
	if matched[3] == "" {
		return number, nil
	}
	// Get the multiplier for the prefix.
	for _, prefix := range humanizer.allPrefixes {
		if prefix.short == matched[3] || prefix.long == matched[3] {
			result := new(big.Float).Mul(number, bigPow(prefix.base, prefix.power))
			return result, nil
		}
	}

	// No prefix was found. This should never happen as the regexp covers all units.
	return new(big.Float), errors.New(fmt.Sprintf("Can't match prefix for '%s'.", matched[3]))
}
