// As ever, there are different ways to complete this exercise.
// Try using using programmatic recursion to generate the verses of the song,
// thus reflecting the song's grammatical recursion.

// While recursion isn't always the simplest or most efficient solution to a problem,
// it's a powerful programming technique nonetheless.
//
// New to recursion? Here's a quick introduction:
// https://www.golang-book.com/books/intro/7#section5

package house

import (
	"strings"
	"testing"
)

func TestVerse(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if got := Verse(tc.verse); got != tc.expected {
				t.Fatalf("Verse(%d)\ngot:\n%q\nwant:\n%q", tc.verse, got, tc.expected)
			}
		})
	}
}

func TestSong(t *testing.T) {
	s := Song()
	if s == strings.Join(song, "\n\n") {
		return
	}
	// a little help in locating an error
	gotStanzas := len(strings.Split(s, "\n\n"))
	wantStanzas := len(song)
	if wantStanzas != gotStanzas {
		t.Fatalf("Song() has %d verse(s), want %d verses", gotStanzas, wantStanzas)
	}
	got := strings.Split(s, "\n")
	want := strings.Split(strings.Join(song, "\n\n"), "\n")
	var g, w string
	var i int
	for i, w = range want {
		if len(got) <= i {
			g = ""
			break
		}
		if g = got[i]; g != w {
			break
		}
	}
	t.Fatalf("Song() line %d =\n%q\n want \n%q", i+1, g, w)
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
