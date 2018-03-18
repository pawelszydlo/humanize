// Package humanize provides methods for formatting and parsing values in human readable form.
package humanize

import (
	"errors"
	"fmt"
	"regexp"
)

// Humanizer is the main struct that provides the public methods.
type Humanizer struct {
	provider      languageProvider
	timeInputRe   *regexp.Regexp
	metricInputRe *regexp.Regexp
}

// New creates a new humanizer for a given language.
func New(language string) (*Humanizer, error) {
	if provider, exists := languages[language]; exists {
		humanizer := &Humanizer{
			provider: provider,
		}
		humanizer.buildTimeInputRe()
		humanizer.buildMetricInputRe()
		return humanizer, nil
	} else {
		return nil, errors.New(fmt.Sprintf("Language not supported: %s", language))
	}
}
