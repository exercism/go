// Package hamming is the solution to http://exercism.io/exercises/go/hamming
package hamming

import (
	"errors"
	"unicode/utf8"
)

// Distance is the hamming distance between two strands of DNA
// strands must be the same length.
func Distance(a, b string) (d int, err error) {
	if utf8.RuneCountInString(b) != utf8.RuneCountInString(a) {
		return 0, errors.New("strings are of unequal length")
	}
	for i := range a {
		if a[i] != b[i] {
			d++
		}
	}
	return
}
