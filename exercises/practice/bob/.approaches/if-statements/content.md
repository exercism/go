# `if` statements

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

	if shout {
		if question {
			return "Calm down, I know what I'm doing!"
		} 
		return "Whoa, chill out!"
	}
	if question {
		return "Sure."
	}
	return "Whatever."
}
```

In this approach you have a series of `if` statements using the private functions to evaluate the conditions.

```exercism/note
A private (also called unexported) function is indicated by starting with a lowercase letter.
More info on exported and unexported functions can be found [here](https://yourbasic.org/golang/public-private/).
```


As soon as the right condition is found, the correct response is returned.

```exercism/note
Note that there are no `else if` or `else` statements.
If an `if` statement can return, then an `else if` or `else` is not needed.
Execution will either return or will continue to the next statement anyway.
```

In the `isShout()` function, the [`strings`][strings] function [`IndexFunc()`][indexfunc] is used to ensure there is at least one letter character in the input.
If not, the function returns `false`, because if the input were only `"123"` it would equal itself uppercased, but without letters it would not be a shout.
If the input does contain a letter, the function returns whether the uppercased input is the same as the input.
The uppercasing is done by using the [`ToUpper()`][toupper] function.

In the `isQuestion()` function, the `strings` function [`HasSuffix()`][hassuffix] is used to determine if the input ends with a question mark.

In the `Hey()` function, the [`TrimSpace()`][trimspace] function is used to eliminate any whitespace at the ends of the input.
If the string has no characters left, it returns the response for saying nothing.

If the string is not empty, the `shout` and `question` variables are set.

If `shout` is `true`, then the function returns the response for if the input is a shouted question os just a shout.

If the input is a question, then the function returns the response for that.

Finally, if the function has not returned by the end, the response for neither a shout nor a question is returned.

[strings]: https://pkg.go.dev/strings@go1.19.4
[indexfunc]: https://pkg.go.dev/strings@go1.19.4#IndexFunc
[toupper]: https://pkg.go.dev/strings@go1.19.4#ToUpper
[hassuffix]: https://pkg.go.dev/strings#example-HasSuffix
[trimspace]: https://pkg.go.dev/strings#TrimSpace
