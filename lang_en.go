package humanize

// English l10n. For description see lang_structs.go.
var langEn = languageProvider{
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
	prefixes: map[string]string{
		// SI.
		"Y":  "yotta",
		"Z":  "zetta",
		"E":  "exa",
		"P":  "peta",
		"T":  "tera",
		"G":  "giga",
		"M":  "mega",
		"k":  "kilo",
		"h":  "hecto",
		"da": "deca",
		"d":  "deci",
		"c":  "centi",
		"m":  "milli",
		"µ":  "micro",
		"n":  "nano",
		"p":  "pico",
		"f":  "femto",
		"a":  "atto",
		"z":  "zepto",
		"y":  "yocto",
		// Bit.
		"Yi": "yobi",
		"Zi": "zebi",
		"Ei": "exbi",
		"Pi": "pebi",
		"Ti": "tebi",
		"Gi": "gibi",
		"Mi": "mebi",
		"Ki": "kibi",
	},
}
