package complexnumbers

// Source: exercism/problem-specifications
// Commit: 24a7bfa Add exponential resulting in a number with real and imaginary part (#2052)

type complexNumber struct {
	a float64
	b float64
}

var realTestCases = []struct {
	description string
	in          complexNumber
	want        float64
}{
	{
		description: "Real part of a purely real number",
		in: complexNumber{
			a: 1.000000,
			b: 0.000000,
		},
		want: 1.000000,
	},
	{
		description: "Real part of a purely imaginary number",
		in: complexNumber{
			a: 0.000000,
			b: 1.000000,
		},
		want: 0.000000,
	},
	{
		description: "Real part of a number with real and imaginary part",
		in: complexNumber{
			a: 1.000000,
			b: 2.000000,
		},
		want: 1.000000,
	},
}

var imaginaryTestCases = []struct {
	description string
	in          complexNumber
	want        float64
}{
	{
		description: "Imaginary part of a purely real number",
		in: complexNumber{
			a: 1.000000,
			b: 0.000000,
		},
		want: 0.000000,
	},
	{
		description: "Imaginary part of a purely imaginary number",
		in: complexNumber{
			a: 0.000000,
			b: 1.000000,
		},
		want: 1.000000,
	},
	{
		description: "Imaginary part of a number with real and imaginary part",
		in: complexNumber{
			a: 1.000000,
			b: 2.000000,
		},
		want: 2.000000,
	},
}

var addTestCases = []struct {
	description string
	n1          complexNumber
	n2          complexNumber
	want        complexNumber
}{
	{
		description: "Add purely real numbers",
		n1: complexNumber{
			a: 1.000000,
			b: 0.000000,
		},
		n2: complexNumber{
			a: 2.000000,
			b: 0.000000,
		},
		want: complexNumber{
			a: 3.000000,
			b: 0.000000,
		},
	},
	{
		description: "Add purely imaginary numbers",
		n1: complexNumber{
			a: 0.000000,
			b: 1.000000,
		},
		n2: complexNumber{
			a: 0.000000,
			b: 2.000000,
		},
		want: complexNumber{
			a: 0.000000,
			b: 3.000000,
		},
	},
	{
		description: "Add numbers with real and imaginary part",
		n1: complexNumber{
			a: 1.000000,
			b: 2.000000,
		},
		n2: complexNumber{
			a: 3.000000,
			b: 4.000000,
		},
		want: complexNumber{
			a: 4.000000,
			b: 6.000000,
		},
	},
	{
		description: "Add real number to complex number",
		n1: complexNumber{
			a: 1.000000,
			b: 2.000000,
		},
		n2: complexNumber{
			a: 5.000000,
			b: 0.000000,
		},
		want: complexNumber{
			a: 6.000000,
			b: 2.000000,
		},
	},
	{
		description: "Add complex number to real number",
		n1: complexNumber{
			a: 5.000000,
			b: 0.000000,
		},
		n2: complexNumber{
			a: 1.000000,
			b: 2.000000,
		},
		want: complexNumber{
			a: 6.000000,
			b: 2.000000,
		},
	},
}

