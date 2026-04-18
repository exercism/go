package twelvedays

import (
	"fmt"
	"strings"
	"testing"
)

// diff compares two multi-line strings and returns a helpful comment
func diff(got, want string) string {
	g := strings.Split(got, "\n")
	w := strings.Split(want, "\n")
	for i := 0; ; i++ {
		switch {
		case i < len(g) && i < len(w):
			if g[i] == w[i] {
				continue
			}
			return fmt.Sprintf("-- first difference in line %d:\n"+
				"-- got : %q\n-- want: %q\n", i+1, g[i], w[i])
		case i < len(g):
			return fmt.Sprintf("-- got %d extra lines after line %d:\n"+
				"-- first extra line: %q\n", len(g)-len(w), i, g[i])
		case i < len(w):
			return fmt.Sprintf("-- got %d correct lines, want %d more lines:\n"+
				"-- want next: %q\n", i, len(w)-i, w[i])
		default:
			return "no differences found"
		}
	}
}

func TestVerse(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			got := Verse(tc.verse)
			if got != tc.expected {
				t.Errorf("Verse(%d)\n got: %q\nwant: %q", tc.verse, got, tc.expected)
			}
		})
	}
}

func TestSong(t *testing.T) {
	expected := strings.Join(song, "\n")
	if actual := Song(); expected != actual {
		t.Fatalf("Song() =\n%s\n  want:\n%s\n%s", actual, expected, diff(actual, expected))
	}
}

func BenchmarkVerse(b *testing.B) {
	for range b.N {
		for _, tc := range testCases {
			Verse(tc.verse)
		}
	}
}

func BenchmarkSong(b *testing.B) {
	for range b.N {
		Song()
	}
}
