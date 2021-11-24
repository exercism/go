# Hints

## General

- [Basic types][basic types] and [Type conversions][type conversions] tutorials.

## 1. Calculate the number of working cars produced per hour

- To calculate, divide the percentage by 100 and multiply the result by the number of cars produced.
- When multiple two numbers together, they both need to be of the same type.

## 2. Calculate the number of working cars produced per minute

- Use the `CalculateProductionRatePerHour` function.
- Remember, there are 60 minutes in an hour.
- Remember to cast the result to an `int`.

## 3. Calculate the cost of production 

- Start by working out how many complete groups of 10 there are.
- Then work out how many cars are remaining (there is an [operator][modulo operator] for this).
- From the remaining, work out how many complete groups of 3 there are and then how many are left ungrouped.

[basic types]: https://tour.golang.org/basics/11
[type conversions]: https://tour.golang.org/basics/13
[modulo operator]: https://golangbyexample.com/remainder-modulus-go-golang/
