// +build first

package slice

import "testing"

func TestFirst(t *testing.T) {
	for _, test := range allTests {
		switch res, ok := First(test.n, test.s); {
		case !ok:
			if len(test.out) > 0 {
				t.Fatalf("First(%d, %q) returned !ok, want ok.",
					test.n, test.s)
			}
		case len(test.out) == 0:
			t.Fatalf("First(%d, %q) = %q, %t.  Expected ok == false",
				test.n, test.s, res, ok)
		case res != test.out[0]:
			t.Fatalf("First(%d, %q) = %q.  Want %q.",
				test.n, test.s, res, test.out[0])
		}
	}
}
