# Design

## Goal

The goal of this exercise is to teach the student the basics of programming in Go.

## Learning objectives

- Know what a package is.
- Know about the package declaration.
- Know what a variable is.
- Know how to define a variable.
- Know how to update a variable.
- Know how to use type inference for variables.
- Know how to define a function.
- Know how to return a value from a function.
- Know how to call a function.
- Know how to export a function.
- Know how to define an integer.
- Know how to use mathematical operators on integers.
- Know how to define single- and multiline comments.

## Out of scope

- Naming rules for identifiers.
- Methods and other function usages.
- Memory and performance characteristics.
- Variadic parameters.
- Named return values.

## Concepts

The Concepts this exercise unlocks are:

- `Basics`: Know what a package is. Know about the package declaration. Know how to import other packages. Know what a variable is. Know how to define a variable. Know how to update a variable. Know how to use type inference for variables. Know how to define a function. Know how to return a value from a function. Know how to call a function. Know how to export a function. Know how to define an integer. Know how to use mathematical operators on integers. Know how to define single and multiline comments.

## Prerequisites

There are no prerequisites.

## Representer

This exercise does not require any specific representation logic to be added to the [representer][representer].

## Analyzer

This exercise could benefit from the following rules added to the [analyzer][analyzer]:

- Verify that the `RemainingOvenTime()` function calls the `OvenTime()` function.
- Verify that the `ElapsedTime()` function calls the `PreparationTime()` function.

[analyzer]: https://github.com/exercism/go-analyzer
[representer]: https://github.com/exercism/go-representer
