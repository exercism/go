# Introduction

In Go, you can iterate over a `slice` using `for` and an index, or you can use `range`.
`range` also allows you to iterate over a `map` or a `channel`.
This concept will cover iterating over a `map` but iterating over a `channel` is out of the scope for this concept.

Every iteration returns two values: the index/key and a copy of the element at that index/key.

## Iterate over a slice

Easy as pie, loops over a slice, ordered as expected.

```go
xi := []int{10, 20, 30}
for i, x := range xi {
  fmt.Println(i, x)
}
// =>
// 0, 10
// 1, 20
// 2, 30
```

## Iterate over a map

Iterating over a map raises a new problem. The order is now random.

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
// =>
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
// =>
// 0
// 1
// 2
```

Last but not least, if you are required to perform some action but you are not
interested in values nor keys of the slice or map, you can omit both index and
value:

```go
xi := []int{10, 20, 30}
count := 0
for range xi {
  count++
}
// count value:
// 3
```
