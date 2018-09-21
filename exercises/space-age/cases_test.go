package space

// Source: exercism/problem-specifications
// Commit: 8d4df79 space-age: Apply new "input" policy
// Problem Specifications Version: 1.1.0

var testCases = []struct {
	description string
	planet      string
	seconds     float64
	expected    float64
}{
	{
		description: "age on Earth",
		planet:      "Earth",
		seconds:     1000000000,
		expected:    31.69,
	},
	{
		description: "age on Mercury",
		planet:      "Mercury",
		seconds:     2134835688,
		expected:    280.88,
	},
	{
		description: "age on Venus",
		planet:      "Venus",
		seconds:     189839836,
		expected:    9.78,
	},
	{
		description: "age on Mars",
		planet:      "Mars",
		seconds:     2329871239,
		expected:    39.25,
	},
	{
		description: "age on Jupiter",
		planet:      "Jupiter",
		seconds:     901876382,
		expected:    2.41,
	},
	{
		description: "age on Saturn",
		planet:      "Saturn",
		seconds:     3000000000,
		expected:    3.23,
	},
	{
		description: "age on Uranus",
		planet:      "Uranus",
		seconds:     3210123456,
		expected:    1.21,
	},
	{
		description: "age on Neptune",
		planet:      "Neptune",
		seconds:     8210123456,
		expected:    1.58,
	},
}
