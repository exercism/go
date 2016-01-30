package triangle

import "math"

type Kind string

const TestVersion = 1

const (
	Equ Kind = "equilateral"
	Iso = "isosceles"
	Sca = "scalene"
	NaT = "not a triangle"
)

func KindFromSides(a, b, c float64) Kind {
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return NaT
	}
	if a > b {
		a, b = b, a
	}
	if b > c {
		b, c = c, b
	}
	if a > b {
		a, b = b, a
	}
	// sides are now sorted
	switch {
	case math.IsInf(c, 1): // largest side is +Inf, guards against (3, +Inf, +Inf)
		return NaT
	case a <= 0:
		return NaT
	case a+b < c: // triangle inequality
		return NaT
	case a == b:
		if b == c {
			return Equ
		}
		return Iso
	case a == c || b == c:
		return Iso
	}
	return Sca
}
