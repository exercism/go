package hamming

import "errors"

const TestVersion = 1

func Distance(a, b string) (d int, err error) {
	if len(b) != len(a) {
		return 0, errors.New("strings of unequal length")
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			d++
		}
	}
	return
}
