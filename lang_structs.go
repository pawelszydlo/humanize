package humanize

// Language definition structures.

// List all the existing language providers here.
var languages = map[string]languageProvider{
	"pl": lang_pl,
	"en": lang_en,
}

// languageProvider is a struct defining all the needed language elements.
type languageProvider struct {
	times times
}

// Time related language elements.
type times struct {
	// Time ranges to humanize time.
	ranges []timeRanges
	// String for formatting time in the future.
	future string
	// String for formatting time in the past.
	past string
	// String to humanize now.
	now string
	// Remainder separator
	remainderSep string
	// Unit values for matching the input. Partial matches are ok.
	units timeUnits
}

// Time unit definitions for input parsing. Use partial matches.
type timeUnits map[string]int64

// Definition of time ranges to match against.
type timeRanges struct {
	upperLimit int64
	divideBy   int64
	ranges     []timeRange
}
type timeRange struct {
	upperLimit int64
	format     string
}
