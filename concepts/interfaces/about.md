# About Interfaces

An **interface** is effectively a set of method signatures.
Any value that defines those methods "implements" the interface type (implicitly).
Here is what an interface definition might look like:

```go
type InterfaceName interface {
    MethodOne() MethodOneReturnType
    ...
}
```

## The `error` interface

For example, here is the built-in `error` interface:

```go
type error interface {
    Error() string
}
```

This means that any type which implements an `Error()` method which returns a `string` implements the `error` interface.
For example:

```go
type MyCustomError struct {}

func (MyCustomError mce) Error() string {
    return "yikes!"
}
```

## The empty interface

There is one very special interface in Go: the **empty interface**.
The empty interface looks like this: `interface{}`.
Since it has no methods, every type implements the empty interface.
This means a method can generically take any value by accepting the empty interface as input.
For example:

```go
func AnythingGoes(i interface{}) string {
    if i == 0 {
        return "zero"
    }
    if i == "" {
        return "empty string"
    }
    return "something else"
}
```
