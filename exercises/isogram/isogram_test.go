package isogram

import "testing"

var testCases = []struct {
	word     string
	expected bool
}{
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

const targetTestVersion = 1

func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Errorf("Found testVersion = %v, want %v.", testVersion, targetTestVersion)
	}
}
