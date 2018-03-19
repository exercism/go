package allyourbase

// Source: exercism/problem-specifications
// Commit: c21ffd7 all-your-base: improve "input base is one" test
// Problem Specifications Version: 2.3.0

var testCases = []struct {
	description string
	inputBase   int
	inputDigits []int
	outputBase  int
	expected    []int
	err         string
}{
	{
		description: "single bit one to decimal",
		inputBase:   2,
		inputDigits: []int{1},
		outputBase:  10,
		expected:    []int{1},
		err:         "",
	},
	{
		description: "binary to single decimal",
		inputBase:   2,
		inputDigits: []int{1, 0, 1},
		outputBase:  10,
		expected:    []int{5},
		err:         "",
	},
	{
		description: "single decimal to binary",
		inputBase:   10,
		inputDigits: []int{5},
		outputBase:  2,
		expected:    []int{1, 0, 1},
		err:         "",
	},
	{
		description: "binary to multiple decimal",
		inputBase:   2,
		inputDigits: []int{1, 0, 1, 0, 1, 0},
		outputBase:  10,
		expected:    []int{4, 2},
		err:         "",
	},
	{
		description: "decimal to binary",
		inputBase:   10,
		inputDigits: []int{4, 2},
		outputBase:  2,
		expected:    []int{1, 0, 1, 0, 1, 0},
		err:         "",
	},
	{
		description: "trinary to hexadecimal",
		inputBase:   3,
		inputDigits: []int{1, 1, 2, 0},
		outputBase:  16,
		expected:    []int{2, 10},
		err:         "",
	},
	{
		description: "hexadecimal to trinary",
		inputBase:   16,
		inputDigits: []int{2, 10},
		outputBase:  3,
		expected:    []int{1, 1, 2, 0},
		err:         "",
	},
	{
		description: "15-bit integer",
		inputBase:   97,
		inputDigits: []int{3, 46, 60},
		outputBase:  73,
		expected:    []int{6, 10, 45},
		err:         "",
	},
	{
		description: "empty list",
		inputBase:   2,
		inputDigits: []int{},
		outputBase:  10,
		expected:    []int{0},
		err:         "",
	},
	{
		description: "single zero",
		inputBase:   10,
		inputDigits: []int{0},
		outputBase:  2,
		expected:    []int{0},
		err:         "",
	},
	{
		description: "multiple zeros",
		inputBase:   10,
		inputDigits: []int{0, 0, 0},
		outputBase:  2,
		expected:    []int{0},
		err:         "",
	},
	{
		description: "leading zeros",
		inputBase:   7,
		inputDigits: []int{0, 6, 0},
		outputBase:  10,
		expected:    []int{4, 2},
		err:         "",
	},
	{
		description: "input base is one",
		inputBase:   1,
		inputDigits: []int{0},
		outputBase:  10,
		expected:    []int(nil),
		err:         "input base must be >= 2",
	},
	{
		description: "input base is zero",
		inputBase:   0,
		inputDigits: []int{},
		outputBase:  10,
		expected:    []int(nil),
		err:         "input base must be >= 2",
	},
	{
		description: "input base is negative",
		inputBase:   -2,
		inputDigits: []int{1},
		outputBase:  10,
		expected:    []int(nil),
		err:         "input base must be >= 2",
	},
	{
		description: "negative digit",
		inputBase:   2,
		inputDigits: []int{1, -1, 1, 0, 1, 0},
		outputBase:  10,
		expected:    []int(nil),
		err:         "all digits must satisfy 0 <= d < input base",
	},
	{
		description: "invalid positive digit",
		inputBase:   2,
		inputDigits: []int{1, 2, 1, 0, 1, 0},
		outputBase:  10,
		expected:    []int(nil),
		err:         "all digits must satisfy 0 <= d < input base",
	},
	{
		description: "output base is one",
		inputBase:   2,
		inputDigits: []int{1, 0, 1, 0, 1, 0},
		outputBase:  1,
		expected:    []int(nil),
		err:         "output base must be >= 2",
	},
	{
		description: "output base is zero",
		inputBase:   10,
		inputDigits: []int{7},
		outputBase:  0,
		expected:    []int(nil),
		err:         "output base must be >= 2",
	},
	{
		description: "output base is negative",
		inputBase:   2,
		inputDigits: []int{1},
		outputBase:  -7,
		expected:    []int(nil),
		err:         "output base must be >= 2",
	},
	{
		description: "both bases are negative",
		inputBase:   -2,
		inputDigits: []int{1},
		outputBase:  -7,
		expected:    []int(nil),
		err:         "input base must be >= 2",
	},
}
