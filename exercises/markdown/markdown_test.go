package markdown

import "testing"

func TestMarkdown(t *testing.T) {
	for _, test := range testCases {
		if html := Render(test.input); html != test.expected {
			t.Fatalf("FAIL: Render(%q) = %q, want %q.", test.input, html, test.expected)
		}
		t.Logf("PASS: %s\n", test.description)
	}
}

func BenchmarkMarkdown(b *testing.B) {
	// Benchmark time to parse all the test cases
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			Render(test.input)
		}
	}
}
