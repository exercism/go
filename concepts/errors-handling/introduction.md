# Introduction

In Go, errors are by convention the last value returned in a multiple returns function.

Always return the zero value for the type, if returning an error:

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

It is best practise to stop the execution when an error is encountered rather
than using `else`. The example below shows a bad practice:

```go
file, err := os.Open("./users.csv")
if err != nil {
  // do something with err
} else {
  // do something with file
}
```

## Scope of error

In Go, it is also a best practice to avoid declaring variables in a scope if
these are not required later.

For the sake of this example, we want to check for errors only:

```go
err := os.Rename(from, to)
if err != nil {
  return nil, err
}
// err remains in-scope here, even though it's not used or needed.
```

So this is the preferred way to consume this error:

```go
if err := os.Rename(from, to); err != nil {
  return nil, err
}
// err is not in-scope here
```
