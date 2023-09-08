# Introduction

There is at least one idomatic approach to solving Strain in Go.
You can use [generics][generics], which were introduced in version 1.18.

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

// Keep filters in elements of a slice.
func Keep[T Slicer](input []T, filter func(T) bool) []T {
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

// Discard filters out elements of a slice.
func Discard[T Slicer](input []T, filter func(T) bool) []T {
	return KeepSlice(input, func(val T) bool { return !filter(val) })
}

```

For more information, check the [using generics approach][approach-using-generics].

[generics]: https://go.dev/blog/intro-generics
[dry]: https://en.wikipedia.org/wiki/Don%27t_repeat_yourself
[approach-using-generics]: https://exercism.org/tracks/go/exercises/strain/approaches/using-generics
