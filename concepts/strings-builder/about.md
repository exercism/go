# About

In Go, strings are immutable.
Concatenating strings with `+` in a loop causes frequent allocations and copies data multiple times.
The [`strings` package](https://pkg.go.dev/strings) provides [`Builder`](https://pkg.go.dev/strings#Builder) to minimize memory overhead by efficiently assembling strings.

## Efficient Text Assembly

The zero-value of a `strings.Builder` is ready to use immediately without any explicit initialization or allocation:

```go
var sb strings.Builder
```

## Appending Content

You can push new data into the builder using specific write methods.
Each write method returns an error value to satisfy standard Go interfaces, but it is guaranteed to always be `nil`.
This error is typically ignored in everyday code.

The most frequently used methods are:

- `WriteString(s string) (int, error)`: Appends a string to the buffer.
- `WriteRune(r rune) (int, error)`: Appends the UTF-8 encoding of a rune.
- `WriteByte(c byte) error`: Appends a single byte.

## Retrieving and Resetting

Once you have finished assembling your text, you can extract the accumulated content as a standard Go string by calling the `String()` method.
Calling `String()` is a very cheap operation.
It returns the underlying byte slice as a string directly without copying the accumulated data.

If you need to reuse the same builder instance to assemble a new string (for example, inside a loop), you can call the `Reset()` method.
This empties the internal buffer.
It allows you to start fresh and build the next string while preserving the already allocated memory capacity under the hood.

```go
sb.Reset()
```

## Core Methods

| Method                                                          | Purpose                                       |
| --------------------------------------------------------------- | --------------------------------------------- |
| [`WriteString`](https://pkg.go.dev/strings#Builder.WriteString) | Appends a string to the buffer                |
| [`WriteRune`](https://pkg.go.dev/strings#Builder.WriteRune)     | Appends the UTF-8 encoding of a rune          |
| [`WriteByte`](https://pkg.go.dev/strings#Builder.WriteByte)     | Appends a single byte                         |
| [`Grow`](https://pkg.go.dev/strings#Builder.Grow)               | Pre-allocates memory for `n` additional bytes |
| [`Cap`](https://pkg.go.dev/strings#Builder.Cap)                 | Returns the current buffer capacity           |
| [`Len`](https://pkg.go.dev/strings#Builder.Len)                 | Returns the number of accumulated bytes       |
| [`String`](https://pkg.go.dev/strings#Builder.String)           | Returns the accumulated text as a string      |
| [`Reset`](https://pkg.go.dev/strings#Builder.Reset)             | Resets the builder to an empty state          |

## Code Examples

### Basic String Construction

The zero-value of a `strings.Builder` is ready to use:

```go
var sb strings.Builder

sb.WriteString("Gopher")
sb.WriteRune(' ')
sb.WriteRune('🚀')

result := sb.String()
fmt.Println(result)
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
