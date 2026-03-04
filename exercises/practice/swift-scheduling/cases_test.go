package swiftscheduling

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: a8ca9c5 Add `swift-scheduling` exercise (#2536)

var testCases = []struct {
	description string
	start       string
	delivery    string
	expected    string
}{
	{
		description: "NOW translates to two hours later",
		start:       "2012-02-13T09:00:00",
		delivery:    "NOW",
		expected:    "2012-02-13T11:00:00",
	},
	{
		description: "ASAP before one in the afternoon translates to today at five in the afternoon",
		start:       "1999-06-03T09:45:00",
		delivery:    "ASAP",
		expected:    "1999-06-03T17:00:00",
	},
	{
		description: "ASAP at one in the afternoon translates to tomorrow at one in the afternoon",
		start:       "2008-12-21T13:00:00",
		delivery:    "ASAP",
		expected:    "2008-12-22T13:00:00",
	},
	{
		description: "ASAP after one in the afternoon translates to tomorrow at one in the afternoon",
		start:       "2008-12-21T14:50:00",
		delivery:    "ASAP",
		expected:    "2008-12-22T13:00:00",
	},
	{
		description: "EOW on Monday translates to Friday at five in the afternoon",
		start:       "2025-02-03T16:00:00",
		delivery:    "EOW",
		expected:    "2025-02-07T17:00:00",
	},
	{
		description: "EOW on Tuesday translates to Friday at five in the afternoon",
		start:       "1997-04-29T10:50:00",
		delivery:    "EOW",
		expected:    "1997-05-02T17:00:00",
	},
	{
		description: "EOW on Wednesday translates to Friday at five in the afternoon",
		start:       "2005-09-14T11:00:00",
		delivery:    "EOW",
		expected:    "2005-09-16T17:00:00",
	},
	{
		description: "EOW on Thursday translates to Sunday at eight in the evening",
		start:       "2011-05-19T08:30:00",
		delivery:    "EOW",
		expected:    "2011-05-22T20:00:00",
	},
	{
		description: "EOW on Friday translates to Sunday at eight in the evening",
		start:       "2022-08-05T14:00:00",
		delivery:    "EOW",
		expected:    "2022-08-07T20:00:00",
	},
	{
		description: "EOW translates to leap day",
		start:       "2008-02-25T10:30:00",
		delivery:    "EOW",
		expected:    "2008-02-29T17:00:00",
	},
	{
		description: "2M before the second month of this year translates to the first workday of the second month of this year",
		start:       "2007-01-02T14:15:00",
		delivery:    "2M",
		expected:    "2007-02-01T08:00:00",
	},
	{
		description: "11M in the eleventh month translates to the first workday of the eleventh month of next year",
		start:       "2013-11-21T15:30:00",
		delivery:    "11M",
		expected:    "2014-11-03T08:00:00",
	},
	{
		description: "4M in the ninth month translates to the first workday of the fourth month of next year",
		start:       "2019-11-18T15:15:00",
		delivery:    "4M",
		expected:    "2020-04-01T08:00:00",
	},
	{
		description: "Q1 in the first quarter translates to the last workday of the first quarter of this year",
		start:       "2003-01-01T10:45:00",
		delivery:    "Q1",
		expected:    "2003-03-31T08:00:00",
	},
	{
		description: "Q4 in the second quarter translates to the last workday of the fourth quarter of this year",
		start:       "2001-04-09T09:00:00",
		delivery:    "Q4",
		expected:    "2001-12-31T08:00:00",
	},
	{
		description: "Q3 in the fourth quarter translates to the last workday of the third quarter of next year",
		start:       "2022-10-06T11:00:00",
		delivery:    "Q3",
		expected:    "2023-09-29T08:00:00",
	},
}
