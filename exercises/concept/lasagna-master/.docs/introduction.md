# Introduction

A function allows you to group code into a reusable unit.
It consists of the `func` keyword, the name of the function, and a comma-separated list of zero or more parameters and types in round brackets.
All parameters must be explicitly typed; there is no type inference for parameters.
There are no default values for parameters so all function parameters are required.

```go
import "fmt"

// No parameters
func PrintHello() {
    fmt.Println("Hello")
}

// Two parameters
func PrintGreetingName(greeting string, name string) {
  fmt.Println(greeting + " " + name)
}
```

Parameters of the same type can be declared together, followed by a single type declaration.

```go
import "fmt"

func PrintGreetingName(greeting, name string) {
  fmt.Println(greeting + " " + name)
}
```

## Return Values

The function arguments are followed by zero or more return values which must also be explicitly typed.
Single return values are left bare, multiple return values are wrapped in parenthesis.
Values are returned to the calling code from functions using the `return` keyword.
There can be multiple `return` statements in a function.
The execution of the function ends as soon as it hits one of those `return`s.
If multiple values are to be returned from a function, they are comma separated.

```go
func Hello(name string) string {
  return "Hello " + name
}

func HelloAndGoodbye(name string) (string, string) {
  return "Hello " + name, "Goodbye " + name
}
```

## Invoking Functions

Invoking a function is done by specifying the function name and passing arguments for each of the function's parameters in parenthesis.

```go
import "fmt"

// No parameters, no return value
func PrintHello() {
    fmt.Println("Hello")
}
// Called like this:
PrintHello()

// One parameter, one return value
func Hello(name string) string {
  return "Hello " + name
}
// Called like this:
greeting := Hello("Dave")

// Multiple parameters, multiple return values
func SumAndMultiply(a, b int) (int, int) {
    return a+b, a*b
}
// Called like this:
aplusb, atimesb := SumAndMultiply(a, b)
```

## Named Return Values and Naked Return

As well as parameters, return values can optionally be named.
If named return values are used, a `return` statement without arguments will return those values.
This is known as a 'naked' return.

```go
func SumAndMultiplyThenMinus(a, b, c int) (sum, mult int) {
    sum, mult = a+b, a*b
    sum -= c
    mult -= c
    return
}
```

## Pointers

Functions in go pass their arguments by value.
This means that the arguments are copied, and any changes to those arguments within the function will not be seen outside the function.

```go
val := 2
func MultiplyByTwo(v int) int {
    v = v * 2
    return v
}
newval := MultiplyByTwo(val)
// newval is 4, val is 2
```

To affect the values a function is called with, it is possible to pass pointer arguments, using the `*` syntax in the function, and the `&` syntax to generate a pointer.

```go
func swap(x, y *int) {
    var temp int
    temp = *x
    *x = *y
    *y = temp
}

var a int = 100
var b int = 200
swap(&a, &b)
// a = 200, b = 100
```
