package humanize

// Polish l10n. For description see lang_structs.go.
var lang_pl = languageProvider{
	times: times{
		ranges: []timeRanges{
			{Minute, 1, false, 20, "sekundę", []timeRange{
				{2, "%d sekund"},
				{5, "%d sekundy"},
				{LongTime, "%d sekund"},
			}},
			{Hour, Minute, false, 20, "minutę", []timeRange{
				{2, "%d minut"},
				{5, "%d minuty"},
				{Hour, "%d minut"},
			}},
			{Day, Hour, false, 20, "godzinę", []timeRange{
				{2, "%d godzin"},
				{5, "%d godziny"},
				{LongTime, "%d godzin"},
			}},
			{Week, Day, false, 20, "1 dzień", []timeRange{
				{LongTime, "%d dni"},
			}},
			{Month, Week, true, 20, "tydzień", []timeRange{
				{2, "%d tygodni"},
				{5, "%d tygodnie"},
				{LongTime, "%d tygodni"},
			}},
			{Year, Month, false, 20, "miesiąc", []timeRange{
				{2, "%d miesięcy"},
				{5, "%d miesiące"},
				{LongTime, "%d miesięcy"},
			}},
			{LongTime, Year, false, 20, "rok", []timeRange{
				{2, "%d lat"},
				{5, "%d lata"},
				{LongTime, "%d lat"},
			}},
		},
		future:       "za %s",
		past:         "%s temu",
		now:          "teraz",
		remainderSep: "i",
		units: inputTimeUnits{
			"sekund": 1,
			"minut":  Minute,
			"godzin": Hour,
			"dzie":   Day,
			"dni":    Day,
			"ty":     Week,
			"miesi":  Month,
			"rok":    Year,
			"lat":    Year,
		},
	},
	prefixes: map[string]string{
		// SI.
		"Y":  "jotta",
		"Z":  "zetta",
		"E":  "eksa",
		"P":  "peta",
		"T":  "tera",
		"G":  "giga",
		"M":  "mega",
		"k":  "kilo",
		"h":  "hekto",
		"da": "deka",
		"d":  "decy",
		"c":  "centy",
		"m":  "mili",
		"µ":  "mikro",
		"n":  "nano",
		"p":  "piko",
		"f":  "femto",
		"a":  "atto",
		"z":  "zepto",
		"y":  "jokto",
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
