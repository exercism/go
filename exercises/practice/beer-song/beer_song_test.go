package beer

import (
	"testing"
)

const verse8 = "8 bottles of beer on the wall, 8 bottles of beer.\nTake one down and pass it around, 7 bottles of beer on the wall.\n"
const verse3 = "3 bottles of beer on the wall, 3 bottles of beer.\nTake one down and pass it around, 2 bottles of beer on the wall.\n"
const verse2 = "2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.\n"
const verse1 = "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n"
const verse0 = "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n"

const verses86 = `8 bottles of beer on the wall, 8 bottles of beer.
Take one down and pass it around, 7 bottles of beer on the wall.

7 bottles of beer on the wall, 7 bottles of beer.
Take one down and pass it around, 6 bottles of beer on the wall.

6 bottles of beer on the wall, 6 bottles of beer.
Take one down and pass it around, 5 bottles of beer on the wall.

`

const verses75 = `7 bottles of beer on the wall, 7 bottles of beer.
Take one down and pass it around, 6 bottles of beer on the wall.

6 bottles of beer on the wall, 6 bottles of beer.
Take one down and pass it around, 5 bottles of beer on the wall.

5 bottles of beer on the wall, 5 bottles of beer.
Take one down and pass it around, 4 bottles of beer on the wall.

`

var verseTestCases = []struct {
	description   string
	verse         int
	expectedVerse string
	expectErr     bool
}{
	{"a typical verse", 8, verse8, false},
	{"another typical verse", 3, verse3, false},
	{"verse 2", 2, verse2, false},
	{"verse 1", 1, verse1, false},
	{"verse 0", 0, verse0, false},
	{"invalid verse", 104, "", true},
}

func TestBottlesVerse(t *testing.T) {
	for _, tc := range verseTestCases {
		t.Run(tc.description, func(t *testing.T) {
			actualVerse, err := Verse(tc.verse)
			if tc.expectErr {
				if err == nil {
					t.Fatalf("Verse(%d) expected an error, but error is nil", tc.verse)
				}
			} else {
				if err != nil {
					t.Fatalf("Verse(%d) returned error: %v, want:%q", tc.verse, err, tc.expectedVerse)
				}

				if actualVerse != tc.expectedVerse {
					t.Fatalf("Verse(%d)\n got:%q\nwant:%q", tc.verse, actualVerse, tc.expectedVerse)
				}
			}
		})
	}
}

var versesTestCases = []struct {
	description   string
	upperBound    int
	lowerBound    int
	expectedVerse string
	expectErr     bool
}{
	{"multiple verses", 8, 6, verses86, false},
	{"a different set of verses", 7, 5, verses75, false},
	{"invalid start", 109, 5, "", true},
	{"invalid stop", 99, -20, "", true},
	{"start less than stop", 8, 14, "", true},
}

func TestSeveralVerses(t *testing.T) {
	for _, tc := range versesTestCases {
		t.Run(tc.description, func(t *testing.T) {
			actualVerse, err := Verses(tc.upperBound, tc.lowerBound)
			if tc.expectErr {
				if err == nil {
					t.Fatalf("Verses(%d,%d) expected an error, but error is nil", tc.upperBound, tc.lowerBound)
				}
			} else {
				if err != nil {
					t.Fatalf("Verses(%d,%d) returned error: %v, want:%q", tc.upperBound, tc.lowerBound, err, tc.expectedVerse)
				}
				if actualVerse != tc.expectedVerse {
					t.Fatalf("Verse(%d,%d)\n got:%q\nwant:%q", tc.upperBound, tc.lowerBound, actualVerse, tc.expectedVerse)
				}
			}
		})
	}
}

func BenchmarkSeveralVerses(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {

		for _, tt := range versesTestCases {
			Verses(tt.upperBound, tt.lowerBound)
		}

	}
}

func TestEntireSong(t *testing.T) {
	expected, err := Verses(99, 0)
	if err != nil {
		t.Fatalf("unexpected error calling Verses(99,0)")
	}
	actual := Song()

	if expected != actual {
		t.Fatalf(`
		  Did not sing the whole song correctly.

			Expected:
			%v

			Actual:
			%v
		`, expected, actual)
	}
}

func BenchmarkEntireSong(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		Song()
	}
}
