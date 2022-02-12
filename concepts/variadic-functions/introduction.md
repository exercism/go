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

`Hence, variadic parameters should be passed at the end`. 

The above function will fail to compile with error syntax error: cannot use `...` with non-final parameter `b`



