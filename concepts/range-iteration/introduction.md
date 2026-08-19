# Introduction

In Go, `range` iterates over slices, arrays, maps, and strings.
Depending on what is being ranged over, each iteration yields one or two values.

## Iterating over a Slice

`range` over a slice yields the index and value of each element in order:

```go
vals := []int{10, 20, 30}
for i, v := range vals {
    fmt.Println(i, v)
}
// 0 10
// 1 20
// 2 30
```

## Iterating over a Map

`range` over a map yields each key and value, but the iteration order is not guaranteed.
Go deliberately randomizes map iteration order, so the same program may produce different output on each run.

```go
hash := map[int]int{9: 10, 99: 20, 999: 30}
for k, v := range hash {
    fmt.Println(k, v)
}
// 99 20
// 999 30
// 9 10
```

## Omitting Index or Value

Go will not compile if a variable is declared but never used.
If only the value is needed, assign the index or key to `_` to ignore it.

```go
vals := []int{10, 20, 30}
for i, v := range vals {
    fmt.Println(v)
}
// Go build failed: declared and not used: i
```

```go
vals := []int{10, 20, 30}
for _, v := range vals {
    fmt.Println(v)
}
// 10
// 20
// 30
```

If only the index is needed, omit it completely:

```go
vals := []int{10, 20, 30}
for i := range vals {
    fmt.Println(i)
}
// 0
// 1
// 2
```

If neither the index nor the value is needed, both can be omitted:

```go
vals := []int{10, 20, 30}
count := 0
for range vals {
    count++
}
// count == 3
```

## Range over an Integer

Since Go 1.22, `range` can iterate over an integer directly, yielding values from `0` up to but not including that integer.
Zero and negative values are valid but produce no iterations.

```go
for n := range 3 {
    fmt.Println(n)
}
// 0
// 1
// 2
```

## Range over an Iterator

Since Go 1.23, `range` accepts an iterator of type `iter.Seq[V]` or `iter.Seq2[K, V]`.
An iterator is a function that produces a sequence of values one at a time.
`range` can then step through those values.

```go
text := "The quick brown fox"
for word := range strings.FieldsSeq(text) {
    fmt.Println(word)
}
// The
// quick
// brown
// fox

names := []string{"Alice", "Bob", "Vera"}
for i, v := range slices.All(names) {
    fmt.Println(i, ":", v)
}
// 0 : Alice
// 1 : Bob
// 2 : Vera
```
