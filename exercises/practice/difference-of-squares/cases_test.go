package differenceofsquares

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 42dd0ce Remove version (#1678)

var testCases = []struct {
	description string
	fn          func(int) int
	fnName      string
	input       int
	expected    int
}{
	{
		description: "sum of squares 1",
		fn:          SumOfSquares,
		fnName:      "SumOfSquares",
		input:       1,
		expected:    1,
	},
	{
		description: "sum of squares 5",
		fn:          SumOfSquares,
		fnName:      "SumOfSquares",
		input:       5,
		expected:    55,
	},
	{
		description: "sum of squares 100",
		fn:          SumOfSquares,
		fnName:      "SumOfSquares",
		input:       100,
		expected:    338350,
	},
	{
		description: "square of sum 1",
		fn:          SquareOfSum,
		fnName:      "SquareOfSum",
		input:       1,
		expected:    1,
	},
	{
		description: "square of sum 5",
		fn:          SquareOfSum,
		fnName:      "SquareOfSum",
		input:       5,
		expected:    225,
	},
	{
		description: "square of sum 100",
		fn:          SquareOfSum,
		fnName:      "SquareOfSum",
		input:       100,
		expected:    25502500,
	},
	{
		description: "difference of squares 1",
		fn:          Difference,
		fnName:      "Difference",
		input:       1,
		expected:    0,
	},
	{
		description: "difference of squares 5",
		fn:          Difference,
		fnName:      "Difference",
		input:       5,
		expected:    170,
	},
	{
		description: "difference of squares 100",
		fn:          Difference,
		fnName:      "Difference",
		input:       100,
		expected:    25164150,
	},
}
