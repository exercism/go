package luhn

import (
	"unicode"
)

//Given a number determine whether or not it is valid per the Luhn checksum formula.
func Valid(id string) bool {
	//custom index of non-space digits
	index := 0
	digitSum := 0
	for _, r := range reverseRunes(id) {
		if unicode.IsSpace(r) {
			continue
		}

		if !unicode.IsDigit(r) {
			return false
		}
		digitSum += luhnDigit(r, index)
		index++
	}

	if index < 2 {
		return false
	}

	return digitSum%10 == 0
}

// Reverse a string into a slice of runes.
func reverseRunes(s string) []rune {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return r
}

// double every **second** digit from the rightmost input position
// so *odd* numbered custom index since this is the reversed input
func luhnDigit(r rune, index int) int {
	//only ascii digits work here
	digit := int(r - '0')

	if index%2 == 1 {
		digit = digit * 2
		if digit > 9 {
			digit -= 9
		}
	}
	return digit
}
