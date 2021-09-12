# Introduction

## Numbers

Go contains basic numeric types that can represent sets of either integer or
floating-point values. There are different types depending on the size of value
you require and the architecture of the computer where the application is
running (e.g. 32-bit and 64-bit).

For the sake of this exercise you will only be dealing with:

- `int`: e.g. `0`, `255`, `2147483647`. A signed integer that is at least 32
  bits in size (value range of: -2147483648 through 2147483647). But this will
  depend on the systems architecture. Most modern computers are 64 bit,
  therefore `int` will be 64 bits in size (value rate of:
  -9223372036854775808 through 9223372036854775807).

- `float64`: e.g. `0.0`, `3.14`. Contains the set of all 64-bit floating-point
  numbers.

For a full list of the available numeric types and more detail see the
following resources:

- [Go builtin type declarations][go-builtins]
- [Digital Ocean - Understanding Data Types in Go][do-understanding-types]
- [Arden Labs - Understanding Type in Go][arden-understanding-types]

Go supports the standard set of arithmetic operators of `+`, `-`, `*`, `/`
and `%` (remainder not modulo).

## Type Conversion

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

## Conditionals If

Conditionals in Go are similar to conditionals in other languages. The underlying type of any conditional operation is the `bool` type, which can have the value of `true` or `false`. Conditionals are often used as flow control mechanisms to check for various conditions.

For checking a particular case an `if` statement can be used, which executes its code if the underlying condition is `true` like this:

```go
var value string

if value == "val" {
    // conditional code
}
```

In scenarios involving more than one case many `if` statements can be chained together using the `else if` and `else` statements.

```go
if value == "val" {
    // conditional code
} else if value == "val2" {
    // conditional code
} else {
    // default code
}
```

[go-builtins]: https://github.com/golang/go/blob/master/src/builtin/builtin.go
[do-understanding-types]: https://www.digitalocean.com/community/tutorials/understanding-data-types-in-go
[arden-understanding-types]: https://www.ardanlabs.com/blog/2013/07/understanding-type-in-go.html
