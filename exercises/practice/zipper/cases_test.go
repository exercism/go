package zipper

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 931f320 Add left, right, up, up case to zipper canon (#1875)

type operation struct {
	funcCall string
	intArg   int
	nodeArg  *Node
}

type want struct {
	wantType string
	intVal   int
	nodeVal  *Node
}

type testCase struct {
	description string
	input       *Node
	operations  []operation
	want        want
}

var testCases = []testCase{
	{
		description: "data is retained",
		input:       &Node{value: 1, left: &Node{value: 2, left: nil, right: &Node{value: 3, left: nil, right: nil}}, right: &Node{value: 4, left: nil, right: nil}},
		operations: []operation{
			operation{funcCall: "to_tree"},
		},
		want: want{
			wantType: "tree",
			nodeVal:  &Node{value: 1, left: &Node{value: 2, left: nil, right: &Node{value: 3, left: nil, right: nil}}, right: &Node{value: 4, left: nil, right: nil}},
		},
	},
	{
		description: "left, right and value",
		input:       &Node{value: 1, left: &Node{value: 2, left: nil, right: &Node{value: 3, left: nil, right: nil}}, right: &Node{value: 4, left: nil, right: nil}},
		operations: []operation{
			operation{funcCall: "left"},
			operation{funcCall: "right"},
			operation{funcCall: "value"},
		},
		want: want{
			wantType: "int",
			intVal:   3,
		},
	},
	{
		description: "dead end",
		input:       &Node{value: 1, left: &Node{value: 2, left: nil, right: &Node{value: 3, left: nil, right: nil}}, right: &Node{value: 4, left: nil, right: nil}},
		operations: []operation{
			operation{funcCall: "left"},
			operation{funcCall: "left"},
		},
		want: want{
			wantType: "zipper",
		},
	},
	{
		description: "tree from deep focus",
		input:       &Node{value: 1, left: &Node{value: 2, left: nil, right: &Node{value: 3, left: nil, right: nil}}, right: &Node{value: 4, left: nil, right: nil}},
		operations: []operation{
			operation{funcCall: "left"},
			operation{funcCall: "right"},
			operation{funcCall: "to_tree"},
		},
		want: want{
			wantType: "tree",
			nodeVal:  &Node{value: 1, left: &Node{value: 2, left: nil, right: &Node{value: 3, left: nil, right: nil}}, right: &Node{value: 4, left: nil, right: nil}},
		},
	},
	{
		description: "traversing up from top",
		input:       &Node{value: 1, left: &Node{value: 2, left: nil, right: &Node{value: 3, left: nil, right: nil}}, right: &Node{value: 4, left: nil, right: nil}},
		operations: []operation{
			operation{funcCall: "up"},
		},
		want: want{
			wantType: "zipper",
		},
	},
	{
		description: "left, right, and up",
		input:       &Node{value: 1, left: &Node{value: 2, left: nil, right: &Node{value: 3, left: nil, right: nil}}, right: &Node{value: 4, left: nil, right: nil}},
		operations: []operation{
			operation{funcCall: "left"},
			operation{funcCall: "up"},
			operation{funcCall: "right"},
			operation{funcCall: "up"},
			operation{funcCall: "left"},
			operation{funcCall: "right"},
			operation{funcCall: "value"},
		},
		want: want{
			wantType: "int",
			intVal:   3,
		},
	},
	{
		description: "test ability to descend multiple levels and return",
		input:       &Node{value: 1, left: &Node{value: 2, left: nil, right: &Node{value: 3, left: nil, right: nil}}, right: &Node{value: 4, left: nil, right: nil}},
		operations: []operation{
			operation{funcCall: "left"},
			operation{funcCall: "right"},
			operation{funcCall: "up"},
			operation{funcCall: "up"},
			operation{funcCall: "value"},
		},
		want: want{
			wantType: "int",
			intVal:   1,
		},
	},
	{
		description: "set_value",
		input:       &Node{value: 1, left: &Node{value: 2, left: nil, right: &Node{value: 3, left: nil, right: nil}}, right: &Node{value: 4, left: nil, right: nil}},
		operations: []operation{
			operation{funcCall: "left"},
			operation{funcCall: "set_value", intArg: 5},
			operation{funcCall: "to_tree"},
		},
		want: want{
			wantType: "tree",
			nodeVal:  &Node{value: 1, left: &Node{value: 5, left: nil, right: &Node{value: 3, left: nil, right: nil}}, right: &Node{value: 4, left: nil, right: nil}},
		},
	},
	{
		description: "set_value after traversing up",
		input:       &Node{value: 1, left: &Node{value: 2, left: nil, right: &Node{value: 3, left: nil, right: nil}}, right: &Node{value: 4, left: nil, right: nil}},
		operations: []operation{
			operation{funcCall: "left"},
			operation{funcCall: "right"},
			operation{funcCall: "up"},
			operation{funcCall: "set_value", intArg: 5},
			operation{funcCall: "to_tree"},
		},
		want: want{
			wantType: "tree",
			nodeVal:  &Node{value: 1, left: &Node{value: 5, left: nil, right: &Node{value: 3, left: nil, right: nil}}, right: &Node{value: 4, left: nil, right: nil}},
		},
	},
	{
		description: "set_left with leaf",
		input:       &Node{value: 1, left: &Node{value: 2, left: nil, right: &Node{value: 3, left: nil, right: nil}}, right: &Node{value: 4, left: nil, right: nil}},
		operations: []operation{
			operation{funcCall: "left"},
			operation{funcCall: "set_left", nodeArg: &Node{value: 5, left: nil, right: nil}},
			operation{funcCall: "to_tree"},
		},
		want: want{
			wantType: "tree",
			nodeVal:  &Node{value: 1, left: &Node{value: 2, left: &Node{value: 5, left: nil, right: nil}, right: &Node{value: 3, left: nil, right: nil}}, right: &Node{value: 4, left: nil, right: nil}},
		},
	},
	{
		description: "set_right with null",
		input:       &Node{value: 1, left: &Node{value: 2, left: nil, right: &Node{value: 3, left: nil, right: nil}}, right: &Node{value: 4, left: nil, right: nil}},
		operations: []operation{
			operation{funcCall: "left"},
			operation{funcCall: "set_right", nodeArg: nil},
			operation{funcCall: "to_tree"},
		},
		want: want{
			wantType: "tree",
			nodeVal:  &Node{value: 1, left: &Node{value: 2, left: nil, right: nil}, right: &Node{value: 4, left: nil, right: nil}},
		},
	},
	{
		description: "set_right with subtree",
		input:       &Node{value: 1, left: &Node{value: 2, left: nil, right: &Node{value: 3, left: nil, right: nil}}, right: &Node{value: 4, left: nil, right: nil}},
		operations: []operation{
			operation{funcCall: "set_right", nodeArg: &Node{value: 6, left: &Node{value: 7, left: nil, right: nil}, right: &Node{value: 8, left: nil, right: nil}}},
			operation{funcCall: "to_tree"},
		},
		want: want{
			wantType: "tree",
			nodeVal:  &Node{value: 1, left: &Node{value: 2, left: nil, right: &Node{value: 3, left: nil, right: nil}}, right: &Node{value: 6, left: &Node{value: 7, left: nil, right: nil}, right: &Node{value: 8, left: nil, right: nil}}},
		},
	},
	{
		description: "set_value on deep focus",
		input:       &Node{value: 1, left: &Node{value: 2, left: nil, right: &Node{value: 3, left: nil, right: nil}}, right: &Node{value: 4, left: nil, right: nil}},
		operations: []operation{
			operation{funcCall: "left"},
			operation{funcCall: "right"},
			operation{funcCall: "set_value", intArg: 5},
			operation{funcCall: "to_tree"},
		},
		want: want{
			wantType: "tree",
			nodeVal:  &Node{value: 1, left: &Node{value: 2, left: nil, right: &Node{value: 5, left: nil, right: nil}}, right: &Node{value: 4, left: nil, right: nil}},
		},
	},
}
