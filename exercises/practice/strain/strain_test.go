package strain

import (
	"reflect"
	"testing"
)

func lt10(x int) bool { return x < 10 }
func gt10(x int) bool { return x > 10 }
func odd(x int) bool  { return x&1 == 1 }
func even(x int) bool { return x&1 == 0 }

type testCase struct {
	description string
	filterFunc  func(int) bool
	list        Ints
	expected    Ints
}

var keepTests = []testCase{
	{
		description: "keep lower than 10 without inputs",
		filterFunc:  lt10,
		list:        nil,
		expected:    nil,
	},
	{
		description: "keep lower than 10 with all 3 inputs kept",
		filterFunc:  lt10,
		list:        Ints{1, 2, 3},
		expected:    Ints{1, 2, 3},
	},
	{
		description: "keep odd with 2 odd an 1 even input",
		filterFunc:  odd,
		list:        Ints{1, 2, 3},
		expected:    Ints{1, 3},
	},
	{
		description: "keep even with 2 even and 3 odd inputs",
		filterFunc:  even,
		list:        Ints{1, 2, 3, 4, 5},
		expected:    Ints{2, 4},
	},
}

var discardTests = []testCase{
	{
		description: "discard lower than 10 without inputs",
		filterFunc:  lt10,
		list:        nil,
		expected:    nil},
	{
		description: "discard greater than 10 without discarded inputs",
		filterFunc:  gt10,
		list:        Ints{1, 2, 3},
		expected:    Ints{1, 2, 3},
	},
	{
		description: "discard odd with 2 odd and 1 even input",
		filterFunc:  odd,
		list:        Ints{1, 2, 3},
		expected:    Ints{2},
	},
	{
		description: "discard even with 2 odd and 3 even inputs",
		filterFunc:  even,
		list:        Ints{1, 2, 3, 4, 5},
		expected:    Ints{1, 3, 5},
	},
}

func TestKeepInts(t *testing.T) {
	for _, tc := range keepTests {
		t.Run(tc.description, func(t *testing.T) {
			// setup here copies test.list, preserving the nil value if it is nil
			// and making a fresh copy of the underlying array otherwise.
			cp := tc.list
			if cp != nil {
				cp = append(Ints{}, cp...)
			}
			got := cp.Keep(tc.filterFunc)
			switch {
			case !reflect.DeepEqual(cp, tc.list):
				t.Fatalf("%#v.Keep() should not modify its receiver. %#v should stay %#v", tc.list, cp, tc.list)
			case !reflect.DeepEqual(got, tc.expected):
				t.Fatalf("%#v.Keep()\n got: %#v\nwant: %#v", tc.list, got, tc.expected)
			}
		})
	}
}

func TestDiscardInts(t *testing.T) {
	for _, tc := range discardTests {
		t.Run(tc.description, func(t *testing.T) {
			cp := tc.list
			if cp != nil {
				cp = append(Ints{}, cp...) // dup underlying array
			}
			got := cp.Discard(tc.filterFunc)
			switch {
			case !reflect.DeepEqual(cp, tc.list):
				t.Fatalf("%#v.Discard() should not modify its receiver. %#v should stay %#v", tc.list, cp, tc.list)
			case !reflect.DeepEqual(got, tc.expected):
				t.Fatalf("%#v.Discard()\n got: %#v\nwant: %#v", tc.list, got, tc.expected)
			}
		})
	}
}

func TestKeepStrings(t *testing.T) {
	zword := func(s string) bool { return len(s) > 0 && s[0] == 'z' }
	list := Strings{"apple", "zebra", "banana", "zombies", "cherimoya", "zealot"}
	want := Strings{"zebra", "zombies", "zealot"}

	cp := append(Strings{}, list...) // make copy, as with TestInts
	got := cp.Keep(zword)
	switch {
	case !reflect.DeepEqual(cp, list):
		t.Fatalf("%#v.Keep() should not modify its receiver. %#v should stay %#v", list, cp, list)
	case !reflect.DeepEqual(got, want):
		t.Fatalf("%#v.Keep()\n got: %#v\nwant: %#v", list, got, want)
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
	got := cp.Keep(has5)
	switch {
	case !reflect.DeepEqual(cp, list):
		t.Fatalf("%#v.Keep() should not modify its receiver. %#v should stay %#v", list, cp, list)
	case !reflect.DeepEqual(got, want):
		t.Fatalf("%#v.Keep()\n got: %#v\nwant: %#v", list, got, want)
	}
}

func BenchmarkKeepInts(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range keepTests {
			test.list.Keep(test.filterFunc)
		}
	}
}

func BenchmarkDiscardInts(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range discardTests {
			test.list.Discard(test.filterFunc)
		}
	}
}
