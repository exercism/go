package cryptosquare

import "testing"

var tests = []struct {
	pt string
	ct string
}{
	{
		"s#$%^&plunk",
		"supnl k",
	},
	{
		"1, 2, 3 GO!",
		"1g2o3",
	},
	{
		"1234",
		"1324",
	},
	{
		"123456789",
		"14725 8369",
	},
	{
		"123456789abc",
		"15926 a37b4 8c",
	},
	{
		"Never vex thine heart with idle woes",
		"neewl exhie vtetw ehaho ririe vntds",
	},
	{
		"ZOMG! ZOMBIES!!!",
		"zzioo emmsg b",
	},
	{
		"Time is an illusion. Lunchtime doubly so.",
		"tasne yinic dsmio hooel ntuil libsu uml",
	},
	{
		"We all know interspecies romance is weird.",
		"wneia weore neaws scili prerl neoid ktcms",
	},
	{
		"Madness, and then illumination.",
		"msemo aanin dninn dlaet ltshu i",
	},
	{
		"Vampires are people too!",
		"vrela epems etpao oirpo",
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
		"12",
	},
	{
		"123",
		"132",
	},
	{
		"12345678",
		"14725 836",
	},
	{
		"123456789a",
		"15926 a3748",
	},
}

func TestEncode(t *testing.T) {
	for _, test := range tests {
		if ct := Encode(test.pt); ct != test.ct {
			t.Fatalf("Encode(%q) = %q, want %q", test.pt, ct, test.ct)
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
