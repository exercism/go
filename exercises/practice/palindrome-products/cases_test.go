package palindromeproducts

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: acdf199 [palindrome-products ]: Added test case where first product is not min product (#2174)

type testCase struct {
	description string
	// input to Products(): range limits for factors of the palindrome
	fmin, fmax int
	// are we checking the min or max
	property string
	product  int
	factors  [][2]int
	errorMsg string // start of text if there is an error, "" otherwise
}

var testCases = []testCase{
	{
		description: "find the smallest palindrome from single digit factors",
		fmin:        1,
		fmax:        9,
		property:    "smallest",
		product:     1,
		factors:     [][2]int{[2]int{1, 1}},
		errorMsg:    "",
	},
	{
		description: "find the smallest palindrome from double digit factors",
		fmin:        10,
		fmax:        99,
		property:    "smallest",
		product:     121,
		factors:     [][2]int{[2]int{11, 11}},
		errorMsg:    "",
	},
	{
		description: "find the smallest palindrome from triple digit factors",
		fmin:        100,
		fmax:        999,
		property:    "smallest",
		product:     10201,
		factors:     [][2]int{[2]int{101, 101}},
		errorMsg:    "",
	},
	{
		description: "find the smallest palindrome from four digit factors",
		fmin:        1000,
		fmax:        9999,
		property:    "smallest",
		product:     1002001,
		factors:     [][2]int{[2]int{1001, 1001}},
		errorMsg:    "",
	},
	{
		description: "empty result for smallest if no palindrome in the range",
		fmin:        1002,
		fmax:        1003,
		property:    "smallest",
		product:     0,
		factors:     [][2]int{},
		errorMsg:    "",
	},
	{
		description: "error result for smallest if min is more than max",
		fmin:        10000,
		fmax:        1,
		property:    "smallest",
		product:     0,
		factors:     [][2]int(nil),
		errorMsg:    "min must be <= max",
	},
	{
		description: "smallest product does not use the smallest factor",
		fmin:        3215,
		fmax:        4000,
		property:    "smallest",
		product:     10988901,
		factors:     [][2]int{[2]int{3297, 3333}},
		errorMsg:    "",
	},
	{
		description: "find the largest palindrome from single digit factors",
		fmin:        1,
		fmax:        9,
		property:    "largest",
		product:     9,
		factors:     [][2]int{[2]int{1, 9}, [2]int{3, 3}},
		errorMsg:    "",
	},
	{
		description: "find the largest palindrome from double digit factors",
		fmin:        10,
		fmax:        99,
		property:    "largest",
		product:     9009,
		factors:     [][2]int{[2]int{91, 99}},
		errorMsg:    "",
	},
	{
		description: "find the largest palindrome from triple digit factors",
		fmin:        100,
		fmax:        999,
		property:    "largest",
		product:     906609,
		factors:     [][2]int{[2]int{913, 993}},
		errorMsg:    "",
	},
	{
		description: "find the largest palindrome from four digit factors",
		fmin:        1000,
		fmax:        9999,
		property:    "largest",
		product:     99000099,
		factors:     [][2]int{[2]int{9901, 9999}},
		errorMsg:    "",
	},
	{
		description: "empty result for largest if no palindrome in the range",
		fmin:        15,
		fmax:        15,
		property:    "largest",
		product:     0,
		factors:     [][2]int{},
		errorMsg:    "",
	},
	{
		description: "error result for largest if min is more than max",
		fmin:        2,
		fmax:        1,
		property:    "largest",
		product:     0,
		factors:     [][2]int(nil),
		errorMsg:    "min must be <= max",
	},
}
