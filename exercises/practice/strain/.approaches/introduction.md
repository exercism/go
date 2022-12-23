# Introduction

There is at least one idomatic approach to solving Strain in Go.
You can use [generics][generics], which were introduced in version 1.18.

```exercism/note
As of this writing, the Go version defined in the `go.mod` file is `1.16`.
To use generics, the `go.mod` file needs to be updated to version `1.18` or higher.
```

## General guidance

The key to solving Strain is to handle the different types that are used for slicing.
Using generics can simplify the amount of code needed, thus keeping the code [DRY][dry].

## Approach: Using generics

```go
// Package strain is a small library for modifying slices of different types.
package strain

type Ints []int
type Lists [][]int
type Strings []string

type Slicer interface {
	int | string | []int
}

// KeepSlice filters in elements of a slice.
func KeepSlice[T Slicer](input []T, filter func(T) bool) []T {
	if input == nil {
		return input
	}
	output := make([]T, 0, len(input))
	for _, value := range input {
		if filter(value) {
			output = append(output, value)
		}
	}
	return output
}

// DiscardSlice filters out elements of a slice.
func DiscardSlice[T Slicer](input []T, filter func(T) bool) []T {
	return KeepSlice(input, func(val T) bool { return !filter(val) })
}

// Keep filters in elements of an int slice.
func (i Ints) Keep(filter func(int) bool) Ints {
	return KeepSlice(i, filter)
}

// Discard filters out elements of an int slice.
func (i Ints) Discard(filter func(int) bool) Ints {
	return DiscardSlice(i, filter)
}

// Keep filters in elements of an string slice.
func (s Strings) Keep(filter func(string) bool) Strings {
	return KeepSlice(s, filter)
}

// Keep filters in elements of a slice of int slices.
func (l Lists) Keep(filter func([]int) bool) Lists {
	return KeepSlice(l, filter)
}
```

For more information, check the [using generics approach][approach-using-generics].

[generics]: https://go.dev/blog/intro-generics
[dry]: https://en.wikipedia.org/wiki/Don%27t_repeat_yourself
[approach-using-generics]: https://exercism.org/tracks/go/exercises/strain/approaches/using-generics
