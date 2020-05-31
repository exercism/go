## Logical Operators

Go supports the three [logical operators][logical_operators] `&&` (AND), `||` (OR), and `!` (NOT).

## If statement

Conditionals in Go are similar to conditionals in other languages. The underlying type of any conditional operation is the `bool` type, which can have the value of `true` or `false`. Conditionals are often used as flow control mechanisms to check for various conditions. For checking a particular case an [`if` statement][if_statement] can be used, which executes its code if the underlying condition is `true` like this:

```go
var value string

if value == "val" {
    // conditional code
}
```

In scenarios involving more than one case many `if` statements can be chained together using the `else if` and `else` statements.

```go
if value == "val" {
    // conditional code
} else if value == "val2" {
    // conditional code
} else {
    // default code
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

## Switch statement

Go also provides a [`switch` statement][switch_statement] for more advanced scenarios. It can be used to switch on a variable's content or as a replacement for `if ... else if` statements. There is a third application for `switch` statements which will be introduced later: the `type switch`. A switch statement can have a `default` case which is executed if no other case applies.

If there are three or more cases in a single `if` (e.g. `if ... else if ... else`), it should be replaced by a `switch` statement. A `switch` with a single case should be replaced by an `if` statement.

```go
var value string

// switch statement on variable content
switch value {
case "val1":
    // conditional code
case "val2", "val3", "foo":
    // conditional code
default:
    // default code
}

// switch statement using conditions (can replace `if ... else if` statements)
switch {
case value == "val1":
    // conditional code
case strings.HasPrefix(value, "val") || strings.HasPrefix(value, "foo"):
    // conditional code
}
```

Coming from other languages one may be tempted to try to use one-line conditionals. Go does not support ternary operators or one-line conditionals. This is a purposeful design decision to limit the amount of possibilities, making code simpler and easier to read.

To learn more about this topic it is recommended to check these sources:
[Go by Example: If/Else][go_by_example_if] and [Go by Example: Switch][go_by_example_switch].

[logical_operators]: https://golang.org/ref/spec#Logical_operators
[if_statement]: https://golang.org/ref/spec#If_statements
[switch_statement]: https://golang.org/ref/spec#Switch_statements
[go_by_example_if]: https://gobyexample.com/if-else
[go_by_example_switch]: https://gobyexample.com/switch
