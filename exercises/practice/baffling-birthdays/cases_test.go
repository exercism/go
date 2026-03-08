package bafflingbirthdays

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 1799950 Add `baffling-birthdays` exercise (#2539)

var sharedTestCase = []struct {
	description string
	input       []string
	expected    bool
}{
	{
		description: "one birthdate",
		input:       []string{"2000-01-01"},
		expected:    false,
	},
	{
		description: "two birthdates with same year, month, and day",
		input:       []string{"2000-01-01", "2000-01-01"},
		expected:    true,
	},
	{
		description: "two birthdates with same year and month, but different day",
		input:       []string{"2012-05-09", "2012-05-17"},
		expected:    false,
	},
	{
		description: "two birthdates with same month and day, but different year",
		input:       []string{"1999-10-23", "1988-10-23"},
		expected:    true,
	},
	{
		description: "two birthdates with same year, but different month and day",
		input:       []string{"2007-12-19", "2007-04-27"},
		expected:    false,
	},
	{
		description: "two birthdates with different year, month, and day",
		input:       []string{"1997-08-04", "1963-11-23"},
		expected:    false,
	},
	{
		description: "multiple birthdates without shared birthday",
		input:       []string{"1966-07-29", "1977-02-12", "2001-12-25", "1980-11-10"},
		expected:    false,
	},
	{
		description: "multiple birthdates with one shared birthday",
		input:       []string{"1966-07-29", "1977-02-12", "2001-07-29", "1980-11-10"},
		expected:    true,
	},
	{
		description: "multiple birthdates with more than one shared birthday",
		input:       []string{"1966-07-29", "1977-02-12", "2001-12-25", "1980-07-29", "2019-02-12"},
		expected:    true,
	},
}

var probabilityTestCase = []struct {
	description string
	input       int
	expected    float64
}{
	{
		description: "for one person",
		input:       1,
		expected:    0,
	},
	{
		description: "among ten people",
		input:       10,
		expected:    11.694818,
	},
	{
		description: "among twenty-three people",
		input:       23,
		expected:    50.729723,
	},
	{
		description: "among seventy people",
		input:       70,
		expected:    99.915958,
	},
}
