package triangle

type Kind string

const (
	Equ Kind = "equilateral"
	Iso Kind = "isosceles"
	Sca Kind = "scalene"
	NaT Kind = "not a triangle"
)

func KindFromSides(a, b, c float64) Kind {
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
	case a <= 0:
		return NaT
	case a+b < c: // triangle inequality
		return NaT
	case a == c:
		return Equ
	case a == b || b == c:
		return Iso
	}
	return Sca
}
