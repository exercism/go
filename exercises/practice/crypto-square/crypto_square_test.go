package cryptosquare

import "testing"

// Tests to add on after the canonical test cases, from cases_test.go
var additionalTests = []testCase{
	{
		input:    "s#$%^&plunk",
		expected: "su pn lk",
	},
	{
		input:    "1, 2, 3 GO!",
		expected: "1g 2o 3 ",
	},
	{
		input:    "1234",
		expected: "13 24",
	},
	{
		input:    "123456789",
		expected: "147 258 369",
	},
	{
		input:    "123456789abc",
		expected: "159 26a 37b 48c",
	},
	{
		input:    "Never vex thine heart with idle woes",
		expected: "neewl exhie vtetw ehaho ririe vntds",
	},
	{
		input:    "ZOMG! ZOMBIES!!!",
		expected: "zzi ooe mms gb ",
	},
	{
		input:    "Time is an illusion. Lunchtime doubly so.",
		expected: "tasney inicds miohoo elntu  illib  suuml ",
	},
	{
		input:    "We all know interspecies romance is weird.",
		expected: "wneiaw eorene awssci liprer lneoid ktcms ",
	},
	{
		input:    "Madness, and then illumination.",
		expected: "msemo aanin dnin  ndla  etlt  shui ",
	},
	{
		input:    "Vampires are people too!",
		expected: "vrel aepe mset paoo irpo",
	},
	{
		input:    "",
		expected: "",
	},
	{
		input:    "1",
		expected: "1",
	},
	{
		input:    "12",
		expected: "1 2",
	},
	{
		input:    "12 3",
		expected: "13 2 ",
	},
	{
		input:    "12345678",
		expected: "147 258 36 ",
	},
	{
		input:    "123456789a",
		expected: "159 26a 37  48 ",
	},
	{
		input:    "Have a nice day. Feed the dog & chill out!",
		expected: "hifei acedl veeol eddgo aatcu nyhht",
	},
}

func TestEncode(t *testing.T) {
	for _, tc := range append(testCases, additionalTests...) {
		description := tc.description
		if len(description) == 0 {
			description = tc.input
		}
		t.Run(description, func(t *testing.T) {
			if got := Encode(tc.input); got != tc.expected {
				t.Fatalf("Encode(%q):\n got:%q\nwant:%q", tc.input, got, tc.expected)
			}
		})
	}
}

func BenchmarkEncode(b *testing.B) {
	for range b.N {
		for _, tc := range append(testCases, additionalTests...) {
			Encode(tc.input)
		}
	}
}
