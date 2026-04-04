# Introduction

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

## Doc comments

Doc comments sit directly before a declaration, with no blank line separating the comment from the declaration.
[`godoc`][godoc] parses them to generate package documentation.
A doc comment should be a complete sentence starting with the name of the identifier and ending with a period.

Package comments begin with the package name:

```go
// Package kelvin provides tools to convert temperatures to and from Kelvin.
package kelvin
```

Function comments begin with the function name:

```go
// CelsiusFreezingTemp returns an integer value equal to the temperature at which water freezes in degrees Celsius.
func CelsiusFreezingTemp() int {
	return 0
}
```

[godoc]: https://go.dev/blog/godoc
