package strand

import "testing"

const testVersion = 2

func TestRNATranscription(t *testing.T) {
	if TestVersion != testVersion {
		t.Fatalf("Found TestVersion = %v, want %v", TestVersion, testVersion)
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
