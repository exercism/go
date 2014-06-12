package cryptosquare

import (
	"math"
	"strings"
)

func norm(r rune) rune {
	switch {
	case r >= 'a' && r <= 'z' || r >= '0' && r <= '9':
		return r
	case r >= 'A' && r <= 'Z':
		return r + 'a' - 'A'
	}
	return -1
}

func Encode(pt string) string {
	pt = strings.Map(norm, pt)
	cols := int(math.Ceil(math.Sqrt(float64(len(pt)))))
	b := make([]byte, len(pt)+(len(pt)-1)/5)
	px := 0
	for bx := range b {
		if (bx+1)%6 == 0 {
			b[bx] = ' '
		} else {
			b[bx] = pt[px]
			px += cols
			if px >= len(pt) {
				px = (px + 1) % cols
			}
		}
	}
	return string(b)
}
