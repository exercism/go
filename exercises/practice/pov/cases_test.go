package pov

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: b820099 Allow prettier to format more files (#1966)

type povTestCase struct {
	description string
	from        string
	tree        *Tree
	expected    *Tree
}

var povTestCases = []povTestCase{
	{
		description: "Results in the same tree if the input tree is a singleton",
		from:        "x",
		tree:        New("x"),
		expected:    New("x"),
	},
	{
		description: "Can reroot a tree with a parent and one sibling",
		from:        "x",
		tree:        New("parent", New("x"), New("sibling")),
		expected:    New("x", New("parent", New("sibling"))),
	},
	{
		description: "Can reroot a tree with a parent and many siblings",
		from:        "x",
		tree:        New("parent", New("a"), New("x"), New("b"), New("c")),
		expected:    New("x", New("parent", New("a"), New("b"), New("c"))),
	},
	{
		description: "Can reroot a tree with new root deeply nested in tree",
		from:        "x",
		tree:        New("level-0", New("level-1", New("level-2", New("level-3", New("x"))))),
		expected:    New("x", New("level-3", New("level-2", New("level-1", New("level-0"))))),
	},
	{
		description: "Moves children of the new root to same level as former parent",
		from:        "x",
		tree:        New("parent", New("x", New("kid-0"), New("kid-1"))),
		expected:    New("x", New("kid-0"), New("kid-1"), New("parent")),
	},
	{
		description: "Can reroot a complex tree with cousins",
		from:        "x",
		tree:        New("grandparent", New("parent", New("x", New("kid-0"), New("kid-1")), New("sibling-0"), New("sibling-1")), New("uncle", New("cousin-0"), New("cousin-1"))),
		expected:    New("x", New("kid-1"), New("kid-0"), New("parent", New("sibling-0"), New("sibling-1"), New("grandparent", New("uncle", New("cousin-0"), New("cousin-1"))))),
	},
	{
		description: "Errors if target does not exist in a singleton tree",
		from:        "nonexistent",
		tree:        New("x"),
		expected:    nil,
	},
	{
		description: "Errors if target does not exist in a large tree",
		from:        "nonexistent",
		tree:        New("parent", New("x", New("kid-0"), New("kid-1")), New("sibling-0"), New("sibling-1")),
		expected:    nil,
	},
}

type pathTestCase struct {
	description string
	from        string
	to          string
	tree        *Tree
	expected    []string
}

var pathTestCases = []pathTestCase{
	{
		description: "Can find path to parent",
		from:        "x",
		to:          "parent",
		tree:        New("parent", New("x"), New("sibling")),
		expected:    []string{"x", "parent"},
	},
	{
		description: "Can find path to sibling",
		from:        "x",
		to:          "b",
		tree:        New("parent", New("a"), New("x"), New("b"), New("c")),
		expected:    []string{"x", "parent", "b"},
	},
	{
		description: "Can find path to cousin",
		from:        "x",
		to:          "cousin-1",
		tree:        New("grandparent", New("parent", New("x", New("kid-0"), New("kid-1")), New("sibling-0"), New("sibling-1")), New("uncle", New("cousin-0"), New("cousin-1"))),
		expected:    []string{"x", "parent", "grandparent", "uncle", "cousin-1"},
	},
	{
		description: "Can find path not involving root",
		from:        "x",
		to:          "sibling-1",
		tree:        New("grandparent", New("parent", New("x"), New("sibling-0"), New("sibling-1"))),
		expected:    []string{"x", "parent", "sibling-1"},
	},
	{
		description: "Can find path from nodes other than x",
		from:        "a",
		to:          "c",
		tree:        New("parent", New("a"), New("x"), New("b"), New("c")),
		expected:    []string{"a", "parent", "c"},
	},
	{
		description: "Errors if destination does not exist",
		from:        "x",
		to:          "nonexistent",
		tree:        New("parent", New("x", New("kid-0"), New("kid-1")), New("sibling-0"), New("sibling-1")),
		expected:    []string(nil),
	},
	{
		description: "Errors if source does not exist",
		from:        "nonexistent",
		to:          "x",
		tree:        New("parent", New("x", New("kid-0"), New("kid-1")), New("sibling-0"), New("sibling-1")),
		expected:    []string(nil),
	},
}
