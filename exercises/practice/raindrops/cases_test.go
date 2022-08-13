package raindrops

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 42dd0ce Remove version (#1678)

var testCases = []struct {
	description string
	input       int
	expected    string
}{
	{
		description: "the sound for 1 is 1",
		input:       1,
		expected:    "1",
	},
	{
		description: "the sound for 3 is Pling",
		input:       3,
		expected:    "Pling",
	},
	{
		description: "the sound for 5 is Plang",
		input:       5,
		expected:    "Plang",
	},
	{
		description: "the sound for 7 is Plong",
		input:       7,
		expected:    "Plong",
	},
	{
		description: "the sound for 6 is Pling as it has a factor 3",
		input:       6,
		expected:    "Pling",
	},
	{
		description: "2 to the power 3 does not make a raindrop sound as 3 is the exponent not the base",
		input:       8,
		expected:    "8",
	},
	{
		description: "the sound for 9 is Pling as it has a factor 3",
		input:       9,
		expected:    "Pling",
	},
	{
		description: "the sound for 10 is Plang as it has a factor 5",
		input:       10,
		expected:    "Plang",
	},
	{
		description: "the sound for 14 is Plong as it has a factor of 7",
		input:       14,
		expected:    "Plong",
	},
	{
		description: "the sound for 15 is PlingPlang as it has factors 3 and 5",
		input:       15,
		expected:    "PlingPlang",
	},
	{
		description: "the sound for 21 is PlingPlong as it has factors 3 and 7",
		input:       21,
		expected:    "PlingPlong",
	},
	{
		description: "the sound for 25 is Plang as it has a factor 5",
		input:       25,
		expected:    "Plang",
	},
	{
		description: "the sound for 27 is Pling as it has a factor 3",
		input:       27,
		expected:    "Pling",
	},
	{
		description: "the sound for 35 is PlangPlong as it has factors 5 and 7",
		input:       35,
		expected:    "PlangPlong",
	},
	{
		description: "the sound for 49 is Plong as it has a factor 7",
		input:       49,
		expected:    "Plong",
	},
	{
		description: "the sound for 52 is 52",
		input:       52,
		expected:    "52",
	},
	{
		description: "the sound for 105 is PlingPlangPlong as it has factors 3, 5 and 7",
		input:       105,
		expected:    "PlingPlangPlong",
	},
	{
		description: "the sound for 3125 is Plang as it has a factor 5",
		input:       3125,
		expected:    "Plang",
	},
}
