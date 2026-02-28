# Introduction

There are at least two idomatic approaches to solving ISBN Verifier in Go.
You can iterate the characters as [bytes][bytes].
Or you can iterate the characters as [runes][runes].

## General guidance

One important aspect to solving ISBN is to process the characters as digits according to their position, while disregarding dashes,
and paying special attention that an `X` character is only a legal value of `10` if it is in the last legal digit position.
Another important aspect is that the number of digit positions must be ten.

## Approach: Iterate as bytes

```go
// Package isbn is a small library for validating input as an isbn.
package isbnverifier

func getDigit(chr byte) (digit int, ok bool) {
	if chr < '0' || chr > '9' {
		return digit, ok
	}
	return int(chr - '0'), true
}

// IsValidISBN calculates the validity of the input as an ISBN.
func IsValidISBN(input string) bool {
	length := len(input)
	pos := 10
	sum := 0

	for i := 0; i < length; i++ {
		chr := input[i]
		if digit, ok := getDigit(chr); ok {
			sum += (digit * pos)
			pos--
		} else if pos == 1 && chr == 'X' {
			sum += 10
			pos--
		} else if chr != '-' {
			return false
		}
	}
	return pos == 0 && sum%11 == 0
}
```

For more information, check the [iterate as bytes approach][approach-iterate-bytes].

## Approach: Iterate as runes

```go
// Package isbn is a small library for validating input as an isbn.
package isbnverifier

func getDigit(chr rune) (digit int, ok bool) {
	if chr < '0' || chr > '9' {
		return digit, ok
	}
	return int(chr - '0'), true
}

// IsValidISBN calculates the validity of the input as an ISBN.
func IsValidISBN(input string) bool {
	pos := 10
	sum := 0

	for _, chr := range input {
		if digit, ok := getDigit(chr); ok {
			sum += (digit * pos)
			pos--
		} else if pos == 1 && chr == 'X' {
			sum += 10
			pos--
		} else if chr != '-' {
			return false
		}
	}
	return pos == 0 && sum%11 == 0
}
```

For more information, check the [iterate as runes approach][approach-iterate-runes].

## Which approach to use?

To iterate as bytes benchmarked the fastest.
As of the time of this writing we can iterate bytes, since all of the characters are [ASCII][ascii].

To compare performance of the approaches, check the [Performance article][article-performance].

[bytes]: https://pkg.go.dev/bytes
[runes]: https://golangdocs.com/rune-in-golang
[approach-iterate-bytes]: https://exercism.org/tracks/go/exercises/isbn-verifier/approaches/iterate-bytes
[approach-iterate-runes]: https://exercism.org/tracks/go/exercises/isbn-verifier/approaches/iterate-runes
[article-performance]: https://exercism.org/tracks/go/exercises/isbn-verifier/articles/performance
[ascii]: https://www.asciitable.com/
