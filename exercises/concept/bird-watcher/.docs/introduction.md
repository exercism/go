# Introduction

## For Loops

A `for` statement repeats a block of code.
Its first line controls how often the block repeats:

```go
for init; condition; post {
    // loop body
}
```

- The **init** statement runs once before the first iteration. It's common to declare and initialize a counter variable used by the loop.
- The **condition** is a Boolean expression evaluated before each iteration. The loop stops when this evaluates to `false`.
- The **post** statement runs after each completed iteration. It can be used to update a counter.

Do not surround `init; condition; post` with parentheses.
However, the braces surrounding the loop body are always required.

This loop starts `i` at `1`, prints the current value of `i`, increases it for the next iteration, and stops when `i` reaches `10`:

```go
for i := 1; i < 10; i++ {
    fmt.Println(i)
}
```
