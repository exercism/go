package palindrome

import (
	"fmt"
	"strconv"
)

func isPal(x int) bool {
	s := strconv.Itoa(x)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

type Product struct {
	Product        int
	Factorizations [][2]int
}

func Products(fmin, fmax int) (pmin, pmax Product, err error) {
	if fmin > fmax {
		err = fmt.Errorf("fmin > fmax: %d > %d", fmin, fmax)
		return
	}
	for x := fmin; x <= fmax; x++ {
		for y := x; y <= fmax; y++ {
			p := x * y
			if !isPal(p) {
				continue
			}
			compare := func(current *Product, better bool) {
				switch {
				case current.Factorizations == nil || better:
					*current = Product{p, [][2]int{{x, y}}}
				case p == current.Product:
					current.Factorizations =
						append(current.Factorizations, [2]int{x, y})
				}
			}
			compare(&pmin, p < pmin.Product)
			compare(&pmax, p > pmax.Product)
		}
	}
	if len(pmin.Factorizations) == 0 {
		err = fmt.Errorf("no palindromes in range [%d, %d]", fmin, fmax)
	}
	return
}
