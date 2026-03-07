package rationalnumbers

type Rational struct {
	numerator, denominator int
}

// Reduce simplifies a Rational, eg changing Rational{4, 2} into Rational{2, 1}.
func Reduce(r Rational) Rational {
	panic("Please implement the Reduce function")
}

func Add(a, b Rational) Rational {
	panic("Please implement the Add function")
}

func Sub(a, b Rational) Rational {
	panic("Please implement the Sub function")
}

func Mul(a, b Rational) Rational {
	panic("Please implement the Mul function")
}

func Div(a, b Rational) Rational {
	panic("Please implement the Div function")
}

func Abs(r Rational) Rational {
	panic("Please implement the Abs function")
}

func Exprational(base Rational, power int) Rational {
	panic("Please implement the Exprational function")
}

func Expreal(base int, power Rational) float64 {
	panic("Please implement the Expreal function")
}
