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

// One parameter
func PrintHelloName(name string) {
  fmt.Println("Hello " + name)
}
```

Parameters of the same type can be declared together, followed by a single type declaration.

```go
import "fmt"

func PrintGreetingName(greeting, name string) {
  fmt.Println(greeting + " " + name)
}
```

## Parameters vs. Arguments

Let's quickly cover two terms that are often confused together: `parameters` and `arguments`. 
Function parameters are the names defined in the function's signature, such as `greeting` and `name` in the function `PrintGreetingName` above.
Function arguments are the concrete values passed to the function parameters when we invoke the function.
For instance, in the example below, `"Hello"` and `"Katrina"` are the arguments passed to the `greeting` and `name` parameters:

```go
PrintGreetingName("Hello", "Katrina")
```

## Pass by Value & Pass by Pointer

It is also important to clarify the concept of passing by value as well as passing by pointer (sometimes referred to as _by reference_).
In the example below, `name`, the `string` parameter to the function `PrintHelloName`, is passed into the function by value.
This means that Go essentially makes a copy of the original value so that only this copy is accessed or modified by our function.

```go
import "fmt"

func PrintHelloName(name string) {
  fmt.Println("Hello " + name)
}
```

Strictly speaking, all arguments are passed by value in Go, i.e. a copy is made of the value or data provided to the function.
But if you don't want to make a copy of the data that you choose to pass to a function and
want to actually access and change the data in the function,
then you should use pointers as arguments, a.k.a. pass by reference.

In Go, not only can we pass primitive values into a function, but we can also pass a `pointer` of these values
(a reference to the data) to the function.
This means that the function receives only a copy to the pointer instead of the underlying data that this pointer refers to.
Note that we say _a copy to the pointer_ since everything is passed by value to functions in Go,
but a pointer copy, in contrast to a value copy, gives us an "address" that points to the data we want to access and modify.

If the concept of `pointer` is confusing, no worries.
We have a dedicated section and exercise to help you understand pointers.

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

