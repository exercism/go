# Hints

## 1. Calculate the interest rate

- By default, any floating-point number defined in Go code is treated as a `float64`.
  As the return type of the `InterestRate` function is `float32`, you don't need to do any type conversions.

## 2. Calculate the interest

- When calculating interest, it might be helpful to notice that `InterestRate` returns a percentage.
- You can use the function defined in the previous task (`InterestRate`) to calculate the interest.
- You might need to convert the result `InterestRate` to a `float64`.

## 3. Calculate the annual balance update

- When calculating the annual balance update, use the functions that have been defined in the previous steps.

## 4. Calculate the years before reaching the desired balance

- To calculate the years, one can keep looping until the desired balance is reached using a [`for` loop](https://gobyexample.com/for).
- There is a special [operator](https://golang.org/ref/spec#IncDec_statements) to increment values by 1.
