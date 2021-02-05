TODO: the content below is copied from the exercise introduction and probably needs rewriting to a proper concept introduction

## packages

## functions

## variables

In Go an application is organized in packages. A package is a collection of source files located in the same folder. All source files in a folder must have the same package name at the top of the file. The package name is preferred to be the same as the folder it is located in.

```go
package greeting
```

Go is a statically-typed language, which means that everything has a type at compile-time. Assigning a value to a name is referred to as defining a variable. A variable can be defined either by explicitly specifying its type, or by assigning a value to have the Go compiler infer its type based on the assigned value.

```go
var explicit int // Explicitly typed
implicit := 10   // Implicitly typed
```

The value of a variable can be assigned using the `:=` and updated using the `=`. Once defined, a variable's type can never change.

```go
count := 1 // Assign initial value
count = 2  // Update to new value

// Compiler error when assigning different type
// count = false
```

A function can have zero or more parameters. All parameters must be explicitly typed, there is no type inference for parameters. A function can also have multiple return values which must also be explicitly typed. Values are returned from functions using the `return` keyword. To allow a function to be called by code in other packages, the name of the function must start with a capital letter.

```go
package greeting

// Hello is a public function
func Hello(name string) string {
    return hello(name)
}

// hello is a private function
func hello(name string) string {
    return "Hello " + name
}
```

Invoking a function is done by specifying the function name and passing arguments for each of the function's parameters.

Go supports two types of comments. Single line comments are preceded by `//` and multiline comments are inserted between `/*` and `*/`.
