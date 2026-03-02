package hamming

import "errors"

func Distance(a, b string) (d int, err error) {
	if len(b) != len(a) {
		return 0, errors.New("strings of unequal length")
	}
	for i := range len(a) {
		if a[i] != b[i] {
			d++
		}
	}
	return d, err
}
