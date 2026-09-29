# Introduction

In Go, you can iterate over a `slice` using `for` and an index, or you can use `range`.
`range` also allows you to iterate over a `map`.

Every iteration returns two values: the index/key and a copy of the element at that index/key.

## Iterate over a slice

Easy as pie, loops over a slice, ordered as expected.

```go
xi := []int{10, 20, 30}
for i, x := range xi {
  fmt.Println(i, x)
}
// outputs:
// 0, 10
// 1, 20
// 2, 30
```

## Iterate over a map

Iterating over a map raises a new problem.
The order is now random.

```go
hash := map[int]int{9: 10, 99: 20, 999: 30}
for k, v := range hash {
  fmt.Println(k, v)
}
// outputs, for example:
// 99 20
// 999 30
// 9 10
```

~~~~exercism/note
It may seem the above output is incorrect, as one would expect the first key/value pair on the declaration of the map `9 10` to be the first one printed and not the last.
However, maps are unordered by nature - there isn't a first or last key/value pair.
Because of that, when iterating over the entries of a map, the order by which entries will be visited will be random and not follow any specific pattern.
This means the above output is possible but might differ from what you get if you try to run this yourself.
To learn more about this see [Go Language Spec: range clause](https://go.dev/ref/spec#RangeClause).
~~~~

## Iteration omitting key or value

In Go an unused variable will raise an error at build time.
Sometimes you only need the value, as per the first example:

```go
xi := []int{10, 20, 30}
for i, x := range xi {
  fmt.Println(x)
}
// Go build failed: i declared but not used
```

You can replace the `i` with `_` which tells the compiler we don't use that value:

```go
xi := []int{10, 20, 30}
for _, x := range xi {
  fmt.Println(x)
}
// outputs:
// 10
// 20
// 30
```

If you want to only print the index, you can replace the `x` with `_`,
or simply omit the declaration:

```go
xi := []int{10, 20, 30}
// for i, _ := range xi {
for i := range xi {
  fmt.Println(i)
}
// outputs:
// 0
// 1
// 2
```

## Defining Types

A type definition creates a new named type, distinct from every other type, even one with the same underlying type.

```go
type File []bool
```

`File` here refers to a single vertical column on a chessboard, commonly labeled A to H.
Each square in this column is represented by a Boolean value marking whether it's occupied or not.
Now that `File` is defined, we can use it inside other definitions:

```go
type Chessboard map[string]File

board := Chessboard{
    "A": File{true, false, true},
    "B": File{false, true, false},
}
```

`Chessboard` maps file identifiers to `File` values.
Each `File` is a slice of Booleans representing the squares in one chess file.
Without these type definitions, `board` would be a `map[string][]bool`, which would make it harder to understand what we're representing.
