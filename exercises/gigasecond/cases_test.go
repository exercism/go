package gigasecond

// Source: exercism/problem-specifications
// Commit: 61e7d70 Fix "Gigasecond: Schema Compliance"
// Problem Specifications Version: 1.0.0

// Add one gigasecond to the input.
var addCases = []struct {
	description string
	in          string
	want        string
}{
	{
		"date only specification of time",
		"2011-04-25",
		"2043-01-01T01:46:40",
	},
	{
		"second test for date only specification of time",
		"1977-06-13",
		"2009-02-19T01:46:40",
	},
	{
		"third test for date only specification of time",
		"1959-07-19",
		"1991-03-27T01:46:40",
	},
	{
		"full time specified",
		"2015-01-24T22:00:00",
		"2046-10-02T23:46:40",
	},
	{
		"full time with day roll-over",
		"2015-01-24T23:59:59",
		"2046-10-03T01:46:39",
	},
}
