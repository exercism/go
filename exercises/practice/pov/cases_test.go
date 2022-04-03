package pov

// source: problem-specification repo - POV exercise

type fromPovTestCase struct {
	description string // same as in canonical-data.json
	tree        *Tree
	from        string
	expected    *Tree
}

var fromPovTestCases = []fromPovTestCase{
	{
		description: "Results in the same tree if the input tree is a singleton",
		tree:        New("x"),
		from:        "x",
		expected:    New("x"),
	},
	{
		description: "Can reroot a tree with a parent and one sibling",
		tree:        New("parent", New("x"), New("sibling")),
		from:        "x",
		expected:    New("x", New("parent", New("sibling"))),
	},
	{
		description: "Can reroot a tree with a parent and many siblings",
		tree:        New("parent", New("a"), New("x"), New("b"), New("c")),
		from:        "x",
		expected:    New("x", New("parent", New("a"), New("b"), New("c"))),
	},
	{
		description: "Can reroot a tree with new root deeply nested in tree",
		tree:        New("level-0", New("level-1", New("level-2", New("level-3", New("x"))))),
		from:        "x",
		expected:    New("x", New("level-3", New("level-2", New("level-1", New("level-0"))))),
	},
	{
		description: "Moves children of the new root to same level as former parent",
		tree:        New("parent", New("x", New("kid-0"), New("kid-1"))),
		from:        "x",
		expected:    New("x", New("kid-0"), New("kid-1"), New("parent")),
	},
	{
		description: "Can reroot a complex tree with cousins",
		tree: New("grandparent", New("parent",
			New("x", New("kid-0"), New("kid-1")), New("sibling-0"),
			New("sibling-1")), New("uncle", New("cousin-0"), New("cousin-1"))),
		from: "x",
		expected: New("x", New("kid-0"), New("kid-1"),
			New("parent", New("sibling-0"), New("sibling-1"),
				New("grandparent", New("uncle", New("cousin-0"), New("cousin-1"))))),
	},
	{
		description: "Errors if target does not exist in a singleton tree",
		tree:        New("x"),
		from:        "nonexistent",
		expected:    nil,
	},
	{
		description: "Errors if target does not exist in a large tree",
		tree: New("parent",
			New("x", New("kid-0"), New("kid-1")), New("sibling-0"), New("sibling-1")),
		from:     "nonexistent",
		expected: nil,
	},
}

type pathToTestCase struct {
	description string
	tree        *Tree
	from        string
	to          string
	expected    []string
}

var pathToTestCases = []pathToTestCase{
	{
		description: "Can find path to parent",
		tree:        New("parent", New("x"), New("sibling")),
		from:        "x",
		to:          "parent",
		expected:    []string{"x", "parent"},
	},
	{
		description: "Can find path to sibling",
		tree:        New("parent", New("a"), New("x"), New("b"), New("c")),
		from:        "x",
		to:          "b",
		expected:    []string{"x", "parent", "b"},
	},
	{
		description: "Can find path to cousin",
		tree: New("grandparent", New("parent",
			New("x", New("kid-0"), New("kid-1")), New("sibling-0"),
			New("sibling-1")), New("uncle", New("cousin-0"), New("cousin-1"))),
		from:     "x",
		to:       "cousin-1",
		expected: []string{"x", "parent", "grandparent", "uncle", "cousin-1"},
	},
	{
		description: "Can find path not involving root",
		tree:        New("grandparent", New("parent", New("x"), New("sibling-0"), New("sibling-1"))),
		from:        "x",
		to:          "sibling-1",
		expected:    []string{"x", "parent", "sibling-1"},
	},
	{
		description: "Can find path from nodes other than x",
		tree:        New("parent", New("a"), New("x"), New("b"), New("c")),
		from:        "a",
		to:          "c",
		expected:    []string{"a", "parent", "c"},
	},
	{
		description: "Errors if destination does not exist",
		tree:        New("parent", New("x", New("kid-0"), New("kid-1")), New("sibling-0"), New("sibling-1")),
		from:        "x",
		to:          "nonexistent",
		expected:    nil,
	},
	{
		description: "Errors if source does not exist",
		tree:        New("parent", New("x", New("kid-0"), New("kid-1")), New("sibling-0"), New("sibling-1")),
		from:        "nonexistent",
		to:          "x",
		expected:    nil,
	},
}
