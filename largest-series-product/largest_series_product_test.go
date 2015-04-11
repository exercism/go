package lsproduct

import "testing"

// Define a function LargestSeriesProduct(string, int) (int64, error).
//
// Also define an exported TestVersion with a value that matches
// the internal testVersion here.

const testVersion = 1

// Retired testVersions
// (none) 1c5d360b98fc3a9f59c1e5e2f3a668ff8438419d

var tests = []struct {
	digits  string
	span    int
	product int64
	ok      bool
}{
	{"0123456789",
		2, 72, true},
	{"12",
		2, 2, true},
	{"19",
		2, 9, true},
	{"576802143",
		2, 48, true},
	{"0123456789",
		3, 504, true},
	{"1027839564",
		3, 270, true},
	{"0123456789",
		5, 15120, true},
	{"73167176531330624919225119674426574742355349194934",
		6, 23520, true},
	{"52677741234314237566414902593461595376319419139427",
		6, 28350, true},
	{"",
		0, 1, true},
	{"123",
		4, 0, false},
	{pe1k, 5, 40824, true},        // original PE problem
	{pe1k, 13, 23514624000, true}, // new PE problem
}

const pe1k = "73167176531330624919225119674426574742355349194934" +
	"96983520312774506326239578318016984801869478851843" +
	"85861560789112949495459501737958331952853208805511" +
	"12540698747158523863050715693290963295227443043557" +
	"66896648950445244523161731856403098711121722383113" +
	"62229893423380308135336276614282806444486645238749" +
	"30358907296290491560440772390713810515859307960866" +
	"70172427121883998797908792274921901699720888093776" +
	"65727333001053367881220235421809751254540594752243" +
	"52584907711670556013604839586446706324415722155397" +
	"53697817977846174064955149290862569321978468622482" +
	"83972241375657056057490261407972968652414535100474" +
	"82166370484403199890008895243450658541227588666881" +
	"16427171479924442928230863465674813919123162824586" +
	"17866458359124566529476545682848912883142607690042" +
	"24219022671055626321111109370544217506941658960408" +
	"07198403850962455444362981230987879927244284909188" +
	"84580156166097919133875499200524063689912560717606" +
	"05886116467109405077541002256983155200055935729725" +
	"71636269561882670428252483600823257530420752963450"

func TestLargestSeriesProduct(t *testing.T) {
	if TestVersion != testVersion {
		t.Fatalf("Found TestVersion = %v, want %v", TestVersion, testVersion)
	}
	for _, test := range tests {
		p, err := LargestSeriesProduct(test.digits, test.span)
		switch {
		case err != nil:
			if test.ok {
				t.Fatalf("LargestSeriesProduct(%s, %d) returned error %q.  "+
					"Error not expected.",
					test.digits, test.span, err)
			}
		case !test.ok:
			t.Fatalf("LargestSeriesProduct(%s, %d) = %d, %v.  Expected error",
				test.digits, test.span, p, err)
		case p != test.product:
			t.Fatalf("LargestSeriesProduct(%s, %d) = %d, want %d",
				test.digits, test.span, p, test.product)
		}
	}
}

func BenchmarkLargestSeriesProduct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			LargestSeriesProduct(test.digits, test.span)
		}
	}
}
