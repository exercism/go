# Introduction

## Slices

Slices in Go are similar to lists or arrays in other languages.
They hold several elements of a specific type (or interface).

Slices in Go are based on arrays.
Arrays have a fixed size.
A slice, on the other hand, is a dynamically-sized, flexible view of the elements of an array.

A slice is written like `[]T` with `T` being the type of the elements in the slice:

```go
var empty []int                 // an empty slice
withData := []int{0,1,2,3,4,5}  // a slice pre-filled with some data
```

You can get or set an element at a given zero-based index using the square-bracket notation:

```go
withData[1] = 5
x := withData[1] // x is now 5
```

You can create a new slice from an existing slice by getting a range of elements.
Once again using square-bracket notation, but specifying both a starting (inclusive) and ending (exclusive) index.
If you don't specify a starting index, it defaults to 0.
If you don't specify an ending index, it defaults to the length of the slice.

```go
newSlice := withData[2:4]
// => []int{2,3}
newSlice := withData[:2]
// => []int{0,1}
newSlice := withData[2:]
// => []int{2,3,4,5}
newSlice := withData[:]
// => []int{0,1,2,3,4,5}
```

You can add elements to a slice using the `append` function.
Below we append `4` and `2` to the `a` slice.

```go
a := []int{1, 3}
a = append(a, 4, 2)
// => []int{1,3,4,2}
```

`append` always returns a new slice, and when we just want to append elements to an existing slice, it's common to reassign it back to the slice variable we pass as the first argument as we did above.

`append` can also be used to merge two slices:

```go
nextSlice := []int{100,101,102}
newSlice  := append(withData, nextSlice...)
// => []int{0,1,2,3,4,5,100,101,102}
```

## Variadic Functions

Usually, functions in Go accept only a fixed number of arguments.
However, it is also possible to write variadic functions in Go.

A variadic function is a function that accepts a variable number of arguments.

If the type of the last parameter in a function definition is prefixed by ellipsis `...`, then the function can accept any number of arguments for that parameter.

```go
func find(a int, b ...int) {
    // ...
}
```

In the above function, parameter `b` is variadic and we can pass 0 or more arguments to `b`.

```go
find(5, 6)
find(5, 6, 7)
find(5)
```

~~~~exercism/caution
The variadic parameter must be the last parameter of the function.
~~~~

The way variadic functions work is by converting the variable number of arguments to a slice of the type of the variadic parameter.

Here is an example of an implementation of a variadic function.

```go
func find(num int, nums ...int) {
    fmt.Printf("type of nums is %T\n", nums)

    for i, v := range nums {
        if v == num {
            fmt.Println(num, "found at index", i, "in", nums)
            return
        }
    }

    fmt.Println(num, "not found in ", nums)
}

func main() {
    find(89, 90, 91, 95)
    // =>
    // type of nums is []int
    // 89 not found in  [90 91 95]

    find(45, 56, 67, 45, 90, 109)
    // =>
    // type of nums is []int
    // 45 found at index 2 in [56 67 45 90 109]

    find(87)
    // =>
    // type of nums is []int
    // 87 not found in  []
}
```

In line `find(89, 90, 91, 95)` of the program above, the variable number of arguments to the find function are `90`, `91` and `95`.
The `find` function expects a variadic int parameter after `num`.
Hence these three arguments will be converted by the compiler to a slice of type `int` `[]int{90, 91, 95}` and then it will be passed to the find function as `nums`.

Sometimes you already have a slice and want to pass that to a variadic function.
This can be achieved by passing the slice followed by `...`.
That will tell the compiler to use the slice as is inside the variadic function.
The step described above where a slice is created will simply be omitted in this case.

```go
list := []int{1, 2, 3}
find(1, list...) // "find" defined as shown above
```

[append-yourbasic]: https://yourbasic.org/golang/append-explained/
