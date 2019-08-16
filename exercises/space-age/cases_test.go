package space

// Source: exercism/problem-specifications
// Commit: 28b3dac0 space-age: restrict seconds to fit within 32-bit int range
// Problem Specifications Version: 1.2.0

var testCases = []struct {
	description string
	planet      Planet
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
		seconds:     2129871239,
		expected:    35.88,
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
		seconds:     2000000000,
		expected:    2.15,
	},
	{
		description: "age on Uranus",
		planet:      "Uranus",
		seconds:     1210123456,
		expected:    0.46,
	},
	{
		description: "age on Neptune",
		planet:      "Neptune",
		seconds:     1821023456,
		expected:    0.35,
	},
}
