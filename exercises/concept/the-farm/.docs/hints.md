# Hints

## 1. Divide the food evenly

- Start by writing the function signature of `DivideFood`.
  It should accept 2 parameters of type `FodderCalculator` and `int` and return two values of types `float64` and `error`.
  Revisit the [functions concept][concept-functions] if you need more information on how to define functions.
- In the function body, call the `FodderAmount` [method][concept-methods] on `FodderCalculator` to fetch the default total amount of fodder for the cows.
  It will return the actual result and an error.
  Handle the error via an if-statement as it was explained in the introduction.
- After that, call the `FatteningFactor` method and handle the error return value as before.
- Now that you have the fodder amount and the factor, you can calculate the final result.
  You need to divide the fodder by the number of cows (revisit [numbers] for hints on type conversion) and multiply with the factor. Check the introduction for what to return as the error value in case of success.

## 2. Check the number of cows

- `ValidateInputAndDivideFood` has the same function signature as `DivideFood`.
- Since you want to return early in case of an error in Go, you first check whether the number of cows is less or equal than 0 with an if-statement.
- If it is, you return an error that you created with `errors.New`.
  Make sure the message matches the instructions.
- If the number of cows is valid, you can proceed to call the existing `DivideFood` function from task 1.

## 3. Improve the error handling

- Start by creating the `InvalidCowsError` [struct][concept-structs] with two unexported fields that hold the number of cows and the message.
- Next, define the `Error` method on that struct (with a pointer receiver). Revisit the exercise introduction for help on how to do this.
- Now you can work on the `ValidateNumberOfCows` function.
  Depending on the number of cows ([if-statement][concept-conditionals]), it should create and return a new instance of the `InvalidCowsError` and set the correct message while doing so.
  If the number of cows was valid, `nil` should be returned.

[concept-methods]: /tracks/go/concepts/methods
[concept-functions]: /tracks/go/concepts/functions
[concept-numbers]: /tracks/go/concepts/numbers
[concept-structs]: /tracks/go/concepts/structs
[concept-conditionals]:  /tracks/go/concepts/conditionals-if