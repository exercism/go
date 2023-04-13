# Design

/*
- define a function that includes an error
(- create inline vs variable)
- return zero values for other things
- check for an error
- custom error type
*/

## Learning objectives

- Using `errors.New` to create a simple error
- Know how and where to return an error, how to signal "no error" and what to set for other return values
- Checking for errors and returning early
- Creating a custom error type

## Out of Scope

The following topics will be introduced later and should therefore not be part of this concept exercise.

- `type-assertions`
- `error-wrapping` (`fmt.Errorf("...%w")`, `errors.Is` vs. `==`, `errors.As` vs. type assertion)

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

## Analyzer

This exercise could benefit from the following rules in the [analyzer][analyzer].

- Check that the student actually implemented a custom error type
- Check that the `Error() string` method has a pointer receiver

[analyzer]: https://github.com/exercism/go-analyzer
