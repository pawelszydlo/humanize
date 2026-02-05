package humanize

import (
	"testing"
	"time"
)

func TestHumanizer_TimeDiffNow(t *testing.T) {
	cases := map[string]map[time.Duration]string{
		"en": {
			time.Duration(0):                           "now",
			time.Duration(1 * time.Second):             "in 1 second",
			time.Duration(15 * time.Minute):            "in 15 minutes",
			time.Duration(2*time.Hour + 5*time.Minute): "in 2 hours",
			time.Duration(3 * 24 * time.Hour):          "in 3 days",
			time.Duration(15 * 24 * time.Hour):         "in 2 weeks",
			time.Duration(40 * 24 * time.Hour):         "in 1 month",
		},
		"pl": {
			time.Duration(0):                           "teraz",
			time.Duration(1 * time.Second):             "za sekundę",
			time.Duration(15 * time.Minute):            "za 15 minut",
			time.Duration(2*time.Hour + 5*time.Minute): "za 2 godziny",
			time.Duration(3 * 24 * time.Hour):          "za 3 dni",
			time.Duration(15 * 24 * time.Hour):         "za 2 tygodnie",
			time.Duration(40 * 24 * time.Hour):         "za miesiąc",
		},
	}

	for lang, caseList := range cases {
		humanizer, err := New(lang)
		if err != nil {
			t.Errorf("Humanizer creation failed with error: %s", err)
		}

		for duration, expected := range caseList {
			humanized := humanizer.TimeDiffNow(time.Now().Add(duration), false)
			if humanized != expected {
				t.Errorf("Expected '%s', got '%s'.", expected, humanized)
			}
		}
	}
}

func TestHumanizer_TimeDiff_PreciseSkip(t *testing.T) {
	humanizer, err := New("en")
	if err != nil {
		t.Errorf("Humanizer creation failed with error: %s", err)
	}
	startDate := time.Date(2000, 6, 1, 12, 0, 0, 0, time.UTC)
	endDate := time.Date(2000, 6, 17, 12, 0, 0, 0, time.UTC)
	// Precise.
	humanized := humanizer.TimeDiff(startDate, endDate, true)
	expected := "in 16 days"
	if humanized != expected {
		t.Errorf("Expected '%s', got '%s'.", expected, humanized)
	}
	// Imprecise.
	humanized = humanizer.TimeDiff(startDate, endDate, false)
	expected = "in 2 weeks"
	if humanized != expected {
		t.Errorf("Expected '%s', got '%s'.", expected, humanized)
	}
}

func TestHumanizer_TimeDiffNow_TZ(t *testing.T) {
	humanizer, err := New("en")
	if err != nil {
		t.Errorf("Humanizer creation failed with error: %s", err)
	}
	// Set arbitrary location.
	loc, _ := time.LoadLocation("Asia/Shanghai")
	date := time.Now().Add(time.Duration(5 * time.Minute)).In(loc)
	// Make sure that TimeDiffNow is TZ agnostic.
	humanized := humanizer.TimeDiffNow(date, false)
	if humanized != "in 5 minutes" {
		t.Errorf("Expected 'in 5 minutes', got '%s'.", humanized)
	}
}

