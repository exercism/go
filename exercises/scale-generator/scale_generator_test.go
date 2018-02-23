package scale

import (
	"fmt"
	"testing"
)

func TestScale(t *testing.T) {
	for _, test := range scaleTestCases {
		actual := Scale(test.tonic, test.interval)
		if fmt.Sprintf("%q", actual) != fmt.Sprintf("%q", test.expected) {
			t.Fatalf("Scale Generator test [%s], expected %s, actual %s", test.description, test.expected, actual)
		}
	}
}

func BenchmarkScale(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range scaleTestCases {
			Scale(test.tonic, test.interval)
		}
	}
}
