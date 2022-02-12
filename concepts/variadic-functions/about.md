# About
Functions in general accept only a fixed number of arguments. 
A variadic function is a function that accepts a variable number of arguments. 
If the last parameter of a function definition is prefixed by ellipsis `...`, then the function can accept any number of arguments for that parameter.

```go
func find(a int, b ...int) {
    // ...
}
```

In the above function, parameter b is variadic and we can pass 0 or more arguments to b.

```go
    find(5, 6)
    find(5, 6, 7)
    find(5)
```

Note:
    Variadic paramter should be the last parameter of the function.

```go
    func find(...a int, b int) {}
```
In the above function, we can pass multiple arguments but all those arguments will assigned to a.
`Hence, variadic parameters should be passed at the end`. The above function will fail to compile with error syntax error: cannot use ... with non-final parameter b

Example:
```go
    func find(num int, nums ...int) {
    fmt.Printf("type of nums is %T\n", nums)
    found := false
    for i, v := range nums {
        if v == num {
            fmt.Println(num, "found at index", i, "in", nums)
            found = true
        }
    }
    if !found {
        fmt.Println(num, "not found in ", nums)
    }
    fmt.Printf("\n")
    }
    func main() {
        find(89, 89, 90, 95)
        find(45, 56, 67, 45, 90, 109)
        find(78, 38, 56, 98)
        find(87)
    }
```
The way variadic functions work is by converting the variable number of arguments to a slice of the type of the variadic parameter. 
For instance, in line no. 22 of the program above, the variable number of arguments to the find function are 89, 90, 95. The find function expects a variadic int argument. 
Hence these three arguments will be converted by the compiler to a `slice of type int []int{89, 90, 95}` and then it will be passed to the find function.
