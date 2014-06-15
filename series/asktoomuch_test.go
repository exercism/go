// +build asktoomuch

package slice

import "testing"

func TestAskTooMuch(t *testing.T) {
	test := allTests[0]
	defer func() {
		if recover() != nil {
			t.Fatalf("Yikes, Frist(%d, %s) panicked!", test.n, test.s)
		}
	}()
	for _, test = range allTests {
		switch res := Frist(test.n, test.s); {
		case len(test.out) > 0: // well, this should work
			if res != test.out[0] {
				t.Fatalf("Yikes, Frist(%d, %s) = %q, want %q.",
					test.n, test.s, res, test.out[0])
			}
		case len(res) != test.n:
			t.Fatalf("Yikes, Frist(%d, %s) = %q, but %q doesn't have %d characters.",
				test.n, test.s, res, res, test.n)
		default:
			t.Fatalf("Yikes, Frist(%d, %s) = %q, but %q isn't in %q",
				test.n, test.s, res, res, test.s)
		}
	}
}
