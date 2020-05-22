package humanize

import (
	"testing"
)

func TestHumanizer_HumanizeParts_NoZeroes(t *testing.T) {
	humanizer, err := New("en")
	if err != nil {
		t.Errorf("Humanizer creation failed with error: %s", err)
	}

	cases := map[float64]string{
		10:       "1,000%",
		0.5:      "50%",
		0.12:     "12%",
		0.0453:   "5%",
		0.0023:   "2‰",
		0.000323: "3‱",
		0.000012: "1pcm",
		0.000003: "3ppm",
	}

	for value, expected := range cases {
		humanized := humanizer.HumanizeParts(value, 0)
		if humanized != expected {
			t.Errorf("Expected '%s', got '%s'.", expected, humanized)
		}
	}
}

func TestHumanizer_HumanizeParts_OneZero(t *testing.T) {
	humanizer, err := New("en")
	if err != nil {
		t.Errorf("Humanizer creation failed with error: %s", err)
	}

	cases := map[float64]string{
		10:       "1,000%",
		0.5:      "50%",
		0.12:     "12%",
		0.0453:   "4.53%",
		0.0023:   "0.23%",
		0.000323: "0.32‰",
		0.000012: "0.12‱",
		0.000003: "0.3pcm",
	}

	for value, expected := range cases {
		humanized := humanizer.HumanizeParts(value, 1)
		if humanized != expected {
			t.Errorf("Expected '%s', got '%s'.", expected, humanized)
		}
	}
}
