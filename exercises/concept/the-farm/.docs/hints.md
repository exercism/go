# Hints

You need to read fodder weight from the given method as parameter. This could
also return an error so you have to return that back to the caller, with no
computed value for division.

Since you can not divide fodder by zero cows, return an error preventing
a division by zero.

A negative number of cows, also indicates an error.

The error for no cows and for negative number of cows must be different.

## What is WeightFodder?

The `WeightFodder` type is a `func() (float64, error)`. The weight scale
sometimes could fail.

## How to check an error?

Try using `errors.Is(error, error) bool`.
