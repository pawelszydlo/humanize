package humanize

import (
	"testing"
)

func TestHumanizer_Prefix(t *testing.T) {
	humanizer, err := New("en")
	if err != nil {
		t.Errorf("Humanizer creation failed with error: %s", err)
	}

	cases := map[string]string {
		"2.9k": humanizer.Prefix(2854, 1, 1000, true),
		"23 mega": humanizer.Prefix(22843853, 0, 1000, false),
		"1.44M": humanizer.Prefix(1440000, 2, 1000, true),
		"5.3µ": humanizer.Prefix(0.00000534, 1, 100, true),
		"2345": humanizer.Prefix(2345, 1, 10000, true),
		"1Y": humanizer.Prefix(1000000000001000000000000, 1, 1000, true),

		"174.5k": humanizer.PrefixFast(174512),
		"28M": humanizer.PrefixFast(28000000),
		"5.1m": humanizer.PrefixFast(0.005123),
		"175": humanizer.PrefixFast(175),
		"1k": humanizer.PrefixFast(1024),
	}

	for expected, humanized := range cases {
		if humanized != expected {
			t.Errorf("Expected '%s', got '%s'.", expected, humanized)
		}
	}
}
