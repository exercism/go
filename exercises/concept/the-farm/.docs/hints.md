# Hints

## 1. Get the amount of fodder from the `weightFodder` function

* read fodder weight and error from the `FodderAmount` method on the supplied `weightFodder`
* if there is an error, use `errors.Is(error, error) bool` to check if it is a `ScaleError`
* for any other error, return that back to the caller, with no computed value for division

## 2. Return an error for negative fodder

* use `errors.New(string)` to return a custom error for this case

## 3. Prevent division by zero

* use `errors.New(string)` to return a custom error for this case

## 4. Return a `SillyNephewError` for a negative number of cows

* return the `SillyNephewError` for this case with the number of cows
