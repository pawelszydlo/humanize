// Package humanize provides methods for formatting and parsing values in human readable form.
package humanize

import (
	"errors"
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"regexp"
)

// Humanizer is the main struct that provides the public methods.
type Humanizer struct {
	provider      languageProvider
	printer       *message.Printer
	timeInputRe   *regexp.Regexp
	metricInputRe *regexp.Regexp
}

// New creates a new humanizer for a given language.
func New(langName string) (*Humanizer, error) {
	if provider, exists := languages[langName]; exists {
		humanizer := &Humanizer{
			provider: provider,
			printer:  message.NewPrinter(language.MustParse(langName)),
		}
		humanizer.buildTimeInputRe()
		humanizer.buildMetricInputRe()
		return humanizer, nil
	} else {
		return nil, errors.New(fmt.Sprintf("Language not supported: %s", langName))
	}
}
