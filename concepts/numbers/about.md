# About

Go contains basic numeric types that can represent sets of either integer or floating-point values.
There a different types depending on the size of value you require and the architecture of the computer where the application is running (e.g. 32-bit and 64-bit).

This concept will concentrate on two number types:

- `int`: e.g. `0`, `255`, `2147483647`.
   A signed integer that is at least 32 bits in size (value range of: -2147483648 through 2147483647).
   But this will depend on the systems architecture.
   Most modern computers are 64 bit, therefore `int` will be 64 bits in size (value rate of: -9223372036854775808 through 9223372036854775807).

- `float64`: e.g. `0.0`, `3.14`.
   Contains the set of all 64-bit floating-point numbers.

For a full list of the available numeric types and more detail see the following resources:

- [Go builtin type declarations][go-builtins]
- [Digital Ocean - Understanding Data Types in Go][do-understanding-types]
- [Arden Labs - Understanding Type in Go][arden-understanding-types]

Go supports the standard set of arithmetic operators of `+`, `-`, `*`, `/` and `%` (remainder not modulo).

### Alternative number notations

Go supports several alternative notations for numeric literals that can improve readability and express values more naturally:

```go
// Underscores improve readability
population := 1_000_000

// Binary literals
permissions := 0b1101

// Hexadecimal literals
color := 0xFF

// Scientific notation
avogadro := 6.022e23
distance := 1.496e11
```

Underscores (`_`) can be used to separate digits for improved readability without affecting the value.
Binary (`0b`) and hexadecimal (`0x`) literals provide convenient representations for bit manipulation and low-level programming.
Scientific notation (`e`) is useful when working with very large or very small floating-point values.

### Complex numbers

Go has built-in support for complex numbers through the `complex64` and `complex128` types:

```go
z1 := complex(2, 3)      // 2 + 3i
z2 := 4 + 5i             // 4 + 5i
sum := z1 + z2           // 6 + 8i
```

Complex numbers are commonly used in scientific computing, signal processing, and mathematical applications.

In Go, assignment of a value between different types requires explicit conversion.
For example, to convert an `int` to a `float64` you would need to do the following:

```go
var x int = 42
f := float64(x)

fmt.Printf("x is of type: %s\n", reflect.TypeOf(x))
// Output: x is of type: int

fmt.Printf("f is of type: %s\n", reflect.TypeOf(f))
// Output: f is of type: float64
```

For more information on Type Conversion please see [A Tour of Go: Type Conversions][type-conversion]

[type-conversion]: https://tour.golang.org/basics/13
[go-builtins]: https://github.com/golang/go/blob/master/src/builtin/builtin.go
[do-understanding-types]: https://www.digitalocean.com/community/tutorials/understanding-data-types-in-go
[arden-understanding-types]: https://www.ardanlabs.com/blog/2013/07/understanding-type-in-go.html
