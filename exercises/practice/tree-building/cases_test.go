package treebuilding

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 884384d [tree-building] Add canonical data where there was none previously (#2643)

type testCase struct {
	description string
	input       []Record
	wantErr     bool
	expected    *Node
}

var testCases = []testCase{
	{
		description: "empty list",
		input:       []Record{},
		expected:    nil,
	},
	{
		description: "single record",
		input: []Record{
			Record{ID: 0, Parent: 0},
		},
		expected: &Node{ID: 0},
	},
	{
		description: "three records in order",
		input: []Record{
			Record{ID: 0, Parent: 0},
			Record{ID: 1, Parent: 0},
			Record{ID: 2, Parent: 0},
		},
		expected: &Node{
			ID: 0,
			Children: []*Node{
				&Node{ID: 1},
				&Node{ID: 2},
			},
		},
	},
	{
		description: "three records in reverse order",
		input: []Record{
			Record{ID: 2, Parent: 0},
			Record{ID: 1, Parent: 0},
			Record{ID: 0, Parent: 0},
		},
		expected: &Node{
			ID: 0,
			Children: []*Node{
				&Node{ID: 1},
				&Node{ID: 2},
			},
		},
	},
	{
		description: "more than two children",
		input: []Record{
			Record{ID: 0, Parent: 0},
			Record{ID: 1, Parent: 0},
			Record{ID: 2, Parent: 0},
			Record{ID: 3, Parent: 0},
		},
		expected: &Node{
			ID: 0,
			Children: []*Node{
				&Node{ID: 1},
				&Node{ID: 2},
				&Node{ID: 3},
			},
		},
	},
	{
		description: "binary tree",
		input: []Record{
			Record{ID: 5, Parent: 1},
			Record{ID: 3, Parent: 2},
			Record{ID: 2, Parent: 0},
			Record{ID: 4, Parent: 1},
			Record{ID: 1, Parent: 0},
			Record{ID: 0, Parent: 0},
			Record{ID: 6, Parent: 2},
		},
		expected: &Node{
			ID: 0,
			Children: []*Node{
				&Node{
					ID: 1,
					Children: []*Node{
						&Node{ID: 4},
						&Node{ID: 5},
					},
				},
				&Node{
					ID: 2,
					Children: []*Node{
						&Node{ID: 3},
						&Node{ID: 6},
					},
				},
			},
		},
	},
	{
		description: "unbalanced tree",
		input: []Record{
			Record{ID: 5, Parent: 2},
			Record{ID: 3, Parent: 2},
			Record{ID: 2, Parent: 0},
			Record{ID: 4, Parent: 1},
			Record{ID: 1, Parent: 0},
			Record{ID: 0, Parent: 0},
			Record{ID: 6, Parent: 2},
		},
		expected: &Node{
			ID: 0,
			Children: []*Node{
				&Node{
					ID: 1,
					Children: []*Node{
						&Node{ID: 4},
					},
				},
				&Node{
					ID: 2,
					Children: []*Node{
						&Node{ID: 3},
						&Node{ID: 5},
						&Node{ID: 6},
					},
				},
			},
		},
	},
	{
		description: "one root node and has parent",
		input: []Record{
			Record{ID: 0, Parent: 1},
		},
		wantErr: true,
	},
	{
		description: "root node has parent",
		input: []Record{
			Record{ID: 0, Parent: 1},
			Record{ID: 1, Parent: 0},
		},
		wantErr: true,
	},
	{
		description: "no root node",
		input: []Record{
			Record{ID: 1, Parent: 0},
			Record{ID: 2, Parent: 0},
		},
		wantErr: true,
	},
	{
		description: "duplicate node",
		input: []Record{
			Record{ID: 0, Parent: 0},
			Record{ID: 1, Parent: 0},
			Record{ID: 1, Parent: 0},
		},
		wantErr: true,
	},
	{
		description: "duplicate root",
		input: []Record{
			Record{ID: 0, Parent: 0},
			Record{ID: 0, Parent: 0},
		},
		wantErr: true,
	},
	{
		description: "non-continuous",
		input: []Record{
			Record{ID: 2, Parent: 0},
			Record{ID: 4, Parent: 2},
			Record{ID: 1, Parent: 0},
			Record{ID: 0, Parent: 0},
		},
		wantErr: true,
	},
	{
		description: "cycle directly",
		input: []Record{
			Record{ID: 5, Parent: 2},
			Record{ID: 3, Parent: 2},
			Record{ID: 2, Parent: 2},
			Record{ID: 4, Parent: 1},
			Record{ID: 1, Parent: 0},
			Record{ID: 0, Parent: 0},
			Record{ID: 6, Parent: 3},
		},
		wantErr: true,
	},
	{
		description: "cycle indirectly",
		input: []Record{
			Record{ID: 5, Parent: 2},
			Record{ID: 3, Parent: 2},
			Record{ID: 2, Parent: 6},
			Record{ID: 4, Parent: 1},
			Record{ID: 1, Parent: 0},
			Record{ID: 0, Parent: 0},
			Record{ID: 6, Parent: 3},
		},
		wantErr: true,
	},
	{
		description: "higher id parent of lower id",
		input: []Record{
			Record{ID: 0, Parent: 0},
			Record{ID: 2, Parent: 0},
			Record{ID: 1, Parent: 2},
		},
		wantErr: true,
	},
}
