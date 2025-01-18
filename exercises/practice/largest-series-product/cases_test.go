package lsproduct

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 7a8722a Reorder keys (#1960)

var testCases = []struct {
	description string
	digits      string
	span        int
	expected    int64
	error       string
}{
	{
		description: "finds the largest product if span equals length",
		digits:      "29",
		span:        2,
		expected:    18,
		error:       "",
	},
	{
		description: "can find the largest product of 2 with numbers in order",
		digits:      "0123456789",
		span:        2,
		expected:    72,
		error:       "",
	},
	{
		description: "can find the largest product of 2",
		digits:      "576802143",
		span:        2,
		expected:    48,
		error:       "",
	},
	{
		description: "can find the largest product of 3 with numbers in order",
		digits:      "0123456789",
		span:        3,
		expected:    504,
		error:       "",
	},
	{
		description: "can find the largest product of 3",
		digits:      "1027839564",
		span:        3,
		expected:    270,
		error:       "",
	},
	{
		description: "can find the largest product of 5 with numbers in order",
		digits:      "0123456789",
		span:        5,
		expected:    15120,
		error:       "",
	},
	{
		description: "can get the largest product of a big number",
		digits:      "73167176531330624919225119674426574742355349194934",
		span:        6,
		expected:    23520,
		error:       "",
	},
	{
		description: "reports zero if the only digits are zero",
		digits:      "0000",
		span:        2,
		expected:    0,
		error:       "",
	},
	{
		description: "reports zero if all spans include zero",
		digits:      "99099",
		span:        3,
		expected:    0,
		error:       "",
	},
	{
		description: "rejects span longer than string length",
		digits:      "123",
		span:        4,
		expected:    0,
		error:       "span must be smaller than string length",
	},
	{
		description: "rejects empty string and nonzero span",
		digits:      "",
		span:        1,
		expected:    0,
		error:       "span must be smaller than string length",
	},
	{
		description: "rejects invalid character in digits",
		digits:      "1234a5",
		span:        2,
		expected:    0,
		error:       "digits input must only contain digits",
	},
	{
		description: "rejects negative span",
		digits:      "12345",
		span:        -1,
		expected:    0,
		error:       "span must not be negative",
	},
}
