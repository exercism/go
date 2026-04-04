# About

Comments let you leave notes in code without affecting how it runs.
A single-line comment starts with `//` anywhere on a line.
The compiler ignores everything from `//` to the end of that line.
When a comment comes after code on the same line, it is often called an inline comment.
A multiline comment starts with `/*` and ends with `*/` anywhere on a line.
The compiler ignores everything in between, spanning one or more lines.

```go
// This is a single-line comment on its own line

x := 1 // This is an single-line comment inline with code
fmt.Println(x) // prints 1

/*
This is a valid
multiline comment.
*/

/* This is also a valid multi-line comment. */

```

Comments serve different purposes depending on where they appear.

## Code comments

Code comments explain why something is done, not what.
Use them when the reason behind a decision isn't obvious from the code:

```go
results = results[1:] // skipping the header row
```

In most cases, code that is clearly written is self-explanatory, and comments are not needed.
Occasionally, it is necessary for the code to do something complex or unexpected; that type of code should be commented.

## Doc comments

Doc comments are used to document a package, variable or function.
This documentation is intended for other coders that would like to know how to _use_ your code.
It should explain, at a high level, what the code does and how to use it.
It should not include details of how the code itself works.
Doc comments sit directly before a declaration, with no blank line separating the comment from the declaration.
[`godoc`][godoc] parses them to generate package documentation, as seen on [pkg.go.dev][go packages].
A doc comment should be a complete sentence starting with the name of the identifier and ending with a period.

Doc comments on exported identifiers (those that begin with an uppercase letter) appear in generated documentation.
Doc comments on unexported identifiers (those that don't begin with an uppercase letter) won't appear, but they can still provide useful context within the code.

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

[godoc]: https://go.dev/blog/godoc
[go packages]: https://pkg.go.dev/
[doc-comments]: https://go.dev/doc/comment
