package proverb

// Source: exercism/problem-specifications
// Commit: 5ae1dba Acronym canonical-data: Remove redundant test case
// Problem Specifications Version: 1.3.0

type proverbTest struct {
	input    []string
	expected string
}

var stringTestCases = []proverbTest{
	{
		input:    []string{},
		expected: "",
	},
	{
		input:    []string{"nail"},
		expected: "And all for the want of a nail.",
	},
	{
		input: []string{"nail", "shoe"},
		expected: "For want of a nail the shoe was lost.\n" +
			"And all for the want of a nail.",
	},
	{
		input: []string{"nail", "shoe", "horse"},
		expected: "For want of a nail the shoe was lost.\n" +
			"For want of a shoe the horse was lost.\n" +
			"And all for the want of a nail.",
	},
	{
		input: []string{"nail", "shoe", "horse", "rider", "message", "battle", "kingdom"},
		expected: "For want of a nail the shoe was lost.\n" +
			"For want of a shoe the horse was lost.\n" +
			"For want of a horse the rider was lost.\n" +
			"For want of a rider the message was lost.\n" +
			"For want of a message the battle was lost.\n" +
			"For want of a battle the kingdom was lost.\n" +
			"And all for the want of a nail.",
	},
	{
		input: []string{"pin", "gun", "soldier", "battle"},
		expected: "For want of a pin the gun was lost.\n" +
			"For want of a gun the soldier was lost.\n" +
			"For want of a soldier the battle was lost.\n" +
			"And all for the want of a pin.",
	},
}
