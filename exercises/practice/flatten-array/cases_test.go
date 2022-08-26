package flatten

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 4b137d6 `flatten-array` Add additional test cases (#1953)

var testCases = []struct {
	description string
	input       interface{}
	expected    []interface{}
}{
	{
		description: "empty",
		input:       []interface{}{},
		expected:    []interface{}{},
	},
	{
		description: "no nesting",
		input:       []interface{}{0, 1, 2},
		expected:    []interface{}{0, 1, 2},
	},
	{
		description: "flattens a nested array",
		input:       []interface{}{[]interface{}{[]interface{}{}}},
		expected:    []interface{}{},
	},
	{
		description: "flattens array with just integers present",
		input:       []interface{}{1, []interface{}{2, 3, 4, 5, 6, 7}, 8},
		expected:    []interface{}{1, 2, 3, 4, 5, 6, 7, 8},
	},
	{
		description: "5 level nesting",
		input:       []interface{}{0, 2, []interface{}{[]interface{}{2, 3}, 8, 100, 4, []interface{}{[]interface{}{[]interface{}{50}}}}, -2},
		expected:    []interface{}{0, 2, 2, 3, 8, 100, 4, 50, -2},
	},
	{
		description: "6 level nesting",
		input:       []interface{}{1, []interface{}{2, []interface{}{[]interface{}{3}}, []interface{}{4, []interface{}{[]interface{}{5}}}, 6, 7}, 8},
		expected:    []interface{}{1, 2, 3, 4, 5, 6, 7, 8},
	},
	{
		description: "null values are omitted from the final result",
		input:       []interface{}{1, 2, interface{}(nil)},
		expected:    []interface{}{1, 2},
	},
	{
		description: "consecutive null values at the front of the list are omitted from the final result",
		input:       []interface{}{interface{}(nil), interface{}(nil), 3},
		expected:    []interface{}{3},
	},
	{
		description: "consecutive null values in the middle of the list are omitted from the final result",
		input:       []interface{}{1, interface{}(nil), interface{}(nil), 4},
		expected:    []interface{}{1, 4},
	},
	{
		description: "6 level nest list with null values",
		input:       []interface{}{0, 2, []interface{}{[]interface{}{2, 3}, 8, []interface{}{[]interface{}{100}}, interface{}(nil), []interface{}{[]interface{}{interface{}(nil)}}}, -2},
		expected:    []interface{}{0, 2, 2, 3, 8, 100, -2},
	},
	{
		description: "all values in nested list are null",
		input:       []interface{}{interface{}(nil), []interface{}{[]interface{}{[]interface{}{interface{}(nil)}}}, interface{}(nil), interface{}(nil), []interface{}{[]interface{}{interface{}(nil), interface{}(nil)}, interface{}(nil)}, interface{}(nil)},
		expected:    []interface{}{},
	},
}
