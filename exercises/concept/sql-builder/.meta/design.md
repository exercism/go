# Design

## Goal

The goal of this exercise is to introduce the student to the `strings.Builder` type from the `strings` package and teach them how to perform efficient string assembly in Go without unnecessary memory allocations.

## Learning Objectives

By completing this exercise, the student should understand:
- Why strings in Go are immutable and why frequent concatenation using `+` or `+=` inside loops is detrimental to performance.
- How to utilize the zero-value of `strings.Builder`.
- How to append text using `WriteString` and individual Unicode characters using `WriteRune`.
- How to convert the accumulated buffer into a final string via `String()`.
- How to reuse an existing builder instance and preserve its allocated underlying memory slice across loop iterations using `Reset()`.
- The basics of capacity pre-allocation using `Grow(n)`.

## Out of Scope

The following concepts are explicitly out of scope for this introductory exercise:
- **`io.Writer` interface deep dive:** Although `strings.Builder` implements `io.Writer`, students do not need to work with generalized interfaces or bytes tracking here.
- **Handling write errors:** The write methods of `strings.Builder` return an error value to satisfy interfaces, but it is guaranteed to always be `nil`. Checking for these errors is out of scope.
- **Concurrency & Thread Safety:** Explaining why `strings.Builder` is not thread-safe (unlike `bytes.Buffer`) is left out to keep the focus on simple, synchronous performance tuning.

## Concepts

The exercise teaches the following concepts:
- `strings-builder`

## Prerequisites

To successfully complete this exercise, the student should already be familiar with:
- `strings` (Basic usage)
- `slices` (Iterating over slices)
- `structs` (Accessing struct fields)
- `loops` (For-range loops)
