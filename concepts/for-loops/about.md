# About

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
However, the braces `{ }` surrounding the loop body are always required.

This loop starts `i` at `1`, prints the current value of `i`, increases it for the next iteration, and stops when `i` reaches `10`:

```go
for i := 1; i < 10; i++ {
    fmt.Println(i)
}
```

## Other Loop Forms

### Condition-Only Loop

The init and post statements may be omitted.
A `for` statement with only a condition serves the same purpose as a `while` loop in other languages.
Go does not have a `while` keyword.

```go
sum := 1
for sum < 1000 {
    sum += sum
}
fmt.Println(sum)
// Output: 1024
```

### Infinite Loop

A `for` statement can be written without a condition:

```go
for {
    // loop body
}
```

This loop does not terminate on its own.
Execution must leave the loop explicitly, for example with `break` or `return`.

## Controlling Loops

### Break and Continue

In this loop, `break` exits the loop:

```go
for n := 0; n <= 5; n++ {
    if n == 3 {
        break
    }
    fmt.Println(n)
}
// Output:
// 0
// 1
// 2
```

`continue` skips the rest of the current loop body and begins the next iteration:

```go
for n := 0; n <= 5; n++ {
    if n % 2 == 0 {
        continue
    }
    fmt.Println(n)
}
// Output:
// 1
// 3
// 5
```

### Labeled Break and Continue

With nested loops, an unlabeled `break` or `continue` applies to the innermost loop.
Label an enclosing loop by placing an identifier followed by a colon on the line directly above it.
Then write the loop's label after `break` or `continue` to direct the action to it:

```go
OuterLoop:
for i := 0; i < 10; i++ {
    for j := 0; j < 10; j++ {
        if i + j == 10 {
            break OuterLoop
        }
    }
}
```
