package palindromeproducts

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

// Bonus curiosities. Can a negative number be a palindrome? Most say no.
// Toggle this to run bonus tests.
const RunBonusTests = false

var bonusData = []testCase{
	// The following test cases have the same input, but different expectations. Uncomment just one or the other.
	// Here you can test that you can reach the limit of the largest palindrome made of two 2-digit numbers.
	/*
		{
			description: "bonus test 1: error for negative limits",
			fmin:        -99,
			fmax:        -10,
			property:    "smallest",
			product:     0,
			factors:     [][2]int{},
			errorMsg:    "Negative limits",
		},
		{
			description: "bonus test 1: error for negative limits",
			fmin:        -99,
			fmax:        -10,
			property:    "largest",
			product:     0,
			factors:     [][2]int{},
			errorMsg:    "Negative limits",
		},
	*/
	// You can still get non-negative products from negative factors.
	{
		description: "bonus test 1: no error for negative limits",
		fmin:        -99,
		fmax:        -10,
		property:    "smallest",
		product:     121,
		factors:     [][2]int{{-11, -11}},
		errorMsg:    "",
	},
	{
		description: "bonus test 1: no error for negative limits",
		fmin:        -99,
		fmax:        -10,
		property:    "largest",
		product:     9009,
		factors:     [][2]int{{-99, -91}},
		errorMsg:    "",
	},
	// The following test cases have the same input, but different expectations for the smallest product.
	// (The largest product is the same.)
	// Uncomment just one or the other.
	// In case you reverse the *digits* you could have the following cases:
	// - the zero has to be considered
	{
		description: "bonus test 2",
		fmin:        -2,
		fmax:        2,
		property:    "smallest",
		product:     0,
		factors:     [][2]int{{-2, 0}, {-1, 0}, {0, 0}, {0, 1}, {0, 2}},
		errorMsg:    "",
	},
	// - you can keep the minus sign in place
	/*
		{
			description: "bonus test 2",
			fmin:        -2,
			fmax:        2,
			property:    "smallest",
			product:     -4,
			factors:     [][2]int{{-2, 2}},
			errorMsg:    "",
		},
	*/
	{
		description: "bonus test 2",
		fmin:        -2,
		fmax:        2,
		property:    "largest",
		product:     4,
		factors:     [][2]int{{-2, -2}, {2, 2}},
		errorMsg:    "",
	},
}

type CacheResult struct {
	min, max Product
	err      error
}
type CachedProducts map[[2]int]CacheResult

// Return a cached result of Products.
func (cp CachedProducts) Products(min, max int) (Product, Product, error) {
	if _, ok := cp[[2]int{min, max}]; !ok {
		pmin, pmax, err := Products(min, max)
		cp[[2]int{min, max}] = CacheResult{pmin, pmax, err}
	}
	result := cp[[2]int{min, max}]
	return result.min, result.max, result.err
}

func TestPalindromeProducts(t *testing.T) {
	if RunBonusTests {
		testCases = append(testCases, bonusData...)
	}
	cp := make(CachedProducts)
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			pmin, pmax, err := cp.Products(tc.fmin, tc.fmax)

			if tc.errorMsg != "" {
				if err == nil {
					t.Fatalf("Products(%d, %d) expected error %q, got nil", tc.fmin, tc.fmax, fmt.Sprintf("%s...", tc.errorMsg))
				}
				if !strings.HasPrefix(err.Error(), tc.errorMsg) {
					t.Fatalf("Products(%d, %d) expected error with prefix %q, got: %q", tc.fmin, tc.fmax, tc.errorMsg, err.Error())
				}
			} else if err != nil {
				t.Fatalf("Products(%d, %d) returned unexpected error: %v", tc.fmin, tc.fmax, err)
			}

			matchProd := func(field string, have, want Product) {
				if len(want.Factorizations) > 0 && // option to skip test
					!reflect.DeepEqual(have, want) {
					t.Fatalf("Products(%d, %d) [%s] = %v, want: %v", tc.fmin, tc.fmax, field, have, want)
				}
			}
			if tc.property == "smallest" {
				matchProd("pmin", pmin, Product{tc.product, tc.factors})
			} else {
				matchProd("pmax", pmax, Product{tc.product, tc.factors})
			}
		})
	}
}

func BenchmarkPalindromeProducts(b *testing.B) {
	for range b.N {
		for _, test := range testCases {
			Products(test.fmin, test.fmax)
		}
	}
}
