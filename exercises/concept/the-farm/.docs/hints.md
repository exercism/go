# Hints

## 1. Get the amount of fodder from the `weightFodder` function

- read fodder weight and error from the `FodderAmount` method on the supplied `weightFodder`
- if there is an error, use the equality operator `==` to check if it is `ErrScaleMalfunction`
- for any other error, return that back to the caller, with no computed value for division

## 2. Return an error for negative fodder

- use `errors.New(string)` to return a custom error for this case

## 3. Prevent division by zero

- use `errors.New(string)` to return a custom error for this case

## 4. Handle negative cows

- start by defining a `SillyNephewError` struct type
- add a field to the struct to hold the number of cows
- implement a method `Error() string` with a pointer receiver that returns the the correct text
- [string formatting][concept-string-formatting] can help with creating the error message
- if negative cows are supplied, return a pointer of a `SillyNephewError` error that contains the number of cows

[concept-string-formatting]: /tracks/go/concepts/string-formatting
