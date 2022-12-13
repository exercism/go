# Introduction

There are at least two idomatic approaches to solving Luhn in Go.
You can scrub out the spaces and test if the input is an integer first.
Or you can ignore a space and test if each character is a digit as you go.

## General guidance

One important aspect to solving Luhn is to allow for spaces in the input and to disallow all other non-numeric characters.
Another important aspect is to handle the value of each digit according to its position in the string.

## Approach: Scrub out the spaces and test if the input is an integer first

```go
// Package luhn is a small library for returning if a phrase is valid according to the Luhn algorithm.
package luhn

import (
	"strconv"
	"strings"
)

func Valid(num string) bool {
	num = strings.ReplaceAll(num, " ", "")
	if _, err := strconv.Atoi(num); err != nil {
		return false
	}
	idx := len(num) - 1
	total := 0
	pos := 0
	for i := idx; i > -1; i-- {
		digit := int(num[i] - '0')
		if pos%2 == 0 {
			total += digit
		} else {
			switch digit {
			case 1, 2, 3, 4:
				total += digit << 1
			case 5, 6, 7, 8:
				total += (digit << 1) - 9
			case 9:
				total += digit
			}
		}
		pos++
	}
	return pos > 1 && total%10 == 0
}
```

For more information, check the [scrub and validate first appproach][approach-scrub-and-validate-first].

## Approach: Validate as you go

```go
// Package luhn is a small library for returning if a phrase is valid according to the Luhn algorithm.
package luhn

func Valid(num string) bool {
	idx := len(num) - 1
	total := 0
	pos := 0
	for i := idx; i > -1; i-- {
		char := num[i]
		if char == ' ' {
			continue
		}
		if char < '0' || char > '9' {
			return false
		}
		digit := int(char - '0')
		if pos%2 == 0 {
			total += digit
		} else {
			switch digit {
			case 1, 2, 3, 4:
				total += digit << 1
			case 5, 6, 7, 8:
				total += (digit << 1) - 9
			case 9:
				total += digit
			}
		}
		pos++
	}
	return pos > 1 && total%10 == 0
}
```

For more information, check the [validate as you go approach][approach-validate-as-you-go].

## Which approach to use?

The validate as you go approach benchmarked the fastest.
To compare performance of the approaches, check the [Performance article][article-performance].

[approach-scrub-and-validate-first]: https://exercism.org/tracks/go/exercises/luhn/approaches/scrub-and-validate-first
[approach-validate-as-you-go]: https://exercism.org/tracks/go/exercises/luhn/approaches/validate-as-you-go
[article-performance]: https://exercism.org/tracks/go/exercises/luhn/articles/performance