func TestHumanizer_TimeDiff_Imprecise(t *testing.T) {
	startDate := time.Date(2000, 6, 15, 12, 0, 0, 0, time.UTC)
	cases := map[string]map[time.Time]string{
		"en": {
			time.Date(2000, 6, 15, 12, 0, 0, 0, time.UTC):   "now",
			time.Date(2000, 6, 15, 12, 0, 1, 0, time.UTC):   "in 1 second",
			time.Date(2000, 6, 15, 11, 59, 30, 0, time.UTC): "30 seconds ago",
			time.Date(2000, 6, 15, 12, 15, 1, 0, time.UTC):  "in 15 minutes",
			time.Date(2000, 6, 15, 11, 49, 1, 0, time.UTC):  "10 minutes ago",
			time.Date(2000, 6, 18, 12, 0, 1, 0, time.UTC):   "in 3 days",
			time.Date(2000, 6, 10, 5, 0, 1, 0, time.UTC):    "5 days ago",
			time.Date(2000, 6, 29, 12, 0, 1, 0, time.UTC):   "in 2 weeks",
			time.Date(2000, 6, 1, 1, 0, 1, 0, time.UTC):     "2 weeks ago",
			time.Date(2000, 7, 15, 12, 0, 1, 0, time.UTC):   "in 1 month",
			time.Date(2000, 5, 15, 12, 0, 1, 0, time.UTC):   "1 month ago",
		},
		"pl": {
			time.Date(2000, 6, 15, 12, 0, 0, 0, time.UTC):   "teraz",
			time.Date(2000, 6, 15, 12, 0, 1, 0, time.UTC):   "za sekundę",
			time.Date(2000, 6, 15, 11, 59, 30, 0, time.UTC): "30 sekund temu",
			time.Date(2000, 6, 15, 12, 15, 1, 0, time.UTC):  "za 15 minut",
			time.Date(2000, 6, 15, 11, 37, 1, 0, time.UTC):  "22 minuty temu",
			time.Date(2000, 6, 18, 12, 0, 1, 0, time.UTC):   "za 3 dni",
			time.Date(2000, 6, 10, 5, 0, 1, 0, time.UTC):    "5 dni temu",
			time.Date(2000, 6, 29, 12, 0, 1, 0, time.UTC):   "za 2 tygodnie",
			time.Date(2000, 6, 1, 1, 0, 1, 0, time.UTC):     "2 tygodnie temu",
			time.Date(2000, 7, 15, 12, 0, 1, 0, time.UTC):   "za miesiąc",
			time.Date(2000, 5, 15, 12, 0, 1, 0, time.UTC):   "miesiąc temu",
		},
	}

	for lang, caseList := range cases {
		humanizer, err := New(lang)
		if err != nil {
			t.Errorf("Humanizer creation failed with error: %s", err)
		}

		for endDate, expected := range caseList {
			humanized := humanizer.TimeDiff(startDate, endDate, false)
			if humanized != expected {
				t.Errorf("Expected '%s', got '%s'.", expected, humanized)
			}
		}
	}

}

func TestHumanizer_TimeDiff_Precise(t *testing.T) {
	startDate := time.Date(2000, 6, 15, 12, 0, 0, 0, time.UTC)
	cases := map[string]map[time.Time]string{
		"en": {
			time.Date(2000, 6, 15, 12, 0, 1, 0, time.UTC):   "in 1 second",
			time.Date(2000, 6, 15, 11, 59, 30, 0, time.UTC): "30 seconds ago",
			time.Date(2000, 6, 15, 12, 15, 1, 0, time.UTC):  "in 15 minutes and 1 second",
			time.Date(2000, 6, 15, 11, 49, 1, 0, time.UTC):  "10 minutes and 59 seconds ago",
			time.Date(2000, 6, 18, 12, 5, 1, 0, time.UTC):   "in 3 days, 5 minutes and 1 second",
			time.Date(2000, 6, 10, 5, 0, 0, 0, time.UTC):    "5 days and 7 hours ago",
			time.Date(2020, 8, 1, 0, 0, 0, 0, time.UTC):     "in 20 years, 5 months, 1 day and 12 hours",
		},
		"pl": {
			time.Date(2000, 6, 15, 12, 0, 1, 0, time.UTC):   "za sekundę",
			time.Date(2000, 6, 15, 11, 59, 30, 0, time.UTC): "30 sekund temu",
			time.Date(2000, 6, 15, 12, 15, 1, 0, time.UTC):  "za 15 minut i sekundę",
			time.Date(2000, 6, 15, 11, 49, 1, 0, time.UTC):  "10 minut i 59 sekund temu",
			time.Date(2000, 6, 18, 12, 5, 1, 0, time.UTC):   "za 3 dni, 5 minut i sekundę",
			time.Date(2000, 6, 10, 5, 0, 0, 0, time.UTC):    "5 dni i 7 godzin temu",
			time.Date(2020, 8, 1, 0, 0, 0, 0, time.UTC):     "za 20 lat, 5 miesięcy, 1 dzień i 12 godzin",
		},
	}

	for lang, caseList := range cases {
		humanizer, err := New(lang)
		if err != nil {
			t.Errorf("Humanizer creation failed with error: %s", err)
		}

		for endDate, expected := range caseList {
			humanized := humanizer.TimeDiff(startDate, endDate, true)
			if humanized != expected {
				t.Errorf("Expected '%s', got '%s'.", expected, humanized)
			}
		}
	}
}

