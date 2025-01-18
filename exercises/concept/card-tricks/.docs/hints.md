# Hints

## General

- Slices in Go are zero-based. The first index in a slice is `0`.
- The builtin [`append`][append-builtin] function is [`variadic`][variadic-gobyexample]
- You can append the elements of one slice to another slice by using the three dots notation:

```go
a := []int{1, 3}
b := []int{4, 2, 6}
c := append(a, b...)
fmt.Println(c)
// Output: [1 3 4 2 6]
```

## 1. Create a slice with certain cards

- To create a slice pre-filled with some data, use the slice literal notation:
  `s := []T{x1, x2, ..., xn}`

## 2. Retrieve a card from a stack

- To get the `n`th item of a slice [use an index][go-slices].
- To check if an item exists in a slice use a conditional and compare the index with the [length of the slice][len-builtin].

## 3. Exchange a card in the stack

- To set the `n`th item in a slice [use an index][go-slices] and assign a new value to it.
- To add a new item to then end of a slice use the `append` function.

## 4. Add cards to the top of the stack

- Adding items to the front of a slice can be done by appending the elements of the original slice to the `value` argument slice.

## 5. Remove a card from the stack

- Removing an item from a slice can be done by appending the part after `index` to the part before index.

[go-slices]: https://blog.golang.org/go-slices-usage-and-internals
[make-builtin]: https://golang.org/pkg/builtin/#make
[len-builtin]: https://golang.org/pkg/builtin/#len
[append-builtin]: https://golang.org/pkg/builtin/#append
[variadic-gobyexample]: https://gobyexample.com/variadic-functions
