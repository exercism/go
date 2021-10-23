# Introduction

In the previous exercise, we saw that there are two ways to write comments in Go: single-line comments that are preceded by `//`, and multiline comment blocks that are wrapped with `/*` and `*/`.

## Documentation comments

In Go, comments play an important role in documenting code. They are used by the `godoc` command, which extracts these comments to create documentation about Go packages. A documentation comment should be a complete sentence that starts with the name of the thing being described and ends with a period.

Comments should precede packages as well as exported identifiers, for example exported functions, methods, package variables, constants, and structs, which you will learn more about in the next exercises.

A package-level variable can look like this:

```go
// TemperatureCelsius represents a certain temperature in degrees Celsius.
var TemperatureCelsius float64
```

## Package comments

Package comments should be written directly before a package clause (`package x`) and begin with `Package x ...` like this:

```go
// Package kelvin provides tools to convert
// temperatures to and from Kelvin.
package kelvin
```

## Function comments

A function comment should be written directly before the function declaration. It should be a full sentence that starts with the function name. For example, an exported comment for the function `Calculate` should take the form `Calculate  ...`. It should also explain what arguments the function takes, what it does with them, and what its return values mean, ending in a period):

```go
// CelsiusFreezingTemp returns an integer value equal to the temperature at which water freezes in degrees Celsius.
func CelsiusFreezingTemp() int {
	return 0
}
```
