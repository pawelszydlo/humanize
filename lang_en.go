package humanize

import "math"

// English l10n. For description see lang_structs.go.
var lang_en = languageProvider{
	times: times{
		ranges: []timeRanges{
			{Minute, 1, false, 0, "1 second", []timeRange{
				{LongTime, "%d seconds"},
			}},
			{Hour, Minute, false, 0, "1 minute", []timeRange{
				{LongTime, "%d minutes"},
			}},
			{Day, Hour, false, 0, "1 hour", []timeRange{
				{LongTime, "%d hours"},
			}},
			{Week, Day, false, 0, "1 day", []timeRange{
				{LongTime, "%d days"},
			}},
			{Month, Week, true, 0, "1 week", []timeRange{
				{LongTime, "%d weeks"},
			}},
			{Year, Month, false, 0, "1 month", []timeRange{
				{LongTime, "%d months"},
			}},
			{LongTime, Year, false, 0, "1 year", []timeRange{
				{LongTime, "%d years"},
			}},
		},
		future:       "in %s",
		past:         "%s ago",
		now:          "now",
		remainderSep: "and",
		units: inputTimeUnits{
			"second": 1,
			"minute": Minute,
			"hour":   Hour,
			"day":    Day,
			"week":   Week,
			"month":  Month,
			"year":   Year,
		},
	},
	siPrefixes: []Prefix{
		{math.Pow10(24), "Y", "yotta"},
		{math.Pow10(21), "Z", "zetta"},
		{math.Pow10(18), "E", "exa"},
		{math.Pow10(15), "P", "peta"},
		{math.Pow10(12), "T", "tera"},
		{math.Pow10(9), "G", "giga"},
		{math.Pow10(6), "M", "mega"},
		{math.Pow10(3), "k", "kilo"},
		{math.Pow10(2), "h", "hecto"},
		{10, "da", "deca"},
		{math.Pow10(-1), "d", "deci"},
		{math.Pow10(-2), "c", "centi"},
		{math.Pow10(-3), "m", "milli"},
		{math.Pow10(-6), "µ", "micro"},
		{math.Pow10(-9), "n", "nano"},
		{math.Pow10(-12), "p", "pico"},
		{math.Pow10(-15), "f", "femto"},
		{math.Pow10(-18), "a", "atto"},
		{math.Pow10(-21), "z", "zepto"},
		{math.Pow10(-24), "y", "yocto"},
	},
	bitPrefixes: []Prefix{
		{math.Pow(2, 80), "Yi", "yobi"},
		{math.Pow(2, 70), "Zi", "zebi"},
		{math.Pow(2, 60), "Ei", "exbi"},
		{math.Pow(2, 50), "Pi", "pebi"},
		{math.Pow(2, 40), "Ti", "tebi"},
		{math.Pow(2, 30), "Gi", "gibi"},
		{math.Pow(2, 20), "Mi", "mebi"},
		{math.Pow(2, 10), "Ki", "kibi"},
	},
}
