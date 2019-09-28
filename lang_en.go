package humanize

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
}
