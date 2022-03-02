package pov

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

func mkTestTree(treeName string) *Tree {
	treeData := testTrees[treeName]
	return New(treeData.root, treeData.children...)
}

type fromPovTestCase struct {
	treeName string
	from     string
	expected *Tree
}

var fromPovTestCases = []fromPovTestCase{
	{
		treeName: "singleton",
		from:     "x",
		expected: New("x"),
	},
	{
		treeName: "parent and one sibling",
		from:     "x",
		expected: New("x", New("parent", New("sibling"))),
	},
	{
		treeName: "parent and many siblings",
		from:     "x",
		expected: New("x", New("parent", New("a"), New("b"), New("c"))),
	},
	{
		treeName: "linear tree",
		from:     "x",
		expected: New("x", New("level-3", New("level-2", New("level-1", New("level-0"))))),
	},
	{
		treeName: "parent and kids",
		from:     "x",
		expected: New("x", New("kid-0"), New("kid-1"), New("parent")),
	},
	{
		treeName: "complex tree with cousins",
		from:     "x",
		expected: New("x", New("kid-0"), New("kid-1"),
			New("parent", New("sibling-0"), New("sibling-1"),
				New("grandparent", New("uncle", New("cousin-0"), New("cousin-1"))))),
	},
	// target does not exist in tree
	{
		treeName: "singleton",
		from:     "nonexistent",
		expected: nil,
	},
	{
		treeName: "parent, kids and siblings",
		from:     "nonexistent",
		expected: nil,
	},
}

type pathToTestCase struct {
	treeName string
	from     string
	to       string
	expected []string
}

var pathToTestCases = []pathToTestCase{
	{
		treeName: "parent and one sibling",
		from:     "x",
		to:       "parent",
		expected: []string{"x", "parent"},
	}, {
		treeName: "parent and many siblings",
		from:     "x",
		to:       "b",
		expected: []string{"x", "parent", "b"},
	},
	{
		treeName: "complex tree with cousins",
		from:     "x",
		to:       "cousin-1",
		expected: []string{"x", "parent", "grandparent", "uncle", "cousin-1"},
	}, {
		treeName: "grandparent, parent and siblings",
		from:     "x",
		to:       "sibling-1",
		expected: []string{"x", "parent", "sibling-1"},
	},
	{
		treeName: "parent and many siblings",
		from:     "a",
		to:       "c",
		expected: []string{"a", "parent", "c"},
	},
	// target does not exist in tree
	{
		treeName: "parent, kids and siblings",
		from:     "x",
		to:       "nonexistent",
		expected: nil,
	},
	// source does not exist in tree
	{
		treeName: "parent, kids and siblings",
		from:     "nonexistent",
		to:       "x",
		expected: nil,
	},
}
