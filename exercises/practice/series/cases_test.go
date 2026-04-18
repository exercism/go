package series

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 1d72ef4 series: Add test for slice length way bigger than series length (#2093)

type simpleTestCase struct {
	description string
	digits      int
	input       string
	expected    string
}

var simpleTestCases = []simpleTestCase{
	{
		description: "slices of one from one",
		digits:      1,
		input:       "1",
		expected:    "1",
	},
	{
		description: "slices of one from two",
		digits:      1,
		input:       "12",
		expected:    "1",
	},
	{
		description: "slices of two",
		digits:      2,
		input:       "35",
		expected:    "35",
	},
	{
		description: "slices of two overlap",
		digits:      2,
		input:       "9142",
		expected:    "91",
	},
	{
		description: "slices can include duplicates",
		digits:      3,
		input:       "777777",
		expected:    "777",
	},
	{
		description: "slices of a long series",
		digits:      5,
		input:       "918493904243",
		expected:    "91849",
	},
}

type testCase struct {
	description string
	digits      int
	input       string
	expected    []string
	errorMsg    string
}

var testCases = []testCase{
	{
		description: "slices of one from one",
		digits:      1,
		input:       "1",
		expected:    []string{"1"},
	},
	{
		description: "slices of one from two",
		digits:      1,
		input:       "12",
		expected:    []string{"1", "2"},
	},
	{
		description: "slices of two",
		digits:      2,
		input:       "35",
		expected:    []string{"35"},
	},
	{
		description: "slices of two overlap",
		digits:      2,
		input:       "9142",
		expected:    []string{"91", "14", "42"},
	},
	{
		description: "slices can include duplicates",
		digits:      3,
		input:       "777777",
		expected:    []string{"777", "777", "777", "777"},
	},
	{
		description: "slices of a long series",
		digits:      5,
		input:       "918493904243",
		expected:    []string{"91849", "18493", "84939", "49390", "93904", "39042", "90424", "04243"},
	},
	{
		description: "slice length is too large",
		digits:      6,
		input:       "12345",
		expected:    []string(nil),
		errorMsg:    "slice length cannot be greater than series length",
	},
	{
		description: "slice length is way too large",
		digits:      42,
		input:       "12345",
		expected:    []string(nil),
		errorMsg:    "slice length cannot be greater than series length",
	},
	{
		description: "slice length cannot be zero",
		digits:      0,
		input:       "12345",
		expected:    []string(nil),
		errorMsg:    "slice length cannot be zero",
	},
	{
		description: "slice length cannot be negative",
		digits:      -1,
		input:       "123",
		expected:    []string(nil),
		errorMsg:    "slice length cannot be negative",
	},
	{
		description: "empty series is invalid",
		digits:      1,
		input:       "",
		expected:    []string(nil),
		errorMsg:    "series cannot be empty",
	},
}