var subtractTestCases = []struct {
	description string
	n1          complexNumber
	n2          complexNumber
	want        complexNumber
}{
	{
		description: "Subtract purely real numbers",
		n1: complexNumber{
			a: 1.000000,
			b: 0.000000,
		},
		n2: complexNumber{
			a: 2.000000,
			b: 0.000000,
		},
		want: complexNumber{
			a: -1.000000,
			b: 0.000000,
		},
	},
	{
		description: "Subtract purely imaginary numbers",
		n1: complexNumber{
			a: 0.000000,
			b: 1.000000,
		},
		n2: complexNumber{
			a: 0.000000,
			b: 2.000000,
		},
		want: complexNumber{
			a: 0.000000,
			b: -1.000000,
		},
	},
	{
		description: "Subtract numbers with real and imaginary part",
		n1: complexNumber{
			a: 1.000000,
			b: 2.000000,
		},
		n2: complexNumber{
			a: 3.000000,
			b: 4.000000,
		},
		want: complexNumber{
			a: -2.000000,
			b: -2.000000,
		},
	},
	{
		description: "Subtract real number from complex number",
		n1: complexNumber{
			a: 5.000000,
			b: 7.000000,
		},
		n2: complexNumber{
			a: 4.000000,
			b: 0.000000,
		},
		want: complexNumber{
			a: 1.000000,
			b: 7.000000,
		},
	},
	{
		description: "Subtract complex number from real number",
		n1: complexNumber{
			a: 4.000000,
			b: 0.000000,
		},
		n2: complexNumber{
			a: 5.000000,
			b: 7.000000,
		},
		want: complexNumber{
			a: -1.000000,
			b: -7.000000,
		},
	},
}

var divideTestCases = []struct {
	description string
	n1          complexNumber
	n2          complexNumber
	want        complexNumber
}{
	{
		description: "Divide purely real numbers",
		n1: complexNumber{
			a: 1.000000,
			b: 0.000000,
		},
		n2: complexNumber{
			a: 2.000000,
			b: 0.000000,
		},
		want: complexNumber{
			a: 0.500000,
			b: 0.000000,
		},
	},
	{
		description: "Divide purely imaginary numbers",
		n1: complexNumber{
			a: 0.000000,
			b: 1.000000,
		},
		n2: complexNumber{
			a: 0.000000,
			b: 2.000000,
		},
		want: complexNumber{
			a: 0.500000,
			b: 0.000000,
		},
	},
	{
		description: "Divide numbers with real and imaginary part",
		n1: complexNumber{
			a: 1.000000,
			b: 2.000000,
		},
		n2: complexNumber{
			a: 3.000000,
			b: 4.000000,
		},
		want: complexNumber{
			a: 0.440000,
			b: 0.080000,
		},
	},
	{
		description: "Divide complex number by real number",
		n1: complexNumber{
			a: 10.000000,
			b: 100.000000,
		},
		n2: complexNumber{
			a: 10.000000,
			b: 0.000000,
		},
		want: complexNumber{
			a: 1.000000,
			b: 10.000000,
		},
	},
	{
		description: "Divide real number by complex number",
		n1: complexNumber{
			a: 5.000000,
			b: 0.000000,
		},
		n2: complexNumber{
			a: 1.000000,
			b: 1.000000,
		},
		want: complexNumber{
			a: 2.500000,
			b: -2.500000,
		},
	},
}

var multiplyTestCases = []struct {
	description string
	n1          complexNumber
	n2          *complexNumber // if n2 is nil it is a multiplication with the factor
	factor      float64
	want        complexNumber
}{
	{
		description: "Imaginary unit",
		n1: complexNumber{
			a: 0.000000,
			b: 1.000000,
		},
		n2: &complexNumber{
			a: 0.000000,
			b: 1.000000,
		},
		factor: 0.000000,
		want: complexNumber{
			a: -1.000000,
			b: 0.000000,
		},
	},
	{
		description: "Multiply purely real numbers",
		n1: complexNumber{
			a: 1.000000,
			b: 0.000000,
		},
		n2: &complexNumber{
			a: 2.000000,
			b: 0.000000,
		},
		factor: 0.000000,
		want: complexNumber{
			a: 2.000000,
			b: 0.000000,
		},
	},
	{
		description: "Multiply purely imaginary numbers",
		n1: complexNumber{
			a: 0.000000,
			b: 1.000000,
		},
		n2: &complexNumber{
			a: 0.000000,
			b: 2.000000,
		},
		factor: 0.000000,
		want: complexNumber{
			a: -2.000000,
			b: 0.000000,
		},
	},
	{
		description: "Multiply numbers with real and imaginary part",
		n1: complexNumber{
			a: 1.000000,
			b: 2.000000,
		},
		n2: &complexNumber{
			a: 3.000000,
			b: 4.000000,
		},
		factor: 0.000000,
		want: complexNumber{
			a: -5.000000,
			b: 10.000000,
		},
	},
	{
		description: "Multiply complex number by real number",
		n1: complexNumber{
			a: 2.000000,
			b: 5.000000,
		},

		n2: nil,

		factor: 5.000000,
		want: complexNumber{
			a: 10.000000,
			b: 25.000000,
		},
	},
	{
		description: "Multiply real number by complex number",
		n1: complexNumber{
			a: 2.000000,
			b: 5.000000,
		},

		n2: nil,

		factor: 5.000000,
		want: complexNumber{
			a: 10.000000,
			b: 25.000000,
		},
	},
}

