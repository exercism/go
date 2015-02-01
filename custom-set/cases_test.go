package stringset

// Source: exercism/x-common
// Commit: 269f498 Merge pull request #48 from soniakeys/custom-set-json

// Test two sets for equality.
var eqCases = []binBoolCase{
	{ // order doesn't matter
		[]string{"a", "c"},
		[]string{"c", "a"},
		true,
	},
	{ // dupicates don't matter
		[]string{"a", "a"},
		[]string{"a"},
		true,
	},
	{ // empty sets are equal
		[]string{},
		[]string{},
		true,
	},
	{ // set with single element is equal to itself
		[]string{"a"},
		[]string{"a"},
		true,
	},
	{ // different sets are not equal
		[]string{"a", "b", "c"},
		[]string{"c", "d", "e"},
		false,
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
	{ // having most in common is not good enough
		[]string{"a", "b", "c", "d"},
		[]string{"b", "c", "d", "e"},
		false,
	},
}

// Add an element to a set.
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
	{ // add existing element
		[]string{"a", "b", "c"},
		"c",
		[]string{"a", "b", "c"},
	},
}

// Delete an element from a set.
var delCases = []eleOpCase{
	{ // delete an element
		[]string{"c", "b", "a"},
		"b",
		[]string{"a", "c"},
	},
	{ // delete an element not in set
		[]string{"c", "b", "a"},
		"d",
		[]string{"a", "b", "c"},
	},
}

// Test if is a set is empty.
var emptyCases = []unaryBoolCase{
	{ // empty
		[]string{},
		true,
	},
	{ // single element
		[]string{"a"},
		false,
	},
	{ // a few elements
		[]string{"a", "b", "c", "b"},
		false,
	},
}

// Return the cardinality of a set.
var lenCases = []unaryIntCase{
	{ // empty set
		[]string{},
		0,
	},
	{ // non-empty set
		[]string{"a", "b", "c"},
		3,
	},
	{ // duplicate element
		[]string{"a", "b", "c", "b"},
		3,
	},
}

// Test if a value is an element of a set.
var hasCases = []eleBoolCase{
	{ // nothing is an element of the empty set
		[]string{},
		"a",
		false,
	},
	{ // 1 is in the set
		[]string{"a", "b", "c", "b"},
		"a",
		true,
	},
	{ // 2 is in the set
		[]string{"a", "b", "c", "b"},
		"b",
		true,
	},
	{ // 3 is in the set
		[]string{"a", "b", "c", "b"},
		"c",
		true,
	},
	{ // 4 not in the set
		[]string{"a", "b", "c", "b"},
		"d",
		false,
	},
}

// Test if set1 is a subset of set2.
var subsetCases = []binBoolCase{
	{ // empty set is subset of itself
		[]string{},
		[]string{},
		true,
	},
	{ // empty set is subset of non-empty set
		[]string{},
		[]string{"a"},
		true,
	},
	{ // non-empty set is not subset of empty set
		[]string{"a"},
		[]string{},
		false,
	},
	{ // non-empty set is subset of itself
		[]string{"a", "b", "c"},
		[]string{"a", "b", "c"},
		true,
	},
	{ // proper subset
		[]string{"a", "b", "c"},
		[]string{"d", "a", "b", "c"},
		true,
	},
	{ // same number of elements
		[]string{"a", "b", "c"},
		[]string{"d", "a", "c"},
		false,
	},
	{ // superset
		[]string{"a", "b", "c", "d", "e"},
		[]string{"b", "c", "d"},
		false,
	},
	{ // fewer elements but not a subset
		[]string{"a", "b", "c", "k"},
		[]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"},
		false,
	},
}

// Test if two sets are disjoint.
var disjointCases = []binBoolCase{
	{ // the empty set is disjoint with itself
		[]string{},
		[]string{},
		true,
	},
	{ // empty set disjoint with non-empty set
		[]string{},
		[]string{"a"},
		true,
	},
	{ // non-empty set disjoint with empty set
		[]string{"a"},
		[]string{},
		true,
	},
	{ // one element in common
		[]string{"a", "b"},
		[]string{"b", "c"},
		false,
	},
	{ // no elements in common
		[]string{"a", "b"},
		[]string{"c", "d"},
		true,
	},
}

