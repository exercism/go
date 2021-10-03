# Introduction to Interfaces

An **interface type** is effectively a set of method signatures.
Any type that defines those methods "implements" the interface implicitly.
There is no `implements` keyword in Go.
Here is what an interface definition might look like:

```go
type InterfaceName interface {
    MethodOne() MethodOneReturnType
    MethodTwo(paramOne ParamOneType, paramTwo ParamTwoType) MethodTwoReturnType 
    ...
}
```

For example, here is the built-in `error` interface:

```go
type error interface {
    Error() string
}
```

This means that any type which implements an `Error()` method which returns a `string` implements the `error` interface.
This allows a function with return type `error` to return values of different types as long as all of them satisfy the `error` interface.

There is one very special interface type in Go: the **empty interface** type that contains zero methods.
The empty interface type is written like this: `interface{}`.
Since it has no methods, every type implements the empty interface type.
This is helpful for defining a function that can generically accept any value.
In that case, the function parameter uses the empty interface type.
