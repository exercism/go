package flowerfield

import (
	"slices"
	"testing"
)

func TestAnnotate(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			got := Annotate(tc.input)
			want := tc.expected
			if !slices.Equal(want, got) {
				t.Fatalf("Annotate(%#v) =\n%#v\nwant:\n%#v", tc.input, got, want)
			}
		})
	}
}

func BenchmarkAnnotate(b *testing.B) {
	board := []string{
		" *  * ",
		"  *   ",
		"    * ",
		"   * *",
		" *  * ",
		"      ",
	}
	for range b.N {
		Annotate(board)
	}
}
