# Answer array

```go
// Package bob is a library for replying to Bob.
package bob

import (
	"strings"
	"unicode"
)

func isShout(phrase string) int {
	if strings.IndexFunc(phrase, unicode.IsLetter) == -1 {
		return 0
	}
	if strings.ToUpper(phrase) == phrase {
		return 2
	}
	return 0
}

func isQuestion(phrase string) int {
	if strings.HasSuffix(phrase, "?") {
		return 1
	}
	return 0
}

func isEmpty(s string) bool {
	return s == ""
}

var answers = [...]string{"Whatever.", "Sure.", "Whoa, chill out!", "Calm down, I know what I'm doing!"}

// Bob replies to what is remarked to Bob.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if isEmpty(remark) {
		return "Fine. Be that way!"
	}
	return answers[isQuestion(remark)+isShout(remark)]
}
```

In this approach you define an array that contains Bob’s answers, and each condition is given a score.
The correct answer is selected from the array by using the score as the array index.

In the `isShout()` function, the [`strings`][strings] function [`IndexFunc()`][indexfunc] is used to ensure there is at least one letter character in the input.
If not, the function returns `0`, because if the input were only `"123"` it would equal itself uppercased, but without letters it would not be a shout.
If the input does contain a letter, the function returns `2` if the uppercased input is the same as the input,
otherwise it returns `0`.
The uppercasing is done by using the [`ToUpper()`][toupper] function.

In the `isQuestion()` function, the `strings` function [`HasSuffix()`][hassuffix] is used to return `1` if the input ends with a question mark,
otherwise it returns `0`.

An [array literal][array-literal] is used to define the responses.

In the `Hey()` function, the [`TrimSpace()`][trimspace] function is used to eliminate any whitespace at the ends of the input.
If the string has no characters left, it returns the response for saying nothing.

The conditions of being a question and being a shout are assigned scores by adding the calls to `isShout` and `isQuestion`.
For example, giving a question a score of `1` and a shout a score of `0` would use an index of `1` to get the element from the answers array, which is `"Sure."`.

| isShout | isQuestion | Index     | Answer                                |
| ------- | ---------- | --------- | ------------------------------------- |
|       0 |          0 | 0 + 0 = 0 | `"Whatever."`                         |
|       0 |          1 | 0 + 1 = 1 | `"Sure."`                             |
|       2 |          0 | 2 + 0 = 2 | `"Whoa, chill out!"`                  |
|       2 |          1 | 2 + 1 = 3 | `"Calm down, I know what I'm doing!"` |


[strings]: https://pkg.go.dev/strings@go1.19.4
[indexfunc]: https://pkg.go.dev/strings@go1.19.4#IndexFunc
[toupper]: https://pkg.go.dev/strings@go1.19.4#ToUpper
[hassuffix]: https://pkg.go.dev/strings#example-HasSuffix
[trimspace]: https://pkg.go.dev/strings#TrimSpace
[array-literal]: https://yourbasic.org/golang/three-dots-ellipsis/
