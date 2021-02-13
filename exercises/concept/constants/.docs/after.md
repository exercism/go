# After

Constants in Go are simple, [unchanging values][const] created with the `const` keyword. Constants may be given an explicit type:

```go
const hi string = "hi" // string
```

or an implicit, default type:

```go
const hello = "hello" // string
```

when the variable needs a type. This gives constants more flexibility in Go's type system and allows them to work in a variety of contexts without triggering a compiler error:

```go
const number2 = 2 // 2 is an untyped numeric constant and does not need to be explicitly given the type float64 as required by the Sqrt method
sqrt2 := math.Sqrt(number2)
```

Go does not have enums like some other languages but the `iota` enumerator can be used to create successive untyped integer constants:

```go
const (
    a = 6 * iota // 0
    b            // 6
    c            // 12
    d            // 18
)
```

Complex types like maps and slices are mutable and cannot be constants; the compiler will throw errors:

```go
func main() {
    const m = map[int]int{2: 8}
    const s = []string{"exercism", "v3"}
}
// Output: const initializer map[int]int literal is not a constant
// Output: const initializer []string literal is not a constant
```

For a fuller explanation please see the resources [here][const2], [here][const3], and [here][const4].

[const]: https://golang.org/ref/spec#Constants
[const2]: https://golang.org/doc/effective_go.html#constants
[const3]: https://blog.golang.org/constants
[const4]: https://yourbasic.org/golang/untyped-constants/
