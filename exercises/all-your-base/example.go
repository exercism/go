package allyourbase

import (
	"errors"
	"math"
)

// ConvertToBase converts given digits in input base to given output base
func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return nil, errors.New("input base must be >= 2")
	} else if outputBase < 2 {
		return nil, errors.New("output base must be >= 2")
	}
	for _, d := range inputDigits {
		if d < 0 || d >= inputBase {
			return nil, errors.New("all digits must satisfy 0 <= d < input base")
		}
	}

	total := toCanonicalBase(inputBase, inputDigits)
	return digitsInBase(total, outputBase), nil
}

func toCanonicalBase(base int, digits []int) int {
	var total int

	l := len(digits)
	for i := l - 1; i >= 0; i-- {
		total += digits[l-1-i] * int(math.Pow(float64(base), float64(i)))
	}

	return total
}

func reverse(a []int) []int {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	return a
}

func digitsInBase(total, base int) []int {
	var digits []int

	for total >= base {
		digits = append(digits, total%base)
		total /= base
	}
	digits = append(digits, total%base)

	return reverse(digits)
}
