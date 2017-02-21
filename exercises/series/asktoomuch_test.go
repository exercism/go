// +build asktoomuch

package slice

import "testing"

func TestAskTooMuch(t *testing.T) {
	test := allTests[0]
	defer func() {
		if recover() != nil {
			t.Fatalf("Yikes, UnsafeFirst(%d, %q) panicked!", test.n, test.s)
		}
	}()
	for _, test = range allTests {
		switch res := UnsafeFirst(test.n, test.s); {
		case len(test.out) > 0: // well, this should work
			if res != test.out[0] {
				t.Fatalf("Yikes, UnsafeFirst(%d, %q) = %q, want %q.",
					test.n, test.s, res, test.out[0])
			}
		case len(res) != test.n:
			t.Fatalf("Yikes, UnsafeFirst(%d, %q) = %q, but %q doesn't have %d characters.",
				test.n, test.s, res, res, test.n)
		default:
			t.Fatalf("Yikes, UnsafeFirst(%d, %q) = %q, but %q isn't in %q",
				test.n, test.s, res, res, test.s)
		}
	}
}
