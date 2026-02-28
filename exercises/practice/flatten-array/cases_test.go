package flattenarray

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 4b137d6 `flatten-array` Add additional test cases (#1953)

var testCases = []struct {
	description string
	input       any
	expected    []any
}{
	{
		description: "empty",
		input:       []any{},
		expected:    []any{},
	},
	{
		description: "no nesting",
		input:       []any{0, 1, 2},
		expected:    []any{0, 1, 2},
	},
	{
		description: "flattens a nested array",
		input:       []any{[]any{[]any{}}},
		expected:    []any{},
	},
	{
		description: "flattens array with just integers present",
		input:       []any{1, []any{2, 3, 4, 5, 6, 7}, 8},
		expected:    []any{1, 2, 3, 4, 5, 6, 7, 8},
	},
	{
		description: "5 level nesting",
		input:       []any{0, 2, []any{[]any{2, 3}, 8, 100, 4, []any{[]any{[]any{50}}}}, -2},
		expected:    []any{0, 2, 2, 3, 8, 100, 4, 50, -2},
	},
	{
		description: "6 level nesting",
		input:       []any{1, []any{2, []any{[]any{3}}, []any{4, []any{[]any{5}}}, 6, 7}, 8},
		expected:    []any{1, 2, 3, 4, 5, 6, 7, 8},
	},
	{
		description: "null values are omitted from the final result",
		input:       []any{1, 2, any(nil)},
		expected:    []any{1, 2},
	},
	{
		description: "consecutive null values at the front of the list are omitted from the final result",
		input:       []any{any(nil), any(nil), 3},
		expected:    []any{3},
	},
	{
		description: "consecutive null values in the middle of the list are omitted from the final result",
		input:       []any{1, any(nil), any(nil), 4},
		expected:    []any{1, 4},
	},
	{
		description: "6 level nest list with null values",
		input:       []any{0, 2, []any{[]any{2, 3}, 8, []any{[]any{100}}, any(nil), []any{[]any{any(nil)}}}, -2},
		expected:    []any{0, 2, 2, 3, 8, 100, -2},
	},
	{
		description: "all values in nested list are null",
		input:       []any{any(nil), []any{[]any{[]any{any(nil)}}}, any(nil), any(nil), []any{[]any{any(nil), any(nil)}, any(nil)}, any(nil)},
		expected:    []any{},
	},
}
