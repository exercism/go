package twelve

import (
	"fmt"
	"strings"
	"testing"
)

type testCase struct {
	input    int
	expected string
}

var testCases = []testCase{
	{1, "On the first day of Christmas my true love gave to me, a Partridge in a Pear Tree."},
	{2, "On the second day of Christmas my true love gave to me, two Turtle Doves, and a Partridge in a Pear Tree."},
	{3, "On the third day of Christmas my true love gave to me, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree."},
	{4, "On the fourth day of Christmas my true love gave to me, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree."},
	{5, "On the fifth day of Christmas my true love gave to me, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree."},
	{6, "On the sixth day of Christmas my true love gave to me, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree."},
	{7, "On the seventh day of Christmas my true love gave to me, seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree."},
	{8, "On the eighth day of Christmas my true love gave to me, eight Maids-a-Milking, seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree."},
	{9, "On the ninth day of Christmas my true love gave to me, nine Ladies Dancing, eight Maids-a-Milking, seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree."},
	{10, "On the tenth day of Christmas my true love gave to me, ten Lords-a-Leaping, nine Ladies Dancing, eight Maids-a-Milking, seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree."},
	{11, "On the eleventh day of Christmas my true love gave to me, eleven Pipers Piping, ten Lords-a-Leaping, nine Ladies Dancing, eight Maids-a-Milking, seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree."},
	{12, "On the twelfth day of Christmas my true love gave to me, twelve Drummers Drumming, eleven Pipers Piping, ten Lords-a-Leaping, nine Ladies Dancing, eight Maids-a-Milking, seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree."},
}

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

func TestSong(t *testing.T) {
	var expected = ""
	for _, test := range testCases {
		expected += test.expected + "\n"
	}
	actual := Song()
	if expected != actual {
		t.Fatalf("Song() =\n%s\n  want:\n%s\n%s", actual, expected, diff(actual, expected))
	}
}

func TestVerse(t *testing.T) {
	for _, test := range testCases {
		actual := Verse(test.input)
		if actual != test.expected {
			t.Errorf("Twelve Days test [%d], expected [%s], actual [%s]", test.input, test.expected, actual)
		}
	}
}
