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

[first-class-functinos]: https://golangbot.com/first-class-functions
[methods]: https://golang.org/ref/spec#Method_declarations
[return]: https://golang.org/ref/spec#Return_statements
[public-vs-private]: https://golang.org/ref/spec#Exported_identifiers
