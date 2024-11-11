# Introduction

`Bitwise operations` in Go are operations that apply to individual bits of a number, using the binary representation of numbers to produce a result.

They are most commonly used in cryptography, flag handling, and network programming.

## Binary Representation

To display a number in [binary format][binary-format-link], you can use the `fmt` package:

```go
a := 23
fmt.Printf("Number %d in binary: %b\n", a, a)  // -> 10111
fmt.Printf("Number %d in 8-bit binary format: %08b\n", a, a)  // -> 0001 0111
```

Understanding how decimal numbers are represented in binary is essential for using bitwise operators effectively.

~~~~exercism/advanced
Arithmetic operations in the base 2 number system - [binary arithmetic][binary-arithmetic-link].

How to convert a number with base 10 to a number with base 2 and back - [cpnversion][conversion-link].
~~~~

## Bitwise Operators

Bitwise operators perform operations on each bit of a number. Some operators are binary, used between two numbers (e.g., `a+b`), and others are unary, applied to a single number (e.g., `-a`).

__Binary Bitwise Operators__:

| Simbol | Name | Description |
|--------|------|-------------|
|&       |AND   |1 if both bits are 1|
|\|      |OR    |1 if at least one of the bits is 1|
|^       |XOR   |1 if the bits differ|

_Examples_:

```go
a := 154     // 1001 1010
b := 83      // 0101 0011

and := a & b // 0001 0010 -> 18
or  := a | b // 1101 1011 -> 219
xor := a ^ b // 1100 1001 -> 201
```

If you need to save the result back into the same variable, you can use the shortened form:

```go
a &= b // equivalent to a = a & b
```

~~~~exercism/note
Applying the XOR (`^`) operator to a number with itself always be 0:

```go
zero := a ^ a  // always 0
```
~~~~

__Unary NOT Operator__:

| Simbol | Name | Description |
|--------|------|-------------|
|^       |NOT   |Inverts the bits of a number|

_Example_:

```go
a := 154   // 1001 1010
not := ^a  // 0110 0101 -> 101
```

__Shift Operations__:

| Simbol | Name | Description |
|--------|------|-------------|
|<<|Shift left  |Shifts bits to the left by a specified number |
|>>|Shift right |Shifts bits to the right by a specified number|

_Examples_:

```go
a := 27         // 0001 1011
left := a << 2  // 0110 1100 -> 108 (two 0 was add)
right := a >> 2 // 0000 0110 -> 6  (two last bit was remove)
```

~~~~exercism/note
Each left shift is equivalent to multiplying by 2, and each right shift is equivalent to dividing by 2.
~~~~

## Usage example

To specify a number in binary format directly, you can use the prefix `0b`:

```go
var a uint8 = 0b0010_1101 // 45
```

Bitwise operations are often used when working with _masks_.
A [mask][mask-link] is a predefined number that can be used to modify incoming numbers.

_Example_: Let's say you need to extract the last 4 bits of a `uint16` number.
Here’s what the function might look like:

```go
LastFourBits(num uint16) uint16 {
    var mask uint16 = 0b0000_0000_0000_1111
    return mask & num
}
```

[binary-format-link]: https://en.wikipedia.org/wiki/Binary_number#Representation
[binary-arithmetic-link]: https://en.wikipedia.org/wiki/Binary_number#Binary_arithmetic
[conversion-link]: https://en.wikipedia.org/wiki/Binary_number#Conversion_to_and_from_other_numeral_systems
[mask-link]: https://en.wikipedia.org/wiki/Mask_(computing)
