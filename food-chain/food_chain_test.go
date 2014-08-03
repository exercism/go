package foodchain

import (
	"strings"
	"testing"
)

const testVersion = 1

func TestTestVersion(t *testing.T) {
	if TestVersion != testVersion {
		t.Errorf("Found TestVersion = %v, want %v.", TestVersion, testVersion)
	}
}

var ref = []string{``,

	`I know an old lady who swallowed a fly.
I don't know why she swallowed the fly. Perhaps she'll die.`,

	`I know an old lady who swallowed a spider.
It wriggled and jiggled and tickled inside her.
She swallowed the spider to catch the fly.
I don't know why she swallowed the fly. Perhaps she'll die.`,

	`I know an old lady who swallowed a bird.
How absurd to swallow a bird!
She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.
She swallowed the spider to catch the fly.
I don't know why she swallowed the fly. Perhaps she'll die.`,

	`I know an old lady who swallowed a cat.
Imagine that, to swallow a cat!
She swallowed the cat to catch the bird.
She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.
She swallowed the spider to catch the fly.
I don't know why she swallowed the fly. Perhaps she'll die.`,

	`I know an old lady who swallowed a dog.
What a hog, to swallow a dog!
She swallowed the dog to catch the cat.
She swallowed the cat to catch the bird.
She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.
She swallowed the spider to catch the fly.
I don't know why she swallowed the fly. Perhaps she'll die.`,

	`I know an old lady who swallowed a goat.
Just opened her throat and swallowed a goat!
She swallowed the goat to catch the dog.
She swallowed the dog to catch the cat.
She swallowed the cat to catch the bird.
She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.
She swallowed the spider to catch the fly.
I don't know why she swallowed the fly. Perhaps she'll die.`,

	`I know an old lady who swallowed a cow.
I don't know how she swallowed a cow!
She swallowed the cow to catch the goat.
She swallowed the goat to catch the dog.
She swallowed the dog to catch the cat.
She swallowed the cat to catch the bird.
She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.
She swallowed the spider to catch the fly.
I don't know why she swallowed the fly. Perhaps she'll die.`,

	`I know an old lady who swallowed a horse.
She's dead, of course!`,
}

func TestVerse(t *testing.T) {
	for v := 1; v <= 8; v++ {
		if ret := Verse(v); ret != ref[v] {
			t.Fatalf("Verse(%d) =\n%s\n  want:\n%s", v, ret, ref[v])
		}
	}
}

func TestVerses(t *testing.T) {
	if ret, want := Verses(1, 2), ref[1]+"\n\n"+ref[2]; ret != want {
		t.Fatalf("Verses(1, 2) =\n%s\n  want:\n%s", ret, want)
	}

}

func TestSong(t *testing.T) {
	if ret, want := Song(), strings.Join(ref[1:], "\n\n"); ret != want {
		t.Fatalf("Song() =\n%s\n  want:\n%s", ret, want)
	}
}

func BenchmarkSong(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Song()
	}
}
