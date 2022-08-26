package armstrong

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: c6884e5 armstrong-numbers: Add hyphens in descriptions (#1687)

var testCases = []struct {
	description string
	input       int
	expected    bool
}{
	{
		description: "Zero is an Armstrong number",
		input:       0,
		expected:    true,
	},
	{
		description: "Single-digit numbers are Armstrong numbers",
		input:       5,
		expected:    true,
	},
	{
		description: "There are no two-digit Armstrong numbers",
		input:       10,
		expected:    false,
	},
	{
		description: "Three-digit number that is an Armstrong number",
		input:       153,
		expected:    true,
	},
	{
		description: "Three-digit number that is not an Armstrong number",
		input:       100,
		expected:    false,
	},
	{
		description: "Four-digit number that is an Armstrong number",
		input:       9474,
		expected:    true,
	},
	{
		description: "Four-digit number that is not an Armstrong number",
		input:       9475,
		expected:    false,
	},
	{
		description: "Seven-digit number that is an Armstrong number",
		input:       9926315,
		expected:    true,
	},
	{
		description: "Seven-digit number that is not an Armstrong number",
		input:       9926314,
		expected:    false,
	},
}
