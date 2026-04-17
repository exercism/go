package tournament

import (
	"bytes"
	"strings"
	"testing"
)

// Define a function Tally(io.Reader, io.Writer) error.
//
// Note that unlike other tracks the Go version of the tally function
// should not ignore errors. It's not idiomatic Go to ignore errors.

// These test what testers call the happy path, where there's no error.
var extraTestCases = []TestCase{
	{
		description: "ignore comments and newlines",
		input: `

Allegoric Alaskians;Blithering Badgers;win
Devastating Donkeys;Allegoric Alaskians;win
# Catastrophic Loss of the Californians
Courageous Californians;Blithering Badgers;loss

Blithering Badgers;Devastating Donkeys;loss
Allegoric Alaskians;Courageous Californians;win
Devastating Donkeys;Courageous Californians;draw


`,
		expected: `
Team                           | MP |  W |  D |  L |  P
Devastating Donkeys            |  3 |  2 |  1 |  0 |  7
Allegoric Alaskians            |  3 |  2 |  0 |  1 |  6
Blithering Badgers             |  3 |  1 |  0 |  2 |  3
Courageous Californians        |  3 |  0 |  1 |  2 |  1
`,
	},
}

type errorTestCase struct {
	description string
	input       string
}

var errorTestCases = []errorTestCase{
	{
		description: "Invalid result 'Bla'",
		input:       "Bla;Bla;Bla",
	},
	{
		description: "Invalid separator, underscore",
		input:       "Devastating Donkeys_Courageous Californians;draw",
	},
	{
		description: "Invalid separator, '@'",
		input:       "Devastating Donkeys@Courageous Californians;draw",
	},
	{
		description: "Invalid result 'dra'",
		input:       "Devastating Donkeys;Allegoric Alaskians;dra",
	},
}

func RunOneTest(t *testing.T, tc TestCase) {
	reader := strings.NewReader(tc.input)
	var buffer bytes.Buffer
	err := Tally(reader, &buffer)
	// We don't expect errors for any of the test cases
	if err != nil {
		t.Fatalf("Tally for input named %q returned unexpected error %q", tc.description, err)
	}
	got := buffer.String()
	expected := strings.TrimLeft(tc.expected, "\n")
	if got != expected {
		t.Fatalf("Tally for input named %q returned unexpected value\ngot: %s\nwant: %s", tc.description, got, expected)
	}
}

func TestTallyCanonical(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			RunOneTest(t, tc)
		})
	}
}
func TestTallyExtra(t *testing.T) {
	for _, tc := range extraTestCases {
		t.Run(tc.description, func(t *testing.T) {
			RunOneTest(t, tc)
		})
	}
}

func TestTallyError(t *testing.T) {
	for _, tc := range errorTestCases {
		t.Run(tc.description, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			var buffer bytes.Buffer
			err := Tally(reader, &buffer)
			if err == nil {
				t.Fatalf("Tally for input %q expected error, got nil", tc.input)
			}
		})
	}
}

func BenchmarkTally(b *testing.B) {
	for range b.N {
		for _, tc := range testCases {
			var buffer bytes.Buffer
			Tally(strings.NewReader(tc.input), &buffer)
		}
		for _, tc := range extraTestCases {
			var buffer bytes.Buffer
			Tally(strings.NewReader(tc.input), &buffer)
		}
		for _, tc := range errorTestCases {
			var buffer bytes.Buffer
			Tally(strings.NewReader(tc.input), &buffer)
		}
	}
}
