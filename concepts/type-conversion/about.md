# About Type Conversion

In Go, assignment of a value between different types requires explicit conversion.
For example, this is a compiler error:

```go
var x int = 42
var f float64 = x // compiler error x is an int not a float64
```

Converting between types is done via a function with the name of the type to convert to.
A value `v` is converted to type `T` using `T(v)`.
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
var i int = strconv.Atoi(intString)

var number int = 12
var s string = strconv.Itoa(number)
```

Here are some of the other common conversion methods:

| Method      | Purpose                                       |
| ----------- | --------------------------------------------- |
| ParseBool   | Convert `string` to `bool`                    |
| FormatBool  | Convert `bool` to `string`                    |
| ParseFloat  | Convert `string` to `float`                   |
| FormatFloat | Convert `float` to `string`                   |
| ParseInt    | Convert `string` to `int` (if base is not 10) |
| FormatInt   | Convert `int` to `string` (if base is not 10) |
