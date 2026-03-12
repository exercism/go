# Instructions append

The `diagram` argument starts each row with a `\n`.
This allows Go's raw string literals to present diagrams in source code nicely as two rows flush left.
For example, the test may contain the following.

```go
        diagram := `
VVCCGG
VVCCGG`
```

If the `children` argument is `nil`, use the list of children defined in the instructions above.
If it is not `nil`, use the given value.
