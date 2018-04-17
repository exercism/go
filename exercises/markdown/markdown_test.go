package markdown

import (
	"fmt"
	"testing"
)

func TestMarkdown(t *testing.T) {
	for _, test := range testCases {
		if html := Render(test.input); html != test.expected {
			t.Fatalf("FAIL: source(%q) = %q, want %q.", test.input, html, test.expected)
		}
		fmt.Printf("PASS: %s\n", test.description)
	}
}
