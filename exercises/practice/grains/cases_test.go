package grains

// Source: exercism/problem-specifications
// Commit: 2ec42ab Grains: Fixed canonical data to have standard error indicator (#1322)
// Problem Specifications Version: 1.2.0

// returns the number of grains on the square
var squareTests = []struct {
	description string
	input       int
	expectedVal uint64
	expectError bool
}{
	{
		description: "1",
		input:       1,
		expectedVal: 1,
	},
	{
		description: "2",
		input:       2,
		expectedVal: 2,
	},
	{
		description: "3",
		input:       3,
		expectedVal: 4,
	},
	{
		description: "4",
		input:       4,
		expectedVal: 8,
	},
	{
		description: "16",
		input:       16,
		expectedVal: 32768,
	},
	{
		description: "32",
		input:       32,
		expectedVal: 2147483648,
	},
	{
		description: "64",
		input:       64,
		expectedVal: 9223372036854775808,
	},
	{
		description: "square 0 returns an error",
		input:       0,
		expectError: true,
	},
	{
		description: "negative square returns an error",
		input:       -1,
		expectError: true,
	},
	{
		description: "square greater than 64 returns an error",
		input:       65,
		expectError: true,
	},
}
