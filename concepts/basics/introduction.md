Go is a statically typed, compiled programming language. Go is syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style concurrency. The language is often referred to as Golang because of its domain name, golang.org, but the proper name is Go.



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



=======================================






## Packages

## Functions

## Variables

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
