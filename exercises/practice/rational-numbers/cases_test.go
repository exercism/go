package rationalnumbers

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 4bcdfc3 rational-numbers: test to reduce abs value (#1938)

var testCasesAbs = []struct {
	description string
	input       Rational
	expected    Rational
}{
	{
		description: "Absolute value of a positive rational number",
		input:       Rational{1, 2},
		expected:    Rational{1, 2},
	},
	{
		description: "Absolute value of a positive rational number with negative numerator and denominator",
		input:       Rational{-1, -2},
		expected:    Rational{1, 2},
	},
	{
		description: "Absolute value of a negative rational number",
		input:       Rational{-1, 2},
		expected:    Rational{1, 2},
	},
	{
		description: "Absolute value of a negative rational number with negative denominator",
		input:       Rational{1, -2},
		expected:    Rational{1, 2},
	},
	{
		description: "Absolute value of zero",
		input:       Rational{0, 1},
		expected:    Rational{0, 1},
	},
	{
		description: "Absolute value of a rational number is reduced to lowest terms",
		input:       Rational{2, 4},
		expected:    Rational{1, 2},
	},
}

var testCasesAdd = []struct {
	description string
	input       [2]Rational
	expected    Rational
}{
	{
		description: "Add two positive rational numbers",
		input:       [2]Rational{{1, 2}, {2, 3}},
		expected:    Rational{7, 6},
	},
	{
		description: "Add a positive rational number and a negative rational number",
		input:       [2]Rational{{1, 2}, {-2, 3}},
		expected:    Rational{-1, 6},
	},
	{
		description: "Add two negative rational numbers",
		input:       [2]Rational{{-1, 2}, {-2, 3}},
		expected:    Rational{-7, 6},
	},
	{
		description: "Add a rational number to its additive inverse",
		input:       [2]Rational{{1, 2}, {-1, 2}},
		expected:    Rational{0, 1},
	},
}

var testCasesSub = []struct {
	description string
	input       [2]Rational
	expected    Rational
}{
	{
		description: "Subtract two positive rational numbers",
		input:       [2]Rational{{1, 2}, {2, 3}},
		expected:    Rational{-1, 6},
	},
	{
		description: "Subtract a positive rational number and a negative rational number",
		input:       [2]Rational{{1, 2}, {-2, 3}},
		expected:    Rational{7, 6},
	},
	{
		description: "Subtract two negative rational numbers",
		input:       [2]Rational{{-1, 2}, {-2, 3}},
		expected:    Rational{1, 6},
	},
	{
		description: "Subtract a rational number from itself",
		input:       [2]Rational{{1, 2}, {1, 2}},
		expected:    Rational{0, 1},
	},
}

var testCasesMul = []struct {
	description string
	input       [2]Rational
	expected    Rational
}{
	{
		description: "Multiply two positive rational numbers",
		input:       [2]Rational{{1, 2}, {2, 3}},
		expected:    Rational{1, 3},
	},
	{
		description: "Multiply a negative rational number by a positive rational number",
		input:       [2]Rational{{-1, 2}, {2, 3}},
		expected:    Rational{-1, 3},
	},
	{
		description: "Multiply two negative rational numbers",
		input:       [2]Rational{{-1, 2}, {-2, 3}},
		expected:    Rational{1, 3},
	},
	{
		description: "Multiply a rational number by its reciprocal",
		input:       [2]Rational{{1, 2}, {2, 1}},
		expected:    Rational{1, 1},
	},
	{
		description: "Multiply a rational number by 1",
		input:       [2]Rational{{1, 2}, {1, 1}},
		expected:    Rational{1, 2},
	},
	{
		description: "Multiply a rational number by 0",
		input:       [2]Rational{{1, 2}, {0, 1}},
		expected:    Rational{0, 1},
	},
}

var testCasesDiv = []struct {
	description string
	input       [2]Rational
	expected    Rational
}{
	{
		description: "Divide two positive rational numbers",
		input:       [2]Rational{{1, 2}, {2, 3}},
		expected:    Rational{3, 4},
	},
	{
		description: "Divide a positive rational number by a negative rational number",
		input:       [2]Rational{{1, 2}, {-2, 3}},
		expected:    Rational{-3, 4},
	},
	{
		description: "Divide two negative rational numbers",
		input:       [2]Rational{{-1, 2}, {-2, 3}},
		expected:    Rational{3, 4},
	},
	{
		description: "Divide a rational number by 1",
		input:       [2]Rational{{1, 2}, {1, 1}},
		expected:    Rational{1, 2},
	},
}

var testCasesExprational = []struct {
	description string
	inputR      Rational
	inputInt    int
	expected    Rational
}{
	{
		description: "Raise a positive rational number to a positive integer power",
		inputR:      Rational{1, 2},
		inputInt:    3,
		expected:    Rational{1, 8},
	},
	{
		description: "Raise a negative rational number to a positive integer power",
		inputR:      Rational{-1, 2},
		inputInt:    3,
		expected:    Rational{-1, 8},
	},
	{
		description: "Raise a positive rational number to a negative integer power",
		inputR:      Rational{3, 5},
		inputInt:    -2,
		expected:    Rational{25, 9},
	},
	{
		description: "Raise a negative rational number to an even negative integer power",
		inputR:      Rational{-3, 5},
		inputInt:    -2,
		expected:    Rational{25, 9},
	},
	{
		description: "Raise a negative rational number to an odd negative integer power",
		inputR:      Rational{-3, 5},
		inputInt:    -3,
		expected:    Rational{-125, 27},
	},
	{
		description: "Raise zero to an integer power",
		inputR:      Rational{0, 1},
		inputInt:    5,
		expected:    Rational{0, 1},
	},
	{
		description: "Raise one to an integer power",
		inputR:      Rational{1, 1},
		inputInt:    4,
		expected:    Rational{1, 1},
	},
	{
		description: "Raise a positive rational number to the power of zero",
		inputR:      Rational{1, 2},
		inputInt:    0,
		expected:    Rational{1, 1},
	},
	{
		description: "Raise a negative rational number to the power of zero",
		inputR:      Rational{-1, 2},
		inputInt:    0,
		expected:    Rational{1, 1},
	},
}

var testCasesExpreal = []struct {
	description string
	inputInt    int
	inputR      Rational
	expected    float64
}{
	{
		description: "Raise a real number to a positive rational number",
		inputInt:    8,
		inputR:      Rational{4, 3},
		expected:    16,
	},
	{
		description: "Raise a real number to a negative rational number",
		inputInt:    9,
		inputR:      Rational{-1, 2},
		expected:    0.3333333333333333,
	},
	{
		description: "Raise a real number to a zero rational number",
		inputInt:    2,
		inputR:      Rational{0, 1},
		expected:    1,
	},
}

var testCasesReduce = []struct {
	description string
	input       Rational
	expected    Rational
}{
	{
		description: "Reduce a positive rational number to lowest terms",
		input:       Rational{2, 4},
		expected:    Rational{1, 2},
	},
	{
		description: "Reduce places the minus sign on the numerator",
		input:       Rational{3, -4},
		expected:    Rational{-3, 4},
	},
	{
		description: "Reduce a negative rational number to lowest terms",
		input:       Rational{-4, 6},
		expected:    Rational{-2, 3},
	},
	{
		description: "Reduce a rational number with a negative denominator to lowest terms",
		input:       Rational{3, -9},
		expected:    Rational{-1, 3},
	},
	{
		description: "Reduce zero to lowest terms",
		input:       Rational{0, 6},
		expected:    Rational{0, 1},
	},
	{
		description: "Reduce an integer to lowest terms",
		input:       Rational{-14, 7},
		expected:    Rational{-2, 1},
	},
	{
		description: "Reduce one to lowest terms",
		input:       Rational{13, 13},
		expected:    Rational{1, 1},
	},
}
