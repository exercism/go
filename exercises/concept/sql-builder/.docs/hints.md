# Hints

## 1. Generate SELECT Queries

- Remember that you don't need to initialize the builder with a constructor or literal. A simple `var builder strings.Builder` is immediately operational.
- To prevent the builder from dynamically resizing its internal buffer multiple times during execution, you can use the `Grow(n int)` method ahead of time if you can estimate or calculate the final size of the SQL query.
- Make sure to handle the case where the `columns` slice is empty or `nil` by appending the wildcard `*` symbol.

## 2. Generate Batch INSERT Statements

- Allocating a brand new `strings.Builder` inside a `for` loop defeats the performance benefits. Instead, declare a single builder *before* the loop starts.
- At the beginning of each loop iteration, call the `Reset()` method on your builder instance. This empties the buffer so you can start a fresh query while preserving the already allocated memory capacity under the hood.
- Don't forget that string values in the SQL `VALUES` clause need to be explicitly enclosed in single quotes `'`. You can use `WriteRune('\'')` to cleanly add them.
