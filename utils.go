package humanize

import (
	"math/big"
	"strings"
)

// Helper functions.

// Power returning a big float.
func bigPow(x int, y int) *big.Float {
	one := new(big.Float).SetInt64(1)
	if y == 0 {
		return one
	}
	bigX := new(big.Float).SetInt64(int64(x))
	product := new(big.Float).Copy(bigX)
	// Always calculate positive power, inverse later.
	isNeg := false
	if y < 0 {
		isNeg = true
		y = -y
	}
	for i := 0; i < y-1; i++ {
		product = new(big.Float).Mul(product, bigX)
	}
	if isNeg {
		return new(big.Float).Quo(one, product)
	}
	return product
}

// Hack to get rid of trailing zeroes (while keeping the precision if necessary)
func trimZeroes(value string) string {
	if strings.ContainsRune(value, '.') {
		value = strings.TrimRight(value, "0")
		value = strings.TrimRight(value, ".")
	}
	return value
}
