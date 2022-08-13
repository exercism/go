package stringset

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 42dd0ce Remove version (#1678)

type (
	// unary function, bool result (IsEmpty)
	unaryBoolCase struct {
		description string
		set         []string
		want        bool
	}
	// set-element function, bool result (Has)
	eleBoolCase struct {
		description string
		set         []string
		element     string
		want        bool
	}
	// binary function, bool result (Subset, Disjoint, Equal)
	binBoolCase struct {
		description string
		set1        []string
		set2        []string
		want        bool
	}
	// set-element operator (Add)
	eleOpCase struct {
		description string
		set         []string
		element     string
		want        []string
	}
	// set-set operator (Intersection, Difference, Union)
	binOpCase struct {
		description string
		set1        []string
		set2        []string
		want        []string
	}
)

// Returns true if the set contains no elements
var emptyCases = []unaryBoolCase{
	{
		description: "sets with no elements are empty",
		set:         []string{},
		want:        true,
	},
	{
		description: "sets with elements are not empty",
		set:         []string{"a"},
		want:        false,
	},
}

// Sets can report if they contain an element
var containsCases = []eleBoolCase{
	{
		description: "nothing is contained in an empty set",
		set:         []string{},
		element:     "a",
		want:        false,
	},
	{
		description: "when the element is in the set",
		set:         []string{"a", "b", "c"},
		element:     "a",
		want:        true,
	},
	{
		description: "when the element is not in the set",
		set:         []string{"a", "b", "c"},
		element:     "d",
		want:        false,
	},
}

// A set is a subset if all of its elements are contained in the other set
var subsetCases = []binBoolCase{
	{
		description: "empty set is a subset of another empty set",
		set1:        []string{},
		set2:        []string{},
		want:        true,
	},
	{
		description: "empty set is a subset of non-empty set",
		set1:        []string{},
		set2:        []string{"a"},
		want:        true,
	},
	{
		description: "non-empty set is not a subset of empty set",
		set1:        []string{"a"},
		set2:        []string{},
		want:        false,
	},
	{
		description: "set is a subset of set with exact same elements",
		set1:        []string{"a", "b", "c"},
		set2:        []string{"a", "b", "c"},
		want:        true,
	},
	{
		description: "set is a subset of larger set with same elements",
		set1:        []string{"a", "b", "c"},
		set2:        []string{"d", "a", "b", "c"},
		want:        true,
	},
	{
		description: "set is not a subset of set that does not contain its elements",
		set1:        []string{"a", "b", "c"},
		set2:        []string{"d", "a", "c"},
		want:        false,
	},
}

// Sets are disjoint if they share no elements
var disjointCases = []binBoolCase{
	{
		description: "the empty set is disjoint with itself",
		set1:        []string{},
		set2:        []string{},
		want:        true,
	},
	{
		description: "empty set is disjoint with non-empty set",
		set1:        []string{},
		set2:        []string{"a"},
		want:        true,
	},
	{
		description: "non-empty set is disjoint with empty set",
		set1:        []string{"a"},
		set2:        []string{},
		want:        true,
	},
	{
		description: "sets are not disjoint if they share an element",
		set1:        []string{"a", "b"},
		set2:        []string{"b", "c"},
		want:        false,
	},
	{
		description: "sets are disjoint if they share no elements",
		set1:        []string{"a", "b"},
		set2:        []string{"c", "d"},
		want:        true,
	},
}

// Sets with the same elements are equal
var equalCases = []binBoolCase{
	{
		description: "empty sets are equal",
		set1:        []string{},
		set2:        []string{},
		want:        true,
	},
	{
		description: "empty set is not equal to non-empty set",
		set1:        []string{},
		set2:        []string{"a", "b", "c"},
		want:        false,
	},
	{
		description: "non-empty set is not equal to empty set",
		set1:        []string{"a", "b", "c"},
		set2:        []string{},
		want:        false,
	},
	{
		description: "sets with the same elements are equal",
		set1:        []string{"a", "b"},
		set2:        []string{"b", "a"},
		want:        true,
	},
	{
		description: "sets with different elements are not equal",
		set1:        []string{"a", "b", "c"},
		set2:        []string{"a", "b", "d"},
		want:        false,
	},
	{
		description: "set is not equal to larger set with same elements",
		set1:        []string{"a", "b", "c"},
		set2:        []string{"a", "b", "c", "d"},
		want:        false,
	},
}

// Unique elements can be added to a set
var addCases = []eleOpCase{
	{
		description: "add to empty set",
		set:         []string{},
		element:     "c",
		want:        []string{"c"},
	},
	{
		description: "add to non-empty set",
		set:         []string{"a", "b", "d"},
		element:     "c",
		want:        []string{"a", "b", "c", "d"},
	},
	{
		description: "adding an existing element does not change the set",
		set:         []string{"a", "b", "c"},
		element:     "c",
		want:        []string{"a", "b", "c"},
	},
}

// Intersection returns a set of all shared elements
var intersectionCases = []binOpCase{
	{
		description: "intersection of two empty sets is an empty set",
		set1:        []string{},
		set2:        []string{},
		want:        []string{},
	},
	{
		description: "intersection of an empty set and non-empty set is an empty set",
		set1:        []string{},
		set2:        []string{"c", "b", "e"},
		want:        []string{},
	},
	{
		description: "intersection of a non-empty set and an empty set is an empty set",
		set1:        []string{"a", "b", "c", "d"},
		set2:        []string{},
		want:        []string{},
	},
	{
		description: "intersection of two sets with no shared elements is an empty set",
		set1:        []string{"a", "b", "c"},
		set2:        []string{"d", "e", "f"},
		want:        []string{},
	},
	{
		description: "intersection of two sets with shared elements is a set of the shared elements",
		set1:        []string{"a", "b", "c", "d"},
		set2:        []string{"c", "b", "e"},
		want:        []string{"b", "c"},
	},
}

// Difference (or Complement) of a set is a set of all elements that are only in the first set
var differenceCases = []binOpCase{
	{
		description: "difference of two empty sets is an empty set",
		set1:        []string{},
		set2:        []string{},
		want:        []string{},
	},
	{
		description: "difference of empty set and non-empty set is an empty set",
		set1:        []string{},
		set2:        []string{"c", "b", "e"},
		want:        []string{},
	},
	{
		description: "difference of a non-empty set and an empty set is the non-empty set",
		set1:        []string{"a", "b", "c", "d"},
		set2:        []string{},
		want:        []string{"a", "b", "c", "d"},
	},
	{
		description: "difference of two non-empty sets is a set of elements that are only in the first set",
		set1:        []string{"c", "b", "a"},
		set2:        []string{"b", "d"},
		want:        []string{"a", "c"},
	},
}

// Union returns a set of all elements in either set
var unionCases = []binOpCase{
	{
		description: "union of empty sets is an empty set",
		set1:        []string{},
		set2:        []string{},
		want:        []string{},
	},
	{
		description: "union of an empty set and non-empty set is the non-empty set",
		set1:        []string{},
		set2:        []string{"b"},
		want:        []string{"b"},
	},
	{
		description: "union of a non-empty set and empty set is the non-empty set",
		set1:        []string{"a", "c"},
		set2:        []string{},
		want:        []string{"a", "c"},
	},
	{
		description: "union of non-empty sets contains all unique elements",
		set1:        []string{"a", "c"},
		set2:        []string{"b", "c"},
		want:        []string{"c", "b", "a"},
	},
}
