package flatten

// Source: exercism/problem-specifications
// Commit: 0290376 flatten-array: Apply new "input" policy
// Problem Specifications Version: 1.2.0

var testCases = []struct {
	description string
	input       interface{}
	expected    []interface{}
}{
	{
		description: "no nesting",
		input:       []interface{}{0, 1, 2},
		expected:    []interface{}{0, 1, 2},
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
