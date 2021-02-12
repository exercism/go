# About

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

TODO: decide what to do with this section

Integer values are defined as one or more (consecutive) digits and support the [default mathematical operators][operators].

[assignment]: https://golang.org/ref/spec#Assignments
[operators]: https://golang.org/ref/spec#Operators
