# Introduction to Type Conversion

In Go, assignment of a value between different types requires explicit conversion.
Converting between types is done via a function with the name of the type to convert to.
A value `v` is converted to type `T` using `T(v)`.
For example, to convert an `int` to a `float64` you would need to do the following:

```go
var x int = 42 // x has type int
f := float64(x) // f has type float64 (ie. 42.0)
```
