// Define two functions:  (Two? Yes, sometimes we ask more out of Go.)
//
//    // All returns a list of all substrings of s with length n.
//    All(n int, s string) []string
//
//    // UnsafeFirst returns the first substring of s with length n.
//    UnsafeFirst(n int, s string) string
//
// At this point you could consider this exercise complete and move on.
// But wait, maybe you ask a reasonable question:
// Why is the function called **Unsafe** First?
// If you are interested, read on for a bonus exercise.
//
// Bonus exercise:
//
// Once you get `go test` passing, try `go test -tags asktoomuch`.
// This uses a *build tag* to enable a test that wasn't enabled before.
// You can read more about those at https://golang.org/pkg/go/build/#hdr-Build_Constraints
//
// You may notice that you can't make this asktoomuch test happy.
// We need a way to signal that in some cases you can't take the first N characters of the string.
// UnsafeFirst can't do that since it only returns a string.
//
// To fix that, let's add another return value to the function.  Define
//
//    First(int, string) (first string, ok bool)
//
// and test with `go test -tags first`.

package series

import (
	"reflect"
	"testing"
)

var allTests = []struct {
	n   int
	s   string
	out []string
}{
	{1, "01234",
		[]string{"0", "1", "2", "3", "4"}},
	{1, "92834",
		[]string{"9", "2", "8", "3", "4"}},
	{2, "01234",
		[]string{"01", "12", "23", "34"}},
	{2, "98273463",
		[]string{"98", "82", "27", "73", "34", "46", "63"}},
	{2, "37103",
		[]string{"37", "71", "10", "03"}},
	{3, "01234",
		[]string{"012", "123", "234"}},
	{3, "31001",
		[]string{"310", "100", "001"}},
	{3, "982347",
		[]string{"982", "823", "234", "347"}},
	{4, "01234",
		[]string{"0123", "1234"}},
	{4, "91274",
		[]string{"9127", "1274"}},
	{5, "01234",
		[]string{"01234"}},
	{5, "81228",
		[]string{"81228"}},
	{6, "01234", nil},
	{len(cx) + 1, cx, nil},
}

var cx = "01032987583"

func TestAll(t *testing.T) {
	for _, test := range allTests {
		switch res := All(test.n, test.s); {
		case len(res) == 0 && len(test.out) == 0:
		case reflect.DeepEqual(res, test.out):
		default:
			t.Fatalf("All(%d, %q) = %q, want %q.",
				test.n, test.s, res, test.out)
		}
	}
}

func TestUnsafeFirst(t *testing.T) {
	for _, test := range allTests {
		if len(test.out) == 0 {
			continue
		}
		if res := UnsafeFirst(test.n, test.s); res != test.out[0] {
			t.Fatalf("UnsafeFirst(%d, %q) = %q, want %q.",
				test.n, test.s, res, test.out[0])
		}
	}
}
