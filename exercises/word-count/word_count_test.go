package wordcount

import (
	"reflect"
	"testing"
)

const targetTestVersion = 2

func TestWordCount(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v", testVersion, targetTestVersion)
	}
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
	for i := 0; i < b.N; i++ {
		for _, tt := range testCases {
			WordCount(tt.input)
		}
	}
}
