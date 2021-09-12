# About

## Indexes in slices

Working with indexes of slices should always be protected in some way by a check that makes sure the index actually exists.
Failing to do so will crash the entire application.

## Empty slices

`Nil`-slices are the default empty slice. They have no drawbacks towards a slice with no values in them.
The `len` function works on `nil`-slices, items can be added without initializing it, and so on.
If creating a new slice prefer `var s []int` (nil-slice) over `s := []int{}` (empty slice).

## Performance

When creating slices to be filled iteratively, there is a low hanging fruit to improve performance, if the final size of the slice is known.
The key is to minimize the amount of times memory has to be allocated, which is rather expensive and happens if the slice grows beyond its allocated memory space.
The safest way to do this is to specify a capacity `cap` for the slice with `s := make([]int, 0, cap)` and then to `append` to the slice as usual.
This way the space for `cap` amount of items is allocated immediately while the slice length is zero.
In practice, `cap` is often the length of another slice: `s := make([]int, 0, len(otherSlice))`.

## Append is not a pure function

The append function of Go is optimized for performance and therefore does not make a copy of the input slice.
This means that the original slice (1st parameter in `append`) will be changed sometimes.
