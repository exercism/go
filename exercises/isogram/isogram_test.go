package isogram

import "testing"

const targetTestVersion = 1

var testCases = []struct {
	word     string
	expected bool
}{
	// converted from x-common by hand
	// x-common version: 0.8.0.20160608
	{"duplicates", true},
	{"eleven", false},
	{"subdermatoglyphic", true},
	{"Alphabet", false},
	{"thumbscrew-japingly", true},
	{"Hjelmqvist-Gryb-Zock-Pfund-Wax", true},
	{"Heizölrückstoßabdämpfung", true},
	{"the quick brown fox", false},
	{"Emily Jung Schwartzkopf", true},
	{"éléphant", false},
}

func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v.", testVersion, targetTestVersion)
	}
}

func TestIsIsogram(t *testing.T) {
	for _, c := range testCases {
		if IsIsogram(c.word) != c.expected {
			t.Fatalf("FAIL: Word %q, expected %v, got %v", c.word, c.expected, !c.expected)
		}

		t.Logf("PASS: Word %q", c.word)
	}
}

func BenchmarkIsIsogram(b *testing.B) {
	b.StopTimer()
	for _, c := range testCases {
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			IsIsogram(c.word)
		}

		b.StopTimer()
	}
}
