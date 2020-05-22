package humanize

import "golang.org/x/text/number"

// Parts of one humanization functions.

// Single part unit definition.
type partUnitDef struct {
	divider float64
	unit    string
}

var partUnits = []partUnitDef{
	{100, "%"},       // Percent
	{1000, "‰"},      // Per mille.
	{10000, "‱"},     // Basis point.
	{100000, "pcm"},  // Per cent mille.
	{1000000, "ppm"}, // Parts per million. Not really matching here, but no better unit.
}

// HumanizeParts makes part of one fractions human readable, e.g.: 0.02 -> 2%, 0.007 -> 7‰.
// Arguments:
//   value - the value to be formatted
//   allowedZeroes - number of leading zeroes allowed in the fraction before moving to a smaller unit
func (humanizer *Humanizer) HumanizeParts(value float64, allowedZeroes int) string {
	foundDivider := 0.0
	foundUnit := ""
	for _, part := range partUnits {
		foundDivider = part.divider
		foundUnit = part.unit
		valueToCheck := value*part.divider
		if allowedZeroes > 0 {
			valueToCheck *= float64(allowedZeroes*10)
		}
		if valueToCheck > 1 {
			break
		}
	}
	finalValue := value * foundDivider
	decimalPoints := 0
	if allowedZeroes > 0 {
		decimalPoints = allowedZeroes + 1
	}
	return humanizer.printer.Sprintf(
		"%g%s", number.Decimal(finalValue, number.MaxFractionDigits(decimalPoints)), foundUnit)
}
