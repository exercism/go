package darts

import "math"

func Score(x, y float64) int {

	dart_location := math.Sqrt(x*x + y*y)

	switch {
	case dart_location <= 1.0:
		return 10
	case dart_location <= 5.0:
		return 5
	case dart_location <= 10.0:
		return 1
	default:
		return 0
	}
}
