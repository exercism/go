# Introduction

## Interfaces

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

## Errors

In Go, errors are by convention the last value returned in a function with multiple return values.
If the function is returning an error, it should always return the zero value for other returned values:

```go
import "errors"

// Do this:
func GoodFoo() (int, error) {
  return 0, errors.New("Error")
}

// Not this:
func BadFoo() (int, error) {
  return 10, errors.New("Error")
}
```

> NOTE: the standard library has an `errors` package can be used to create and check errors.

Return `nil` for the error when there are no errors:

```go
func Foo() (int, error) {
  return 10, nil
}
```

Errors can be checked against `nil`:

```go
file, err := os.Open("./users.csv")
if err != nil {
    // do something with err
    return // or continue
}
// do something with file
```

You can also compare an error using the `errors.Is` function:

```go
var ErrCustom = errors.New("custom error")
// ..
if errors.Is(someErr, ErrCustom) {
  // do something about the custom error
}
```

And you can convert an error to a particular error type using the `errors.As` function:

```go
var someErr *SomeError
// returns true if err is a SomeError
if errors.As(err, &someErr) {
  // someErr now holds the value
}
```

It is best practice to stop the execution when an error is encountered rather than using `else`.
The example below shows a bad practice:

```go
file, err := os.Open("./users.csv")
if err != nil {
  // do something with err
} else {
  // do something with file
}
```

## Scope of error

In Go, it is also a best practice to avoid declaring variables in a scope if they are not required later.

For the sake of this example, we want to check for errors only:

```go
_, err := os.Getwd()
if err != nil {
  return nil, err
}
// at this point, err is still visible.
```

So this is the preferred way to consume this error:

```go
if _, err := os.Getwd(); err != nil {
  return
}
// err is now out of this scope.
```
