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

var (
	// song copied from README
	expectedSong = `This is the house that Jack built.

This is the malt
that lay in the house that Jack built.

This is the rat
that ate the malt
that lay in the house that Jack built.

This is the cat
that killed the rat
that ate the malt
that lay in the house that Jack built.

This is the dog
that worried the cat
that killed the rat
that ate the malt
that lay in the house that Jack built.

This is the cow with the crumpled horn
that tossed the dog
that worried the cat
that killed the rat
that ate the malt
that lay in the house that Jack built.

This is the maiden all forlorn
that milked the cow with the crumpled horn
that tossed the dog
that worried the cat
that killed the rat
that ate the malt
that lay in the house that Jack built.

This is the man all tattered and torn
that kissed the maiden all forlorn
that milked the cow with the crumpled horn
that tossed the dog
that worried the cat
that killed the rat
that ate the malt
that lay in the house that Jack built.

This is the priest all shaven and shorn
that married the man all tattered and torn
that kissed the maiden all forlorn
that milked the cow with the crumpled horn
that tossed the dog
that worried the cat
that killed the rat
that ate the malt
that lay in the house that Jack built.

This is the rooster that crowed in the morn
that woke the priest all shaven and shorn
that married the man all tattered and torn
that kissed the maiden all forlorn
that milked the cow with the crumpled horn
that tossed the dog
that worried the cat
that killed the rat
that ate the malt
that lay in the house that Jack built.

This is the farmer sowing his corn
that kept the rooster that crowed in the morn
that woke the priest all shaven and shorn
that married the man all tattered and torn
that kissed the maiden all forlorn
that milked the cow with the crumpled horn
that tossed the dog
that worried the cat
that killed the rat
that ate the malt
that lay in the house that Jack built.

This is the horse and the hound and the horn
that belonged to the farmer sowing his corn
that kept the rooster that crowed in the morn
that woke the priest all shaven and shorn
that married the man all tattered and torn
that kissed the maiden all forlorn
that milked the cow with the crumpled horn
that tossed the dog
that worried the cat
that killed the rat
that ate the malt
that lay in the house that Jack built.`

	expectedVerses = strings.Split(expectedSong, "\n\n")
)

func TestVerse(t *testing.T) {
	for v := 0; v < len(expectedVerses); v++ {
		if ret := Verse(v + 1); ret != expectedVerses[v] {
			t.Fatalf("Verse(%d) =\n%q\n  want:\n%q", v+1, ret, expectedVerses[v])
		}
	}
}

func TestSong(t *testing.T) {
	s := Song()
	if s == expectedSong {
		return
	}
	// a little help in locating an error
	gotStanzas := len(strings.Split(s, "\n\n"))
	wantStanzas := len(expectedVerses)
	if wantStanzas != gotStanzas {
		t.Fatalf("Song() has %d verse(s), want %d verses", gotStanzas, wantStanzas)
	}
	got := strings.Split(s, "\n")
	want := strings.Split(expectedSong, "\n")
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
