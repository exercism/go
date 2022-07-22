package markdown

import "testing"

func TestMarkdown(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual := Render(tc.input); actual != tc.expected {
				t.Fatalf("Render(%q)\n got:%q\nwant:%q", tc.input, actual, tc.expected)
			}
		})
	}
}

func BenchmarkMarkdown(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}

	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			Render(test.input)
		}
	}
}
