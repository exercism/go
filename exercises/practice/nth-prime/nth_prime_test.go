package prime

import "testing"

func TestNth(t *testing.T) {
	for _, test := range tests {
		switch p, err := Nth(test.n); {
		case err != nil:
			if test.err == nil {
				t.Fatalf("FAIL %s\nNth(%d) returned error.  Expecting no error.", test.description, test.n)
			}
		case test.err != nil:
			t.Fatalf("FAIL %s\nNth(%d) = %d, err = %t.  Expecting error.", test.description, test.n, p, err)
		case p != test.p:
			t.Fatalf("FAIL %s\nNth(%d) = %d, want %d.", test.description, test.n, p, test.p)
		}
		t.Logf("PASS %s", test.description)
	}
}

func BenchmarkNth(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		Nth(10001)
	}
}
