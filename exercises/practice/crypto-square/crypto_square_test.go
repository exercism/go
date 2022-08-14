package cryptosquare

import "testing"

var tests = []struct {
	input    string // plain text
	expected string // cipher text
}{
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
		input:    "If man was meant to stay on the ground god would have given us roots",
		expected: "imtgdvs fearwer mayoogo anouuio ntnnlvt wttddes aohghn  sseoau ",
	},
	{
		input:    "Have a nice day. Feed the dog & chill out!",
		expected: "hifei acedl veeol eddgo aatcu nyhht",
	},
}

func TestEncode(t *testing.T) {
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			if got := Encode(test.input); got != test.expected {
				t.Fatalf("Encode(%q):\n got:%q\nwant:%q", test.input, got, test.expected)
			}
		})
	}
}

func BenchmarkEncode(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			Encode(test.input)
		}
	}
}
