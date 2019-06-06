package darts

// Source: exercism/problem-specifications
// Commit: efcb2c8 Darts: Flatten sub-cases to single level - to simplify code generation (#1510)
// Problem Specifications Version: 2.0.0

var testCases = []struct {
	description string
	x           float64
	y           float64
	expected    int
}{
	{
		description: "A dart lands outside the target",
		x:           -9.0,
		y:           9.0,
		expected:    0,
	},
	{
		description: "A dart lands just in the border of the target",
		x:           0.0,
		y:           10.0,
		expected:    1,
	},
	{
		description: "A dart lands in the outer circle",
		x:           4.0,
		y:           4.0,
		expected:    1,
	},
	{
		description: "A dart lands right in the border between outer and middle circles",
		x:           5.0,
		y:           0.0,
		expected:    5,
	},
	{
		description: "A dart lands in the middle circle",
		x:           0.8,
		y:           -0.8,
		expected:    5,
	},
	{
		description: "A dart lands right in the border between middle and inner circles",
		x:           0.0,
		y:           -1.0,
		expected:    10,
	},
	{
		description: "A dart lands in the inner circle",
		x:           -0.1,
		y:           -0.1,
		expected:    10,
	},
	{
		description: "A dart whose coordinates sum to > 1 but whose radius to origin is <= 1 is scored in the inner circle",
		x:           0.4,
		y:           0.8,
		expected:    10,
	},
	{
		description: "A dart whose coordinates sum to > 5 but whose radius to origin is <= 5 is scored in the middle circle",
		x:           2.0,
		y:           4.0,
		expected:    5,
	},
	{
		description: "A dart whose coordinates sum to > 10 but whose radius to origin is <= 10 is scored in the outer circle",
		x:           4.0,
		y:           8.0,
		expected:    1,
	},
}
