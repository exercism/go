package beer

import (
	"bytes"
	"fmt"
)

// Song returns the full lyrics for 99 bottles of beer
func Song() (result string) {
	result, _ = Verses(99, 0)
	return
}

// Verses returns an excerpt of the lyrics of 99 bottles of beer
// between verse start and stop. The verse numbers count backwards, so the
// first verse sung will be verse 99, and the last will be verse 0
func Verses(start, stop int) (string, error) {
	switch {
	case 0 > start || start > 99:
		return "", fmt.Errorf("start value[%d] is not a valid verse", start)
	case 0 > stop || stop > 99:
		return "", fmt.Errorf("stop value[%d] is not a valid verse", stop)
	case start < stop:
		return "", fmt.Errorf("start value[%d] is less than stop value[%d]", start, stop)
	}

	var buff bytes.Buffer
	for i := start; i >= stop; i-- {
		v, _ := Verse(i)
		buff.WriteString(v)
		buff.WriteString("\n")
	}
	return buff.String(), nil
}

// Verse returns a single verse of the lyrics of 99 bottles of beer. The verse
// numbers count backwards, so the first verse sung will be verse 99, and the
// last will be verse 0
func Verse(n int) (string, error) {
	result := ""
	switch {
	case 0 > n || n > 99:
		return "", fmt.Errorf("%d is not a valid verse", n)
	case n == 0:
		result = "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n"
	case n == 1:
		result = "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n"
	case n == 2:
		result = "2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.\n"
	default:
		result = fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n", n, n, n-1)
	}
	return result, nil
}
