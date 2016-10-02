// Short cut!  Have you already done the matrix exercise?
// If it seems helpful, copy your code from the matrix exercise
// to this directory so you have access to it.  You can leave it as matrix.go
// and start a new file saddle_points.go with additional code that completes
// this exercise.  Submit only saddle_points.go for this exercise.

package matrix

import "testing"

const targetTestVersion = 1

var tests = []struct {
	m  string
	sp []Pair
}{
	{"2 1\n1 2", nil},
	{"1 2\n3 4", []Pair{{0, 1}}},
	{"18 3 39 19 91\n38 10 8 77 320\n3 4 8 6 7", []Pair{{2, 2}}},
	{"4 5 4\n3 5 5\n1 5 4", []Pair{{0, 1}, {1, 1}, {2, 1}}},
}

func TestSaddle(t *testing.T) {
	for _, test := range tests {
		m, err := New(test.m)
		if err != nil {
			var _ error = err
			t.Fatalf("TestSaddle needs working New.  "+
				"New(%s) returned %s.  Error not expected.",
				test.m, err)
		}
		if sp := m.Saddle(); !eq(sp, test.sp) {
			t.Fatalf("%v.Saddle() = %v, want %v", m, sp, test.sp)
		}
	}
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

func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Errorf("Found testVersion = %v, want %v.", testVersion, targetTestVersion)
	}
}

func BenchmarkSaddle(b *testing.B) {
	ms := make([]*Matrix, len(tests))
	var err error
	for i, test := range tests {
		if ms[i], err = New(test.m); err != nil {
			b.Fatalf("BenchmarkSaddle needs working New.  "+
				"New(%s) returned %s.  Error not expected.",
				test.m, err)
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, m := range ms {
			m.Saddle()
		}
	}
}
