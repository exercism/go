package strand

import "testing"

const testVersion = 1

// Retired testVersions
// (none) 3e164bf0858b6bf7dbed8ff8a8d487105a41ada6

func TestRnaTranscription(t *testing.T) {
	if TestVersion != testVersion {
		t.Fatalf("Found TestVersion = %v, want %v", TestVersion, testVersion)
	}
	for _, test := range rnaTests {
		if actual := ToRna(test.input); actual != test.expected {
			t.Errorf("ToRna(%s): %s, expected %s",
				test.input, actual, test.expected)
		}
	}
}

func BenchmarkRnaTranscription(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range rnaTests {
			ToRna(test.input)
		}
	}
}
