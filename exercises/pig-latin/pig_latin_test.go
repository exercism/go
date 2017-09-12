package igpay

import "testing"

var tests = []struct{ pl, in string }{
	{"appleay", "apple"},
	{"earay", "ear"},
	{"igpay", "pig"},
	{"oalakay", "koala"},
	{"airchay", "chair"},
	{"eenquay", "queen"},
	{"aresquay", "square"},
	{"erapythay", "therapy"},
	{"ushthray", "thrush"},
	{"oolschay", "school"},
	{"ickquay astfay unray", "quick fast run"},
	{"ellowyay", "yellow"},
	{"yttriaay", "yttria"},
	{"enonxay", "xenon"},
	{"xrayay", "xray"},
}

func TestPigLatin(t *testing.T) {
	for _, test := range tests {
		if pl := PigLatin(test.in); pl != test.pl {
			t.Fatalf("PigLatin(%q) = %q, want %q.", test.in, pl, test.pl)
		}
	}
}
