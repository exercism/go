package complexnumbers

import (
	"math"
)

func Addition(c1 complex128, c2 complex128) complex128 {
	realPart := real(c1) + real(c2)
	imagPart := imag(c1) + imag(c2)
	return complex(realPart, imagPart)
}

func Subtraction(c1 complex128, c2 complex128) complex128 {
	realPart := real(c1) - real(c2)
	imagPart := imag(c1) - imag(c2)
	return complex(realPart, imagPart)
}

func Multiply(c1 complex128, c2 complex128) complex128 {
	realPart := (real(c1) * real(c2)) - (imag(c1) * imag(c2))
	imagPart := (imag(c1) * real(c2)) + (real(c1) * imag(c2))
	return complex(realPart, imagPart)
}

func Division(c1 complex128, c2 complex128) complex128 {
	realPart := ((real(c1) * real(c2)) + (imag(c1) * imag(c2))) / ((real(c2) * real(c2)) + (imag(c2) * imag(c2)))
	imagPart := ((imag(c1) * real(c2)) - (real(c1) * imag(c2))) / ((real(c2) * real(c2)) + (imag(c2) * imag(c2)))
	return complex(realPart, imagPart)
}

func Conjugate(c complex128) complex128 {
	return complex(real(c), -imag(c))
}

func Absolute(c complex128) float64 {
	return math.Sqrt((real(c) * real(c)) + (imag(c) * imag(c)))
}
func Exponent(c complex128) complex128 {
	realPart := math.Exp(real(c)) * math.Cos(imag(c))
	imagPart := math.Exp(real(c)) * math.Sin(imag(c))
	return complex(realPart, imagPart)
}
