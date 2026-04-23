package dotdsl

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: d6451b0 [dot-dsl] Make test case descriptions unique

type testCase struct {
	description string
	input       string
	nodes       map[string]Properties
	edges       map[string]Properties
	attrs       Properties
	expectedErr string
}

var testCases = []testCase{
	{
		description: "empty graph",
		input:       "graph {\n}",
		nodes:       nil,
		edges:       nil,
		attrs:       nil,
	},
	{
		description: "graph with one node",
		input:       "graph {\n    a;\n}",
		nodes:       map[string]Properties{"a": nil},
		edges:       nil,
		attrs:       nil,
	},
	{
		description: "graph with one node with attribute",
		input:       "graph {\n    a [color=green];\n}",
		nodes:       map[string]Properties{"a": {"color": "green"}},
		edges:       nil,
		attrs:       nil,
	},
	{
		description: "graph with one edge",
		input:       "graph {\n    a -- b;\n}",
		nodes:       map[string]Properties{"a": nil, "b": nil},
		edges:       map[string]Properties{"{a b}": nil},
		attrs:       nil,
	},
	{
		description: "graph with one attribute",
		input:       "graph {\n    [foo=1];\n}",
		nodes:       nil,
		edges:       nil,
		attrs:       Properties{"foo": 1},
	},
	{
		description: "graph with comments",
		input:       "graph {\n    // a C-like comment\n    # a shell-like comment\n    [foo=1];\n}",
		nodes:       nil,
		edges:       nil,
		attrs:       Properties{"foo": 1},
	},
	{
		description: "graph with nodes, edges, and attributes",
		input:       "graph {\n    [foo=1];\n    [title=\"Testing Attrs\"];\n    a [color=green];\n    b [label=\"Beta!\"];\n    b -- c;\n    a -- b [color=blue];\n    [bar=true];\n}",
		nodes:       map[string]Properties{"a": {"color": "green"}, "b": {"label": "Beta!"}, "c": nil},
		edges:       map[string]Properties{"{a b}": {"color": "blue"}, "{b c}": nil},
		attrs:       Properties{"bar": true, "foo": 1, "title": "Testing Attrs"},
	},
	{
		description: "multiple edges on one line",
		input:       "graph {\n    a -- b -- c -- d [style=dotted];\n}",
		nodes:       map[string]Properties{"a": nil, "b": nil, "c": nil, "d": nil},
		edges:       map[string]Properties{"{a b}": {"style": "dotted"}, "{b c}": {"style": "dotted"}, "{c d}": {"style": "dotted"}},
		attrs:       nil,
	},
	{
		description: "only 1 edge between nodes",
		input:       "graph {\n    a -- b;\n    a -- b;\n    b -- a [color=blue];\n}",
		nodes:       map[string]Properties{"a": nil, "b": nil},
		edges:       map[string]Properties{"{a b}": {"color": "blue"}},
		attrs:       nil,
	},
	{
		description: "malformed input",
		input:       "graphical {\n}",
		expectedErr: "invalid graph",
	},
	{
		description: "malformed edge",
		input:       "graph {\n    a --;\n}",
		expectedErr: "invalid edge",
	},
	{
		description: "malformed edge 2",
		input:       "graph {\n    a b;\n}",
		expectedErr: "invalid edge",
	},
	{
		description: "invalid edge type",
		input:       "graph {\n    a == b;\n} ",
		expectedErr: "invalid edge",
	},
	{
		description: "multiple edges missing a node",
		input:       "graph {\n    a -- b --;\n}",
		expectedErr: "invalid edge",
	},
	{
		description: "multiple edges missing a connector",
		input:       "graph {\n    a -- b c;\n}",
		expectedErr: "invalid edge",
	},
	{
		description: "empty attribute",
		input:       "graph {\n    a [];\n}",
		expectedErr: "invalid attribute",
	},
	{
		description: "malformed attribute",
		input:       "graph {\n    a [name];\n}",
		expectedErr: "invalid attribute",
	},
	{
		description: "empty attribute name",
		input:       "graph {\n    a [=value];\n}",
		expectedErr: "invalid attribute",
	},
	{
		description: "non-alphanumeric node name",
		input:       "graph {\n    ? [key=value];\n}",
		expectedErr: "node name must be alphanumeric",
	},
}
