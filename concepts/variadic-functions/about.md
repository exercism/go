# About

Typically, functions accept only a fixed number of arguments.
Prefix the last parameter's type with `...` to accept any number of trailing arguments:

```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}
```

Inside the function, the variadic parameter is a slice:

```go
sum(1, 2, 3)    // nums is []int{1, 2, 3}
sum(1, 2, 3, 4) // nums is []int{1, 2, 3, 4}
sum()           // nums is []int{}
```

A function can have non-variadic parameters before the variadic one:

```go
func greet(greeting string, names ...string) {
    for _, name := range names {
        fmt.Printf("%s, %s!\n", greeting, name)
    }
}
```

## Spreading a slice

To pass a slice into the variadic parameter, follow it with `...`:

```go
nums := []int{1, 2, 3}
sum(nums...) // equivalent to sum(1, 2, 3)
```

`...` is only valid when passing a slice to a variadic parameter.