// Produce the union of two sets.
var unionCases = []binOpCase{
	{ // union of empty sets
		[]string{},
		[]string{},
		[]string{},
	},
	{ // union of empty set with set of one element
		[]string{},
		[]string{"b"},
		[]string{"b"},
	},
	{ // union of empty set with non-empty set
		[]string{},
		[]string{"c", "b", "e"},
		[]string{"b", "c", "e"},
	},
	{ // union of non-empty set with empty set
		[]string{"a", "c"},
		[]string{},
		[]string{"a", "c"},
	},
	{ // union of a set with itself
		[]string{"a", "c"},
		[]string{"c", "a"},
		[]string{"a", "c"},
	},
	{ // union with one element
		[]string{"a", "c"},
		[]string{"b"},
		[]string{"a", "b", "c"},
	},
	{ // one element in common, one different
		[]string{"a", "c"},
		[]string{"b", "c"},
		[]string{"c", "b", "a"},
	},
	{ // two elements in common
		[]string{"a", "b", "c", "d"},
		[]string{"c", "b", "e"},
		[]string{"a", "b", "c", "d", "e"},
	},
}

// Intersect two sets.
var intersectionCases = []binOpCase{
	{ // intersect empty sets
		[]string{},
		[]string{},
		[]string{},
	},
	{ // intersect empty set with non-empty set
		[]string{},
		[]string{"c", "b", "e"},
		[]string{},
	},
	{ // intersect non-empty set with empty set
		[]string{"a", "b", "c", "d"},
		[]string{},
		[]string{},
	},
	{ // intersect one element with itself
		[]string{"c"},
		[]string{"c"},
		[]string{"c"},
	},
	{ // one element in common, extra elements in both sets
		[]string{"a", "b", "c"},
		[]string{"c", "e", "d"},
		[]string{"c"},
	},
	{ // two elements in common, extras in both sets
		[]string{"a", "b", "c", "d"},
		[]string{"c", "b", "e"},
		[]string{"b", "c"},
	},
	{ // intersect with subset
		[]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"},
		[]string{"e", "f", "g", "h", "i", "j"},
		[]string{"e", "f", "g", "h", "i", "j"},
	},
	{ // nothing in common
		[]string{"a", "b", "c"},
		[]string{"d", "e", "f"},
		[]string{},
	},
}

// Produce the set difference (set1 - set2)
// or more specifically, (set1 ∖ set2)
var differenceCases = []binOpCase{
	{ // difference of two empty sets
		[]string{},
		[]string{},
		[]string{},
	},
	{ // difference of empty set and non-empty set
		[]string{},
		[]string{"c", "b", "e"},
		[]string{},
	},
	{ // difference of non-empty set and empty set
		[]string{"a", "b", "c", "d"},
		[]string{},
		[]string{"a", "b", "c", "d"},
	},
	{ // no elements in common
		[]string{"a", "b", "c"},
		[]string{"d"},
		[]string{"a", "b", "c"},
	},
	{ // one element in common, one extra
		[]string{"c", "b", "a"},
		[]string{"b", "d"},
		[]string{"a", "c"},
	},
	{ // two elements in common, one extra
		[]string{"a", "b", "c", "d"},
		[]string{"c", "b", "e"},
		[]string{"a", "d"},
	},
}

// Produce the symmetric difference of two sets.  The symmetric
// difference consists of elements in one or the other but not both.
var symmetricDifferenceCases = []binOpCase{
	{ // two empty sets
		[]string{},
		[]string{},
		[]string{},
	},
	{ // empty set and non-empty set
		[]string{},
		[]string{"c", "b", "e"},
		[]string{"c", "b", "e"},
	},
	{ // non-empty set and empty set
		[]string{"a", "b", "c", "d"},
		[]string{},
		[]string{"a", "b", "c", "d"},
	},
	{ // no elements in common
		[]string{"a", "b", "c"},
		[]string{"d"},
		[]string{"a", "b", "c", "d"},
	},
	{ // one element in common
		[]string{"c", "b", "a"},
		[]string{"b", "d"},
		[]string{"a", "c", "d"},
	},
}
