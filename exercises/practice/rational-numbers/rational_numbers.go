package rationalnumbers

type Rational struct {
	numerator, denominator int
}

// Reduce simplifies a Rational, eg changing Rational{4, 2} into Rational{2, 1}.
func (r Rational) Reduce() Rational {
	panic("Please implement the Reduce function")
}

func (r Rational) Add(s Rational) Rational {
	panic("Please implement the Add function")
}

func (r Rational) Sub(s Rational) Rational {
	panic("Please implement the Sub function")
}

func (r Rational) Mul(s Rational) Rational {
	panic("Please implement the Mul function")
}

func (r Rational) Div(s Rational) Rational {
	panic("Please implement the Div function")
}

func (r Rational) Abs() Rational {
	panic("Please implement the Abs function")
}

// Compute r ^ power, a rational raised to an int exponent.
func (r Rational) Exprational(power int) Rational {
	panic("Please implement the Exprational function")
}

// Compute base ^ r, an int raised to a rational.
func (r Rational) Expreal(base int) float64 {
	panic("Please implement the Expreal function")
}
