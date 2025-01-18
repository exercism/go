# `switch` statement

```go
// Package bob is a library for replying to Bob.
package bob

import (
	"strings"
	"unicode"
)

func isShout(phrase string) bool {
	if strings.IndexFunc(phrase, unicode.IsLetter) == -1 {
		return false
	}
	return strings.ToUpper(phrase) == phrase
}

func isQuestion(phrase string) bool {
	return strings.HasSuffix(phrase, "?")
}

func isEmpty(s string) bool {
	return s == ""
}

// Bob replies to what is remarked to Bob.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if isEmpty(remark) {
		return "Fine. Be that way!"
	}

	question := isQuestion(remark)
	shout := isShout(remark)

	switch true {
	case question && shout:
		return "Calm down, I know what I'm doing!"
	case shout:
		return "Whoa, chill out!"
	case question:
		return "Sure."
	default:
		return "Whatever."
	}
}
```

In this approach you use a [`switch`][switch] statement to test if there is a question or a shout.
The `switch`, with the help of some private functions to evaluate the conditions,
returns the right response for a question, shout, shouted question, or for not being a shout or a question.

~~~~exercism/note
A private (also called unexported) function is indicated by starting with a lowercase letter.
More info on exported and unexported functions can be found [here](https://yourbasic.org/golang/public-private/).
~~~~

In the `isShout()` function, the [`strings`][strings] function [`IndexFunc()`][indexfunc] is used to ensure there is at least one letter character in the input.
If not, the function returns `false`, because if the input were only `"123"` it would equal itself uppercased, but without letters it would not be a shout.
If the input does contain a letter, the function returns whether the uppercased input is the same as the input.
The uppercasing is done by using the [`ToUpper()`][toupper] function.

In the `isQuestion()` function, the `strings` function [`HasSuffix()`][hassuffix] is used to determine if the input ends with a question mark.

In the `Hey()` function, the [`TrimSpace()`][trimspace] function is used to eliminate any whitespace at the ends of the input.
If the string has no characters left, it returns the response for saying nothing.

If the string is not empty, the `shout` and `question` variables are set.

The `question` and `shout` values are tested in a `switch`.
The [logical AND operator][logical-ops] (`&&`) is used to check if both `question` and `shout` are `true`.
If a switch arm evaluates to `true`, the appropriate response is returned.
If neither `question` nor `shout` is `true`, the `default` arm of the `switch` returns the response when the input is neither a question nor a shout.

[switch]: https://go.dev/tour/flowcontrol/9
[strings]: https://pkg.go.dev/strings
[trimspace]: https://pkg.go.dev/strings#TrimSpace
[indexfunc]: https://pkg.go.dev/strings@go1.19.4#IndexFunc
[toupper]: https://pkg.go.dev/strings@go1.19.4#ToUpper
[hassuffix]: https://pkg.go.dev/strings#example-HasSuffix
[logical-ops]: https://go.dev/ref/spec#Logical_operators
