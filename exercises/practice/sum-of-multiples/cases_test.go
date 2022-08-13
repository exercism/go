package summultiples

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 42dd0ce Remove version (#1678)

var testCases = []struct {
	description string
	divisors    []int
	limit       int
	expected    int
}{
	{
		description: "no multiples within limit",
		divisors:    []int{3, 5},
		limit:       1,
		expected:    0,
	},
	{
		description: "one factor has multiples within limit",
		divisors:    []int{3, 5},
		limit:       4,
		expected:    3,
	},
	{
		description: "more than one multiple within limit",
		divisors:    []int{3},
		limit:       7,
		expected:    9,
	},
	{
		description: "more than one factor with multiples within limit",
		divisors:    []int{3, 5},
		limit:       10,
		expected:    23,
	},
	{
		description: "each multiple is only counted once",
		divisors:    []int{3, 5},
		limit:       100,
		expected:    2318,
	},
	{
		description: "a much larger limit",
		divisors:    []int{3, 5},
		limit:       1000,
		expected:    233168,
	},
	{
		description: "three factors",
		divisors:    []int{7, 13, 17},
		limit:       20,
		expected:    51,
	},
	{
		description: "factors not relatively prime",
		divisors:    []int{4, 6},
		limit:       15,
		expected:    30,
	},
	{
		description: "some pairs of factors relatively prime and some not",
		divisors:    []int{5, 6, 8},
		limit:       150,
		expected:    4419,
	},
	{
		description: "one factor is a multiple of another",
		divisors:    []int{5, 25},
		limit:       51,
		expected:    275,
	},
	{
		description: "much larger factors",
		divisors:    []int{43, 47},
		limit:       10000,
		expected:    2203160,
	},
	{
		description: "all numbers are multiples of 1",
		divisors:    []int{1},
		limit:       100,
		expected:    4950,
	},
	{
		description: "no factors means an empty sum",
		divisors:    []int{},
		limit:       10000,
		expected:    0,
	},
	{
		description: "the only multiple of 0 is 0",
		divisors:    []int{0},
		limit:       1,
		expected:    0,
	},
	{
		description: "the factor 0 does not affect the sum of multiples of other factors",
		divisors:    []int{3, 0},
		limit:       4,
		expected:    3,
	},
	{
		description: "solutions using include-exclude must extend to cardinality greater than 3",
		divisors:    []int{2, 3, 5, 7, 11},
		limit:       10000,
		expected:    39614537,
	},
}
