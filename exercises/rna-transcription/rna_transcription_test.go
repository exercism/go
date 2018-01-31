package strand

import "testing"

func TestRNATranscription(t *testing.T) {
	for _, test := range rnaTests {
		if actual := ToRNA(test.input); actual != test.expected {
			t.Fatalf("FAIL: %s - ToRNA(%q): %q, expected %q",
				test.description, test.input, actual, test.expected)
		}
		t.Logf("PASS: %s", test.description)
	}
}

func BenchmarkRNATranscription(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range rnaTests {
			ToRNA(test.input)
		}
	}
}
