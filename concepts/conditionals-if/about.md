# About

Conditionals in Go are similar to conditionals in other languages. The underlying type of any conditional operation is the `bool` type, which can have the value of `true` or `false`. Conditionals are often used as flow control mechanisms to check for various conditions. For checking a particular case an [`if` statement][if_statement] can be used, which executes its code if the underlying condition is `true` like this:

```go
var value string

if value == "val" {
    return "was val"
}
```

In scenarios involving more than one case many `if` statements can be chained together using the `else if` and `else` statements.

```go
var number int
result := "This number is "

if number > 0 {
    result += "positive"
} else if number < 0 {
    result += "negative"
} else {
    result += "zero"
}
```

However, it is convention to avoid `else` statements as Go promotes early returns:

```go
function getVal(connected bool) int {
    // The exceptional case should be in the `if` statemeent.
    // In this case being `connected` is the default, `readLocalVal` the fallback.
    if !connected {
        // using an `early return` to remove the need for an `else` case
        return readLocalVal()
    }

    return readVal()
}
```

Coming from other languages one may be tempted to try to use one-line conditionals. Go does not support ternary operators or one-line conditionals. This is a purposeful design decision to limit the amount of possibilities, making code simpler and easier to read.

To learn more about this topic it is recommended to check this source: [Go by Example: If/Else][go_by_example_if]

[if_statement]: https://golang.org/ref/spec#If_statements
[go_by_example_if]: https://gobyexample.com/if-else
