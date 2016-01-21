package queenattack

import "testing"

// Arguments to CanQueenAttack are in algebraic notation.
// See http://en.wikipedia.org/wiki/Algebraic_notation_(chess)

var tests = []struct {
	w, b   string
	attack bool
	ok     bool
}{
	{"b4", "b4", false, false},      // same square
	{"a8", "b9", false, false},      // off board
	{"here", "there", false, false}, // invalid
	{"", "", false, false},

	{"b3", "d7", false, true}, // no attack
	{"b4", "b7", true, true},  // same file
	{"e4", "b4", true, true},  // same rank
	{"a1", "f6", true, true},  // common diagonals
	{"a6", "b7", true, true},
	{"d1", "f3", true, true},
	{"f1", "a6", true, true},
}

func TestCanQueenAttack(t *testing.T) {
	for _, test := range tests {
		switch attack, err := CanQueenAttack(test.w, test.b); {
		case err != nil:
			if test.ok {
				t.Fatalf("CanQueenAttack(%s, %s) returned error %q.  "+
					"Error not expected.",
					test.w, test.b, err)
			}
		case !test.ok:
			t.Fatalf("CanQueenAttack(%s, %s) = %t, %v.  Expected error.",
				test.w, test.b, attack, err)
		case attack != test.attack:
			t.Fatalf("CanQueenAttack(%s, %s) = %t, want %t.",
				test.w, test.b, attack, test.attack)
		}
	}
}

// Benchmark combined time for all test cases
func BenchmarkCanQueenAttack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			CanQueenAttack(test.w, test.b)
		}
	}
}
