// Package humanize provides methods for printing and reading values
// in human readable form.
package humanize

import (
	"errors"
	"fmt"
)

// List all the language providers here.
var languages = map[string]languageProvider{
	"pl": lang_pl,
	"en": lang_en,
}

// languageProvider is a struct defining all the needed language elements.
type languageProvider struct {
	// Slice of time ranges to humanize time.
	timeRanges []timeRange
	// String for formatting time in the future.
	timeFuture string
	// string for formatting time in the past.
	timePast string
}

// Humanizer is the main struct that provides the public methods.
type Humanizer struct {
	provider languageProvider
}

// New creates a new humanizer for a given language.
func New(language string) (*Humanizer, error) {
	if provider, exists := languages[language]; exists {
		return &Humanizer{provider: provider}, nil
	} else {
		return nil, errors.New(fmt.Sprintf("Language not supported: %s", language))
	}
}
