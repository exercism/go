package prime

import "testing"

func TestNth(t *testing.T) {
	for _, test := range tests {
		switch p, ok := Nth(test.n); {
		case !ok:
			if test.ok {
				t.Fatalf("FAIL %s\nNth(%d) returned !ok.  Expecting ok.", test.description, test.n)
			}
		case !test.ok:
			t.Fatalf("FAIL %s\nNth(%d) = %d, ok = %t.  Expecting !ok.", test.description, test.n, p, ok)
		case p != test.p:
			t.Fatalf("FAIL %s\nNth(%d) = %d, want %d.", test.description, test.n, p, test.p)
		}
		t.Logf("PASS %s", test.description)
	}
}

func BenchmarkNth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Nth(10001)
	}
}
