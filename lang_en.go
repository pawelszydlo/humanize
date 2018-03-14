package humanize

var lang_en = languageProvider{
	times: times{
		ranges: []timeRanges{
			{Minute, 1, []timeRange{
				{2, "1 second"},
				{60, "%d seconds"},
			}},
			{Hour, Minute, []timeRange{
				{2 * Minute, "1 minute"},
				{Hour, "%d minutes"},
			}},
			{Day, Hour, []timeRange{
				{2 * Hour, "1 hour"},
				{Day, "%d hours"},
			}},
			{Week, Day, []timeRange{
				{2 * Day, "1 day"},
				{Week, "%d days"},
			}},
			{Month, Week, []timeRange{
				{2 * Week, "1 week"},
				{Month, "%d weeks"},
			}},
			{Year, Month, []timeRange{
				{2 * Month, "1 month"},
				{Year, "%d months"},
			}},
			{LongTime, Year, []timeRange{
				{2 * Year, "1 year"},
				{LongTime, "%d years"},
			}},
		},
		future:       "in %s",
		past:         "%s ago",
		now:          "now",
		remainderSep: "and",
		units: timeUnits{
			"second": 1,
			"minute": Minute,
			"hour":   Hour,
			"day":    Day,
			"week":   Week,
			"month":  Month,
			"year":   Year,
		},
	},
}
