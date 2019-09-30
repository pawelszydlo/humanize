package humanize

import (
	"math/big"
	"testing"
)

func TestHumanizer_Prefix(t *testing.T) {
	humanizer, err := New("en")
	if err != nil {
		t.Errorf("Humanizer creation failed with error: %s", err)
	}

	cases := map[string]string{
		// Metric prefixes.
		"2.9k":    humanizer.SiPrefix(2854, 1, 1000, true),
		"23 mega": humanizer.SiPrefix(22843853, 0, 1000, false),
		"1.44M":   humanizer.SiPrefix(1440000, 2, 1000, true),
		"5.3µ":    humanizer.SiPrefix(0.00000534, 1, 100, true),
		"2345":    humanizer.SiPrefix(2345, 1, 10000, true),
		"1Y":      humanizer.SiPrefix(1000000000001000000000000, 1, 1000, true),
		// Too low threshold.
		"1": humanizer.SiPrefix(1, 1, 1, true),
		// Fast.
		"174.5k": humanizer.SiPrefixFast(174512),
		"28M":    humanizer.SiPrefixFast(28000000),
		"5.1m":   humanizer.SiPrefixFast(0.005123),
		"175":    humanizer.SiPrefixFast(175),
		"1k":     humanizer.SiPrefixFast(1024),
		// Bit prefixes.
		"2.8Ki":   humanizer.BitPrefix(2854, 1, 1000, true),
		"21 tebi": humanizer.BitPrefix(22823452343853, 0, 1000, false),
		"1.44Mi":  humanizer.BitPrefix(1509949, 2, 1000, true),
		// Fast bit prefixes.
		"1001":   humanizer.BitPrefixFast(1001), // Too small.
		"26.7Mi": humanizer.BitPrefixFast(28000000),
		"0.01":   humanizer.BitPrefixFast(0.005123), // prefixDef not found.
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

	// Huge cases need to be loaded from string.
	case5y, _ := new(big.Float).SetString("5000000000000000000000000")
	case736Yi, _ := new(big.Float).SetString("889769403236367072583745536")

	cases := map[string]*big.Float{
		// SI.
		"2.9k":       new(big.Float).SetInt64(2900),
		"13.":        new(big.Float).SetInt64(13),
		"0.5":        new(big.Float).SetFloat64(0.5),
		"13.5 kilo ": new(big.Float).SetInt64(13500),
		"20m":        new(big.Float).SetFloat64(0.020),
		" 20M ":      new(big.Float).SetInt64(20000000),
		"5yotta":     case5y,
		"15 µ":       new(big.Float).SetFloat64(0.000015),
		"3 y":        new(big.Float).SetFloat64(0.000000000000000000000003),
		// Bit.
		"3Mi":     new(big.Float).SetInt64(3145728),
		"50 tebi": new(big.Float).SetInt64(54975581388800),
		"0.5 Gi":  new(big.Float).SetInt64(536870912),
		"736 Yi":  case736Yi,
	}

	for input, expected := range cases {
		parsed, err := humanizer.ParsePrefix(input)
		if err != nil {
			t.Errorf("Error parsing '%s': %s", input, err)
		}

		// If string representation is the same we are close enough.
		if parsed.String() != expected.String() {
			t.Errorf("Expected '%f', got '%f'.", expected, parsed)
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
