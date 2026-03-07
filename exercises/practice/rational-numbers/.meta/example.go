package rationalnumbers

import "math"

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// gcd calculates the Greatest Common Divisor of two integers using the Euclidean algorithm.
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b // This swaps a and b and updates b with the remainder
	}
	return a
}

type Rational struct {
	numerator, denominator int
}

func Reduce(r Rational) Rational {
	gcd := abs(gcd(r.numerator, r.denominator))
	if r.denominator < 0 {
		gcd *= -1
	}
	return Rational{r.numerator / gcd, r.denominator / gcd}
}

func Add(a, b Rational) Rational {
	numerator := a.numerator*b.denominator + b.numerator*a.denominator
	denominator := a.denominator * b.denominator
	return Reduce(Rational{numerator, denominator})
}

func Sub(a, b Rational) Rational {
	numerator := a.numerator*b.denominator - b.numerator*a.denominator
	denominator := a.denominator * b.denominator
	return Reduce(Rational{numerator, denominator})
}

func Mul(a, b Rational) Rational {
	return Reduce(Rational{a.numerator * b.numerator, a.denominator * b.denominator})
}

func Div(a, b Rational) Rational {
	return Reduce(Rational{a.numerator * b.denominator, b.numerator * a.denominator})
}

func Abs(r Rational) Rational {
	return Reduce(Rational{abs(r.numerator), abs(r.denominator)})
}

func Exprational(base Rational, power int) Rational {
	if power >= 0 {
		return Reduce(Rational{pow(base.numerator, power), pow(base.denominator, power)})
	}
	power = -power
	return Reduce(Rational{pow(base.denominator, power), pow(base.numerator, power)})
}

func Expreal(base int, power Rational) float64 {
	return math.Pow(math.Pow(float64(base), float64(power.numerator)), 1/float64(power.denominator))
}
