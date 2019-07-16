package darts

// Source: exercism/problem-specifications
// Commit: 212baa3 [Darts] Slim down generated method names (#1530)
// Problem Specifications Version: 2.2.0

var testCases = []struct {
	description string
	x           float64
	y           float64
	expected    int
}{
	{
		description: "Missed target",
		x:           -9.0,
		y:           9.0,
		expected:    0,
	},
	{
		description: "On the outer circle",
		x:           0.0,
		y:           10.0,
		expected:    1,
	},
	{
		description: "On the middle circle",
		x:           -5.0,
		y:           0.0,
		expected:    5,
	},
	{
		description: "On the inner circle",
		x:           0.0,
		y:           -1.0,
		expected:    10,
	},
	{
		description: "Exactly on centre",
		x:           0.0,
		y:           0.0,
		expected:    10,
	},
	{
		description: "Near the centre",
		x:           -0.1,
		y:           -0.1,
		expected:    10,
	},
	{
		description: "Just within the inner circle",
		x:           0.7,
		y:           0.7,
		expected:    10,
	},
	{
		description: "Just outside the inner circle",
		x:           0.8,
		y:           -0.8,
		expected:    5,
	},
	{
		description: "Just within the middle circle",
		x:           -3.5,
		y:           3.5,
		expected:    5,
	},
	{
		description: "Just outside the middle circle",
		x:           -3.6,
		y:           -3.6,
		expected:    1,
	},
	{
		description: "Just within the outer circle",
		x:           -7.0,
		y:           7.0,
		expected:    1,
	},
	{
		description: "Just outside the outer circle",
		x:           7.1,
		y:           -7.1,
		expected:    0,
	},
	{
		description: "Asymmetric position between the inner and middle circles",
		x:           0.5,
		y:           -4.0,
		expected:    5,
	},
}
