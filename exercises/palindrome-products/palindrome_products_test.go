package palindrome

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

/* API to implement:

type Product struct {
	Product int // palindromic, of course
	Factorizations [][2]int //list of all possible two-factor factorizations of Product, within given limits, in order
 }

 func Products(fmin, fmax int) (pmin, pmax Product, error)
*/

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
	{4, 10, Product{}, Product{}, "no palindromes"},
	{10, 4, Product{}, Product{}, "fmin > fmax"},
}

// Bonus curiosities. Can a negative number be a palindrome? Most say no.
var bonusData = []struct {
	fmin, fmax int
	pmin, pmax Product
	errPrefix  string
}{
	// The following two test cases have the same input, but different expectations. Uncomment just one or the other.

	/* Here you can test that you can reach the limit of the largest palindrome made of two 2-digit numbers.
	{-99, -10, Product{}, Product{}, "Negative limits"}, */

	// You can still get non-negative products from negative factors.
	{-99, -10,
		Product{121, [][2]int{{-11, -11}}},
		Product{9009, [][2]int{{-99, -91}}},
		""},

	// The following two test cases have the same input, but different expectations. Uncomment just one or the other.

	/*In case you reverse the *digits* you could have the following cases:
	- the zero has to be considered
	{-2, 2,
		Product{0, [][2]int{{-2, 0}, {-1, 0}, {0, 0}, {0, 1}, {0, 2}}},
		Product{4, [][2]int{{-2, -2}, {2, 2}}},
		""}, */

	// - you can keep the minus sign in place
	{-2, 2,
		Product{-4, [][2]int{{-2, 2}}},
		Product{4, [][2]int{{-2, -2}, {2, 2}}},
		""},
}

func TestPalindromeProducts(t *testing.T) {
	// Uncomment the following line to add the bonus test to the default tests
	// testData = append(testData, bonusData...)
	for _, test := range testData {
		// common preamble for test failures
		ret := fmt.Sprintf("Products(%d, %d) returned",
			test.fmin, test.fmax)
		// test
		pmin, pmax, err := Products(test.fmin, test.fmax)
		// we check if err is of error type
		var _ error = err
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
