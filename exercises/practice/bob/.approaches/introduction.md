# Introduction

There are various idiomatic approaches to solve Bob.
A basic approach can use a series of `if` statements to test the conditions.
Or a [switch][switch] statement can be used.
An array can contain answers from which the right response is selected by an index calculated from scores given to the conditions.

## General guidance

Regardless of the approach used, some things you could look out for include

- If the input is trimmed, [TrimSpace][trimspace] only once.

- Use the [HasSuffix][hassuffix] `string` method instead of checking the last character by index for `?`.

- Don't copy/paste the logic for determining a shout and for determing a question into determing a shouted question.
  Combine the two determinations instead of copying them.
  Not duplicating the code will keep the code [DRY][dry].

- Perhaps consider using `IsQuestion` and `IsShout` to set values once instead of being functions that are possibly called twice.

- If an `if` statement can return, then an `else if` or `else` is not needed.
  Execution will either return or will continue to the next statement anyway.


## Approach: `if` statements

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

For more information, check the [`if` statements approach][approach-if].

## Approach: `switch` statement

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

For more information, check the [`switch` statement approach][approach-switch].


## Approach: Answer array

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

For more information, check the [Answer array approach][approach-answer-array].


## Which approach to use?

Since each approach sometimes gave results slower than the other approaches when benchmarking, which to use could be a matter of stylistic choice. 

- To compare the performance of the approaches, check out the [Performance article][article-performance].

[switch]: https://go.dev/tour/flowcontrol/9
[trimspace]: https://pkg.go.dev/strings#TrimSpace
[hassuffix]: https://pkg.go.dev/strings#example-HasSuffix
[dry]: https://en.wikipedia.org/wiki/Don%27t_repeat_yourself
[approach-if]: https://exercism.org/tracks/go/exercises/bob/approaches/if-statements
[approach-switch]: https://exercism.org/tracks/go/exercises/bob/approaches/switch-statement
[approach-answer-array]: https://exercism.org/tracks/go/exercises/bob/approaches/answer-array
[article-performance]: https://exercism.org/tracks/go/exercises/bob/articles/performance
