# Introduction to Arithmetic Operators

## Operators

Go supports many standard arithmetic operators:

| Operator | Example        |
| -------- | -------------- |
| `+`      | `4 + 6 == 10`  |
| `-`      | `15 - 10 == 5` |
| `*`      | `2 * 3 == 6`   |
| `/`      | `13 / 3 == 4`  |
| `%`      | `13 % 3 == 1`  |

For integer division, the remainder is dropped (eg. `5 / 2 == 2`).

## Shorthand Assignments

These can be used in shorthand assignments to update and assign a variable using the operator:

```go
a := 1
a += 2 // same as a = a + 2 == 3
```

## Increment and Decrement

There are also two special statements: increment (`++`) and decrement (`--`).
They modify the value of a variable by increasing (or decreasing) the value by 1.
For example:

```go
a := 10
a++ // same as a += 1, a == 11
```
