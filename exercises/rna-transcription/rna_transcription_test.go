package strand

import "testing"

const targetTestVersion = 3

func TestRNATranscription(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v", testVersion, targetTestVersion)
	}
	for _, test := range rnaTests {
		if actual := ToRNA(test.input); actual != test.expected {
			t.Errorf("ToRNA(%s): %s, expected %s",
				test.input, actual, test.expected)
		}
	}
}

func BenchmarkRNATranscription(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range rnaTests {
			ToRNA(test.input)
		}
	}
}
