//nolint:gosec // In the context of this exercise, it is fine to use math.Rand instead of crypto.Rand.
package binarysearch

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestSearchInts(t *testing.T) {
	for _, test := range testCases {
		t.Run(test.description, func(t *testing.T) {
			if actual := SearchInts(test.inputList, test.inputKey); actual != test.expectedKey {
				t.Errorf("SearchInts(%#v, %d) = %d, want %d", test.inputList, test.inputKey, actual, test.expectedKey)
			}
		})
	}
}

// Benchmarks also test searching larger random slices
type query struct {
	slice []int
	x     int
}

func newQuery(n int) (query, error) {
	q := query{slice: make([]int, n)}
	for i := 0; i < n; i++ {
		q.slice[i] = i
	}
	q.x = rand.Intn(n)
	if res := SearchInts(q.slice, q.x); res != q.x {
		return q, fmt.Errorf("search of %d values gave different answer", n)
	}
	return q, nil
}

func runBenchmark(n int, b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	q, err := newQuery(n)
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SearchInts(q.slice, q.x)
	}
}

func Benchmark1e2(b *testing.B) { runBenchmark(1e2, b) }
func Benchmark1e4(b *testing.B) { runBenchmark(1e4, b) }
func Benchmark1e6(b *testing.B) { runBenchmark(1e6, b) }
func Benchmark1e8(b *testing.B) { runBenchmark(1e8, b) }
