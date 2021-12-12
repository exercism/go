# Design

## Learning objectives

- Know how to create an error variable
- Know how to create a custom error type
- Know how to return an error
- Know how to check for errors
- Learn about the "return early" best practice

## Out of Scope

The following topics will be introduced later and should therefore not be part of this concept exercise.

- `error-wrapping` (`errors.Is`, `errors.As`)
- `type-assertions`

## Concepts

The Concepts this exercise unlocks are:

- `interfaces`
- `errors`

## Prerequisites

- `conditionals-if` for the error checking
- `functions` for learning about multiple return values
- `structs` to be able to create a custom error type
- `methods` to understand interfaces
- `string-formatting` to implement the `Error() string` method
- `packages` to understand how to import `errors`

## Analyzer

This exercise could benefit from the following rules in the [analyzer][analyzer].

- Check that the student actually implemented a custom error type `SillyNephewError`
- Check that the `Error() string` method has a pointer receiver

[analyzer]: https://github.com/exercism/go-analyzer
