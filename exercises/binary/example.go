package binary

import (
	"fmt"
)

const testVersion = 1

// ParseBinary converts a binary string to an integer value
func ParseBinary(bin string) (int, error) {
	val := 0
	for _, digit := range bin {
		switch digit {
		case '1':
			// multiply val by 2 and add 1
			val = (val << 1) + 1
		case '0':
			// multiply val by 2 and add 0
			val = (val << 1) + 0
		default:
			// if the character was not 1 or 0, it must be invalid
			return 0, fmt.Errorf("unexpected rune '%c'", digit)
		}
	}
	return val, nil
}
