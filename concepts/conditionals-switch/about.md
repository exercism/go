# About

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

To learn more about this topic it is recommended to check this source: [Go by Example: Switch][go_by_example_switch]

[switch_statement]: https://golang.org/ref/spec#Switch_statements
[go_by_example_switch]: https://gobyexample.com/switch
