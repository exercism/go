# Introduction

In Go, strings are immutable.
Concatenating strings with `+` in a loop causes frequent allocations and copies data multiple times.
The [`strings` package](https://pkg.go.dev/strings) provides [`Builder`](https://pkg.go.dev/strings#Builder) to minimize memory overhead by efficiently assembling strings.

## Building Strings

The zero-value of a `strings.Builder` is ready to use immediately without any explicit initialization or allocation.
You can push new data into the builder using specific write methods.
Each write method returns an error value to satisfy standard Go interfaces, but it is guaranteed to always be `nil` and is typically ignored in everyday code.

The most frequently used methods are:

- `WriteString(s string) (int, error)`: Appends a string to the buffer.
- `WriteRune(r rune) (int, error)`: Appends the UTF-8 encoding of a rune.
- `WriteByte(c byte) error`: Appends a single byte.

```go
var sb strings.Builder

sb.WriteString("Hello")
var r rune = ' '
sb.WriteRune(r)
var b byte = '.'
sb.WriteByte(b)
```

## Retrieving and Resetting

Once you have finished assembling your text, you can extract the accumulated content as a standard Go string by calling the `String()` method.
Calling `String()` is a very cheap operation because it returns the underlying byte slice as a string directly without copying the accumulated data.

If you need to reuse the same builder instance to assemble a new string (for example, inside a loop), you can call the `Reset()` method.
This empties the internal buffer, allowing you to start fresh and build the next string while preserving the already allocated memory capacity under the hood.

```go
var sb strings.Builder

sb.WriteString("Task One")
fmt.Println(sb.String()) // Output: Task One

// Clear the buffer for resource reuse
sb.Reset()

sb.WriteString("Task Two")
fmt.Println(sb.String()) // Output: Task Two
```
