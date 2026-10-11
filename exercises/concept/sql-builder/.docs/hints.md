# Hints

## 1. Generate SELECT Queries

- Remember that you don't need to initialize the builder with a constructor or literal.
  A simple `var sb strings.Builder` is immediately operational.
- Make sure to handle the case where the `columns` slice is empty or `nil` by appending the wildcard `*` symbol.
- Note that you are not required nor expected to handle any meta characters or SQL sanitization here. You can safely append the column names directly using `sb.WriteString(col)`.

## 2. Generate Batch INSERT Statements

- Allocating a brand new `strings.Builder` inside a `for` loop defeats the performance benefits.
  Instead, declare a single builder *before* the loop starts.
- At the beginning of each loop iteration, call the `Reset()` method on your builder instance.
  This empties the buffer so you can start a fresh query while preserving the already allocated memory capacity under the hood.
- Don't forget that string values in the SQL `VALUES` clause need to be explicitly enclosed in single quotes `'`.
  You can use `WriteByte('\'')` to cleanly add them.
- Just like with the SELECT queries, do not worry about escaping special characters or preventing SQL injections within the column values. Treat the input strings as safe to use directly.
