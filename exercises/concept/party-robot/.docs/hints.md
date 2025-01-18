# Hints

## General

- The `fmt` package of the standard library has some [formatting functionality for strings][fmt-package].
- There is a way to [concatenate strings][string-concatenation].

## 1. Welcome a new guest to the party

- `fmt.Sprintf` lets you use a [template to create a string][sprintf].

## 2. Welcome a new guest to the party whose birthday is today

- `fmt.Sprintf` can interpolate [more than one value into your string][sprintf-multiple-values]!

## 3. Give directions

- You can use [newline characters in your string][string-newline].
- Padding with zeroes could be achieved with one of the [formatting verbs][formatting-verbs].
- The shorthand assignment operator `+=` can help you build up your result via string concatenation.

[fmt-package]: https://golang.org/pkg/fmt/
[string-concatenation]: https://golang.org/ref/spec#String_concatenation
[sprintf]: https://pkg.go.dev/fmt#Sprintf
[sprintf-multiple-values]: https://www.geeksforgeeks.org/fmt-sprintf-function-in-golang-with-examples/
[string-newline]: https://yourbasic.org/golang/multiline-string/#interpreted-string-literals
[formatting-verbs]: https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/#cheat-sheet
