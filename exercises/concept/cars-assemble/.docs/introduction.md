# Introduction

## Numbers

Go contains basic numeric types that can represent sets of either integer or floating-point values.
There are different types depending on the size of value you require and the architecture of the computer where the application is running (e.g. 32-bit and 64-bit).

For the sake of this exercise you will only be dealing with:

- `int`: e.g. `0`, `255`, `2147483647`. A signed integer that is at least 32 bits in size (value range of: -2147483648 through 2147483647).
  But this will depend on the systems architecture.
  Most modern computers are 64 bit, therefore `int` will be 64 bits in size (value rate of: -9223372036854775808 through 9223372036854775807).

- `float64`: e.g. `0.0`, `3.14`. Contains the set of all 64-bit floating-point numbers.

Go supports the standard set of arithmetic operators of `+`, `-`, `*`, `/` and `%` (remainder not modulo).

## Arithmetic Operators

Go supports many standard arithmetic operators:

| Operator | Example        |
| -------- | -------------- |
| `+`      | `4 + 6 == 10`  |
| `-`      | `15 - 10 == 5` |
| `*`      | `2 * 3 == 6`   |
| `/`      | `13 / 3 == 4`  |
| `%`      | `13 % 3 == 1`  |

For integer division, the remainder is dropped (eg. `5 / 2 == 2`).

## Shorthand Assignments

These can be used in shorthand assignments to update and assign a variable using the operator:

```go
a := 1
a += 2 // same as a = a + 2 == 3

b := 2
b -= 2 // same as b = b - 2 == 0

c := 4
c *= 2 // same as c = c * 2 == 8

d := 8
d /= 2 // same as d = d / 2 == 4

e := 16
e %= 2 // same as e = e % 2 == 0
```

## Increment and Decrement

There are also two special statements: increment (`++`) and decrement (`--`).
They modify the value of a variable by increasing (or decreasing) the value by 1.
For example:

```go
a := 10
a++ // same as a += 1, a == 11

b := 10
b-- // same as b -= 1, b == 9
```

## Type Conversion

In Go, assignment of a value between different types requires explicit conversion.
Converting between types is done via a function with the name of the type to convert to.
For example, to convert an `int` to a `float64` you would need to do the following:

```go
var x int = 42 // x has type int
f := float64(x) // f has type float64 (ie. 42.0)
```

## Conditionals If

Conditionals in Go are similar to conditionals in other languages.
The underlying type of any conditional operation is the `bool` type, which can have the value of `true` or `false`.
Conditionals are often used as flow control mechanisms to check for various conditions.

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
