package bottlesong

import (
	"fmt"
)

func Recite(startBottles, takeDown int) []string {
	verses := []string{}
	for i := startBottles; i > startBottles-takeDown; i -= 1 {
		verses = append(verses, verse(i)...)

		if i > startBottles-takeDown+1 {
			verses = append(verses, "")
		}
	}
	return verses
}

var numberToWord = map[int]string{
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
}

func verse(n int) []string {
	switch {
	case n == 1:
		return []string{
			"One green bottle hanging on the wall,",
			"One green bottle hanging on the wall,",
			"And if one green bottle should accidentally fall,",
			"There'll be no green bottles hanging on the wall.",
		}
	case n == 2:
		return []string{
			"Two green bottles hanging on the wall,",
			"Two green bottles hanging on the wall,",
			"And if one green bottle should accidentally fall,",
			"There'll be one green bottle hanging on the wall.",
		}
	default:
		return []string{
			fmt.Sprintf("%s green bottles hanging on the wall,", Title(numberToWord[n])),
			fmt.Sprintf("%s green bottles hanging on the wall,", Title(numberToWord[n])),
			"And if one green bottle should accidentally fall,",
			fmt.Sprintf("There'll be %s green bottles hanging on the wall.", numberToWord[n-1]),
		}
	}
}
