package genericset

import "testing"

func TestNew(t *testing.T) {
	if got := NewSet[int](); got == nil {
		t.Errorf("NewSet() = nil, want a Set.")
	}
}

func TestEmptyEqual(t *testing.T) {
	setOne := NewSet[int]()
	setTwo := NewSet[int]()
	if !setOne.Equal(setTwo) {
		t.Errorf("setOne.Equal(setTwo) = false, want true.")
	}
}

type addCases[T comparable] struct {
	description string
	addOne      []T
	addTwo      []T
	want        bool
}

var addStringCases = []addCases[string]{
	{
		description: "Two empty Set[string] should be equal.",
		addOne:      []string{},
		addTwo:      []string{},
		want:        true,
	},
	{
		description: "Two Set[string] with one element each should be equal.",
		addOne:      []string{"one"},
		addTwo:      []string{"one"},
		want:        true,
	},
	{
		description: "Two Set[string] with two elements each should be equal.",
		addOne:      []string{"one", "two"},
		addTwo:      []string{"one", "two"},
		want:        true,
	},
	{
		description: "Changing the Add order should not make Sets unequal.",
		addOne:      []string{"one", "two"},
		addTwo:      []string{"two", "one"},
		want:        true,
	},
	{
		description: "Two Set[string] with one different element each should be unequal.",
		addOne:      []string{"one"},
		addTwo:      []string{"two"},
		want:        false,
	},
	{
		description: "Two Set[string] with more elements in the second set should be unequal.",
		addOne:      []string{"one"},
		addTwo:      []string{"one", "two"},
		want:        false,
	},
	{
		description: "Two Set[string] with more elements in the first set should be unequal.",
		addOne:      []string{"one", "two"},
		addTwo:      []string{"two"},
		want:        false,
	},
	{
		description: "Two Set[string] with partial overlap should be unequal.",
		addOne:      []string{"one", "two"},
		addTwo:      []string{"zero", "two"},
		want:        false,
	},
}

func TestAddString(t *testing.T) {
	for _, tc := range addStringCases {
		t.Run(tc.description, func(t *testing.T) {
			one := NewSet[string]()
			for _, i := range tc.addOne {
				one.Add(i)
			}
			two := NewSet[string]()
			for _, i := range tc.addTwo {
				two.Add(i)
			}
			if got := one.Equal(two); got != tc.want {
				t.Errorf("%v.Equal(%v) = %t, want %t", one, two, got, tc.want)
			}
		})
	}
}

var addIntCases = []addCases[int]{
	{
		description: "Two empty Set[int] should be equal.",
		addOne:      []int{},
		addTwo:      []int{},
		want:        true,
	},
	{
		description: "Two Set[int] with one element each should be equal.",
		addOne:      []int{1},
		addTwo:      []int{1},
		want:        true,
	},
	{
		description: "Two Set[int] with two elements each should be equal.",
		addOne:      []int{1, 2},
		addTwo:      []int{1, 2},
		want:        true,
	},
	{
		description: "Changing the Add order should not make Sets unequal.",
		addOne:      []int{1, 2},
		addTwo:      []int{2, 1},
		want:        true,
	},
	{
		description: "Two Set[int] with one different element each should be unequal.",
		addOne:      []int{1},
		addTwo:      []int{2},
		want:        false,
	},
	{
		description: "Two Set[int] with more elements in the second set should be unequal.",
		addOne:      []int{1},
		addTwo:      []int{1, 2},
		want:        false,
	},
	{
		description: "Two Set[int] with more elements in the first set should be unequal.",
		addOne:      []int{1, 2},
		addTwo:      []int{2},
		want:        false,
	},
	{
		description: "Two Set[int] with partial overlap should be unequal.",
		addOne:      []int{1, 2},
		addTwo:      []int{0, 2},
		want:        false,
	},
}

func TestAddInt(t *testing.T) {
	for _, tc := range addIntCases {
		t.Run(tc.description, func(t *testing.T) {
			one := NewSet[int]()
			for _, i := range tc.addOne {
				one.Add(i)
			}
			two := NewSet[int]()
			for _, i := range tc.addTwo {
				two.Add(i)
			}
			if got := one.Equal(two); got != tc.want {
				t.Errorf("%v.Equal(%v) = %t, want %t", one, two, got, tc.want)
			}
		})
	}
}

type hasCases[T comparable] struct {
	description string
	add         []T
	check       T
	want        bool
}

