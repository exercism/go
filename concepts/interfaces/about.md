# About Interfaces

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

Since interfaces are implemented implicitly, a type can easily implement multiple different interfaces.
If all function parameters and return values needed to use concrete (ie. non-interface) types, that would strongly limit their re-usability.
By using interface types for the parameters, only the behavior that is needed for the function to do its job is defined.
Now the function can be used with parameters of different concrete types, as long as those concrete types implement the interface.
The same logic applies when using interface types as return values.

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

This allows a function with return type `error` to return values of different types as long as all of them satisfy the `error` interface.

## The empty interface

There is one very special interface type in Go: the **empty interface** type that contains zero methods.
The empty interface type is written like this: `interface{}`.
Since it has no methods, every type implements the empty interface type.
This is helpful for defining a function that can generically accept any value.
In that case, the function parameter uses the empty interface type.
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
