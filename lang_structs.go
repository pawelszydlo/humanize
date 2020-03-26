package humanize

// Language definition structures.

// List all the existing language providers here.
var languages = map[string]languageProvider{
	"pl": langPl,
	"en": langEn,
}

// languageProvider is a struct defining all the needed language elements.
type languageProvider struct {
	times    times
	prefixes map[string]string
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
	units inputTimeUnits
}

// Time unit definitions for input parsing. Use partial matches.
type inputTimeUnits map[string]int64

// Definition of time ranges to match against.
type timeRanges struct {
	upperLimit         int64 // Range end.
	divideBy           int64
	skipWhenPrecise    bool   // Skip this range in precise mode (useful for skipping "weeks")
	onlyLastDigitAfter int64  // Consider only the last digit for the unit after this number. 0 to disable.
	singular           string // Most languages need special treatment for singular units.
	ranges             []timeRange
}

// Represents a single time range.
type timeRange struct {
	upperLimit int64 // Limit in the units of the range!
	format     string
}
