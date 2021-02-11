TODO: the content below is copied from the exercise introduction and probably needs rewriting to a proper concept introduction

## conditionals-if

## conditionals-switch

## booleans

Go supports the three logical operators `&&` (AND), `||` (OR), and `!` (NOT).

Conditionals in Go are similar to conditionals in other languages. The underlying type of any conditional operation is the `bool` type, which can have the value of `true` or `false`. Conditionals are often used as flow control mechanisms to check for various conditions. For checking a particular case an `if` statement can be used, which executes its code if the underlying condition is `true` like this:

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

Go also provides a `switch` statement for more advanced scenarios.

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
