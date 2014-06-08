package wordy

import "testing"

var tests = []struct {
	q  string
	a  int
	ok bool
}{
	{"What is 1 plus 1?", 2, true},
	{"What is 53 plus 2?", 55, true},
	{"What is -1 plus -10?", -11, true},
	{"What is 123 plus 45678?", 45801, true},
	{"What is 4 minus -12?", 16, true},
	{"What is -3 multiplied by 25?", -75, true},
	{"What is 33 divided by -3?", -11, true},
	{"What is 1 plus 1 plus 1?", 3, true},
	{"What is 1 plus 5 minus -2?", 8, true},
	{"What is 20 minus 4 minus 13?", 3, true},
	{"What is 17 minus 6 plus 3?", 14, true},
	{"What is 2 multiplied by -2 multiplied by 3?", -12, true},
	{"What is -3 plus 7 multiplied by -2?", -8, true},
	{"What is -12 divided by 2 divided by -3?", 2, true},
	{"What is 53 cubed?", 0, false},
	{"Who is the president of the United States?", 0, false},
}

func TestAnswer(t *testing.T) {
	for _, test := range tests {
		switch a, ok := Answer(test.q); {
		case !ok:
			if test.ok {
				t.Errorf("Answer(%q) returned ok = false, expecting true.", test.q)
			}
		case !test.ok:
			t.Errorf("Answer(%q) = %d, %t, expecting ok = false.", test.q, a, ok)
		case a != test.a:
			t.Errorf("Answer(%q) = %d, want %d.", test.q, a, test.a)
		}
	}
}

// Benchmark combined time to answer all questions.
func BenchmarkAnswer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			Answer(test.q)
		}
	}
}
