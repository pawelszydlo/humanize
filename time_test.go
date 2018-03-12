package humanize

import (
	"testing"
	"time"
)

func TestHumanizer_TimeDiff(t *testing.T) {
	humanizer, err := New("en")
	if err != nil {
		t.Errorf("Humanizer creation failed with error: %s", err)
	}

	startDate := time.Date(2000, 6, 15, 12, 0,0,0, time.UTC)
	cases := map[time.Time]string {
		time.Date(2000, 6, 15, 12, 0,1,0, time.UTC): "in a moment",
		time.Date(2000, 6, 15, 11, 59,30,0, time.UTC): "a moment ago",
		time.Date(2000, 6, 15, 12, 15,1,0, time.UTC): "in 15 minutes",
		time.Date(2000, 6, 15, 11, 49,1,0, time.UTC): "10 minutes ago",
		time.Date(2000, 6, 18, 12, 0,1,0, time.UTC): "in 3 days",
		time.Date(2000, 6, 10, 5, 0,1,0, time.UTC): "5 days ago",
		time.Date(2000, 6, 29, 12, 0,1,0, time.UTC): "in 2 weeks",
		time.Date(2000, 6, 1, 1, 0,1,0, time.UTC): "2 weeks ago",
		time.Date(2000, 7, 15, 12, 0,1,0, time.UTC): "in 1 month",
		time.Date(2000, 5, 15, 12, 0,1,0, time.UTC): "1 month ago",
	}

	for endDate, expected := range cases {
		humanized := humanizer.TimeDiff(startDate, endDate)
		if humanized != expected {
			t.Errorf("Expected '%s', got '%s'.", expected, humanized)
		}
	}
}