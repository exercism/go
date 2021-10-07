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
// outputs:
// 0, 10
// 1, 20
// 2, 30
```

## Iterate over a map

Iterating over a map raises a new problem. The order is now random.

```go
hash := map[int]int{0: 10, 1: 20, 2: 30}
for k, v := range hash {
  fmt.Println(k, v)
}
// outputs, for example:
// 1 20
// 2 30
// 0 10
```

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
