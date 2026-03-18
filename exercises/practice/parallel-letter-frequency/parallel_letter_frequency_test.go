package parallelletterfrequency

import (
	"fmt"
	"maps"
	"slices"
	"strings"
	"testing"
)

// In the separate file parallel_letter_frequency.go, you are given a function,
// Frequency(), to sequentially count letter frequencies in a single text.
//
// Perform this exercise on parallelism using Go concurrency features.
// Make concurrent calls to Frequency and combine results to obtain the answer.

func fmtMap(m FreqMap) string {
	var items []string
	for k, v := range m {
		items = append(items, fmt.Sprintf("'%c': %d", k, v))
	}
	slices.Sort(items)
	return fmt.Sprintf("FreqMap{%s}", strings.Join(items, ", "))
}

func TestCalculateFrequencies(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual := ConcurrentFrequency(tc.input); !maps.Equal(actual, tc.expected) {
				t.Fatalf("ConcurrentFrequency(%#v),\ngot: %#v,\nwant: %#v", tc.input, fmtMap(actual), fmtMap(tc.expected))
			}
		})
	}
}

func BenchmarkSequentialFrequency(b *testing.B) {
	for range b.N {
	}
}

func BenchmarkConcurrentFrequency(b *testing.B) {
	for range b.N {
	}
}
