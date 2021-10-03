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

## Non-struct types

You've previously seen defining struct types.
It is also possible to define non-struct types which you can use as an alias for a built-in type declaration, and you can define receiver functions on them to extend them in the same way as struct types.

```go
type Name string
func SayHello(n Name) {
  fmt.Printf("Hello %s\n", n)
}
n := Name("Fred")
SayHello(n)
// Output: Hello Fred
```

You can also define non-struct types composed of arrays and maps.

```go
type Names []string
func SayHello(n Names) {
  for _, name := range n {
    fmt.Printf("Hello %s\n", name)
  }
}
n := Names([]string{"Fred", "Bill"})
SayHello(n)
// Output:
// Hello Fred
// Hello Bill
```
