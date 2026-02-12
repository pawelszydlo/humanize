// Package humanize provides methods for formatting and parsing values in human readable form.
package humanize

import (
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
	prefixInputRe *regexp.Regexp
	allPrefixes   []prefixDef // Helper slice of all prefixes.
}

// New creates a new humanizer for a given language.
func New(langName string) (*Humanizer, error) {
	if provider, exists := languages[langName]; exists {
		humanizer := &Humanizer{
			provider:    provider,
			printer:     message.NewPrinter(language.MustParse(langName)),
			allPrefixes: make([]prefixDef, len(siPrefixes)+len(bitPrefixes)),
		}
		humanizer.buildTimeInputRe()
		humanizer.preparePrefixes()
		return humanizer, nil
	}
	return nil, fmt.Errorf("language not supported: %s", langName)
}
