package humanize

var lang_pl = languageProvider{
	timeRanges: []timeRange{
		{60, "chwilkę", 1},
		{2 * Minute, "1 minutę", 1},
		{5 * Minute, "%d minuty", Minute},
		{Hour, "%d minut", Minute},
		{2 * Hour, "1 godzinę", 1},
		{5 * Hour, "%d godziny", Hour},
		{Day, "%d godzin", Hour},
		{2 * Day, "1 dzień", 1},
		{Week, "%d dni", Day},
		{2 * Week, "1 tydzień", 1},
		{5 * Week, "%d tygodnie", Week},
		{Month, "%d tygodni", Week},
		{2 * Month, "1 miesiąc", 1},
		{5 * Month, "%d miesiące", Month},
		{Year, "%d miesięcy", Month},
		{18 * Month, "1 rok", 1},
		{5 * Year, "%d lata", Year},
		{LongTime, "%d lat", Year},
	},
	timeFuture: "za %s",
	timePast:   "%s temu",
}
