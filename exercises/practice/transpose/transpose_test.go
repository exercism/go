package transpose

import (
	"reflect"
	"testing"
)

func TestTranspose(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := Transpose(tc.input)
			if len(actual) == 0 && len(tc.expected) == 0 {
				return
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				// let's make the error more specific and find the row it's on
				min := min(len(tc.expected), len(actual))
				for i := 0; i < min; i++ {
					if tc.expected[i] != actual[i] {
						t.Fatalf("Transpose(%#v)\n got:%#v\nwant:%#v\n row:%d\n got:%q\nwant:%q", tc.input, actual, tc.expected, i, actual[i], tc.expected[i])
					}
				}
				t.Fatalf("Transpose(%#v)\n got:%#v\nwant:%#v", tc.input, actual, tc.expected)
			}
		})
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func BenchmarkTranspose(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			Transpose(test.input)
		}
	}
}
