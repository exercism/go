package space

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: d137db1 Format using prettier (#1917)

var testCases = []struct {
	description string
	planet      Planet
	seconds     float64
	expected    float64
}{
	{
		description: "age on Earth",
		planet:      "Earth",
		seconds:     1000000000.0,
		expected:    31.69,
	},
	{
		description: "age on Mercury",
		planet:      "Mercury",
		seconds:     2134835688.0,
		expected:    280.88,
	},
	{
		description: "age on Venus",
		planet:      "Venus",
		seconds:     189839836.0,
		expected:    9.78,
	},
	{
		description: "age on Mars",
		planet:      "Mars",
		seconds:     2129871239.0,
		expected:    35.88,
	},
	{
		description: "age on Jupiter",
		planet:      "Jupiter",
		seconds:     901876382.0,
		expected:    2.41,
	},
	{
		description: "age on Saturn",
		planet:      "Saturn",
		seconds:     2000000000.0,
		expected:    2.15,
	},
	{
		description: "age on Uranus",
		planet:      "Uranus",
		seconds:     1210123456.0,
		expected:    0.46,
	},
	{
		description: "age on Neptune",
		planet:      "Neptune",
		seconds:     1821023456.0,
		expected:    0.35,
	},
	{
		description: "invalid planet causes error",
		planet:      "Sun",
		seconds:     680804807.0,
		expected:    -1.00,
	},
}
