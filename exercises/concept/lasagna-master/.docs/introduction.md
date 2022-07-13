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

## Parameters vs. Arguments

Let's quickly cover two terms that are often confused together: `parameters` and `arguments`. 
Function parameters are the names defined in the function's signature, such as `greeting` and `name` in the function `PrintGreetingName` above.
Function arguments are the concrete values passed to the function parameters when we invoke the function.
For instance, in the example below, `"Hello"` and `"Katrina"` are the arguments passed to the `greeting` and `name` parameters:

```go
PrintGreetingName("Hello", "Katrina")
```

## Return Values

The function parameters are followed by zero or more return values which must also be explicitly typed.
Single return values are left bare, multiple return values are wrapped in parenthesis.
Values are returned to the calling code from functions using the `return` keyword.
There can be multiple `return` statements in a function.
The execution of the function ends as soon as it hits one of those `return` statements.
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

## Pass by Value vs. Pass by Reference

It is also important to clarify the concept of passing by value and passing by reference.

First, let's clarify passing by value.
In the example below, when we pass the variable `val` to the function `MultiplyByTwo`, we passed a copy of `val`.
Because of this, `newVal` has the updated value `4` but the original variable `val` is still `2`.
Behind the scene, Go essentially makes a copy of the original value so that only this copy, a.k.a. `v`, is modified by the function.

```go
val := 2
func MultiplyByTwo(v int) int {
    v = v * 2
    return v
}
newVal := MultiplyByTwo(val)
// newval is 4, val is still 2 because only a copy of its value was passed into the function
```

Strictly speaking, all arguments are passed by value in Go, i.e. a copy is made of the value or data provided to the function.
But if you don't want to make a copy of the data that is passed to a function and want to change the data in the function,
then you should use `pointers` as arguments, a.k.a. pass by reference.

## Pointers

We use a `pointer` to achieve passing by reference.
By passing `pointer` arguments into a function,
we could modify the underlying data passed into the function instead of only operating on a copy of the data.

For now, it is sufficient to know that pointer types can be recognized by the `*` in front of the type in the function signature.

```go
func HandlePointers(x, y *int) {
    // Some logic to handle integer values referenced by pointers x and y
}
```

If the concept of `pointer` is confusing, no worries.
We have a dedicated section later in the syllabus to help you master pointers.

## Exceptions

Note that `slices` and `maps` are exceptions to the above-mentioned rule. 
When we pass a `slice` or a `map` as arguments into a function, they are treated as pointer types even though there is no explicit `*` in the type.
This means that if we pass a slice or map into a function and modify its underlying data,
the changes will be reflected on the original slice or map.