func TestHumanizer_ParseDuration(t *testing.T) {
	cases := map[string]map[string]time.Duration{
		"en": {
			"3 minutes":              time.Duration(3 * Minute * time.Second),
			"2.5 hours":              time.Duration(2.5 * Hour * time.Second),
			"70 days":                time.Duration(70 * Day * time.Second),
			"5 weeks":                time.Duration(5 * Week * time.Second),
			"3.3 months":             time.Duration(3.3 * Month * time.Second),
			"10 years":               time.Duration(10 * Year * time.Second),
			"2 days and 5 hours":     time.Duration(2*Day*time.Second + 5*Hour*time.Second),
			"2 years, 19 months":     time.Duration(2*Year*time.Second + 19*Month*time.Second),
			"2 days and then 2 days": time.Duration(4 * Day * time.Second),
			"-2 days":                time.Duration(-2 * Day * time.Second),
			"-2 months and 10 days":  time.Duration(-2*Month*time.Second - 10*Day*time.Second),
		},
		"pl": {
			"3 minuty":              time.Duration(3 * Minute * time.Second),
			"2.5 godziny":           time.Duration(2.5 * Hour * time.Second),
			"70 dni":                time.Duration(70 * Day * time.Second),
			"5 tygodni":             time.Duration(5 * Week * time.Second),
			"3.3 miesiąca":          time.Duration(3.3 * Month * time.Second),
			"10 lat":                time.Duration(10 * Year * time.Second),
			"2 dni i 5 godzin":      time.Duration(2*Day*time.Second + 5*Hour*time.Second),
			"2 lata, 19 miesięcy":   time.Duration(2*Year*time.Second + 19*Month*time.Second),
			"2 dni i jeszcze 2 dni": time.Duration(4 * Day * time.Second),
		},
	}

	for lang, caseList := range cases {
		humanizer, err := New(lang)
		if err != nil {
			t.Errorf("Humanizer creation failed with error: %s", err)
		}

		for input, expected := range caseList {
			humanized, err := humanizer.ParseDuration(input)
			if err != nil {
				t.Errorf("Humanization failed: %s", err)
			}
			if humanized != expected {
				t.Errorf("Expected '%s', got '%s'.", expected, humanized)
			}
		}
	}
}

func TestHumanizer_ParseDuration_Incorrect(t *testing.T) {
	humanizer, err := New("en")
	if err != nil {
		t.Errorf("Humanizer creation failed with error: %s", err)
	}
	_, err = humanizer.ParseDuration("wrong duration")
	if err == nil {
		t.Error("Humanization succeeded where it should have failed.")
	}
}

func TestHumanizer_SecondsToTimeString(t *testing.T) {
	humanizer, err := New("en")
	if err != nil {
		t.Errorf("Humanizer creation failed with error: %s", err)
	}
	cases := map[int64]string{
		5:    "00:05",
		67:   "01:07",
		127:  "02:07",
		0:    "00:00",
		9999: "2:46:39",
	}

	for input, expected := range cases {
		humanized := humanizer.SecondsToTimeString(input)
		if humanized != expected {
			t.Errorf("Expected '%s', got '%s'.", expected, humanized)
		}
	}
}
