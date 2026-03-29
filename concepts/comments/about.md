# About

Comments let you leave notes in code without affecting how it runs.
Single-line comments start with `//`, and multiline comments are wrapped in `/*` and `*/`:

```go
// This is a single-line comment.

/*
This is a
multiline comment.
*/
```

Comments serve different purposes depending on where they appear.

## Code comments

Code comments explain why something is done, not what.
Use them when the reason behind a decision isn't obvious from the code:

```go
results = results[1:] // skipping the header row
```

Prefer clear code over explanatory comments.

## Doc comments

Doc comments sit directly before a declaration, with no blank line separating the comment from the declaration.
[`godoc`][godoc] parses them to generate package documentation, as seen on [pkg.go.dev][go packages].
A doc comment should be a complete sentence starting with the name of the identifier and ending with a period.

Doc comments on exported identifiers appear in generated documentation.
Doc comments on unexported identifiers won't appear, but they can still provide useful context within the code.

Comments documenting packages begin with the package name and describe what the package does:

```go
// Package kelvin provides tools to convert temperatures to and from Kelvin.
package kelvin
```

Comments documenting functions begin with the function name and describe what the function does, including its arguments and return values:

```go
// CelsiusFreezingTemp returns an integer value equal to the temperature at which water freezes in degrees Celsius.
func CelsiusFreezingTemp() int {
	return 0
}
```

Comments documenting variables begin with the variable name and describe what the variable represents:

```go
// TemperatureFahrenheit represents a certain temperature in degrees Fahrenheit.
var TemperatureFahrenheit float64
```

For the full doc comment specification, see the [Go documentation guide][doc-comments].

[`golint`][golint] was a popular tool for linting doc comments, but [`revive`][revive] is its drop-in replacement.


[godoc]: https://go.dev/blog/godoc
[go packages]: https://pkg.go.dev/
[revive]: https://github.com/mgechev/revive
[golint]: https://github.com/golang/lint
[doc-comments]: https://go.dev/doc/comment