var hasIntCases = []hasCases[int]{
	{
		description: "Empty Set does not have 1.",
		add:         []int{},
		check:       1,
		want:        false,
	},
	{
		description: "Set {1} does have 1.",
		add:         []int{1},
		check:       1,
		want:        true,
	},
	{
		description: "Set {1, 2, 3, 4, 5} does have 3.",
		add:         []int{1, 2, 3, 4, 5},
		check:       3,
		want:        true,
	},
	{
		description: "Set {1, 2, 3, 4, 5} does not have 7.",
		add:         []int{1, 2, 3, 4, 5},
		check:       7,
		want:        false,
	},
}

var hasStringCases = []hasCases[string]{
	{
		description: "Empty Set does not have 'dog'.",
		add:         []string{},
		check:       "dog",
		want:        false,
	},
	{
		description: "Set {'dog'} does have 'dog'.",
		add:         []string{"dog"},
		check:       "dog",
		want:        true,
	},
	{
		description: "Set {'cat', 'dog', 'mouse'} does have 'cat'.",
		add:         []string{"cat", "dog", "mouse"},
		check:       "cat",
		want:        true,
	},
	{
		description: "Set {'cat', 'dog', 'mouse'} does not have 'bird'.",
		add:         []string{"cat", "dog", "mouse"},
		check:       "bird",
		want:        false,
	},
}

func TestHasString(t *testing.T) {
	for _, tc := range hasStringCases {
		t.Run(tc.description, func(t *testing.T) {
			set := NewSet[string]()
			for _, i := range tc.add {
				set.Add(i)
			}
			if got := set.Has(tc.check); got != tc.want {
				t.Errorf("%v.Has(%v) = %t, want %t", set, tc.check, got, tc.want)
			}
		})
	}
}

func TestHasInt(t *testing.T) {
	for _, tc := range hasIntCases {
		t.Run(tc.description, func(t *testing.T) {
			set := NewSet[int]()
			for _, i := range tc.add {
				set.Add(i)
			}
			if got := set.Has(tc.check); got != tc.want {
				t.Errorf("%v.Has(%d) = %t, want %t", set, tc.check, got, tc.want)
			}
		})
	}
}

var removeCases = []struct {
	description string
	add         []int
	remove      []int
	check       int
	want        bool
}{
	{
		description: "Remove 1 from {1}, does not have 1.",
		add:         []int{1},
		remove:      []int{1},
		check:       1,
		want:        false,
	},
	{
		description: "Remove 1 from {1, 2}, does not have 1.",
		add:         []int{1, 2},
		remove:      []int{1},
		check:       1,
		want:        false,
	},
	{
		description: "Remove 1 from {1, 2}, does have 2.",
		add:         []int{1, 2},
		remove:      []int{1},
		check:       2,
		want:        true,
	},
	{
		description: "Remove 1, 2 from {1, 2}, does not have 1.",
		add:         []int{1, 2},
		remove:      []int{1, 2},
		check:       1,
		want:        false,
	},
	{
		description: "Remove 1, 2 from {1, 2}, does not have 2.",
		add:         []int{1, 2},
		remove:      []int{1, 2},
		check:       2,
		want:        false,
	},
}

func TestRemove(t *testing.T) {
	for _, tc := range removeCases {
		t.Run(tc.description, func(t *testing.T) {
			set := NewSet[int]()
			for _, i := range tc.add {
				set.Add(i)
			}
			for _, i := range tc.remove {
				set.Remove(i)
			}
			if got := set.Has(tc.check); got != tc.want {
				t.Errorf("%v.Has(%d) = %t, want %t", set, tc.check, got, tc.want)
			}
		})
	}
}

var isEmptyCases = []struct {
	description string
	add         []int
	remove      []int
	want        bool
}{
	{
		description: "Empty Set is empty.",
		add:         []int{},
		remove:      []int{1},
		want:        true,
	},
	{
		description: "Remove 1 from {1, 2}, not empty.",
		add:         []int{1, 2},
		remove:      []int{1},
		want:        false,
	},
	{
		description: "Remove 2 from {1, 2}, not empty.",
		add:         []int{1, 2},
		remove:      []int{2},
		want:        false,
	},
	{
		description: "Remove 1, 2 from {1, 2}, is empty.",
		add:         []int{1, 2},
		remove:      []int{1, 2},
		want:        true,
	},
}

func TestEmpty(t *testing.T) {
	for _, tc := range isEmptyCases {
		t.Run(tc.description, func(t *testing.T) {
			set := NewSet[int]()
			for _, i := range tc.add {
				set.Add(i)
			}
			for _, i := range tc.remove {
				set.Remove(i)
			}
			if got := set.IsEmpty(); got != tc.want {
				t.Errorf("%v.IsEmpty() = %t, want %t", set, got, tc.want)
			}
		})
	}
}
