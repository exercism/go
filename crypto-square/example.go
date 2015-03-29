package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

const TestVersion = 1

func norm(r rune) rune {
	if unicode.IsLetter(r) || unicode.IsDigit(r) {
		return unicode.ToLower(r)
	} else {
		return -1
	}
}

func Encode(pt string) string {
	pt = strings.Map(norm, pt)
	numCols := int(math.Ceil(math.Sqrt(float64(len(pt)))))
	cols := make([]string, numCols)
	for i, r := range pt {
		cols[i%numCols] += string(r)
	}
	return strings.Join(cols, " ")
}
