package grains

import "math"

func Square(number int) uint64 {
	return uint64(math.Exp2(float64(number - 1)))
}

func Total() uint64 {
	var total uint64
	for i := 1; i <= 64; i++ {
		total += Square(i)
	}
	return total
}
