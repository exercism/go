# Introduction

A `string` in Go is an immutable sequence of bytes. Strings may contain arbitrary bytes but usually they contain human-readable text.
Text strings are conventionally interpreted as UTF-8 encoded sequence of Unicode code points (runes) which will be explained in a future exercise.
A `string` value can be written as a string literal, which is a sequence of bytes enclosed in double quotes:

```go
s := "Hello World!"
```

Go provides an in-built package called `fmt` (format package) which offers a variety of functions to manipulate the format of input and output.
The most commonly used function is `Sprinf`, which uses verbs like `%s` to interpolate values into a string and returns that string.

```go
food := "taco"
fmt.Sprintf("Bring me a %s", food)
// Returns: Bring me a taco
```

In Go floating point values are conveniently formatted with Sprintf's verbs: `%g` (compact representation), `%e` (exponent) or `%f` (non exponent).
All three verbs allow the field's width and numeric position to be controlled.

```go
f := 4.3242
fmt.Sprintf("%.2f", f)
// Returns: 4.32
```

`fmt` contains other functions for working with strings, such as `Println` which simply prints the input received and `Printf` which formats the input in the same was as `Sprintf` beford print it.
