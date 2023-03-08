# Introduction

## The error interface

Error handling is **not** done via exceptions in Go.
Instead, errors are normal _values_ of types that implement the built-in `error` interface.
The `error` interface is very minimal.
It contains only one method `Error()` that returns the error message as a string.

```go
type error interface {
  Error() string
}
```

Every time you define a function in which an error could happen during the execution that needs to reach the caller, you need to include `error` as one of the return types.
If the function has multiple return values, by convention `error` is always the last one.

```go
func DoSomething() (int, error) {
  // ...
}
```

## Creating and returning a simple error

You do not have to always implement the error interface yourself.
To create a simple error, you can use the `errors.New()` function that is part of the standard library package `errors`.
The only thing you need to pass in is the error message as a string.

If the function returns an error, it is good practice to return the zero value for all other return parameters:

```go
func DoSomething() (SomeStruct, int, error) {
  // ...
  return SomeStruct{}, 0, errors.New("failed to calculate result")
}
```

~~~~exercism/caution
You should not assume that all functions return zero values for other return values if an error is present.
It is best practice to assume that it is not safe to use any of the other return values if an error occurred.
The only exceptions are cases where the documentation clearly states that other returns values are meaningful in case of an error.
~~~~

If you want to use such a simple error in multiple places, you should declare a variable for the error instead of using `errors.New` in-line.
By convention, the name of the variable should start with `Err` or `err` (depending on whether it is exported or not).
These error variables are often called _sentinel errors_.

```go
import "errors"

var ErrNotFound = errors.New("resource was not found")

func DoSomething() error {
  // ...
  return ErrNotFound
}
```

Return `nil` for the error to signal that there were no errors during the function execution:

```go
func Foo() (int, error) {
  return 10, nil
}
```

## Error checking

If you call a function that returns an error, it is common to store the error value in a variable called `err`.
Before you use the actual result of the function, you need to check that there was no error.

To avoid nesting the "happy path" of your code, error cases should be handled first.
We can use `==` and `!=` to compare the error against `nil` and we know there was an error when `err` is not `nil`.

```go
func processUserFile() error {
	file, err := os.Open("./users.csv")
	if err != nil {
		return err
	}

	// do something with file
}
```

Most of the time, the error will be returned up the function stack as shown in the example above.
Another way of handling the error could be to log it and continue with some other operation.
It is good practice to either return or log the error, never both.

Since most functions in Go include an error as one of the return values, you will see/use the `if err != nil` pattern all over the place in Go code.

## Custom error types

If you want your error to include more information than just the error message string, you can create a custom error type.
As mentioned before, everything that implements the `error` interface (i.e. has an `Error() string` method) can serve as an error in Go.

Usually, a struct is used to create a custom error type.
By convention, custom error type names should end with `Error`.
Also, it is best to set up the `Error() string` method with a pointer receiver, see this [Stackoverflow comment][stackoverflow-errors] to learn about the reasoning.
Note that this means you need to return a pointer to your custom error otherwise it will not count as `error` because the non-pointer value does not provide the `Error() string` method.

```go
type MyCustomError struct {
  message string
  details string
}

func (e *MyCustomError) Error() string {
  return fmt.Sprintf("%s, details: %s", e.message, e.details)
}

func someFunction() error {
  // ...
  return &MyCustomError{
    message: "...",
    details: "...",
  }
}
```

[stackoverflow-errors]: https://stackoverflow.com/a/50333850
