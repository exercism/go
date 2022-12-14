# Iterate as bytes

```go
// Package hamming is a small library for determining the hamming distance between two strands.
package hamming

import "errors"

// Distance accepts two strands and returns their hamming distance.
func Distance(strand1, strand2 string) (distance int, e error) {

	if len(strand1) != len(strand2) {
		return distance, errors.New("strands must be of equal length")
	}
	for i := 0; i < len(strand1); i++ {
		if strand1[i] != strand2[i] {
			distance++
		}
	}
	return distance, nil
}
```
