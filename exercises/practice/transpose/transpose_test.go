package transpose

import (
	"slices"
	"testing"
)

func TestTranspose(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := Transpose(tc.input)
			if len(actual) == 0 && len(tc.expected) == 0 {
				return
			}
			if !slices.Equal(actual, tc.expected) {
				// let's make the error more specific and find the row it's on
				smallest := min(len(tc.expected), len(actual))
				for i := range smallest {
					if tc.expected[i] != actual[i] {
						t.Fatalf("Transpose(%#v)\ngot: %#v\nwant: %#v\n row:%d\ngot: %q\nwant: %q", tc.input, actual, tc.expected, i, actual[i], tc.expected[i])
					}
				}
				t.Fatalf("Transpose(%#v)\ngot: %#v\nwant: %#v", tc.input, actual, tc.expected)
			}
		})
	}
}

func BenchmarkTranspose(b *testing.B) {
	for range b.N {
		for _, test := range testCases {
			Transpose(test.input)
		}
	}
}
