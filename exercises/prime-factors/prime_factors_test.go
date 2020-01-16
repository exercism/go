package prime

import (
	"reflect"
	"sort"
	"testing"
)

func TestPrimeFactors(t *testing.T) {
	for _, test := range tests {
		actual := Factors(test.input)
		sort.Sort(ascending(actual))
		if !reflect.DeepEqual(actual, test.expected) {
			t.Fatalf("FAIL %s\nFactors(%d) = %#v;\nexpected %#v",
				test.description, test.input,
				actual, test.expected)
		}
		t.Logf("PASS %s", test.description)
	}
}

func BenchmarkPrimeFactors(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			Factors(test.input)
		}
	}
}

type ascending []int64

func (l ascending) Len() int             { return len(l) }
func (l ascending) Swap(ii, jj int)      { l[ii], l[jj] = l[jj], l[ii] }
func (l ascending) Less(ii, jj int) bool { return l[ii] < l[jj] }
