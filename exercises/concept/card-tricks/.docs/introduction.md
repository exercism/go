# Introduction

## Slices

Slices in Go are similar to lists or arrays in other languages. They hold a number of elements of a specific type (or interface).

Slices in Go are based on arrays. Arrays have a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array.

A slice is written like `[]T` with `T` being the type of the elements in the slice:

```go
var empty []int                 // an empty slice
withData := []int{0,1,2,3,4,5}  // a slice pre-filled with some data
```

You can get/set an element at a given zero-based index using square-bracket notation:

```go
withData[1] = 5
x := withData[1] // x is now 5
```

You can create a new slice from an existing slice by getting a range of elements, once again using square-bracket notation, but specifying both a starting (inclusive) and ending (exclusive) index.
If you don't specify a starting index, it defaults to 0.
If you don't specify an ending index, it defaults to the length of the slice.

```go
newSlice := withData[2:4] // newSlice == []int{2,3}
```

## Multiple Return Values

Go functions and methods can return multiple values.
Very often, a second return value is used to return an error.
For example:

```go
func GetCard() (Card, error) { ... }
```

The assignment for multiple return values just uses a comma to separate the variables:

```go
card, err := GetCard()
```

If statements can use an initializer before the condition separated by a semicolon.
This is a common idiom seen for error handling:

```go
if card, err := GetCard(); err != nil {
    // handle the error
}
```
