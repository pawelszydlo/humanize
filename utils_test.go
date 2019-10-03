package humanize

import (
	"math/big"
	"testing"
)

func TestBigPow(t *testing.T) {
	huge, _ := new(big.Float).SetString("12063348350820368238715343")
	cases := map[*big.Float][]int{
		new(big.Float).SetFloat64(1):          {1234, 0},
		new(big.Float).SetFloat64(1000000):    {10, 6},
		huge:                                  {47, 15},
		new(big.Float).SetFloat64(0.00000256): {5, -8},
		new(big.Float).SetFloat64(0):          {0, 9},
	}

	for expected, arguments := range cases {
		result := bigPow(arguments[0], arguments[1])
		if result.String() != expected.String() {
			t.Errorf("Expected '%s', got '%s'.", expected.String(), result.String())
		}
	}
}

func TestTrimZeros(t *testing.T) {
	cases := map[string]string{
		"3.30000":    "3.3",
		"7seven":     "7seven",
		"not.number": "not.number",
		"345.":       "345",
		"23.00000":   "23",
	}

	for input, expected := range cases {
		result := trimZeroes(input)
		if result != expected {
			t.Errorf("Expected '%s', got '%s'.", expected, result)
		}
	}
}
