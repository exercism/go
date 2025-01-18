# Introduction

## Numbers

Go contains basic numeric types that can represent sets of either integer or floating-point values.
There are different types depending on the size of value you require and the architecture of the computer where the application is running (e.g. 32-bit and 64-bit).

For the sake of this exercise you will only be dealing with:

- `int`: e.g. `0`, `255`, `2147483647`. A signed integer that is at least 32 bits in size (value range of: -2147483648 through 2147483647).
  But this will depend on the systems architecture.
  Most modern computers are 64 bit, therefore `int` will be 64 bits in size (value rate of: -9223372036854775808 through 9223372036854775807).

- `float64`: e.g. `0.0`, `3.14`. Contains the set of all 64-bit floating-point numbers.

- `uint`: e.g. `0`, `255`. An unsigned integer that is the same size as `int` (value range of: 0 through 4294967295 for 32 bits and 0 through 18446744073709551615 for 64 bits)

Numbers can be converted to other numeric types through Type Conversion.

## Arithmetic Operators

Go supports many standard arithmetic operators:

| Operator | Example        |
| -------- | -------------- |
| `+`      | `4 + 6 == 10`  |
| `-`      | `15 - 10 == 5` |
| `*`      | `2 * 3 == 6`   |
| `/`      | `13 / 3 == 4`  |
| `%`      | `13 % 3 == 1`  |

For integer division, the remainder is dropped (e.g. `5 / 2 == 2`).

Go has shorthand assignment for the operators above (e.g. `a += 5` is short for `a = a + 5`).
Go also supports the increment and decrement statements `++` and `--` (e.g. `a++`).

## Converting between types 

Converting between types is done via a function with the name of the type to convert to.
For example:

```go
var x int = 42 // x has type int
f := float64(x) // f has type float64 (ie. 42.0)
var y float64 = 11.9 // y has type float64
i := int(y) // i has type int (ie. 11)
```
## Arithmetic operations on different types

In many languages you can perform arithmetic operations on different types of variables, but in Go this gives an error.
For example:

```go
var x int = 42

// this line produces an error
value := float32(2.0) * x // invalid operation: mismatched types float32 and int

// you must convert int type to float32 before performing arithmetic operation
value := float32(2.0) * float32(x)
```
