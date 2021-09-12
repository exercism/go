# Basics About

The Basics concept introduced [Go](https://golang.org) as a statically typed, compiled programming language. 
The language is often referred to as Golang because of its domain name, golang.org, but the proper name is Go.

The Basics concept and corresponding exercises introduced three major language features: Packages, Functions, and Variables.

## Packages

Go applications are organized in packages.
A package is a collection of source files located in the same directory.
All source files in a directory must share the same package name.
It is conventional for the package name to be the last directory in the import path. For example, the files in the ["math/rand" package](https://golang.org/src/math/rand/) begin with the statement `package rand`.

```go
package lasagna
```

## Functions

Go functions accept zero or more parameters.
Parameters must be explicitly typed, there is no type inference.

Functions may have multiple (explicitly typed) return values.
Values are returned from functions using the `return` keyword. 

A function is invoked by specifying the function name and passing arguments for each of the function's parameters.
When a package is imported, only functions starting with a capital letter can
be used / accessed.

Note that Go supports two types of comments.
Single line comments are preceded by `//` and multiline comments are inserted between `/*` and `*/`.

```go
package greeting

// Hello is a public function
func Hello (name string) string {
    return hi(name)
}

// hi is a private function
func hi (name string) string {
    return "hi " + name
}
```

## Variables

Go is statically-typed, which means all variables [must have a defined type](https://en.wikipedia.org/wiki/Type_system) at compile-time.

Variables can be defined by explicitly specifying a type:

```go
var explicit int // Explicitly typed
```

You can also use an initializer, and the compiler will assign the variable type to match the type of the initializer.

```go
implicit := 10   // Implicitly typed as an int
```

Once declared, variables can be assigned values using the `=` operator.
Once declared, a variable's type can never change.

```go
count := 1 // Assign initial value
count = 2  // Update to new value

count = false // This throws a compiler error due to assigning a non `int` type
```
