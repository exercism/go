package tournament

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

// Define a function Tally(io.Reader, io.Writer) error.
//
// Note that unlike other tracks the Go version of the tally function
// should not ignore errors. It's not idiomatic Go to ignore errors.

// These test what testers call the happy path, where there's no error.
var happyTestCases = []struct {
	description string
	input       string
	expected    string
}{
	{
		description: "good",
		input: `
Allegoric Alaskians;Blithering Badgers;win
Devastating Donkeys;Courageous Californians;draw
Devastating Donkeys;Allegoric Alaskians;win
Courageous Californians;Blithering Badgers;loss
Blithering Badgers;Devastating Donkeys;loss
Allegoric Alaskians;Courageous Californians;win
`,
		expected: `
Team                           | MP |  W |  D |  L |  P
Devastating Donkeys            |  3 |  2 |  1 |  0 |  7
Allegoric Alaskians            |  3 |  2 |  0 |  1 |  6
Blithering Badgers             |  3 |  1 |  0 |  2 |  3
Courageous Californians        |  3 |  0 |  1 |  2 |  1
`[1:], // [1:] = strip initial readability newline
	},
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
`[1:],
	},
	{
		// A complete competition has all teams play eachother once or twice.
		description: "incomplete competition",
		input: `
Allegoric Alaskians;Blithering Badgers;win
Devastating Donkeys;Allegoric Alaskians;win
Courageous Californians;Blithering Badgers;loss
Allegoric Alaskians;Courageous Californians;win
`,
		expected: `
Team                           | MP |  W |  D |  L |  P
Allegoric Alaskians            |  3 |  2 |  0 |  1 |  6
Blithering Badgers             |  2 |  1 |  0 |  1 |  3
Devastating Donkeys            |  1 |  1 |  0 |  0 |  3
Courageous Californians        |  2 |  0 |  0 |  2 |  0
`[1:],
	},
	{
		description: "tie for first and last place",
		input: `
Courageous Californians;Devastating Donkeys;win
Allegoric Alaskians;Blithering Badgers;win
Devastating Donkeys;Allegoric Alaskians;loss
Courageous Californians;Blithering Badgers;win
Blithering Badgers;Devastating Donkeys;draw
Allegoric Alaskians;Courageous Californians;draw
`,
		expected: `
Team                           | MP |  W |  D |  L |  P
Allegoric Alaskians            |  3 |  2 |  1 |  0 |  7
Courageous Californians        |  3 |  2 |  1 |  0 |  7
Blithering Badgers             |  3 |  0 |  1 |  2 |  1
Devastating Donkeys            |  3 |  0 |  1 |  2 |  1
`[1:],
	},
}

var errorTestCases = []string{
	"Bla;Bla;Bla",
	"Devastating Donkeys_Courageous Californians;draw",
	"Devastating Donkeys@Courageous Californians;draw",
	"Devastating Donkeys;Allegoric Alaskians;dra",
}

func TestTallyHappy(t *testing.T) {
	for _, tc := range happyTestCases {
		t.Run(tc.description, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			var buffer bytes.Buffer
			err := Tally(reader, &buffer)
			// We don't expect errors for any of the test cases
			if err != nil {
				t.Fatalf("Tally for input named %q returned unexpected error %v", tc.description, err)
			}
			got := buffer.String()
			if got != tc.expected {
				t.Fatalf("Tally for input named %q returned unexpected value\ngot: %s\nwant: %s", tc.description, got, tc.expected)
			}
		})
	}
}

func TestTallyError(t *testing.T) {
	for i, errorInput := range errorTestCases {
		t.Run(fmt.Sprintf("Error Input %d", i), func(t *testing.T) {
			reader := strings.NewReader(errorInput)
			var buffer bytes.Buffer
			err := Tally(reader, &buffer)
			if err == nil {
				t.Fatalf("Tally for input %q expected error, got nil", errorInput)
			}
		})
	}
}

func BenchmarkTally(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tt := range happyTestCases {
			var buffer bytes.Buffer
			Tally(strings.NewReader(tt.input), &buffer)
		}
		for _, s := range errorTestCases {
			var buffer bytes.Buffer
			Tally(strings.NewReader(s), &buffer)
		}
	}
}