var conjugateTestCases = []struct {
	description string
	in          complexNumber
	want        complexNumber
}{
	{
		description: "Conjugate a purely real number",
		in: complexNumber{
			a: 5.000000,
			b: 0.000000,
		},
		want: complexNumber{
			a: 5.000000,
			b: 0.000000,
		},
	},
	{
		description: "Conjugate a purely imaginary number",
		in: complexNumber{
			a: 0.000000,
			b: 5.000000,
		},
		want: complexNumber{
			a: 0.000000,
			b: -5.000000,
		},
	},
	{
		description: "Conjugate a number with real and imaginary part",
		in: complexNumber{
			a: 1.000000,
			b: 1.000000,
		},
		want: complexNumber{
			a: 1.000000,
			b: -1.000000,
		},
	},
}

var absTestCases = []struct {
	description string
	in          complexNumber
	want        float64
}{
	{
		description: "Absolute value of a positive purely real number",
		in: complexNumber{
			a: 5.000000,
			b: 0.000000,
		},
		want: 5.000000,
	},
	{
		description: "Absolute value of a negative purely real number",
		in: complexNumber{
			a: -5.000000,
			b: 0.000000,
		},
		want: 5.000000,
	},
	{
		description: "Absolute value of a purely imaginary number with positive imaginary part",
		in: complexNumber{
			a: 0.000000,
			b: 5.000000,
		},
		want: 5.000000,
	},
	{
		description: "Absolute value of a purely imaginary number with negative imaginary part",
		in: complexNumber{
			a: 0.000000,
			b: -5.000000,
		},
		want: 5.000000,
	},
	{
		description: "Absolute value of a number with real and imaginary part",
		in: complexNumber{
			a: 3.000000,
			b: 4.000000,
		},
		want: 5.000000,
	},
}

var expTestCases = []struct {
	description string
	in          complexNumber
	want        complexNumber
}{
	{
		description: "Euler's identity/formula",
		in: complexNumber{
			a: 0.000000,
			b: 3.141593,
		},
		want: complexNumber{
			a: -1.000000,
			b: 0.000000,
		},
	},
	{
		description: "Exponential of 0",
		in: complexNumber{
			a: 0.000000,
			b: 0.000000,
		},
		want: complexNumber{
			a: 1.000000,
			b: 0.000000,
		},
	},
	{
		description: "Exponential of a purely real number",
		in: complexNumber{
			a: 1.000000,
			b: 0.000000,
		},
		want: complexNumber{
			a: 2.718282,
			b: 0.000000,
		},
	},
	{
		description: "Exponential of a number with real and imaginary part",
		in: complexNumber{
			a: 0.693147,
			b: 3.141593,
		},
		want: complexNumber{
			a: -2.000000,
			b: 0.000000,
		},
	},
	{
		description: "Exponential resulting in a number with real and imaginary part",
		in: complexNumber{
			a: 0.346574,
			b: 0.785398,
		},
		want: complexNumber{
			a: 1.000000,
			b: 1.000000,
		},
	},
}
