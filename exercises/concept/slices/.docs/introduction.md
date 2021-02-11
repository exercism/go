Slices in Go are similar to lists or arrays in other languages. They hold a number of elements of a specific type (or interface).

Slices in Go are based on arrays. Arrays have a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array.
Technically a slice consists of a pointer to a start element in an array and a length. This makes for example creating a sub-slice `s[3:7]` very efficient.
But don't worry: most of the time this knowledge is not needed in order to work with slices.

A slice is written like `[]T` with `T` being the type of the elements in the slice:

```go
var s []int
```
