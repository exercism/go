# About

Error handling is **not** done via exceptions in Go.
Instead, errors are normal values of the interface type `error`.

## Creating and Returning Errors

You do not have to always implement the error interface yourself.
To create a simple error, you can use the `errors.New()` function that is part of the standard library package `errors`.
These error variables are also called _sentinel errors_ and by convention their names should start with `Err` or `err` (depending on whether they are exported or not).
You should use error variables instead of directly writing `errors.New` in cases where you use an error multiple times or where you want the consumer of your package to be able to check for the specific error.

```go
import "errors"

var ErrSomethingWrong = errors.New("something went wrong")
ErrSomethingWrong.Error() // returns "something went wrong"
```

An error is by convention the last value returned in a function with multiple return values.
If the function returns an error, it should always return the zero value for other returned values:

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

Return `nil` for the error when there are no errors:

```go
func Foo() (int, error) {
  return 10, nil
}
```

## Custom Error Types

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
  return fmt.Sprintf("%s, Details: %s", e.message, e.details)
}

func someFunction() error {
  // ...
  return &MyCustomError{
    message: "...",
    details: "...",
  }
}
```

## Checking for Errors

Errors can be checked against `nil`:

```go
func myFunc() error {
	file, err := os.Open("./users.csv")
	if err != nil {
		// handle the error
		return err // or e.g. log it and continue
	}
	// do something with file
}
```

Since most functions in Go include an error as one of the return values, you will see/use this `if err != nil` pattern all over the place in Go code.

You can compare error variables with the equality operator `==`:

```go
var ErrResourceNotFound = errors.New("resource not found")
// ...
if err == ErrResourceNotFound {
  // do something about the resource-not-found error
}
```

How to check for errors of a specific custom error type will be covered in later concepts.

## Best Practice: Return Early

It is best practice to return early when an error is encountered.
That avoids nesting the "happy path" code, see this [article by Mat Ryer][line-of-sight] for more details.

```go
// Do this:
func myFunc() error {
  file, err := os.Open("./users.csv")
  if err != nil {
    return err
  }
  // do something with file
}

// Not this:
func myFunc() error {
  file, err := os.Open("./users.csv")
  if err != nil {
    // handle err
  } else {
    // do something with file
  }
}

// Also not this:
func myFunc() error {
  file, err := os.Open("./users.csv")
  if err == nil {
    // do something with file
  }
  // handle err
}
```

## Scope of Errors

In Go, it is also a best practice to avoid declaring variables in a scope if these are not required later.

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

[stackoverflow-errors]: https://stackoverflow.com/a/50333850
[line-of-sight]: https://medium.com/@matryer/line-of-sight-in-code-186dd7cdea88

