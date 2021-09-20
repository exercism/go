# About

Like other languages, Go also provides a [`switch` statement][switch_statement]. Switch statements are a shorter way to write long `if ... else if` statements.

## Basic usage

To make a switch, we start by using the keyword `switch` followed by a value or expression. We then declare each one of the conditions with the `case` keyword. We can also declare a `default` case, that will run when none of the previous `case` conditions matched:

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

If we want to run the same piece of code for several cases, we can group them together in a single `case`, separating them with a `,`:

```go
operatingSystem := "windows"

switch operatingSystem {
case "windows", "linux":
    // do something if the operating system is windows or linux
case "macos":
    // do something if the operating system is macos
default:
    // do something if the operating system is none of the above
} 
```

## Cases with boolean expressions 

One interesting thing about switch statements, is that the value after the `switch` keyword can be omitted, and we can have boolean conditions for each `case`. This effectively can be a shorter way to write complex `if ... else` statements:

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

## Fallthrough

When the condition in a `case` matches, the corresponding code will run and Go will not evaluate the other `case` conditions by default. We can make use of the `fallthrough` keyword to tell Go to evaluate the other `case` conditions. Take this example:

```go
age := 21

switch {
case age > 20:
    // do something if age is greater than 20
case age > 30:
    // WARNING: code here will never run. If age is greater than 30,
    // it is also greater than 20, which means only the previous case will run
default:
    // do something else for every other case
}
```

We can correct this code by using the `fallthrough` keyword:

```go
age := 21

switch {
case age > 20:
    // do something if age is greater than 20
    fallthrough
case age > 30:
    // Since the previous case uses 'fallthrough',
    // this code will now run if age is also greater than 30
default:
    // do something else for every other case
}
```

To learn more about this topic, check [Go by Example: Switch][go_by_example_switch] or [Tour of Go: Switch][switch_statement].

[switch_statement]: https://tour.golang.org/flowcontrol/9
[go_by_example_switch]: https://gobyexample.com/switch
