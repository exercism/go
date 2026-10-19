package resistorcolortolerance

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: dc0574a Adding a fourth canonical exercise to the 'resistor color' series (#2682)

type labelTestCase struct {
	description string
	input       []string
	expected    string
}

var labelTestCases = []labelTestCase{
	{
		description: "Orange, orange, black and red",
		input:       []string{"orange", "orange", "black", "red"},
		expected:    "33 ohms ±2%",
	},
	{
		description: "Blue, gray, brown and violet",
		input:       []string{"blue", "grey", "brown", "violet"},
		expected:    "680 ohms ±0.1%",
	},
	{
		description: "Red, black, red and green",
		input:       []string{"red", "black", "red", "green"},
		expected:    "2 kiloohms ±0.5%",
	},
	{
		description: "Green, brownm orange and grey",
		input:       []string{"green", "brown", "orange", "grey"},
		expected:    "51 kiloohms ±0.05%",
	},
	{
		description: "Only one band",
		input:       []string{"black"},
		expected:    "0 ohms",
	},
	{
		description: "Orange, orange, yellow, black and brown",
		input:       []string{"orange", "orange", "yellow", "black", "brown"},
		expected:    "334 ohms ±1%",
	},
	{
		description: "Red, green, yellow, yellow and brown",
		input:       []string{"red", "green", "yellow", "yellow", "brown"},
		expected:    "2.54 megaohms ±1%",
	},
	{
		description: "Blue, gray, white, brown and brown",
		input:       []string{"blue", "grey", "white", "brown", "brown"},
		expected:    "6.89 kiloohms ±1%",
	},
	{
		description: "Violet, orange, red and gray",
		input:       []string{"violet", "orange", "red", "grey"},
		expected:    "7.3 kiloohms ±0.05%",
	},
	{
		description: "Brown, red, orange, green and blue",
		input:       []string{"brown", "red", "orange", "green", "blue"},
		expected:    "12.3 megaohms ±0.25%",
	},
	{
		description: "Brown, black, brown, yellow and violet",
		input:       []string{"brown", "black", "brown", "yellow", "violet"},
		expected:    "1.01 megaohms ±0.1%",
	},
	{
		description: "Brown, black, red and red",
		input:       []string{"brown", "black", "red", "red"},
		expected:    "1 kiloohms ±2%",
	},
}
