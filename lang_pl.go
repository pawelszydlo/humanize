package humanize

var lang_pl = languageProvider{
	times: times{
		ranges: []timeRanges{
			{Minute, 1, []timeRange{
				{2, "1 sekundę"},
				{5, "%d sekundy"},
				{Minute, "%d sekund"},
			}},
			{Hour, Minute, []timeRange{
				{2 * Minute, "minutę"},
				{5 * Minute, "%d minuty"},
				{Hour, "%d minut"},
			}},
			{Day, Hour, []timeRange{
				{2 * Hour, "1 godzinę"},
				{5 * Hour, "%d godziny"},
				{Day, "%d godzin"},
			}},
			{Week, Day, []timeRange{
				{2 * Day, "1 dzień"},
				{Week, "%d dni"},
			}},
			{Month, Week, []timeRange{
				{2 * Week, "1 tydzień"},
				{5 * Week, "%d tygodnie"},
				{Month, "%d tygodni"},
			}},
			{Year, Month, []timeRange{
				{2 * Month, "1 miesiąc"},
				{5 * Month, "%d miesiące"},
				{Year, "%d miesięcy"},
			}},
			{LongTime, Year, []timeRange{
				{2 * Year, "1 rok"},
				{5 * Year, "%d lata"},
				{LongTime, "%d lat"},
			}},
		},
		future:       "za %s",
		past:         "%s temu",
		now:          "teraz",
		remainderSep: "i",
		units: timeUnits{
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
}
