## General

- [Constants][const]

## 1. Represent the fixed interest rate on a loan

- The constant should be untyped, but needs to be acceptable anywhere that a float can be used. The function should return the constant.

## 2. Represent the number of days in a year

- The function should return the constant.

## 3. Represent the months of the year

- There is an [identifier][iota] for creating enumerated constants. A block looks like [this][block]. The function should take one argument and return it.

## 4. Represent an account number

- The constant should be untyped, but needs to be acceptable anywhere that a string can be used. The function should return the constant.

[const]: https://golang.org/doc/effective_go.html#constants
[iota]: https://golang.org/ref/spec#Iota
[block]: https://golang.org/pkg/debug/macho/#pkg-constants
