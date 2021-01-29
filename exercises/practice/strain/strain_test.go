// Collections, hm?  For this exercise in Go you'll work with slices as
// collections.  Define the following in your solution:
//
//    type Ints []int
//    type Lists [][]int
//    type Strings []string
//
// Then complete the exercise by implementing these methods:
//
//    (Ints) Keep(func(int) bool) Ints
//    (Ints) Discard(func(int) bool) Ints
//    (Lists) Keep(func([]int) bool) Lists
//    (Strings) Keep(func(string) bool) Strings

package strain

import (
	"reflect"
	"testing"
)

func lt10(x int) bool { return x < 10 }
func gt10(x int) bool { return x > 10 }
func odd(x int) bool  { return x&1 == 1 }
func even(x int) bool { return x&1 == 0 }

var keepTests = []struct {
	pred func(int) bool
	list Ints
	want Ints
}{
	{lt10,
		nil,
		nil},
	{lt10,
		Ints{1, 2, 3},
		Ints{1, 2, 3}},
	{odd,
		Ints{1, 2, 3},
		Ints{1, 3}},
	{even,
		Ints{1, 2, 3, 4, 5},
		Ints{2, 4}},
}

var discardTests = []struct {
	pred func(int) bool
	list Ints
	want Ints
}{
	{lt10,
		nil,
		nil},
	{gt10,
		Ints{1, 2, 3},
		Ints{1, 2, 3}},
	{odd,
		Ints{1, 2, 3},
		Ints{2}},
	{even,
		Ints{1, 2, 3, 4, 5},
		Ints{1, 3, 5}},
}

func TestKeepInts(t *testing.T) {
	for _, test := range keepTests {
		// setup here copies test.list, preserving the nil value if it is nil
		// and making a fresh copy of the underlying array otherwise.
		cp := test.list
		if cp != nil {
			cp = append(Ints{}, cp...)
		}
		switch res := cp.Keep(test.pred); {
		case !reflect.DeepEqual(cp, test.list):
			t.Fatalf("%#v.Keep() should not modify its receiver.  "+
				"Found %#v, receiver should stay %#v",
				test.list, cp, test.list)
		case !reflect.DeepEqual(res, test.want):
			t.Fatalf("%#v.Keep()\ngot: %#v\nwant: %#v",
				test.list, res, test.want)
		}
	}
}

func TestDiscardInts(t *testing.T) {
	for _, test := range discardTests {
		cp := test.list
		if cp != nil {
			cp = append(Ints{}, cp...) // dup underlying array
		}
		switch res := cp.Discard(test.pred); {
		case !reflect.DeepEqual(cp, test.list):
			t.Fatalf("%#v.Discard() should not modify its receiver.  "+
				"Found %#v, receiver should stay %#v",
				test.list, cp, test.list)
		case !reflect.DeepEqual(res, test.want):
			t.Fatalf("%#v.Discard()\ngot: %#v\nwant: %#v",
				test.list, res, test.want)
		}
	}
}

func TestKeepStrings(t *testing.T) {
	zword := func(s string) bool { return len(s) > 0 && s[0] == 'z' }
	list := Strings{"apple", "zebra", "banana", "zombies", "cherimoya", "zealot"}
	want := Strings{"zebra", "zombies", "zealot"}

	cp := append(Strings{}, list...) // make copy, as with TestInts
	switch res := cp.Keep(zword); {
	case !reflect.DeepEqual(cp, list):
		t.Fatalf("%#v.Keep() should not modify its receiver.  "+
			"Found %#v, receiver should stay %#v",
			list, cp, list)
	case !reflect.DeepEqual(res, want):
		t.Fatalf("%#v.Keep()\ngot: %#v\nwant: %#v", list, res, want)
	}
}

func TestKeepLists(t *testing.T) {
	has5 := func(l []int) bool {
		for _, e := range l {
			if e == 5 {
				return true
			}
		}
		return false
	}
	list := Lists{
		{1, 2, 3},
		{5, 5, 5},
		{5, 1, 2},
		{2, 1, 2},
		{1, 5, 2},
		{2, 2, 1},
		{1, 2, 5},
	}
	want := Lists{
		{5, 5, 5},
		{5, 1, 2},
		{1, 5, 2},
		{1, 2, 5},
	}
	cp := append(Lists{}, list...)
	switch res := cp.Keep(has5); {
	case !reflect.DeepEqual(cp, list):
		t.Fatalf("%#v.Keep() should not modify its receiver.  "+
			"Found %#v, receiver should stay %#v",
			list, cp, list)
	case !reflect.DeepEqual(res, want):
		t.Fatalf("%#v.Keep()\ngot: %#v\nwant: %#v", list, res, want)
	}
}

func BenchmarkKeepInts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range keepTests {
			test.list.Keep(test.pred)
		}
	}
}

func BenchmarkDiscardInts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range discardTests {
			test.list.Discard(test.pred)
		}
	}
}
