# Introduction

Go functions and methods can return multiple values:

```go
func GetCard() (Card, error) { ... }
```

Every return value must be assigned in order and separated by commas:

```go
card, err := GetCard()
```

Each position in the assignment declares a variable.
Go will not compile if any declared variable is never used.
The blank identifier `_` takes the place of a variable name to skip that return value:

```go
card, _ := GetCard()
```

`error` is a built-in interface.
A `nil` error means no error occurred; any non-nil value means something went wrong.
By convention, functions that can fail should return an `error` as their last value.
The caller must check it before using any other return values:

```go
card, err := GetCard()
if err != nil {
    return err // do not use card here
}
// safe to use card here
```
