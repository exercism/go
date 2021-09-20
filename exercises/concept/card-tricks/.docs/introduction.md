# Introduction

Slices in Go are similar to lists or arrays in other languages. They hold a number of elements of a specific type (or interface).

Slices in Go are based on arrays. Arrays have a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array.

A slice is written like `[]T` with `T` being the type of the elements in the slice:

```go
var s []int
```
