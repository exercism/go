# Introduction

## Numbers

Go contains basic numeric types that can represent sets of either integer or floating-point values.
There are different types depending on the size of value you require and the architecture of the computer where the application is running (e.g. 32-bit and 64-bit).

For the sake of this exercise you will only be dealing with:

- `int`: e.g. `0`, `255`, `2147483647`. A signed integer that is at least 32 bits in size (value range of: -2147483648 through 2147483647).
  But this will depend on the systems architecture.
  Most modern computers are 64 bit, therefore `int` will be 64 bits in size (value rate of: -9223372036854775808 through 9223372036854775807).

- `float64`: e.g. `0.0`, `3.14`. Contains the set of all 64-bit floating-point numbers.

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

Go has shorthand assignment for the operators above (e.g. `a += 5` is short for `a = a + 5`).
Go also supports the increment and decrement statements `++` and `--` (e.g. `a++`).

## Converting between int and float64

Converting between types is done via a function with the name of the type to convert to.
For example:

```go
var x int = 42 // x has type int
f := float64(x) // f has type float64 (ie. 42.0)
var y float64 = 11.9 // y has type float64
i := int(y) // i has type int (ie. 11)
```

## If Statements

Conditionals in Go are similar to conditionals in other languages.
The underlying type of any conditional operation is the `bool` type, which can have the value of `true` or `false`.
Conditionals are often used as flow control mechanisms to check for various conditions.

For checking a particular case an `if` statement can be used, which executes its code if the underlying condition is `true` like this:

```go
var value string

if value == "val" {
    return "was val"
}
```

In scenarios involving more than one case many `if` statements can be chained together using the `else if` and `else` statements.

```go
var number int
result := "This number is "

if number > 0 {
    result += "positive"
} else if number < 0 {
    result += "negative"
} else {
    result += "zero"
}
```

If statements can also use a short initialization statement that can be used to initialize one or more variables for the if statement.
For example:

```go
num := 7
if v := 2 * num; v > 10 {
    fmt.Println(v)
} else {
    fmt.Println(num)
}
// Output: 7
```

> Note: any variables created in the initialization statement go out of scope after the end of the if statement.
