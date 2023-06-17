# Design

## Learning objectives

- Understand what an error is in Go and that error handling is not done via exceptions
- Know how and where to return an error from a function and how to signal "no error" and what to set for other return values
- Using `errors.New` to create a simple error
- Checking for errors and returning early
- Creating a custom error type

## Out of Scope

This exercise is deliberately kept simple (many aspects left for later exercises) to make sure students properly digest what errors are and the basic mechanics of how to work with them. This is very different in Go compared to many other languages so it is important is sticks. The exercise still has enough content to unlock 

The following topics will be introduced later and should therefore not be part of this concept exercise:

- checking for specific errors or error types
- error wrapping (`fmt.Errorf("...%w")`, `errors.Is` vs. `==`, `errors.As` vs. type assertion)

## Concepts

The Concepts this exercise unlocks are:

- `errors`

## Prerequisites

- `interfaces` to understand the error interface
- `conditionals-if` for the error checking
- `functions` for learning about multiple return values
- `structs` to be able to create a custom error type
- `methods` to understand interfaces
- `string-formatting` to implement the `Error() string` method
- `packages` to understand how to import `errors`
