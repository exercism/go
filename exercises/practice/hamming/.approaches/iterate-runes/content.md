# Iterate as runes

```go
// Package hamming is a small library for determining the hamming distance between two strands.
package hamming

import (
	"errors"
	"unicode/utf8"
)

// Distance accepts two strands and returns their hamming distance.
func Distance(strand1 string, strand2 string) (distance int, e error) {
	if utf8.RuneCountInString(strand1) != utf8.RuneCountInString(strand2) {
		return distance, errors.New("strands must be of equal length")
	}

	strand1Runes := []rune(strand1)
	strand2Runes := []rune(strand2)
	for idx, rooney := range strand1Runes {
		if rooney != strand2Runes[idx] {
			distance++
		}
	}
	return distance, e
}
```
