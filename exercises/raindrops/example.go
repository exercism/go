package raindrops

import "strconv"

// Convert converts a number to a string, the contents
// of which depend on the number's factors.
func Convert(number int) string {
	s := ""
	if number%3 == 0 {
		s += "Pling"
	}
	if number%5 == 0 {
		s += "Plang"
	}
	if number%7 == 0 {
		s += "Plong"
	}
	if len(s) == 0 {
		s = strconv.Itoa(number)
	}
	return s
}

// This test versioning is specific to Exercism,
// you don't need to worry about this now.
const testVersion = 3
