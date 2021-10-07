# About

A floating-point number is a number with zero or more digits behind the decimal separator.
Examples are `-2.4`, `0.1`, `3.14`, `16.984025` and `1024.0`.

Different floating-point types can store different numbers of digits after the digit separator - this is referred to as its precision.

Go has two floating-point types:

- `float32`: 32 bits (~6-9 digits precision).
- `float64`: 64 bits (~15-17 digits precision).
  This is the default floating-point type.

As can be seen, both types can store a different number of digits.
This means that trying to store PI in a `float32` will only store the first 6 to 9 digits (with the last digit being rounded).

By default, Go will use `float64` for floating-point numbers, unless the floating-point number is:

1. assigned to a variable with type `float32`, or
2. returned from a function which return type is `float32`, or
3. passed as an argument to the `float32()` function.

Always be careful when checking the values of floating-point types for equality, as values that can appear to represent the same value could actually be different.

You can find a short introduction to floating-point numbers at [0.30000000000000004.com][0.30000000000000004.com]. The [Float Toy page][evanw.github.io-float-toy] has a nice, graphical explanation how a floating-point numbers' bits are converted to an actual floating-point value.

[0.30000000000000004.com]: https://0.30000000000000004.com/
[evanw.github.io-float-toy]: https://evanw.github.io/float-toy/
