package beer

import (
	"testing"
)

const targetTestVersion = 1

func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Errorf("Found testVersion = %v, want %v.", testVersion, targetTestVersion)
	}
}

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
	for _, tt := range verseTestCases {
		actualVerse, err := Verse(tt.verse)
		if tt.expectErr {
			// check if err is of error type
			var _ error = err

			// if we expect an error and there isn't one
			if err == nil {
				t.Errorf("Verse(%d): expected an error, but error is nil", tt.verse)
			}
		} else {
			if actualVerse != tt.expectedVerse {
				t.Fatalf("Verse(%d):\nexpected\n%s\nactual\n%s", tt.verse, tt.expectedVerse, actualVerse)
			}

			// if we don't expect an error and there is one
			if err != nil {
				t.Errorf("Verse(%d): expected no error, but error is: %s", tt.verse, err)
			}
		}
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

	for _, tt := range versesTestCases {
		actualVerse, err := Verses(tt.upperBound, tt.lowerBound)
		if tt.expectErr {
			// check if err is of error type
			var _ error = err

			// if we expect an error and there isn't one
			if err == nil {
				t.Errorf("Verses(%d, %d): expected an error, but error is nil", tt.upperBound, tt.lowerBound)
			}
		} else {
			if actualVerse != tt.expectedVerse {
				t.Fatalf("Verses(%d, %d):\nexpected\n%s\nactual\n%s", tt.upperBound, tt.lowerBound, tt.expectedVerse, actualVerse)
			}

			// if we don't expect an error and there is one
			if err != nil {
				t.Errorf("Verses(%d, %d): expected no error, but error is: %s", tt.upperBound, tt.lowerBound, err)
			}
		}
	}
}

func BenchmarkSeveralVerses(b *testing.B) {
	b.StopTimer()
	for _, tt := range versesTestCases {
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			Verses(tt.upperBound, tt.lowerBound)
		}

		b.StopTimer()
	}
}

func TestEntireSong(t *testing.T) {
	expected, err := Verses(99, 0)
	var _ error = err
	if err != nil {
		t.Fatalf("unexpected error calling Verses(99,0)")
	}
	actual := Song()

	if expected != actual {
		msg := `
		  Did not sing the whole song correctly.

			Expected:

			%v

			Actual:

			%v
		`
		t.Fatalf(msg, expected, actual)
	}
}

func BenchmarkEntireSong(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Song()
	}
}
