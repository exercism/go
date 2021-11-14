package wordcount

import (
	"reflect"
	"testing"
)

func TestWordCount(t *testing.T) {
	for _, tt := range testCases {
		expected := tt.output
		actual := WordCount(tt.input)
		if !reflect.DeepEqual(actual, expected) {
			t.Fatalf("%s\n\tExpected: %v\n\tGot: %v", tt.description, expected, actual)
		} else {
			t.Logf("PASS: %s - WordCount(%s)", tt.description, tt.input)
		}
	}
}

func BenchmarkWordCount(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tt := range testCases {
			WordCount(tt.input)
		}
	}
}
