package twelve

import (
	"fmt"
	"strings"
)

var verse = map[string]string{
	"first":    "a Partridge in a Pear Tree.",
	"twelfth":  "twelve Drummers Drumming",
	"eleventh": "eleven Pipers Piping",
	"tenth":    "ten Lords-a-Leaping",
	"ninth":    "nine Ladies Dancing",
	"eighth":   "eight Maids-a-Milking",
	"seventh":  "seven Swans-a-Swimming",
	"sixth":    "six Geese-a-Laying",
	"fifth":    "five Gold Rings",
	"fourth":   "four Calling Birds",
	"third":    "three French Hens",
	"second":   "two Turtle Doves",
}

var wording = map[int]string{
	1:  "first",
	2:  "second",
	3:  "third",
	4:  "fourth",
	5:  "fifth",
	6:  "sixth",
	7:  "seventh",
	8:  "eighth",
	9:  "ninth",
	10: "tenth",
	11: "eleventh",
	12: "twelfth",
}

func Verse(i int) string {
	var gifts = ""
	for from := i; from > 0; from-- {
		if i != 1 && from == 1 {
			gifts = fmt.Sprintf("%s, and %s", gifts, verse[wording[from]])
		} else if i != from {
			gifts = fmt.Sprintf("%s, %s", gifts, verse[wording[from]])
		} else {
			gifts = fmt.Sprintf("%s: %s", gifts, verse[wording[from]])
		}
	}
	return fmt.Sprintf("On the %s day of Christmas my true love gave to me%s", wording[i], gifts)
}

func Song() string {
	var song []string
	for i := 1; i <= 12; i++ {
		song = append(song, Verse(i))
	}
	return strings.Join(song, "\n")
}
