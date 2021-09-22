# Hints

## General

- The [strings package][strings-package] contains many useful `string` functions.
- The `strings` package needs to be imported before it can be used.
- To call functions within the `strings` package, prefix them with `strings.`.

## 1. Create the welcome message

- Strings can be concatenated using the `+` operator.
- There is a function to [convert a `string` to upper case][to-upper].

## 2. Add a fancy border

- There is a function to [create a `string` with a specific character repeated a number of times][repeat].
- A newline is a special [escape character][escape-characters].

## 3. Clean up old marketing messages

- There is a function to [replace strings within a `string`][replace].
- There is a function to [trim leading and trailing spaces from a `string`][trim-space].

[strings-package]: https://pkg.go.dev/strings
[to-upper]: https://pkg.go.dev/strings#ToUpper
[repeat]: https://pkg.go.dev/strings#Repeat
[replace-all]: https://pkg.go.dev/strings#ReplaceAll
[trim-space]: https://pkg.go.dev/strings#TrimSpace
[escape-characters]: https://yourbasic.org/golang/multiline-string/#all-escape-characters
