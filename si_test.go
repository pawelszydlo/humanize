package humanize

import (
	"math"
	"testing"
)

// almostEqual does float comparison ignoring their least significant bit.
func almostEqual(a, b float64) bool {
	ai, bi := int64(math.Float64bits(a)), int64(math.Float64bits(b))
	return a == b || -1 <= ai-bi && ai-bi <= 1
}

func TestHumanizer_Prefix(t *testing.T) {
	humanizer, err := New("en")
	if err != nil {
		t.Errorf("Humanizer creation failed with error: %s", err)
	}

	cases := map[string]string{
		"2.9k":    humanizer.Prefix(2854, 1, 1000, true),
		"23 mega": humanizer.Prefix(22843853, 0, 1000, false),
		"1.44M":   humanizer.Prefix(1440000, 2, 1000, true),
		"5.3µ":    humanizer.Prefix(0.00000534, 1, 100, true),
		"2345":    humanizer.Prefix(2345, 1, 10000, true),
		"1Y":      humanizer.Prefix(1000000000001000000000000, 1, 1000, true),
		// Too low threshold.
		"1": humanizer.Prefix(1, 1, 1, true),
		// Fast.
		"174.5k": humanizer.PrefixFast(174512),
		"28M":    humanizer.PrefixFast(28000000),
		"5.1m":   humanizer.PrefixFast(0.005123),
		"175":    humanizer.PrefixFast(175),
		"1k":     humanizer.PrefixFast(1024),
		// Integer.
		"2k": humanizer.PrefixFastInt(2000),
	}

	for expected, humanized := range cases {
		if humanized != expected {
			t.Errorf("Expected '%s', got '%s'.", expected, humanized)
		}
	}
}

func TestHumanizer_ParsePrefix(t *testing.T) {
	humanizer, err := New("en")
	if err != nil {
		t.Errorf("Humanizer creation failed with error: %s", err)
	}

	cases := map[string]float64{
		"2.9k":       2900,
		"13":         13,
		"0.5":        0.5,
		"13.5 kilo ": 13500,
		"20m":        0.020,
		" 20M ":      20000000,
		"5yotta":     5000000000000000000000000,
		"15 µ":       0.000015,
	}

	for input, expected := range cases {
		humanized, err := humanizer.ParsePrefix(input)
		if err != nil {
			t.Errorf("Error parsing '%s': %s", input, err)
		}

		if !almostEqual(humanized, expected) {
			t.Errorf("Expected '%f', got '%f'.", expected, humanized)
		}
	}
}

func TestHumanizer_ParsePrefix_Incorrect(t *testing.T) {
	humanizer, err := New("en")
	if err != nil {
		t.Errorf("Humanizer creation failed with error: %s", err)
	}
	_, err = humanizer.ParsePrefix("no prefix at all")
	if err == nil {
		t.Error("Humanization succeeded where it should have failed.")
	}
	_, err = humanizer.ParsePrefix("15 flobbers")
	if err == nil {
		t.Error("Humanization succeeded where it should have failed.")
	}
}
