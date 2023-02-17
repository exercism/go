# About

Error handling is **not** done via exceptions in Go.
Instead, errors are normal values of the interface type `error`.

```go
type error interface {
    Error() string
}
```

This means that any type which implements an one simple method `Error()` that returns a `string` implements the `error` interface.
This allows a function with return type `error` to return values of different types as long as all of them satisfy the `error` interface.

## Creating and Returning Errors

You do not have to always implement the error interface yourself.
To create a simple error, you can use the `errors.New()` function that is part of the standard library package `errors`.
These error variables are also called _sentinel errors_ and by convention their names should start with `Err` or `err` (depending on whether they are exported or not).
You should use error variables instead of directly writing `errors.New` in cases where you use an error multiple times or where you want the consumer of your package to be able to check for the specific error.

```go
import "errors"

var ErrSomethingWrong = errors.New("something went wrong")
ErrSomethingWrong.Error()
// => "something went wrong"
```

An error is by convention the last value returned in a function with multiple return values.
If the function returns an error, it is good practice to return the zero value for all other return parameters:

```go
import "errors"

// Do this:
func GoodFoo() (int, error) {
  return 0, errors.New("something went wrong")
}

// Not this:
func BadFoo() (int, error) {
  return 10, errors.New("something went wrong")
}
```

~~~~exercism/caution
You should not assume that all functions return zero values if an error is present.
It is best practice to assume that it is not safe to use any of the other return values if an error is returned.

As an example, imagine a function was trying to read from a file and returns a half-filled byte slice and an error.
If you would re-use that returned byte slice in your code, assuming it is always empty because there was error, you can run into subtle bugs.

The only exceptions are cases where the function documentation clearly states that other returns values are meaningful in case of an error.
Look at the [documentation for `Write` on a buffer][buffer-write] for an example.
~~~~

Return `nil` as the error value to indicate that there was no error, i.e. the function did its job successfully:

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


## Wrapping an Error

By default, an error in Go does not contain a stack trace.
In many cases, you can picture an error in Go more like a glorified string.
If the same error (same message string) can happen in different places in your application and you log it further up in your function chain, it can be really hard to figure out what went wrong just from the log entry.

To make it easier to understand the context in which an error happened, you can wrap it.
In its simplest form, this can be achieved by using `fmt.Errorf` with the `%w` formatting verb to add some more text to the error.

```go
func UserByID(id int) error {
  err := db.User(id)
  if err != nil {
    return fmt.Errorf("retrieving user from database: %w", err)
  }
  // ...
}

// Let's assume db.User returns errors.New("entry not found").
err := UserByID()
err.Error()
// => "retrieving user from database: entry not found"
```




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

[stackoverflow-errors]: https://stackoverflow.com/a/50333850
[line-of-sight]: https://medium.com/@matryer/line-of-sight-in-code-186dd7cdea88
[buffer-write]: https://pkg.go.dev/bytes#Buffer.Write

