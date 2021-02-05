## Zero Values

Unlike some other languages, Go has no empty/null/undefined state for uninitialized variables. Every type has a [zero value][zero_values] that uninitialized variables will hold upon declaration.

These default values are called the zero values for their respective types:

| Type      | Zero Value |
| --------- | ---------- |
| bool      | false      |
| numeric   | 0          |
| string    | ""         |
| pointer   | nil        |
| func      | nil        |
| interface | nil        |
| slice     | nil        |
| channel   | nil        |
| map       | nil        |

The predeclared identifier `nil`, meaning zero, is the zero value for the more complex types in Go. Comparing a type whose zero value is not `nil` to `nil`, like string, is an error:

```go
    func main() {
        myString := "exercism"
        if myString != nil { // invalid operation: myString != nil (mismatched types string and nil)
            fmt.Println(myString)
        }
    }
```

[zero_values]: https://golang.org/ref/spec#The_zero_value
