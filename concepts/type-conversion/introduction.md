# Introduction to Type Conversion

Go requires explicit conversion between different types.
Converting between types (also known as **type casting**) is done via a function with the name of the type to convert to.
For example, to convert an `int` to a `float64` you would need to do the following:

```go
var x int = 42 // x has type int
f := float64(x) // f has type float64 (ie. 42.0)
```

## Converting between primitive types and strings

There is a `strconv` package for converting between primitive types (like `int`) and `string`.

```go
import "strconv"

var intString string = "42"
var i, err = strconv.Atoi(intString)

var number int = 12
var s string = strconv.Itoa(number)
```
