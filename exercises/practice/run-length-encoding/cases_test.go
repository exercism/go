package encode

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: a757698 run-length-encoding: 'lower case' -> 'lowercase' (#1708)

type testCase struct {
	description string
	input       string
	expected    string
}

// run-length encode a string
var encodeTests = []testCase{
	{
		description: "empty string",
		input:       "",
		expected:    "",
	},
	{
		description: "single characters only are encoded without count",
		input:       "XYZ",
		expected:    "XYZ",
	},
	{
		description: "string with no single characters",
		input:       "AABBBCCCC",
		expected:    "2A3B4C",
	},
	{
		description: "single characters mixed with repeated characters",
		input:       "WWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWB",
		expected:    "12WB12W3B24WB",
	},
	{
		description: "multiple whitespace mixed in string",
		input:       "  hsqq qww  ",
		expected:    "2 hs2q q2w2 ",
	},
	{
		description: "lowercase characters",
		input:       "aabbbcccc",
		expected:    "2a3b4c",
	},
}

// run-length decode a string
var decodeTests = []testCase{
	{
		description: "empty string",
		input:       "",
		expected:    "",
	},
	{
		description: "single characters only",
		input:       "XYZ",
		expected:    "XYZ",
	},
	{
		description: "string with no single characters",
		input:       "2A3B4C",
		expected:    "AABBBCCCC",
	},
	{
		description: "single characters with repeated characters",
		input:       "12WB12W3B24WB",
		expected:    "WWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWB",
	},
	{
		description: "multiple whitespace mixed in string",
		input:       "2 hs2q q2w2 ",
		expected:    "  hsqq qww  ",
	},
	{
		description: "lowercase string",
		input:       "2a3b4c",
		expected:    "aabbbcccc",
	},
}

// encode and then decode
var encodeDecodeTests = []testCase{
	{
		description: "encode followed by decode gives original string",
		input:       "zzz ZZ  zZ",
		expected:    "zzz ZZ  zZ",
	},
}
