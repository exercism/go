package palindrome

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

// API to impliment:
// type Product struct {
// 	Product int // palindromic, of course
// 	// list of all possible two-factor factorizations of Product, within
// 	// given limits, in order
// 	Factorizations [][2]int
// }
// func Products(fmin, fmax int) (pmin, pmax Product, error)

var testData = []struct {
	// input to Products(): range limits for factors of the palindrome
	fmin, fmax int
	// output from Products():
	pmin, pmax Product // min and max palandromic products
	errPrefix  string  // start of text if there is an error, "" otherwise
}{
	{1, 9,
		Product{}, // zero value means don't bother to test it
		Product{9, [][2]int{{1, 9}, {3, 3}}},
		""},
	{10, 99,
		Product{121, [][2]int{{11, 11}}},
		Product{9009, [][2]int{{91, 99}}},
		""},
	{100, 999,
		Product{10201, [][2]int{{101, 101}}},
		Product{906609, [][2]int{{913, 993}}},
		""},
	{4, 10, Product{}, Product{}, "No palindromes"},
	{10, 4, Product{}, Product{}, "fmin > fmax"},
	/* bonus curiosities.  (can a negative number be a palindrome?
	// most say no
	{-99, -10, Product{}, Product{}, "Negative limits"},
	// but you can still get non-negative products from negative factors.
	{-99, -10,
		Product{121, [][2]int{{-11, -11}}},
		Product{9009, [][2]int{{-99, -91}}},
		""},
	{-2, 2,
		Product{0, [][2]int{{-2, 0}, {-1, 0}, {0, 0}, {0, 1}, {0, 2}}},
		Product{4, [][2]int{{-2, -2}, {2, 2}}},
		""},
	// or you could reverse the *digits*, keeping the minus sign in place.
	{-2, 2,
		Product{-4, [][2]int{{-2, 2}}},
		Product{4, [][2]int{{-2, -2}, {2, 2}}},
		""},
	{
	{0, (^uint(0))>>1, Product{}, Product{}, "This one's gonna overflow"},
	*/
}

func TestPalindromeProducts(t *testing.T) {
	for _, test := range testData {
		// common preamble for test failures
		ret := fmt.Sprintf("Products(%d, %d) returned",
			test.fmin, test.fmax)
		// test
		pmin, pmax, err := Products(test.fmin, test.fmax)
		switch {
		case err == nil:
			if test.errPrefix > "" {
				t.Fatalf(ret+" err = nil, want %q", test.errPrefix+"...")
			}
		case test.errPrefix == "":
			t.Fatalf(ret+" err = %q, want nil", err)
		case !strings.HasPrefix(err.Error(), test.errPrefix):
			t.Fatalf(ret+" err = %q, want %q", err, test.errPrefix+"...")
		default:
			continue // correct error, no further tests for this test case
		}
		matchProd := func(ww string, rp, wp Product) {
			if len(wp.Factorizations) > 0 && // option to skip test
				!reflect.DeepEqual(rp, wp) {
				t.Fatal(ret, ww, "=", rp, "want", wp)
			}
		}
		matchProd("pmin", pmin, test.pmin)
		matchProd("pmax", pmax, test.pmax)
	}
}

func BenchmarkPalindromeProducts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range testData {
			Products(test.fmin, test.fmax)
		}
	}
}
