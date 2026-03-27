# About

Conditionals are statements that execute different code depending on whether a condition is `true` or `false`.

An [`if` statement][if_statement] executes its body only when its condition is `true`.
Go has no concept of truthy or falsy values so the condition must be of type `bool`.
Expressions like `if mySlice { ... }` or `if 1 { ... }` are compile errors.

```go
func describe(val string) string {
    if val == "val" {
        return "was val"
    }
    return "was not val"
}
```

Use `else if` to check additional conditions, and `else` as a fallback if none match.
Conditions are evaluated top to bottom, and only the first branch whose condition is true will run.

```go
number := 0
result := "This number is "

if number > 0 {
    result += "positive"
} else if number < 0 {
    result += "negative"
} else {
    result += "zero"
}
```

Generally, Go favors returning early from the exceptional case and avoiding `else`.

```go
func getVal(connected bool) int {
    if !connected {
        return readLocalVal()
    }

    return readVal()
}
```

An `if` statement can also include a short initialization statement.
This is commonly used to check an error immediately after a call:

```go
num := 7
if v := 2 * num; v > 10 {
    fmt.Println(v)
} else {
    fmt.Println(num)
}
// Output: 14
```

Variables declared in the initialization statement are only in scope within the entire `if`/`else` chain.
Notably, Go has no ternary operator, a deliberate design choice to keep code easy to scan.

To learn more, see [Go by Example: If/Else][go_by_example_if].

[if_statement]: https://go.dev/ref/spec#If_statements
[go_by_example_if]: https://gobyexample.com/if-else
