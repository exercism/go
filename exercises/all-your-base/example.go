package allyourbase

import (
	"errors"
	"math"
)

var (
	ErrInvalidBase  = errors.New("invalid base given")
	ErrInvalidDigit = errors.New("invalid digit given")
)

func ConvertToBase(inputBase uint64, inputDigits []uint64, outputBase uint64) ([]uint64, error) {
	if inputBase < 2 || outputBase < 2 {
		return nil, ErrInvalidBase
	}

	for _, d := range inputDigits {
		if d >= inputBase {
			return nil, ErrInvalidDigit
		}
	}

	total := toCanonicalBase(inputBase, inputDigits)
	return digitsInBase(total, outputBase), nil
}

func toCanonicalBase(inputBase uint64, inputDigits []uint64) uint64 {
	var total uint64

	digits := len(inputDigits)
	for i := digits - 1; i >= 0; i-- {
		total += inputDigits[digits-1-i] * uint64(math.Pow(float64(inputBase), float64(i)))
	}

	return total
}

func reverse(digits []uint64) []uint64 {
	var output = make([]uint64, len(digits))

	for i, d := range digits {
		output[len(digits)-1-i] = d
	}

	return output
}

func digitsInBase(total uint64, outputBase uint64) []uint64 {
	var digits []uint64
	for total >= outputBase {
		digits = append(digits, total%outputBase)
		total /= outputBase
	}
	digits = append(digits, total%outputBase)

	return reverse(digits)
}
