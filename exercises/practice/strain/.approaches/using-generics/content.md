# Using generics

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
	return Keep(input, func(val T) bool { return !filter(val) })
}
```

This approach starts be defining some [types][types] that will be used in the package.
It then defines an [interface type][interface-type] which will be used to constain which types are to be used for the [generics][generics]:

>Generics are a way of writing code that is independent of the specific types being used. Functions and types may now be written to use any of a set of types.

To keep the code [DRY][dry], there will be one function that does most of the work.

The `Keep()` function takes a generic type parameter of `[T Slicer]`.
`T` is the name of the parameter.
It does not have to be `T`, but it is a common convention in Go (and  other languages) to start with `T`,
and if a second name is needed, to use `U`, and so on.
The type `T` is constrained to be one of the types defined in the `Slicer` interface.
The arguments to the function are a slice of the generic type `T`, and a function which takes a value of type `T` and returns a boolean value.

~~~~exercism/note
If you haven't passed a function as a parameter before, a short explanation can be found [here](https://golangbyexamples.com/func-as-func-argument-go/).
~~~~

The `Keep()` function is defined to return a slice of type `T`.

If the input is `nil` it is simply returned.
Otherwise an output slice is made of the same type as `T`, given a length of `0`, and given a capacity that is the length of the input slice.

~~~~exercism/note
Since the returned slice can be no bigger than the input slice,
the output slice is given a capacity of the length of the input slice to prevent a lot of reallocating of the output slice.
However, if the usual expected input was large and the usual expected output was small,
then a lower capacity might be used to define the output slice to save allocating more memory than needed.
Since this exercise uses a small length for the input slices, then it is okay to make the capacity of the output slice the same length.
For more information on the length and capacity of slices, see [here](https://go.dev/blog/slices-intro).
~~~~

[`range`][range] is used to iterate the elements of the input slice.
Each element is passed to the input `filter` function.
If the function returns `true`, the the element is appended to the output slice.
After the iteration of all the input elements is finished, the output slice is returned from the function.
So, the `Keep()` function is intended to _keep_ the elements for which the `filter` function returns `true`.

The `Discard()` function is defined with the same signature (i.e. the same generic type parameter, arguments and return type)
as `Keep()`, but it is intended to _leave out_ the elements for which the `filter` function returns `true`.
The `Keep()` function can be re-used for this by logically negating the `filter` function passed in to `Discard()`.
This is done by wrapping it in another function created in `Discard()` like so

```go
func(val T) bool { return !filter(val) }
```

This defines an anonymous function which takes an argument of the same type as the `filter` function and also returns a boolean value.
Inside the function, it passes the argument to the `filter` function and returns the result of the `filter` function
logically negated by the [logical NOT operator][not-operator] (`!`).
So, if the `filter` function returns `true` for `Discard()` to leave out the element,
the anonymous function will return `false` when passed in to the `Keep()` function, so the element will not be kept in by `Keep()`.
`Keep()` can be called by `Discard()`, since not keeping an element in with `Keep()` is the same as leaving it out in `Discard()`.
`Discard()` passes the input slice and the anonymous function to `Keep()` and returns the expected result.

Being able to use the same function for multiple types, instead of having to write a separate function for each type,
saves a lot of work!

[types]: https://go.dev/ref/spec#Types
[interface-type]: https://go.dev/ref/spec#Interface_types
[generics]: https://go.dev/blog/intro-generics
[dry]: https://en.wikipedia.org/wiki/Don%27t_repeat_yourself
[range]: https://go.dev/tour/moretypes/16
[not-operator]: https://www.tutorialkart.com/golang-tutorial/golang-not/
[method]: https://go.dev/tour/methods/1
