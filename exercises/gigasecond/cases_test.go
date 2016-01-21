package gigasecond

// Source: exercism/x-common
// Commit: 1e9e232 Merge pull request #45 from soniakeys/gigasecond-tests

// Add one gigasecond to the input.
var addCases = []struct {
	in   string
	want string
}{
	{
		"2011-04-25",
		"2043-01-01T01:46:40",
	},
	{
		"1977-06-13",
		"2009-02-19T01:46:40",
	},
	{
		"1959-07-19",
		"1991-03-27T01:46:40",
	},
	{
		"2015-01-24T22:00:00",
		"2046-10-02T23:46:40",
	},
	{
		"2015-01-24T23:59:59",
		"2046-10-03T01:46:39",
	},
}
