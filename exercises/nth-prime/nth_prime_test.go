package prime

import "testing"

var tests = []struct {
	n  int
	p  int
	ok bool
}{
	{1, 2, true},
	{2, 3, true},
	{3, 5, true},
	{4, 7, true},
	{5, 11, true},
	{6, 13, true},
	{10001, 104743, true},
	{0, 0, false},
}

func TestNth(t *testing.T) {
	for _, test := range tests {
		switch p, ok := Nth(test.n); {
		case !ok:
			if test.ok {
				t.Fatalf("Nth(%d) returned !ok.  Expecting ok.", test.n)
			}
		case !test.ok:
			t.Fatalf("Nth(%d) = %d, ok = %t.  Expecting !ok.", test.n, p, ok)
		case p != test.p:
			t.Fatalf("Nth(%d) = %d, want %d.", test.n, p, test.p)
		}
	}
}

func BenchmarkNth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Nth(10001)
	}
}
