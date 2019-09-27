package humanize

import (
	"testing"
	"time"
)

func TestHumanizer_TimeDiffNow(t *testing.T) {
	humanizer, err := New("en")
	if err != nil {
		t.Errorf("Humanizer creation failed with error: %s", err)
	}
	cases := map[time.Duration]string{
		time.Duration(0):                           "now",
		time.Duration(1 * time.Second):             "in 1 second",
		time.Duration(15 * time.Minute):            "in 15 minutes",
		time.Duration(2*time.Hour + 5*time.Minute): "in 2 hours",
		time.Duration(3 * 24 * time.Hour):          "in 3 days",
		time.Duration(15 * 24 * time.Hour):         "in 2 weeks",
		time.Duration(40 * 24 * time.Hour):         "in 1 month",
	}

	for duration, expected := range cases {
		humanized := humanizer.TimeDiffNow(time.Now().Add(duration), false)
		if humanized != expected {
			t.Errorf("Expected '%s', got '%s'.", expected, humanized)
		}
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
	humanizer, err := New("en")
	if err != nil {
		t.Errorf("Humanizer creation failed with error: %s", err)
	}
	startDate := time.Date(2000, 6, 15, 12, 0, 0, 0, time.UTC)
	cases := map[time.Time]string{
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
	}

	for endDate, expected := range cases {
		humanized := humanizer.TimeDiff(startDate, endDate, false)
		if humanized != expected {
			t.Errorf("Expected '%s', got '%s'.", expected, humanized)
		}
	}
}

func TestHumanizer_TimeDiff_Precise(t *testing.T) {
	humanizer, err := New("en")
	if err != nil {
		t.Errorf("Humanizer creation failed with error: %s", err)
	}
	startDate := time.Date(2000, 6, 15, 12, 0, 0, 0, time.UTC)
	cases := map[time.Time]string{
		time.Date(2000, 6, 15, 12, 0, 1, 0, time.UTC):   "in 1 second",
		time.Date(2000, 6, 15, 11, 59, 30, 0, time.UTC): "30 seconds ago",
		time.Date(2000, 6, 15, 12, 15, 1, 0, time.UTC):  "in 15 minutes and 1 second",
		time.Date(2000, 6, 15, 11, 49, 1, 0, time.UTC):  "10 minutes and 59 seconds ago",
		time.Date(2000, 6, 18, 12, 5, 1, 0, time.UTC):   "in 3 days, 5 minutes and 1 second",
		time.Date(2000, 6, 10, 5, 0, 0, 0, time.UTC):    "5 days and 7 hours ago",
		time.Date(2020, 8, 1, 0, 0, 0, 0, time.UTC):     "in 20 years, 5 months, 1 day and 12 hours",
	}

	for endDate, expected := range cases {
		humanized := humanizer.TimeDiff(startDate, endDate, true)
		if humanized != expected {
			t.Errorf("Expected '%s', got '%s'.", expected, humanized)
		}
	}
}

func TestHumanizer_ParseDuration(t *testing.T) {
	humanizer, err := New("en")
	if err != nil {
		t.Errorf("Humanizer creation failed with error: %s", err)
	}
	cases := map[string]time.Duration{
		"3 minutes":              time.Duration(3 * Minute * time.Second),
		"2.5 hours":              time.Duration(2.5 * Hour * time.Second),
		"70 days":                time.Duration(70 * Day * time.Second),
		"5 weeks":                time.Duration(5 * Week * time.Second),
		"3.3 months":             time.Duration(3.3 * Month * time.Second),
		"10 years":               time.Duration(10 * Year * time.Second),
		"2 days and 5 hours":     time.Duration(2*Day*time.Second + 5*Hour*time.Second),
		"2 years, 19 months":     time.Duration(2*Year*time.Second + 19*Month*time.Second),
		"2 days and then 2 days": time.Duration(4 * Day * time.Second),
	}

	for input, expected := range cases {
		humanized, err := humanizer.ParseDuration(input)
		if err != nil {
			t.Errorf("Humanization failed: %s", err)
		}
		if humanized != expected {
			t.Errorf("Expected '%s', got '%s'.", expected, humanized)
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
		5:              "00:05",
		67:             "01:07",
		127:            "02:07",
		0:              "00:00",
		9999:           "2:46:39",
	}

	for input, expected := range cases {
		humanized := humanizer.SecondsToTimeString(input)
		if humanized != expected {
			t.Errorf("Expected '%s', got '%s'.", expected, humanized)
		}
	}
}
