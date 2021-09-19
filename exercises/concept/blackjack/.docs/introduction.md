# Introduction

Like other languages, Go also provides a `switch` statement. Switch statements are a shorter way to write long `if ... else if` statements. To make a switch, we start by using the keyword `switch` followed by a value or expression. We then declare each one of the conditions with the `case` keyword. We can also declare a `default` case, that will run when none of the previous `case` conditions matched:

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

One interesting thing about switch statements, is that the value after the `switch` keyword can be omitted, and we can have boolean conditions for each `case`:

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
