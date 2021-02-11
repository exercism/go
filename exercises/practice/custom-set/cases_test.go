package stringset

// Source: exercism/problem-specifications
// Commit: 1ef368e custom-set: apply "input" policy
// Problem Specifications Version: 1.3.0

// Returns true if the set contains no elements
var emptyCases = []unaryBoolCase{
	{ // sets with no elements are empty
		[]string{},
		true,
	},
	{ // sets with elements are not empty
		[]string{"a"},
		false,
	},
}

// Sets can report if they contain an element
var containsCases = []eleBoolCase{

	{ // nothing is contained in an empty set
		[]string{},
		"a",
		false,
	},
	{ // when the element is in the set
		[]string{"a", "b", "c"},
		"a",
		true,
	},
	{ // when the element is not in the set
		[]string{"a", "b", "c"},
		"d",
		false,
	},
}

// A set is a subset if all of its elements are contained in the other set
var subsetCases = []binBoolCase{

	{ // empty set is a subset of another empty set
		[]string{},
		[]string{},
		true,
	},
	{ // empty set is a subset of non-empty set
		[]string{},
		[]string{"a"},
		true,
	},
	{ // non-empty set is not a subset of empty set
		[]string{"a"},
		[]string{},
		false,
	},
	{ // set is a subset of set with exact same elements
		[]string{"a", "b", "c"},
		[]string{"a", "b", "c"},
		true,
	},
	{ // set is a subset of larger set with same elements
		[]string{"a", "b", "c"},
		[]string{"d", "a", "b", "c"},
		true,
	},
	{ // set is not a subset of set that does not contain its elements
		[]string{"a", "b", "c"},
		[]string{"d", "a", "c"},
		false,
	},
}

// Sets are disjoint if they share no elements
var disjointCases = []binBoolCase{

	{ // the empty set is disjoint with itself
		[]string{},
		[]string{},
		true,
	},
	{ // empty set is disjoint with non-empty set
		[]string{},
		[]string{"a"},
		true,
	},
	{ // non-empty set is disjoint with empty set
		[]string{"a"},
		[]string{},
		true,
	},
	{ // sets are not disjoint if they share an element
		[]string{"a", "b"},
		[]string{"b", "c"},
		false,
	},
	{ // sets are disjoint if they share no elements
		[]string{"a", "b"},
		[]string{"c", "d"},
		true,
	},
}

// Sets with the same elements are equal
var equalCases = []binBoolCase{

	{ // empty sets are equal
		[]string{},
		[]string{},
		true,
	},
	{ // empty set is not equal to non-empty set
		[]string{},
		[]string{"a", "b", "c"},
		false,
	},
	{ // non-empty set is not equal to empty set
		[]string{"a", "b", "c"},
		[]string{},
		false,
	},
	{ // sets with the same elements are equal
		[]string{"a", "b"},
		[]string{"b", "a"},
		true,
	},
	{ // sets with different elements are not equal
		[]string{"a", "b", "c"},
		[]string{"a", "b", "d"},
		false,
	},
	{ // set is not equal to larger set with same elements
		[]string{"a", "b", "c"},
		[]string{"a", "b", "c", "d"},
		false,
	},
}

// Unique elements can be added to a set
var addCases = []eleOpCase{

	{ // add to empty set
		[]string{},
		"c",
		[]string{"c"},
	},
	{ // add to non-empty set
		[]string{"a", "b", "d"},
		"c",
		[]string{"a", "b", "c", "d"},
	},
	{ // adding an existing element does not change the set
		[]string{"a", "b", "c"},
		"c",
		[]string{"a", "b", "c"},
	},
}

// Intersection returns a set of all shared elements
var intersectionCases = []binOpCase{

	{ // intersection of two empty sets is an empty set
		[]string{},
		[]string{},
		[]string{},
	},
	{ // intersection of an empty set and non-empty set is an empty set
		[]string{},
		[]string{"c", "b", "e"},
		[]string{},
	},
	{ // intersection of a non-empty set and an empty set is an empty set
		[]string{"a", "b", "c", "d"},
		[]string{},
		[]string{},
	},
	{ // intersection of two sets with no shared elements is an empty set
		[]string{"a", "b", "c"},
		[]string{"d", "e", "f"},
		[]string{},
	},
	{ // intersection of two sets with shared elements is a set of the shared elements
		[]string{"a", "b", "c", "d"},
		[]string{"c", "b", "e"},
		[]string{"b", "c"},
	},
}

// Difference (or Complement) of a set is a set of all elements that are only in the first set
var differenceCases = []binOpCase{

	{ // difference of two empty sets is an empty set
		[]string{},
		[]string{},
		[]string{},
	},
	{ // difference of empty set and non-empty set is an empty set
		[]string{},
		[]string{"c", "b", "e"},
		[]string{},
	},
	{ // difference of a non-empty set and an empty set is the non-empty set
		[]string{"a", "b", "c", "d"},
		[]string{},
		[]string{"a", "b", "c", "d"},
	},
	{ // difference of two non-empty sets is a set of elements that are only in the first set
		[]string{"c", "b", "a"},
		[]string{"b", "d"},
		[]string{"a", "c"},
	},
}

// Union returns a set of all elements in either set
var unionCases = []binOpCase{

	{ // union of empty sets is an empty set
		[]string{},
		[]string{},
		[]string{},
	},
	{ // union of an empty set and non-empty set is the non-empty set
		[]string{},
		[]string{"b"},
		[]string{"b"},
	},
	{ // union of a non-empty set and empty set is the non-empty set
		[]string{"a", "c"},
		[]string{},
		[]string{"a", "c"},
	},
	{ // union of non-empty sets contains all unique elements
		[]string{"a", "c"},
		[]string{"b", "c"},
		[]string{"c", "b", "a"},
	},
}
