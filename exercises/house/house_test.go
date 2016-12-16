// Embed places a phrase with another phrase.
//
//    func Embed(prefixPhrase, suffixPhrase string) string

// Verse generates a verse with relative phrases that have a recursive structure.
//
//    func Verse(prefixPhrase string, relPhrases []string, suffixPhrase string) string

// As ever, there are different ways to do this. Try using Embed as a
// subroutine of Verse, and using programmatic recursion to generate the return
// value, reflecting the grammatical recursion of the song.

// Song generates the full text of "The House That Jack Built". Of course, you
// could return a string literal, but humor us; try using Verse as a subroutine.
//
//    func Song() string

package house

import (
	"strings"
	"testing"
)

var (
	s = "This is"
	r = []string{
		"the cat that broke",
		"the vase that was on",
	}
	p    = "the shelf."
	last = len(r) - 1
	// song copied from readme
	song = `This is the house that Jack built.

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
)

func TestEmbed(t *testing.T) {
	l := r[last]
	want := l + " " + p
	if e := Embed(l, p); e != want {
		t.Fatalf("Embed(%q, %q) = %q, want %q.", l, p, e, want)
	}
}

func TestVerse(t *testing.T) {
	for i := len(r); i >= 0; i-- {
		ri := r[i:]
		want := s + " " + strings.Join(append(ri, p), " ")
		if v := Verse(s, ri, p); v != want {
			t.Fatalf("Verse(%q, %q, %q) = %q, want %q.", s, ri, p, v, want)
		}
	}
}

func TestSong(t *testing.T) {
	s := Song()
	if s == song {
		return
	}
	// a little help in locating an error
	got := strings.Split(s, "\n")
	want := strings.Split(song, "\n")
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
	t.Fatalf("Song() line %d = %q, want %q", i+1, g, w)
}
