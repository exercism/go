# Introduction

Instead of using an `if` statement for conditional logic, Go also provides a `switch` statement for more advanced scenarios.

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
