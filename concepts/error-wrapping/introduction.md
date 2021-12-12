# Introduction

TODO

You can compare an error using the `errors.Is` function:

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