package isogram

import "testing"

func TestIsIsogram(t *testing.T) {
	for _, c := range testCases {
		if IsIsogram(c.input) != c.expected {
			t.Fatalf("FAIL: %s\nWord %q, expected %t, got %t", c.description, c.input, c.expected, !c.expected)
		}

		t.Logf("PASS: Word %q", c.input)
	}
}

func BenchmarkIsIsogram(b *testing.B) {
	b.StopTimer()
	for _, c := range testCases {
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			IsIsogram(c.input)
		}

		b.StopTimer()
	}
}
