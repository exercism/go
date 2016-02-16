package raindrops

import "strconv"

const testVersion = 2

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
