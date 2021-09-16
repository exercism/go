package trinary

import "testing"

var tests = []struct {
	arg  string
	want int64
	ok   bool
}{
	{"0", 0, true},
	{"1", 1, true},
	{"2", 2, true},
	{"10", 3, true},
	{"201", 19, true},
	{"0201", 19, true},
	{"0000000000000000000000000000000000000000201", 19, true},
	{"2021110011022210012102010021220101220221", 9223372036854775807, true},
	{"2021110011022210012102010021220101220222", 0, false},
}

func TestParseTrinary(t *testing.T) {
	for _, test := range tests {
		switch res, err := ParseTrinary(test.arg); {
		case err != nil:
			var _ error = err
			if test.ok {
				t.Errorf("ParseTrinary(%q) returned error %q, Error not expected",
					test.arg, err)
			}
		case !test.ok:
			t.Errorf("ParseTrinary(%q) = %d, %v, expected error", test.arg, res, err)
		case res != test.want:
			t.Errorf("ParseTrinary(%q) = %d, want %d", test.arg, res, test.want)
		}
	}
}

func BenchmarkParseTrinary(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			ParseTrinary(test.arg)
		}
	}
}
