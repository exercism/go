# About

In Go, strings are immutable.
Concatenating strings with `+` in a loop causes frequent allocations and copies data multiple times.
The [`strings` package](https://pkg.go.dev/strings) provides [`Builder`](https://pkg.go.dev/strings#Builder) to minimize memory overhead by efficiently assembling strings.

## Core Methods

| Method | Purpose |
| --- | --- |
| [`WriteString`](https://pkg.go.dev/strings#Builder.WriteString) | Appends a string to the buffer |
| [`WriteRune`](https://pkg.go.dev/strings#Builder.WriteRune) | Appends the UTF-8 encoding of a rune |
| [`WriteByte`](https://pkg.go.dev/strings#Builder.WriteByte) | Appends a single byte |
| [`Grow`](https://pkg.go.dev/strings#Builder.Grow) | Pre-allocates memory for `n` additional bytes |
| [`Cap`](https://pkg.go.dev/strings#Builder.Cap) | Returns the current buffer capacity |
| [`Len`](https://pkg.go.dev/strings#Builder.Len) | Returns the number of accumulated bytes |
| [`String`](https://pkg.go.dev/strings#Builder.String) | Returns the accumulated text as a string |
| [`Reset`](https://pkg.go.dev/strings#Builder.Reset) | Resets the builder to an empty state |

## Code Examples

### Basic String Construction

The zero-value of a `strings.Builder` is ready to use:

```go
var sb strings.Builder

sb.WriteString("Gopher")
sb.WriteRune(' ')
sb.WriteRune('🚀')

fmt.Println(sb.String())
// Output: Gopher 🚀
```

### Resource Reuse

Use `Reset()` to clear the buffer and reuse the instance:

```go
var sb strings.Builder

sb.WriteString("Task One")
fmt.Println(sb.String()) // Output: Task One

sb.Reset()

sb.WriteString("Task Two")
fmt.Println(sb.String()) // Output: Task Two
```

### Capacity Optimization

Use `Grow()` upfront if the final string size is predictable:

```go
var sb strings.Builder

sb.Grow(100)
for i := 0; i < 10; i++ {
    sb.WriteString("item...")
}
```

## Important Safety Restriction

**Do not copy a non-zero `strings.Builder`.**
It internally maintains a pointer to its growable buffer.
Copying the structure duplicates this internal pointer, causing multiple instances to share the same memory.
Modifying a copied builder triggers an immediate **runtime panic** to prevent data corruption.
