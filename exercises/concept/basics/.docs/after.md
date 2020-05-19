## Packages

In Go an application is organized in packages. A package is a collection of source files located in the same folder. All source files in a folder must have the same package name at the top of the file. The package is preferred to be the same as the folder it is located in.

```go
package greeting
```

Standard library packages are imported using their name. Other packages are imported using the url of the package's location. A package can be given an alias when importing. This can be used to avoid package name conflicts:

```go
package greeting

import (
	"errors"

	errs "github.com/pkg/errors"
)
```

An imported package is then addressed with the package name or alias:

```go
// using the internal errors package
errors.New("Connection not established")

// using the errors package located at github.com/pkg/errors by alias
errs.New("Connection not established")
```

## Variables

Go is a statically-typed language, which means that everything has a type at compile-time. Assigning a value to a name is referred to as defining a variable. A variable can be defined either by explicitly specifying its type, or by assigning a value to have the Go compiler infer its type based on the assigned value.

```go
var explicitVar int // Explicitly typed
implicitVar := 10   // Implicitly typed
```

The value of a variable can be assigned using the [`:=` operator][assignment] and updated using the [`=` operator][assignment]. Once defined, a variable's type can never change.

```go
count := 1 // Assign initial value
count = 2  // Update to new value

// Compiler error when assigning different type
// count = false
```

### Integers

Integer values are defined as one or more (consecutive) digits and support the [default mathematical operators][operators].

## Functions

Functions in Go are considered [first class citizens][first-class-functions] making them very powerful. They can be used as:

- [public and private][public-vs-private] functions of a package
- attached to custom types also called a [method][methods] ([public and private][public-vs-private])
- assigned to variables
- passed into to functions as parameters and returned from functions (higher order functions)
- create custom types
- be used anonymously
- be used as closures

A function can have zero or more parameters. All parameters must be explicitly typed, there is no type inference for parameters. A function can also have multiple return values which must also be explicitly typed. Values are returned from functions using the [`return` keyword][return]. To allow a function, or a method to be [called by code in other packages][public-vs-private], the name of the function must start with a capital letter.

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

Invoking a function from inside the same package is done by specifying the function name and passing arguments for each of the function's parameters.

Invoking a function from another package is done by specifying the package name as well as the function name.

```go
phrase := greeting.Hello("John")
```

## Comments

Go supports two types of [comments][comments]. Single line comments are preceded by `//` and multiline comments are inserted between `/*` and `*/`.

[assignment]: https://golang.org/ref/spec#Assignments
[first-class-functinos]: https://golangbot.com/first-class-functions
[methods]: https://golang.org/ref/spec#Method_declarations
[return]: https://golang.org/ref/spec#Return_statements
[operators]: https://golang.org/ref/spec#Operators
[comments]: https://golang.org/ref/spec#Comments
[public-vs-private]: https://golang.org/ref/spec#Exported_identifiers
