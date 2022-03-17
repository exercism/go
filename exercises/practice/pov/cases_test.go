package pov

// source: problem-specification repo - POV exercise

type TestTreeData struct {
	description string
	root        string
	children    []*Tree
}

var testTrees = map[string]TestTreeData{
	"singleton": {
		root:     "x",
		children: nil,
	},
	"parent and one sibling": {
		root:     "parent",
		children: []*Tree{New("x"), New("sibling")},
	},
	"parent and many siblings": {
		root:     "parent",
		children: []*Tree{New("a"), New("x"), New("b"), New("c")},
	},
	"parent and kids": {
		root:     "parent",
		children: []*Tree{New("x", New("kid-0"), New("kid-1"))},
	},
	"parent, kids and siblings": {
		root:     "parent",
		children: []*Tree{New("x", New("kid-0"), New("kid-1")), New("sibling-0"), New("sibling-1")},
	},
	"grandparent, parent and siblings": {
		root:     "grandparent",
		children: []*Tree{New("parent", New("x"), New("sibling-0"), New("sibling-1"))},
	},
	"linear tree": {
		root:     "level-0",
		children: []*Tree{New("level-1", New("level-2", New("level-3", New("x"))))},
	},
	"complex tree with cousins": {
		root: "grandparent",
		children: []*Tree{New("parent", New("x", New("kid-0"), New("kid-1")), New("sibling-0"),
			New("sibling-1")), New("uncle", New("cousin-0"), New("cousin-1"))},
	},
}

var newValueChildrenTestTrees = []string{"singleton", "parent and one sibling", "parent and kids"}

func mkTestTree(treeName string) *Tree {
	treeData := testTrees[treeName]
	return New(treeData.root, treeData.children...)
}

type fromPovTestCase struct {
	description string // same as in canonical-data.json
	treeName    string
	from        string
	expected    *Tree
}

var fromPovTestCases = []fromPovTestCase{
	{
		description: "Results in the same tree if the input tree is a singleton",
		treeName:    "singleton",
		from:        "x",
		expected:    New("x"),
	},
	{
		description: "Can reroot a tree with a parent and one sibling",
		treeName:    "parent and one sibling",
		from:        "x",
		expected:    New("x", New("parent", New("sibling"))),
	},
	{
		description: "Can reroot a tree with a parent and many siblings",
		treeName:    "parent and many siblings",
		from:        "x",
		expected:    New("x", New("parent", New("a"), New("b"), New("c"))),
	},
	{
		treeName:    "linear tree",
		description: "Can reroot a tree with new root deeply nested in tree",
		from:        "x",
		expected:    New("x", New("level-3", New("level-2", New("level-1", New("level-0"))))),
	},
	{
		description: "Moves children of the new root to same level as former parent",
		treeName:    "parent and kids",
		from:        "x",
		expected:    New("x", New("kid-0"), New("kid-1"), New("parent")),
	},
	{
		description: "Can reroot a complex tree with cousins",
		treeName:    "complex tree with cousins",
		from:        "x",
		expected: New("x", New("kid-0"), New("kid-1"),
			New("parent", New("sibling-0"), New("sibling-1"),
				New("grandparent", New("uncle", New("cousin-0"), New("cousin-1"))))),
	},
	{
		description: "Errors if target does not exist in a singleton tree",
		treeName:    "singleton",
		from:        "nonexistent",
		expected:    nil,
	},
	{
		description: "Errors if target does not exist in a large tree",
		treeName:    "parent, kids and siblings",
		from:        "nonexistent",
		expected:    nil,
	},
}

type pathToTestCase struct {
	description string
	treeName    string
	from        string
	to          string
	expected    []string
}

var pathToTestCases = []pathToTestCase{
	{
		description: "Can find path to parent",
		treeName:    "parent and one sibling",
		from:        "x",
		to:          "parent",
		expected:    []string{"x", "parent"},
	},
	{
		description: "Can find path to sibling",
		treeName:    "parent and many siblings",
		from:        "x",
		to:          "b",
		expected:    []string{"x", "parent", "b"},
	},
	{
		description: "Can find path to cousin",
		treeName:    "complex tree with cousins",
		from:        "x",
		to:          "cousin-1",
		expected:    []string{"x", "parent", "grandparent", "uncle", "cousin-1"},
	},
	{
		description: "Can find path not involving root",
		treeName:    "grandparent, parent and siblings",
		from:        "x",
		to:          "sibling-1",
		expected:    []string{"x", "parent", "sibling-1"},
	},
	{
		description: "Can find path from nodes other than x",
		treeName:    "parent and many siblings",
		from:        "a",
		to:          "c",
		expected:    []string{"a", "parent", "c"},
	},
	{
		description: "Errors if destination does not exist",
		treeName:    "parent, kids and siblings",
		from:        "x",
		to:          "nonexistent",
		expected:    nil,
	},
	{
		description: "Errors if source does not exist",
		treeName:    "parent, kids and siblings",
		from:        "nonexistent",
		to:          "x",
		expected:    nil,
	},
}
