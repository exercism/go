package cryptosquare

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: e7055f9 Crypto Square: fix test case description (#2556)

type testCase struct {
	description string
	input       string
	expected    string
}

var testCases = []testCase{
	{
		description: "empty plaintext results in an empty ciphertext",
		input:       "",
		expected:    "",
	},
	{
		description: "normalization results in empty plaintext",
		input:       "... --- ...",
		expected:    "",
	},
	{
		description: "Lowercase",
		input:       "A",
		expected:    "a",
	},
	{
		description: "Remove spaces",
		input:       "  b ",
		expected:    "b",
	},
	{
		description: "Remove punctuation",
		input:       "@1,%!",
		expected:    "1",
	},
	{
		description: "9 character plaintext results in 3 chunks of 3 characters",
		input:       "This is fun!",
		expected:    "tsf hiu isn",
	},
	{
		description: "8 character plaintext results in 3 chunks, the last one with a trailing space",
		input:       "Chill out.",
		expected:    "clu hlt io ",
	},
	{
		description: "54 character plaintext results in 8 chunks, the last two with trailing spaces",
		input:       "If man was meant to stay on the ground, god would have given us roots.",
		expected:    "imtgdvs fearwer mayoogo anouuio ntnnlvt wttddes aohghn  sseoau ",
	},
}
