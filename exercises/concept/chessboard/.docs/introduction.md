# Introduction

In Go, you can iterate a `slice` using `for` but you could also use
`range`. You will find this useful to iterate over a `map`.

Every iteration returns two values: the index/key and a copy of the element at
that index/key.

## Iterate a slice

Easy as pie, loops a slice, ordered as expected.

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

## Iterate a map

Iterating a map raises a new problem. The order is now random.

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

In Go an unused variable will raise an error at build time, so if you only
need to use the value, as per the first example:

```go
xi := []int{10, 20, 30}
for i, x := range xi {
  fmt.Println(x)
}
// Go build failed: i declared but not used
```

So, let's replace the `i` with `_` telling the compiler we don't use that
value at all:

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

Now, if you want to only print the index, you can replace the `x` with `_`,
or simply omit the declaration at all:

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

You've previously seen defining struct types, but it's also possible to define non-struct types which you can use as an alias for a built in type declaration, and you can define reciever functions on them to extend them in the same way as struct types.

```go
type Name string
func (n Name) SayHello() {
  fmt.Printf("Hello %s\n", n)
}
n := Name("Fred")
n.SayHello()
// Output: Hello Fred
```

You can also define non-struct types composed of arrays and maps.

```go
type Names []string
func (n Names) SayHello() {
  for _, name := range n {
    fmt.Printf("Hello %s\n", name)
  }
}
n := Names([]string{"Fred", "Bill"})
n.SayHello()
// Output:
// Hello Fred
// Hello Bill
```
