# About

A [`switch` statement][switch_statement] evaluates an expression and runs the first `case` that matches its value.
It is an often cleaner alternative to long `if ... else if` chains.

## Basic Usage

```go
operatingSystem := "windows"

switch operatingSystem {
case "windows":
    // do something if the operating system is windows
case "linux":
    // do something if the operating system is linux
case "macos":
    // do something if the operating system is macos
default:
    // do something if the operating system is none of the above
}
```

Multiple values can be grouped into a single `case` with a comma.
The case matches if the expression equals any one of them:

```go
switch operatingSystem {
case "windows", "linux":
    // do something if the operating system is windows or linux
case "macos":
    // do something if the operating system is macos
default:
    // do something if the operating system is none of the above
}
```

## Expressionless Switch

The expression after `switch` can also be omitted entirely.
Each `case` then holds a boolean expression, and the first one that evaluates to `true` runs.

```go
age := 21

switch {
case age > 20 && age < 30:
    // do something if age is between 20 and 30
case age == 10:
    // do something if age is equal to 10
default:
    // do something else for every other case
}
```

## Initialization Statement

Like `if`, a `switch` can include a short initialization statement before the expression.
Variables declared there are scoped to the `switch` block:

```go
switch v := getValue(); {
case v < 0:
    // do something if v is less than 0
case v == 0:
    // do something if v is 0
default:
    // do something else for every other case
}
// v will not be accessible here
```

## Fallthrough

`fallthrough` causes the next case's body to run regardless of whether its expression would match.
It must be the last statement in the case's body, and there must be a next case to run.

`fallthrough` is uncommon in idiomatic Go because it makes adjacent cases dependent on each other.
As a result, it's harder to reason about the cases in isolation, making the switch harder to read.
When cases should have the same behavior, grouping them with a comma is clearer.
When logic is shared across cases for other reasons, consider extracting the shared logic into a helper function.

```go
s := "hello"

switch s {
case "hello":
    fmt.Println("hello")
    fallthrough
case "world":
    fmt.Println("world") // also runs, even though s != "world"
}
// Output:
// hello
// world
```

To learn more, check [Go by Example: Switch][go_by_example_switch] or the [Go language specification][switch_statement].

[switch_statement]: https://go.dev/ref/spec#Switch_statements
[go_by_example_switch]: https://gobyexample.com/switch
