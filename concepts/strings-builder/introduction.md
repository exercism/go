# Introduction

In Go, strings are immutable.
Concatenating strings with `+` in a loop causes frequent allocations and copies data multiple times.
The [`strings` package](https://pkg.go.dev/strings) provides [`Builder`](https://pkg.go.dev/strings#Builder) to minimize memory overhead by efficiently assembling strings.

## Basic String Construction

The zero-value of a `strings.Builder` is ready to use:

```go
var sb strings.Builder

sb.WriteString("Gopher")
sb.WriteRune(' ')
sb.WriteRune('🚀')

fmt.Println(sb.String())
// Output: Gopher 🚀
```

## Resource Reuse

Use `Reset()` to clear the buffer and reuse the instance:

```go
var sb strings.Builder

sb.WriteString("Task One")
fmt.Println(sb.String()) // Output: Task One

sb.Reset()

sb.WriteString("Task Two")
fmt.Println(sb.String()) // Output: Task Two
```
