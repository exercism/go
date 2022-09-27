package twofer

import "testing"

type testCase struct {
	description, input, expected string
}

var testCases = []testCase{
	{
		description: "empty name",
		input:       "",
		expected:    "One for you, one for me.",
	},
	{
		description: "name is Alice",
		input:       "Alice",
		expected:    "One for Alice, one for me.",
	},
	{
		description: "name is Bob",
		input:       "Bob",
		expected:    "One for Bob, one for me.",
	},
}

func TestShareWith(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			got := ShareWith(tc.input)
			if got != tc.expected {
				t.Fatalf("ShareWith(%q)\n got: %q\nwant: %q", tc.input, got, tc.expected)
			}
		})
	}
}

func BenchmarkShareWith(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {

		for _, test := range testCases {
			ShareWith(test.input)
		}

	}
}
