# About Type Conversion

Go requires explicit conversion between different types.
For example, this causes a compiler error:

```go
var x int = 42
var f float64 = x // compiler error x is an int not a float64
```

Converting between types (also know as **type casting**) is done via a function with the name of the type to convert to.
A value `v` is converted to type `T` using `T(v)`.
For example, to convert an `int` to a `float64` you would need to do the following:

```go
var x int = 42 // x has type int
f := float64(x) // f has type float64 (ie. 42.0)
```

The same applies to any custom types:

```go
type Id int
var number int = 121 // number has type int
userId := Id(int) // userId now has type Id
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

WARNING: using the standard casting method does not have this behavior.
For example:

```go
var num int = 65
str := string(num) // str is now "A" not "65"
```

Here are some of the other common conversion methods:

| Method      | Purpose                     |
| ----------- | --------------------------- |
| ParseBool   | Convert `string` to `bool`  |
| FormatBool  | Convert `bool` to `string`  |
| ParseFloat  | Convert `string` to `float` |
| FormatFloat | Convert `float` to `string` |
| ParseInt    | Convert `string` to `int`   |
| FormatInt   | Convert `int` to `string`   |
