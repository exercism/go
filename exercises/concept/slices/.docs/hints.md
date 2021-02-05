## General

- slices in Go are zero-based. The first index in a slice is `0`.

## 1. Retrieve a card from a stack

- To get the `n`th item of a slice [use an index][go-slices].
- To check if an item exists in a slice use a conditional and compare the index with the [length of the slice][len-builtin].

Note: a slice always starts with index `0`, `1`, `2`, etc.

## 2. Exchange a card in the stack

- To set the `n`th item in a slice [use an index][go-slices] and asign a new value to it.
- To add a new item to a list use the builtin [`append`][append-builtin] function.

## 3. Create a stack of cards

- There are several ways to create a new slice. The simplest is `var s []int`. More options you get with the [`make`][make-builtin] function.
- Use the techniques from 2. to fill the slice with the correct values in a `for` loop.

## 4. Remove a card from the stack

- Removing an item from a slice can be done by appending the part after `index` to the part before index.

[go-slices]: https://blog.golang.org/go-slices-usage-and-internals
[make-builtin]: https://golang.org/pkg/builtin/#make
[len-builtin]: https://golang.org/pkg/builtin/#len
[append-builtin]: https://golang.org/pkg/builtin/#append
