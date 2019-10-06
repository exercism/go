package complex

import "math"

// Complex is a struct that represents a complex number
// with its real and imaginary parts
type Complex struct {
	real, imag float64
}

// NewComplex returns a new complex number given its real
// and imaginary parts
func NewComplex(real, imag float64) Complex {
	return Complex{real, imag}
}

// Real returns the real part of the complex number
func (c Complex) Real() float64 { return c.real }

// Imag returns the imaginary part of the complex number
func (c Complex) Imag() float64 { return c.imag }

// Add returns the sum of two complex numbers
func Add(c1, c2 Complex) Complex {
	return NewComplex(c1.Real()+c2.Real(), c1.Imag()+c2.Imag())
}

// Sub returns the subtraction of two complex numbers
func Sub(c1, c2 Complex) Complex {
	return NewComplex(c1.Real()-c2.Real(), c1.Imag()-c2.Imag())
}

// Mult returns the multiplication of two complex numbers
func Mult(c1, c2 Complex) Complex {
	r := c1.Real()*c2.Real() - c1.Imag()*c2.Imag()
	i := c1.Imag()*c2.Real() + c1.Real()*c2.Imag()
	return NewComplex(r, i)
}

// Div returns the quotient of two complex numbers
func Div(c1, c2 Complex) Complex {
	f := c2.Real()*c2.Real() + c2.Imag()*c2.Imag()
	r := (c1.Real()*c2.Real() + c1.Imag()*c2.Imag()) / f
	i := (c1.Imag()*c2.Real() - c1.Real()*c2.Imag()) / f
	return NewComplex(r, i)
}

// Conj returns the complex conjugate of the input complex number
func Conj(c Complex) Complex {
	return NewComplex(c.Real(), -c.Imag())
}

// Abs returns the absolute value of the input complex number
func Abs(c Complex) float64 {
	return math.Sqrt(c.Real()*c.Real() + c.Imag()*c.Imag())
}

// Exp returns 'e' raised to the power of the input complex number
func Exp(c Complex) Complex {
	a := math.Exp(c.Real())
	bR := math.Cos(c.Imag())
	bI := math.Sin(c.Imag())
	return NewComplex(a*bR, a*bI)
}
