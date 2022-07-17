package prime

// Source: exercism/problem-specifications
// Commit: 42dd0ce Remove version (#1678)

var tests = []struct {
	description string
	input       int
	expected    int
	err         string
}{
	{
		description: "first prime",
		input:       1,
		expected:    2,
		err:         "",
	},
	{
		description: "second prime",
		input:       2,
		expected:    3,
		err:         "",
	},
	{
		description: "sixth prime",
		input:       6,
		expected:    13,
		err:         "",
	},
	{
		description: "big prime",
		input:       10001,
		expected:    104743,
		err:         "",
	},
	{
		description: "there is no zeroth prime",
		input:       0,
		expected:    0,
		err:         "there is no zeroth prime",
	},
}
