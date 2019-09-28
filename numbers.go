package humanize

import "golang.org/x/text/number"

// HumanizeNumber makes the number easily readable by adding decimal separators.
// Arguments:
//   value - the value to be formatted
//   digits - number of (max) fraction digits to be shown
func (humanizer *Humanizer) HumanizeNumber(value float64, digits int) string {
	return humanizer.printer.Sprintf("%g", number.Decimal(value, number.MaxFractionDigits(digits)))
}
