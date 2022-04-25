# Introduction

A function allows you to group code into a reusable unit.
It consists of the `func` keyword, the name of the function, and a comma-separated list of zero or more parameters and types in round brackets.

## Function Parameters

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

In the example above both `greeting` and `name` are parameters of type `string`.
When we pass primitive types such as `string` as arguments into a function, 
we essentially make a copy of the original value or underlying data so that only this copy is accessed or modified by our function.

In Go, not only can we pass primitive values into a function, but we can also pass a `pointer`, as in a reference to some data, to the function.
This means that the function receives only a copy to the pointer instead of the underlying data that this pointer refers to.

If the concept of `pointer` is confusing, no worries.
We have more details below as well as a dedicated section to help you understand pointers.

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

## Parameters vs. Arguments

Let's quickly cover two terms that are often confused together: `parameters` and `arguments`. 
Function parameters are the names defined in the function's signature, such as `greeting` and `name` in the function `PrintGreetingName` above.
Function arguments are the concrete values passed to the function parameters when we invoke the function.
For instance, in the example below, `"Hello"` and `"Katrina"` are the arguments passed to the `greeting` and `name` parameters:

```go
PrintGreetingName("Hello", "Katrina")
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

Functions in Go pass their arguments by value.
This means that the arguments are copied and any changes to the arguments passed to the function will stay inside the function.

```go
val := 2
func MultiplyByTwo(v int) int {
    v = v * 2
    return v
}
newval := MultiplyByTwo(val)
// newval is 4, val is still 2 because only a copy of its value was passed into the function
```

To modify the underlying data passed via an argument, we must pass pointer arguments into the function.
For now, it is sufficient to know that pointer types can be recognized by the `*` in front of the type in the function signature.

```go
func HandlePointers(x, y *int) {
    // Some logic to handle integer pointers x and y
}
```

## Exceptions

Note that `slices` and `maps` are exceptions to the above-mentioned rule. 
When we pass a `slice` or a `map` as arguments into a function, they are treated as pointer types even though there is no explicit `*` in the type.
This means that if we pass a slice or map into a function and modify its underlying data,
the changes will be reflected on the original slice or map.
