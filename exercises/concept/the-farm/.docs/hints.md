# Hints

## 1. Get the amount of fodder from the `weightFodder` function

* read fodder weight and error from the supplied `weightFodder` function
* use `errors.Is(error, error) bool` to check if it is an `ErrWeight` error
* for any other error, return that back to the caller, with no computed value for division

## 2. Return an error for negative fodder

* use `errors.New(string)` to return a custom error for this case

## 3. Prevent division by zero

* use `errors.New(string)` to return a custom error for this case

## 4. Return a `SillyNephew` error for a negative number of cows

* return the custom `SillyNephew` error for this case
