# Hints

## General

This [article][gobyexample] gives a nice introduction to package `math/rand`.

If you are working locally and use Go 1.19 or lower, make sure to seed your random number generator with the current time.

## 1. Roll a die

This [article][yourbasic] shows how to generate integers in a certain range.

## 2. Generate wand energy

Function [rand.Float64][float64] returns a random `float64` number between 0.0 and 1.0.

## 3. Shuffle a slice

Create a slice with the eight animal strings, then call [rand.Shuffle][shuffle] to put it into a random order.

[gobyexample]: https://gobyexample.com/random-numbers
[yourbasic]: https://yourbasic.org/golang/generate-number-random-range
[shuffle]: https://pkg.go.dev/math/rand#Rand.Shuffle
[float64]: https://pkg.go.dev/math/rand#Float64
