package humanize

import "math"

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
	siPrefixes: []Prefix{
		{math.Pow10(24), "Y", "yotta"},
		{math.Pow10(21), "Z", "zetta"},
		{math.Pow10(18), "E", "exa"},
		{math.Pow10(15), "P", "peta"},
		{math.Pow10(12), "T", "tera"},
		{math.Pow10(9), "G", "giga"},
		{math.Pow10(6), "M", "mega"},
		{math.Pow10(3), "k", "kilo"},
		{math.Pow10(2), "h", "hekto"},
		{10, "da", "deka"},
		{math.Pow10(-1), "d", "decy"},
		{math.Pow10(-2), "c", "centy"},
		{math.Pow10(-3), "m", "mili"},
		{math.Pow10(-6), "µ", "mikro"},
		{math.Pow10(-9), "n", "nano"},
		{math.Pow10(-12), "p", "piko"},
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
