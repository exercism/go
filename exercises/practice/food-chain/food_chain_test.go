package foodchain

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
	for _, tc := range verseTestCases {
		t.Run(tc.description, func(t *testing.T) {
			want := strings.Join(tc.expected, "\n")
			if got := Verse(tc.verse); got != want {
				t.Fatalf("Verse(%d)\ngot: %s\nwant: %s\nhelp: %s", tc.verse, got, want, diff(got, want))
			}
		})
	}
}

func TestVerses(t *testing.T) {
	for _, tc := range multiVerseTestCases {
		t.Run(tc.description, func(t *testing.T) {
			want := strings.Join(tc.expected, "\n")
			if got := Verses(tc.start, tc.end); got != want {
				t.Fatalf("Verses(%d, %d)\ngot: %s\nwant: %s\nhelp: %s", tc.start, tc.end, got, want, diff(got, want))
			}
		})
	}
}

func TestSong(t *testing.T) {
	want := strings.Join(song, "\n")
	if got := Song(); got != want {
		t.Fatalf("Song() =\n%s\n  want:\n%s\n%s", got, want, diff(got, want))
	}
}

func BenchmarkSong(b *testing.B) {
	for range b.N {
		Song()
	}
}
