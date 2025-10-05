// Short cut!  Have you already done the matrix exercise?
// If it seems helpful, copy your code from the matrix exercise
// to this directory so you have access to it.  You can leave it as matrix.go
// and start a new file saddle_points.go with additional code that completes
// this exercise. If you do copy the matrix.go file, do not forget to submit both
// saddle_points.go and matrix.go as part of your solution.

package matrix

import (
	"strconv"
	"strings"
	"testing"
)

func TestSaddle(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			in := generateString(tc.input)
			m, err := New(in)
			if err != nil {
				t.Fatalf("TestSaddle needs working New. New(%s) returned %q.  Error not expected.", in, err)
			}
			if got := m.Saddle(); !eq(got, tc.expectedOutput) {
				t.Fatalf("%v.Saddle() = %v, want %v", m, got, tc.expectedOutput)
			}
		})
	}
}

func generateString(in [][]int) string {
	var parts []string
	for _, numbersPerLine := range in {
		var lineParts []string
		for _, number := range numbersPerLine {
			lineParts = append(lineParts, strconv.Itoa(number))
		}
		parts = append(parts, strings.Join(lineParts, " "))
	}
	return strings.Join(parts, "\n")
}

func eq(got, exp []Pair) bool {
	if len(got) != len(exp) {
		return false
	}
exp:
	for _, e := range exp {
		for _, g := range got {
			if g == e {
				continue exp
			}
		}
		return false
	}
	return true
}

func BenchmarkSaddle(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	ms := make([]*Matrix, len(testCases))
	var err error
	for i, tc := range testCases {
		in := generateString(tc.input)
		if ms[i], err = New(in); err != nil {
			b.Fatalf("BenchmarkSaddle needs working New. New(%s) returned %q.  Error not expected.", in, err)
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, m := range ms {
			m.Saddle()
		}
	}
}
