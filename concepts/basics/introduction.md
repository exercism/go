# Basics Introduction

[Go](https://golang.org) is a statically typed, compiled programming language. Go is syntactically similar to C. The language is often referred to as Golang because of its domain name, golang.org, but the proper name is Go.

The Basics concept and corresponding [Gopher's Gorgeous Lasagna exercise](tracks/go/exercises/gophers-gorgeous-lasagna) introduce three major language features: Packages, Functions, and Variables.

## Packages

Go applications are organized in packages. A package is a collection of source files located in the same folder. It is conventional for the package name to be he last directory in the import path. For example, the files in the ["math/rand" package](https://golang.org/src/math/rand/) begin with the statement `package rand`.


```go
package lasagna
```




--------------------
C# is a statically-typed language, which means that everything has a type at compile-time. Assigning a value to a name is referred to as defining a variable. A variable can be defined either by explicitly specifying its type, or by letting the C# compiler infer its type based on the assigned value (known as _type inference_). Therefore, the following two variable definitions are equivalent:

```csharp
int explicitVar = 10; // Explicitly typed
var implicitVar = 10; // Implicitly typed
```

Updating a variable's value is done through the `=` operator. Once defined, a variable's type can never change.

```csharp
var count = 1; // Assign initial value
count = 2;     // Update to new value

// Compiler error when assigning different type
// count = false;
```

C# is an [object-oriented language][object-oriented-programming] and requires all functions to be defined in a _class_. The `class` keyword is used to define a class. Objects (or _instances_) are created by using the `new` keyword.

```csharp
class Calculator
{
    // ...
}

var calculator = new Calculator();
```

A function within a class is referred to as a _method_. Each method can have zero or more parameters. All parameters must be explicitly typed, there is no type inference for parameters. Similarly, the return type must also be made explicit. Values are returned from methods using the `return` keyword. To allow a method to be called by code in other files, the `public` access modifier must be added.

```csharp
class Calculator
{
    public int Add(int x, int y)
    {
        return x + y;
    }
}
```

Methods are invoked using dot (`.`) syntax on an instance, specifying the method name to call and passing arguments for each of the method's parameters. Arguments can optionally specify the corresponding parameter's name.

```csharp
var calculator = new Calculator();
var sum_v1 = calculator.Add(1, 2);
var sum_v2 = calculator.Add(x: 1, y: 2);
```

Scope in C# is defined between the `{` and `}` characters.

C# supports two types of comments. Single line comments are preceded by `//` and multiline comments are inserted between `/*` and `*/`.

[object-oriented-programming]: https://docs.microsoft.com/en-us/dotnet/csharp/programming-guide/concepts/object-oriented-programming

