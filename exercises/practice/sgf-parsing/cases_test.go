package sgfparsing

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 3c33c24 sgf-parsing: Add individual tests of escaping/whitespace behaviour (#1889)

var testCases = []struct {
	description   string
	encoded       string
	expectedError string
	expected      *Node
}{
	{
		description:   "empty input",
		encoded:       "",
		expectedError: "tree missing",
		expected:      nil,
	},
	{
		description:   "tree with no nodes",
		encoded:       "()",
		expectedError: "tree with no nodes",
		expected:      nil,
	},
	{
		description:   "node without tree",
		encoded:       ";",
		expectedError: "tree missing",
		expected:      nil,
	},
	{
		description:   "node without properties",
		encoded:       "(;)",
		expectedError: "",
		expected: &Node{
			Properties: map[string][]string{},
			Children:   []*Node{},
		},
	},
	{
		description:   "single node tree",
		encoded:       "(;A[B])",
		expectedError: "",
		expected: &Node{
			Properties: map[string][]string{"A": {"B"}},
			Children:   []*Node{},
		},
	},
	{
		description:   "multiple properties",
		encoded:       "(;A[b]C[d])",
		expectedError: "",
		expected: &Node{
			Properties: map[string][]string{"A": {"b"}, "C": {"d"}},
			Children:   []*Node{},
		},
	},
	{
		description:   "properties without delimiter",
		encoded:       "(;A)",
		expectedError: "properties without delimiter",
		expected:      nil,
	},
	{
		description:   "all lowercase property",
		encoded:       "(;a[b])",
		expectedError: "property must be in uppercase",
		expected:      nil,
	},
	{
		description:   "upper and lowercase property",
		encoded:       "(;Aa[b])",
		expectedError: "property must be in uppercase",
		expected:      nil,
	},
	{
		description:   "two nodes",
		encoded:       "(;A[B];B[C])",
		expectedError: "",
		expected: &Node{
			Properties: map[string][]string{"A": {"B"}},
			Children: []*Node{
				&Node{
					Properties: map[string][]string{"B": {"C"}},
					Children:   []*Node{},
				},
			},
		},
	},
	{
		description:   "two child trees",
		encoded:       "(;A[B](;B[C])(;C[D]))",
		expectedError: "",
		expected: &Node{
			Properties: map[string][]string{"A": {"B"}},
			Children: []*Node{
				&Node{
					Properties: map[string][]string{"B": {"C"}},
					Children:   []*Node{},
				},
				&Node{
					Properties: map[string][]string{"C": {"D"}},
					Children:   []*Node{},
				},
			},
		},
	},
	{
		description:   "multiple property values",
		encoded:       "(;A[b][c][d])",
		expectedError: "",
		expected: &Node{
			Properties: map[string][]string{"A": {"b", "c", "d"}},
			Children:   []*Node{},
		},
	},
	{
		description:   "within property values, whitespace characters such as tab are converted to spaces",
		encoded:       "(;A[hello\t\tworld])",
		expectedError: "",
		expected: &Node{
			Properties: map[string][]string{"A": {"hello  world"}},
			Children:   []*Node{},
		},
	},
	{
		description:   "within property values, newlines remain as newlines",
		encoded:       "(;A[hello\n\nworld])",
		expectedError: "",
		expected: &Node{
			Properties: map[string][]string{"A": {"hello\n\nworld"}},
			Children:   []*Node{},
		},
	},
	{
		description:   "escaped closing bracket within property value becomes just a closing bracket",
		encoded:       "(;A[\\]])",
		expectedError: "",
		expected: &Node{
			Properties: map[string][]string{"A": {"]"}},
			Children:   []*Node{},
		},
	},
	{
		description:   "escaped backslash in property value becomes just a backslash",
		encoded:       "(;A[\\\\])",
		expectedError: "",
		expected: &Node{
			Properties: map[string][]string{"A": {"\\"}},
			Children:   []*Node{},
		},
	},
	{
		description:   "opening bracket within property value doesn't need to be escaped",
		encoded:       "(;A[x[y\\]z][foo]B[bar];C[baz])",
		expectedError: "",
		expected: &Node{
			Properties: map[string][]string{"A": {"x[y]z", "foo"}, "B": {"bar"}},
			Children: []*Node{
				&Node{
					Properties: map[string][]string{"C": {"baz"}},
					Children:   []*Node{},
				},
			},
		},
	},
	{
		description:   "semicolon in property value doesn't need to be escaped",
		encoded:       "(;A[a;b][foo]B[bar];C[baz])",
		expectedError: "",
		expected: &Node{
			Properties: map[string][]string{"A": {"a;b", "foo"}, "B": {"bar"}},
			Children: []*Node{
				&Node{
					Properties: map[string][]string{"C": {"baz"}},
					Children:   []*Node{},
				},
			},
		},
	},
	{
		description:   "parentheses in property value don't need to be escaped",
		encoded:       "(;A[x(y)z][foo]B[bar];C[baz])",
		expectedError: "",
		expected: &Node{
			Properties: map[string][]string{"A": {"x(y)z", "foo"}, "B": {"bar"}},
			Children: []*Node{
				&Node{
					Properties: map[string][]string{"C": {"baz"}},
					Children:   []*Node{},
				},
			},
		},
	},
	{
		description:   "escaped tab in property value is converted to space",
		encoded:       "(;A[hello\\\tworld])",
		expectedError: "",
		expected: &Node{
			Properties: map[string][]string{"A": {"hello world"}},
			Children:   []*Node{},
		},
	},
	{
		description:   "escaped newline in property value is converted to nothing at all",
		encoded:       "(;A[hello\\\nworld])",
		expectedError: "",
		expected: &Node{
			Properties: map[string][]string{"A": {"helloworld"}},
			Children:   []*Node{},
		},
	},
	{
		description:   "escaped t and n in property value are just letters, not whitespace",
		encoded:       "(;A[\\t = t and \\n = n])",
		expectedError: "",
		expected: &Node{
			Properties: map[string][]string{"A": {"t = t and n = n"}},
			Children:   []*Node{},
		},
	},
	{
		description:   "mixing various kinds of whitespace and escaped characters in property value",
		encoded:       "(;A[\\]b\nc\\\nd\t\te\\\\ \\\n\\]])",
		expectedError: "",
		expected: &Node{
			Properties: map[string][]string{"A": {"]b\ncd  e\\ ]"}},
			Children:   []*Node{},
		},
	},
}
