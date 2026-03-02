package scale

import (
	"fmt"
	"testing"
)

func TestScale(t *testing.T) {
	for _, tc := range scaleTestCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := Scale(tc.tonic, tc.interval)
			if fmt.Sprintf("%q", actual) != fmt.Sprintf("%q", tc.expected) {
				t.Fatalf("Scale(%q, %q)\n got:%#v\nwant:%#v", tc.tonic, tc.interval, actual, tc.expected)
			}
		})
	}
}

func BenchmarkScale(b *testing.B) {
	for range b.N {
		for _, test := range scaleTestCases {
			Scale(test.tonic, test.interval)
		}
	}
}
