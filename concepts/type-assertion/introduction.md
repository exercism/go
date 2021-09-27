# Introduction Type Assertions

Interfaces in Go can introduce ambiguity about the underlying type.
A type assertion allows us to extract the interface value's underlying concrete value using this syntax: `interfaceVariable.(concreteType)`.

For example:

```go
var input interface{} = 12
number := input.(int)
```

NOTE: this will cause a panic if the interface variable does not hold a value of the concrete type.

We can test whether an interface value holds a specific concrete type by returning two values: the underlying value and a boolean value that reports whether the assertion succeeded.
For example:

```go
str, ok := input.(string) // no panic if input is a string
```

If `input` holds a `string`, then `str` will be the underlying value and `ok` will be true.
If `input` does not hold a `string`, then `str` will be the zero value of type `string` (ie. `""` - the empty string) and `ok` will be false.
No panic occurs in any case.
