# Introduction

In Go, errors are by convention the last value returned in a multiple returns function.

Always return the default value for the type, if returning an error:

```go
func Foo() (int, error) {
  // Ok!
  return 0, errors.New("Error")

  // Don't
  return 10, errors.New("Error")
}
```

When no errors have to be returned, return `nil`:

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

You can also compare an error using the built-in `errors.Is()`

```go
ErrCustom := errors.New("custom error")
// ..
errors.Is(err, ErrCustom)
// true || false
```

It is best practice to stop the execution when an error is encountered rather
than using `else`. The example below shows a bad practice:

```go
file, err := os.Open("./users.csv")
if err != nil {
  // do something wit err
} else {
  // do something with file
}
```

## Scope of error

In Go, it is also a best practice to avoid declaring variables in a scope if
these are not required later.

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
