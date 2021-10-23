# Introduction

Go contains basic numeric types that can represent sets of either integer or
floating-point values.

This concept will concentrate on two number types:

- `int`: e.g. `0`, `255`, `2147483647`. A signed integer that is at least 32
  bits in size (value range of: -2147483648 through 2147483647). But this will
  depend on the systems architecture. Most modern computers are 64 bit,
  therefore `int` will be 64 bits in size (value rate of:
  -9223372036854775808 through 9223372036854775807).

- `float64`: e.g. `0.0`, `3.14`. Contains the set of all 64-bit floating-point
  numbers.

Go supports the standard set of arithmetic operators of `+`, `-`, `*`, `/`
and `%` (remainder not modulo).

In Go, assignment of a value between different types requires explicit
conversion. For example, to convert an `int` to a `float64` you would need to
do the following:

```go
var x int = 42
f := float64(x)

fmt.Printf("x is of type: %s\n", reflect.TypeOf(x))
// Output: x is of type: int

fmt.Printf("f is of type: %s\n", reflect.TypeOf(f))
// Output: f is of type: float64
```
