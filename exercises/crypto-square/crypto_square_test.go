package cryptosquare

import "testing"

var tests = []struct {
	pt string // plain text
	ct string // cipher text
}{
	{
		"s#$%^&plunk",
		"su pn lk",
	},
	{
		"1, 2, 3 GO!",
		"1g 2o 3 ",
	},
	{
		"1234",
		"13 24",
	},
	{
		"123456789",
		"147 258 369",
	},
	{
		"123456789abc",
		"159 26a 37b 48c",
	},
	{
		"Never vex thine heart with idle woes",
		"neewl exhie vtetw ehaho ririe vntds",
	},
	{
		"ZOMG! ZOMBIES!!!",
		"zzi ooe mms gb ",
	},
	{
		"Time is an illusion. Lunchtime doubly so.",
		"tasney inicds miohoo elntu  illib  suuml ",
	},
	{
		"We all know interspecies romance is weird.",
		"wneiaw eorene awssci liprer lneoid ktcms ",
	},
	{
		"Madness, and then illumination.",
		"msemo aanin dnin  ndla  etlt  shui ",
	},
	{
		"Vampires are people too!",
		"vrel aepe mset paoo irpo",
	},
	{
		"",
		"",
	},
	{
		"1",
		"1",
	},
	{
		"12",
		"1 2",
	},
	{
		"12 3",
		"13 2 ",
	},
	{
		"12345678",
		"147 258 36 ",
	},
	{
		"123456789a",
		"159 26a 37  48 ",
	},
	{
		"If man was meant to stay on the ground god would have given us roots",
		"imtgdvs fearwer mayoogo anouuio ntnnlvt wttddes aohghn  sseoau ",
	},
	{
		"Have a nice day. Feed the dog & chill out!",
		"hifei acedl veeol eddgo aatcu nyhht",
	},
}

func TestEncode(t *testing.T) {
	for _, test := range tests {
		if ct := Encode(test.pt); ct != test.ct {
			t.Fatalf(`Encode(%q):
got  %q
want %q`, test.pt, ct, test.ct)
		}
	}
}

func BenchmarkEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			Encode(test.pt)
		}
	}
}
