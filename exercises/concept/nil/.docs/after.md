Unlike some other languages, Go has no empty/null/undefined state for uninitialized variables. Every type has a [zero value][zero_values] that uninitialized variables will hold upon declaration. The predeclared identifier `nil` has no type, so comparing a type whose zero value is not `nil`, like string, is an error:

```go
    func main() {
        myString := "exercism"
        if myString != nil { // invalid operation: myString != nil (mismatched types string and nil)
            fmt.Println(myString)
        }
    }
```

[zero_values]: https://golang.org/ref/spec#The_zero_value
