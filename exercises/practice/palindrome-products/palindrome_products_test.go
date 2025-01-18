package palindrome

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

type testCase struct {
	description string
	// input to Products(): range limits for factors of the palindrome
	fmin, fmax int
	// output from Products():
	pmin, pmax Product // min and max palandromic products
	errPrefix  string  // start of text if there is an error, "" otherwise
}

var testCases = []testCase{
	{
		description: "valid limits 1-9",
		fmin:        1,
		fmax:        9,
		pmin:        Product{}, // zero value means don't bother to test it
		pmax:        Product{9, [][2]int{{1, 9}, {3, 3}}},
		errPrefix:   "",
	},
	{
		description: "valid limits 10-99",
		fmin:        10,
		fmax:        99,
		pmin:        Product{121, [][2]int{{11, 11}}},
		pmax:        Product{9009, [][2]int{{91, 99}}},
		errPrefix:   "",
	},
	{
		description: "valid limits 100-999",
		fmin:        100,
		fmax:        999,
		pmin:        Product{10201, [][2]int{{101, 101}}},
		pmax:        Product{906609, [][2]int{{913, 993}}},
		errPrefix:   "",
	},
	{
		description: "no palindromes",
		fmin:        4,
		fmax:        10,
		pmin:        Product{},
		pmax:        Product{},
		errPrefix:   "no palindromes",
	},
	{
		description: "fmin > fmax",
		fmin:        10,
		fmax:        4,
		pmin:        Product{},
		pmax:        Product{},
		errPrefix:   "fmin > fmax",
	},
}

// Bonus curiosities. Can a negative number be a palindrome? Most say no.
/*
var bonusData = []testCase{
	// The following two test cases have the same input, but different expectations. Uncomment just one or the other.
	// Here you can test that you can reach the limit of the largest palindrome made of two 2-digit numbers.
	//{
	//	description: "bonus test 1: error for negative limits",
	//	fmin:        -99,
	//	fmax:        -10,
	//	pmin:        Product{},
	//	pmax:        Product{},
	//	errPrefix:   "Negative limits",
	//},
	// You can still get non-negative products from negative factors.
	{
		description: "bonus test 1: no error for negative limits",
		fmin:        -99,
		fmax:        -10,
		pmin:        Product{121, [][2]int{{-11, -11}}},
		pmax:        Product{9009, [][2]int{{-99, -91}}},
		errPrefix:   "",
	},
	// The following two test cases have the same input, but different expectations. Uncomment just one or the other.
	//In case you reverse the *digits* you could have the following cases:
	//- the zero has to be considered
	//{
	//	description: "bonus test 2",
	//	fmin:        -2,
	//	fmax:        2,
	//	pmin:        Product{0, [][2]int{{-2, 0}, {-1, 0}, {0, 0}, {0, 1}, {0, 2}}},
	//	pmax:        Product{4, [][2]int{{-2, -2}, {2, 2}}},
	//	errPrefix:   "",
	//},
	// - you can keep the minus sign in place
	{
		description: "bonus test 2",
		fmin:        -2,
		fmax:        2,
		pmin:        Product{-4, [][2]int{{-2, 2}}},
		pmax:        Product{4, [][2]int{{-2, -2}, {2, 2}}},
		errPrefix:   "",
	},
}
*/

func TestPalindromeProducts(t *testing.T) {
	// Uncomment the following line and the bonusData var above to add the bonus test to the default tests
	// testData = append(testData, bonusData...)
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			pmin, pmax, err := Products(tc.fmin, tc.fmax)

			switch {
			case tc.errPrefix != "":
				if err == nil {
					t.Fatalf("Products(%d, %d) expected error %q, got nil", tc.fmin, tc.fmax, fmt.Sprintf("%s...", tc.errPrefix))
				}
				if !strings.HasPrefix(err.Error(), tc.errPrefix) {
					t.Fatalf("Products(%d, %d) expected error with prefix %q, got: %q", tc.fmin, tc.fmax, tc.errPrefix, err.Error())
				}
			case err != nil:
				t.Fatalf("Products(%d, %d) returned unexpected error: %v", tc.fmin, tc.fmax, err)
			}

			matchProd := func(field string, have, want Product) {
				if len(want.Factorizations) > 0 && // option to skip test
					!reflect.DeepEqual(have, want) {
					t.Fatalf("Products(%d, %d) [%s] = %v, want: %v", tc.fmin, tc.fmax, field, have, want)
				}
			}
			matchProd("pmin", pmin, tc.pmin)
			matchProd("pmax", pmax, tc.pmax)
		})
	}
}

func BenchmarkPalindromeProducts(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			Products(test.fmin, test.fmax)
		}
	}
}
