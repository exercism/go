Booleans in Go are represented by the predeclared boolean type `bool`, which values can be either `true` or `false`.
It's a defined type.

```go
var a bool
```

Go supports three logical operators that can evaluate expressions down to Boolean values, returning either `true` or `false`.

| Operator    | What it means                                 |
| ----------- | --------------------------------------------- |
| `&&` (and)  | It is true if both statements are true.       |
| `\|\|` (or) | It is true if at least one statement is true. |
| `!` (not)   | It is true only if the statement is false.    |
