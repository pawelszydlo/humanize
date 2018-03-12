package humanize

var lang_en = languageProvider{
	timeRanges: []timeRange{
		{60, "a moment", 1},
		{2 * Minute, "1 minute", 1},
		{Hour, "%d minutes", Minute},
		{2 * Hour, "1 hour", 1},
		{Day, "%d hours", Hour},
		{2 * Day, "1 day", 1},
		{Week, "%d days", Day},
		{2 * Week, "1 week", 1},
		{Month, "%d weeks", Week},
		{2 * Month, "1 month", 1},
		{Year, "%d months", Month},
		{18 * Month, "1 year", 1},
		{LongTime, "%d years", Year},
	},
	timeFuture: "in %s",
	timePast:   "%s ago",
}
