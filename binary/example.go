package binary

import (
	"fmt"
)

// ParseBinary converts a binary string to a decimal
func ParseBinary(bin string) (int, error) {
	decimal := 0
	for _, digit := range bin {
		switch digit {
		case '1':
			// multiply decimal by 2 and add 1
			decimal = (decimal << 1) + 1
		case '0':
			// multiply decimal by 2 and add 0
			decimal = (decimal << 1) + 0
		default:
			// if the character was not 1 or 0, it must be invalid
			return 0, fmt.Errorf("unexpected rune '%c'", bin)
		}
	}
	return decimal, nil
}
