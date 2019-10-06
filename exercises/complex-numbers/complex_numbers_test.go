package complex

import (
	"math"
	"math/cmplx"
	"testing"
)

func TestImaginaryUnit(t *testing.T) {
	expected := 1i
	imgUnit := NewComplex(0, 1)
	if actual := complex(imgUnit.Real(), imgUnit.Imag()); expected != actual {
		t.Fatalf("Expecting the imaginary unit to be %g but was %g", expected, actual)
	}
}

func testFuncToFloat64(t *testing.T, resName string, fActual func(Complex) float64, fExpected func(complex128) float64) {
	for _, test := range testDataForUnOp {
		n := complex(test.r, test.i)
		expected := fExpected(n)
		actual := fActual(NewComplex(test.r, test.i))
		if compareFloat64(expected, actual) {
			t.Fatalf("Expecting the %s of %g to be %g but was %g", resName, n, expected, actual)
		}
	}
}

func TestRealPart(t *testing.T) {
	testFuncToFloat64(t, "real part",
		func(c Complex) float64 { return c.Real() },
		func(c complex128) float64 { return real(c) })
}

func TestImagPart(t *testing.T) {
	testFuncToFloat64(t, "imaginary part",
		func(c Complex) float64 { return c.Imag() },
		func(c complex128) float64 { return imag(c) })
}

func TestComplexAbs(t *testing.T) {
	testFuncToFloat64(t, "absolute value", Abs, cmplx.Abs)
}

func compareFloat64(expected, actual float64) bool {
	diff := math.Abs(expected - actual)
	const delta = 1E-10
	return diff > delta
}

func compareComplex(expected, actual complex128) bool {
	return compareFloat64(real(expected), real(actual)) ||
		compareFloat64(imag(expected), imag(actual))
}

func testUnaryOp(t *testing.T, resName string, opActual func(Complex) Complex, opExpected func(complex128) complex128) {
	for _, test := range testDataForUnOp {
		n := complex(test.r, test.i)
		expected := opExpected(n)
		result := opActual(NewComplex(test.r, test.i))
		actual := complex(result.Real(), result.Imag())
		if compareComplex(expected, actual) {
			t.Fatalf("Expecting the %s of %g to be %g but was %g", resName, n, expected, actual)
		}
	}
}

func TestComplexConj(t *testing.T) {
	testUnaryOp(t, "complex conjugate", Conj, cmplx.Conj)
}

func TestExpComplex(t *testing.T) {
	testUnaryOp(t, "exponential", Exp, cmplx.Exp)
}

func testBinaryOperation(t *testing.T, resName string, opActual func(Complex, Complex) Complex, opExpected func(complex128, complex128) complex128) {
	for _, test := range testDataForBinOp {
		a := complex(test.aReal, test.aImag)
		b := complex(test.bReal, test.bImag)
		expected := opExpected(a, b)
		res := opActual(NewComplex(test.aReal, test.aImag), NewComplex(test.bReal, test.bImag))
		actual := complex(res.Real(), res.Imag())
		if compareComplex(expected, actual) {
			t.Fatalf("Expecting the %s of %g and %g to be %g but was %g", resName, a, b, expected, actual)
		}
	}
}

func TestAddComplex(t *testing.T) {
	testBinaryOperation(t, "sum", Add, func(a, b complex128) complex128 { return a + b })
}

func TestSubtractComplex(t *testing.T) {
	testBinaryOperation(t, "difference", Sub, func(a, b complex128) complex128 { return a - b })
}

func TestMultiplyComplex(t *testing.T) {
	testBinaryOperation(t, "product", Mult, func(a, b complex128) complex128 { return a * b })
}

func TestDivisionComplex(t *testing.T) {
	testBinaryOperation(t, "quotient", Div, func(a, b complex128) complex128 { return a / b })
}
