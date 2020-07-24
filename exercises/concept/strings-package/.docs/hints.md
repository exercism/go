## General

- Strings in Go are based on `[]byte`. For example getting a substring works the same as getting a subslice.
- The `Go by Example` tutorial has two topics that are helpful here: [String Functions][string-functions] and [String Formatting][string-formatting]
- The `strings` package of the standard library has many useful [built-in functions][strings-package].
- The `fmt` package of the standard library has some [formatting functionality for strings][fmt-package].
- Code that does not check if an index exists in a `slice`, will `panic` if that index does not exist.
  A `panic` should be considered a bug in the code to be fixed. There will be more on panics in a later exercise.

## 1. Get message from a log line

- Splitting a string at a certain substring is available in the [`strings` package][split-function].
- Removing white space is [built-in][trimspace-function].

## 2. Get the message length in characters

- A string can be converted into a slice of characters (`[]rune`).

## 3. Get log level from a log line

- Changing a `string`'s casing to lower can be done with [this function][lowercase-function].

## 4. Reformat a log line

- You can just add strings to one another with the plus (`+`) operator. The preferred way is to use the [`fmt` packages][sprintf-function] functionality.

[strings-package]: https://golang.org/pkg/strings/
[string-functions]: https://gobyexample.com/string-functions
[string-formatting]: https://gobyexample.com/string-formatting
[fmt-package]: https://golang.org/pkg/fmt/
[split-function]: https://golang.org/pkg/strings/#Split
[trimspace-function]: https://golang.org/pkg/strings/#TrimSpace
[lowercase-function]: https://golang.org/pkg/strings/#ToLower
[sprintf-function]: https://golang.org/pkg/fmt/#Sprintf
