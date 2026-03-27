# Introduction

A `switch` statement evaluates an expression and runs the first `case` that matches its value.
It is an often cleaner alternative to long `if ... else if` chains.
For example,

```go
operatingSystem := "windows"

if operatingSystem == "windows" {
    // do something if the operating system is windows
} else if operatingSystem == "linux" {
    // do something if the operating system is linux
} else if operatingSystem == "macos" {
    // do something if the operating system is macos
} else {
    // do something if the operating system is none of the above
}
```

can be rewritten as:

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
That case will match if the expression equals any of the values:

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

The expression after `switch` can also be omitted entirely.
Each `case` then holds a boolean expression, and the first case being true is run.

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
