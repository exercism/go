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

// Reduce simplifies a Rational, eg changing Rational{4, 2} into Rational{2, 1}.
func (r Rational) Reduce() Rational {
	gcd := abs(gcd(r.numerator, r.denominator))
	if r.denominator < 0 {
		gcd *= -1
	}
	return Rational{r.numerator / gcd, r.denominator / gcd}
}

func (r Rational) Add(s Rational) Rational {
	numerator := r.numerator*s.denominator + s.numerator*r.denominator
	denominator := r.denominator * s.denominator
	return Rational{numerator, denominator}.Reduce()
}

func (r Rational) Sub(s Rational) Rational {
	numerator := r.numerator*s.denominator - s.numerator*r.denominator
	denominator := r.denominator * s.denominator
	return Rational{numerator, denominator}.Reduce()
}

func (r Rational) Mul(s Rational) Rational {
	return Rational{r.numerator * s.numerator, r.denominator * s.denominator}.Reduce()
}

func (r Rational) Div(s Rational) Rational {
	return Rational{r.numerator * s.denominator, s.numerator * r.denominator}.Reduce()
}

func (r Rational) Abs() Rational {
	return Rational{abs(r.numerator), abs(r.denominator)}.Reduce()
}

// Compute r ^ power, a rational raised to an int exponent.
func (r Rational) Exprational(power int) Rational {
	if power >= 0 {
		return Rational{pow(r.numerator, power), pow(r.denominator, power)}.Reduce()
	}
	power = -power
	return Rational{pow(r.denominator, power), pow(r.numerator, power)}.Reduce()
}

// Compute base ^ r, an int raised to a rational.
func (r Rational) Expreal(base int) float64 {
	return math.Pow(math.Pow(float64(base), float64(r.numerator)), 1/float64(r.denominator))
}
