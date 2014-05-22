package hexadecimal

import (
	"fmt"
)

// ToDecimal converts a hexadecimal string to decimal integer.
// If the string is invalid, it returns 0 and an error.
func ToDecimal(hex string) (int64, error) {
	decimal := int64(0)
	digitVal := int64(0)
	// for each character in the hex string
	for _, c := range hex {
		// get the value of the character
		switch {
		// numbers
		case '0' <= c && c <= '9':
			digitVal = int64(c - '0')
			// valid lowercase letters
		case 'a' <= c && c <= 'f':
			digitVal = int64(c - 'a' + 10)
			// valid uppercase letters
		case 'A' <= c && c <= 'F':
			digitVal = int64(c - 'A' + 10)
		default:
			// if the character is invalid, return 0
			return 0, fmt.Errorf("%s is an invalid hexadecimal string", hex)
		}
		// multiply the current number by 16 (left shift by 4) and add the digit
		decimal = decimal<<4 + digitVal
	}
	return decimal, nil
}
