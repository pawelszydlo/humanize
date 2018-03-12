// Package humanize provides methods for printing and reading values
// in human readable form.
package humanize

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// List all the language providers here.
var languages = map[string]languageProvider{
	"pl": lang_pl,
	"en": lang_en,
}

type timeUnits map[string]int64

// languageProvider is a struct defining all the needed language elements.
type languageProvider struct {
	// Slice of time ranges to humanize time.
	timeRanges []timeRange
	// String for formatting time in the future.
	timeFuture string
	// String for formatting time in the past.
	timePast string
	// Regular expression for matching time input.
	timeUnits timeUnits
}

// Humanizer is the main struct that provides the public methods.
type Humanizer struct {
	provider    languageProvider
	timeInputRe *regexp.Regexp
}

// buildTimeInputRe will build a regular expression to match all possible time inputs.
func (humanizer *Humanizer) buildTimeInputRe() {
	// Get all possible time units.
	units := make([]string, 0, len(humanizer.provider.timeUnits))
	for unit := range humanizer.provider.timeUnits {
		units = append(units, unit)
	}
	// Regexp will match: number, optional coma or dot, optional second number, unit name
	humanizer.timeInputRe = regexp.MustCompile("([0-9]+)[.,]?([0-9]*?) (" + strings.Join(units, "|") + ")")
}

// New creates a new humanizer for a given language.
func New(language string) (*Humanizer, error) {
	if provider, exists := languages[language]; exists {
		humanizer := &Humanizer{
			provider: provider,
		}
		humanizer.buildTimeInputRe()
		return humanizer, nil
	} else {
		return nil, errors.New(fmt.Sprintf("Language not supported: %s", language))
	}
}
