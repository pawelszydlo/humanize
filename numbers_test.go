package humanize

import (
	"testing"
)

func TestHumanizer_HumanizeNumber(t *testing.T) {
	humanizer, err := New("en")
	if err != nil {
		t.Errorf("Humanizer creation failed with error: %s", err)
	}

	cases := map[float64]string{
		10:        "10",
		743:       "743",
		1289:      "1,289",
		34.00001:  "34",
		4324.2894: "4,324.29",
	}

	for value, expected := range cases {
		humanized := humanizer.HumanizeNumber(value, 2)
		if humanized != expected {
			t.Errorf("Expected '%s', got '%s'.", expected, humanized)
		}
	}
}
