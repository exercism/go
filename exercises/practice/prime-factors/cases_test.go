package prime

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: d137db1 Format using prettier (#1917)

var testCases = []struct {
	description string
	input       int64
	expected    []int64
}{
	{
		description: "no factors",
		input:       1,
		expected:    []int64{},
	},
	{
		description: "prime number",
		input:       2,
		expected:    []int64{2},
	},
	{
		description: "another prime number",
		input:       3,
		expected:    []int64{3},
	},
	{
		description: "square of a prime",
		input:       9,
		expected:    []int64{3, 3},
	},
	{
		description: "product of first prime",
		input:       4,
		expected:    []int64{2, 2},
	},
	{
		description: "cube of a prime",
		input:       8,
		expected:    []int64{2, 2, 2},
	},
	{
		description: "product of second prime",
		input:       27,
		expected:    []int64{3, 3, 3},
	},
	{
		description: "product of third prime",
		input:       625,
		expected:    []int64{5, 5, 5, 5},
	},
	{
		description: "product of first and second prime",
		input:       6,
		expected:    []int64{2, 3},
	},
	{
		description: "product of primes and non-primes",
		input:       12,
		expected:    []int64{2, 2, 3},
	},
	{
		description: "product of primes",
		input:       901255,
		expected:    []int64{5, 17, 23, 461},
	},
	{
		description: "factors include a large prime",
		input:       93819012551,
		expected:    []int64{11, 9539, 894119},
	},
}
