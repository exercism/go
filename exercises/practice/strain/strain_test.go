package strain

import (
	"reflect"
	"strings"
	"testing"
)

type testCase[T any] struct {
	description string
	filterFunc  func(T) bool
	list        []T
	expected    []T
}

var keepIntTests = []testCase[int]{
	{
		description: "keep on empty list returns empty list",
		filterFunc:  func(int) bool { return true },
		list:        []int{},
		expected:    []int{},
	},
	{
		description: "keeps everything",
		filterFunc:  func(int) bool { return true },
		list:        []int{1, 3, 5},
		expected:    []int{1, 3, 5},
	},
	{
		description: "keeps nothing",
		filterFunc:  func(int) bool { return false },
		list:        []int{1, 3, 5},
		expected:    []int{},
	},
	{
		description: "keeps first and last",
		filterFunc:  func(x int) bool { return x%2 == 1 },
		list:        []int{1, 2, 3},
		expected:    []int{1, 3},
	},
	{
		description: "keeps neither first nor last",
		filterFunc:  func(x int) bool { return x%2 == 0 },
		list:        []int{1, 2, 3},
		expected:    []int{2},
	},
}

func TestKeep(t *testing.T) {
	for _, tc := range keepIntTests {
		t.Run(tc.description, func(t *testing.T) {
			// setup here copies test.list, preserving the nil value if it is nil
			// and making a fresh copy of the underlying array otherwise.
			cp := tc.list
			if len(cp) != 0 {
				cp = append([]int{}, cp...)
			}

			got := Keep(cp, tc.filterFunc)

			switch {
			case len(tc.expected) == 0 && len(got) == 0:
				// We don't want to make this exercise about nil vs empty slice so we accept both.
				return
			case !reflect.DeepEqual(cp, tc.list):
				t.Fatalf("Keep(%#v, fn) should not modify its argument. %#v should stay %#v", tc.list, cp, tc.list)
			case !reflect.DeepEqual(got, tc.expected):
				t.Fatalf("Keep(%#v, fn)\n got: %#v\nwant: %#v", tc.list, got, tc.expected)
			}
		})
	}

	t.Run("keeps strings", func(t *testing.T) {
		zword := func(x string) bool { return strings.HasPrefix(x, "z") }
		list := []string{"apple", "zebra", "banana", "zombies", "cherimoya", "zealot"}
		expected := []string{"zebra", "zombies", "zealot"}

		cp := append([]string{}, list...) // make copy, as with ints
		got := Keep(cp, zword)

		switch {
		case !reflect.DeepEqual(cp, list):
			t.Fatalf("Keep(%#v, fn) should not modify its argument. %#v should stay %#v", list, cp, list)
		case !reflect.DeepEqual(got, expected):
			t.Fatalf("Keep(%#v, fn)\n got: %#v\nwant: %#v", list, got, expected)
		}
	})

	t.Run("keeps lists", func(t *testing.T) {
		has5 := func(list []int) bool {
			for _, entry := range list {
				if entry == 5 {
					return true
				}
			}
			return false
		}

		list := [][]int{
			{1, 2, 3},
			{5, 5, 5},
			{5, 1, 2},
			{2, 1, 2},
			{1, 5, 2},
			{2, 2, 1},
			{1, 2, 5},
		}
		expected := [][]int{
			{5, 5, 5},
			{5, 1, 2},
			{1, 5, 2},
			{1, 2, 5},
		}

		cp := append([][]int{}, list...)
		got := Keep(cp, has5)

		switch {
		case !reflect.DeepEqual(cp, list):
			t.Fatalf("Keep(%#v, fn) should not modify its argument. %#v should stay %#v", list, cp, list)
		case !reflect.DeepEqual(got, expected):
			t.Fatalf("Keep(%#v, fn)\n got: %#v\nwant: %#v", list, got, expected)
		}
	})
}

var discardIntTests = []testCase[int]{
	{
		description: "discard on empty list returns empty list",
		filterFunc:  func(int) bool { return true },
		list:        []int{},
		expected:    []int{},
	},
	{
		description: "discards everything",
		filterFunc:  func(int) bool { return true },
		list:        []int{1, 3, 5},
		expected:    []int{},
	},
	{
		description: "discards nothing",
		filterFunc:  func(int) bool { return false },
		list:        []int{1, 3, 5},
		expected:    []int{1, 3, 5},
	},
	{
		description: "discards first and last",
		filterFunc:  func(x int) bool { return x%2 == 1 },
		list:        []int{1, 2, 3},
		expected:    []int{2},
	},
	{
		description: "discards neither first nor last",
		filterFunc:  func(x int) bool { return x%2 == 0 },
		list:        []int{1, 2, 3},
		expected:    []int{1, 3},
	},
}

func TestDiscard(t *testing.T) {
	for _, tc := range discardIntTests {
		t.Run(tc.description, func(t *testing.T) {
			// setup here copies test.list, preserving the nil value if it is nil
			// and making a fresh copy of the underlying array otherwise.
			cp := tc.list
			if len(cp) != 0 {
				cp = append([]int{}, cp...)
			}

			got := Discard(cp, tc.filterFunc)

			switch {
			case len(tc.expected) == 0 && len(got) == 0:
				// We don't want to make this exercise about nil vs empty slice so we accept both.
				return
			case !reflect.DeepEqual(cp, tc.list):
				t.Fatalf("Discard(%#v, fn) should not modify its argument. %#v should stay %#v", tc.list, cp, tc.list)
			case !reflect.DeepEqual(got, tc.expected):
				t.Fatalf("Discard(%#v, fn)\n got: %#v\nwant: %#v", tc.list, got, tc.expected)
			}
		})
	}

	t.Run("discards strings", func(t *testing.T) {
		zword := func(x string) bool { return strings.HasPrefix(x, "z") }
		list := []string{"apple", "zebra", "banana", "zombies", "cherimoya", "zealot"}
		expected := []string{"apple", "banana", "cherimoya"}

		cp := append([]string{}, list...) // make copy, as with ints
		got := Discard(cp, zword)

		switch {
		case !reflect.DeepEqual(cp, list):
			t.Fatalf("Discard(%#v, fn) should not modify its argument. %#v should stay %#v", list, cp, list)
		case !reflect.DeepEqual(got, expected):
			t.Fatalf("Discard(%#v, fn)\n got: %#v\nwant: %#v", list, got, expected)
		}
	})

	t.Run("discards lists", func(t *testing.T) {
		has5 := func(list []int) bool {
			for _, entry := range list {
				if entry == 5 {
					return true
				}
			}
			return false
		}

		list := [][]int{
			{1, 2, 3},
			{5, 5, 5},
			{5, 1, 2},
			{2, 1, 2},
			{1, 5, 2},
			{2, 2, 1},
			{1, 2, 5},
		}
		expected := [][]int{
			{1, 2, 3},
			{2, 1, 2},
			{2, 2, 1},
		}

		cp := append([][]int{}, list...)
		got := Discard(cp, has5)

		switch {
		case !reflect.DeepEqual(cp, list):
			t.Fatalf("Discard(%#v, fn) should not modify its argument. %#v should stay %#v", list, cp, list)
		case !reflect.DeepEqual(got, expected):
			t.Fatalf("Discard(%#v, fn)\n got: %#v\nwant: %#v", list, got, expected)
		}
	})
}

func BenchmarkKeep(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range keepIntTests {
			Keep(test.list, test.filterFunc)
		}
	}
}

func BenchmarkDiscard(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range keepIntTests {
			Discard(test.list, test.filterFunc)
		}
	}
}
