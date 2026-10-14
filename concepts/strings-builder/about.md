# About

In Go, strings are immutable.
Concatenating strings with `+` in a loop causes frequent allocations and copies data multiple times.
The [`strings` package](https://pkg.go.dev/strings) provides [`Builder`](https://pkg.go.dev/strings#Builder) to minimize memory overhead by efficiently assembling strings.

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

## Important Safety Restriction

**Do not copy a non-zero `strings.Builder`.**
It internally maintains a pointer to its growable buffer.
Copying the structure duplicates this internal pointer.
This causes multiple instances to share the same memory structure under the hood.
Modifying a copied builder triggers an immediate **runtime panic** to prevent data corruption.

```go
func main() {
    var sb strings.Builder
    sb.WriteString("test")
    
    // This function call hides a copy-by-value operation
    addStr(sb, "test-2") 
}

func addStr(sb strings.Builder, s string) {
    // This line triggers a panic: strings: illegal use of non-zero Builder copied by value
    sb.WriteString(s) 
}
```

### Safe Passing Alternative

If you need to pass a `strings.Builder` to another function, you must always pass it by pointer.
Passing by pointer avoids duplicating the internal data structures and safely modifies the original builder.

```go
func main() {
    var sb strings.Builder
    sb.WriteString("test")
    
    // Always pass the address using the & operator
    addStr(&sb, "test-2") 
    
    fmt.Println(sb.String())
    // Output: testtest-2
}

func addStr(sb *strings.Builder, s string) {
    // This is safe because we use a pointer to the original instance
    sb.WriteString(s) 
}
```

## Performance Comparison

To better understand the differences between standard concatenation and `strings.Builder`, consider the following implementation examples.
The first function uses a builder to assemble text through loops.
The second function uses standard string accumulation.

```go
func ConcatWithStrBuilder(iterations int) string {
    var sb strings.Builder

    for i := 0; i < iterations; i++ {
        sb.WriteString("Gopher")
        sb.WriteRune(' ')
        sb.WriteString("run")
        sb.WriteString("\n")
    }

    return sb.String()
}

func ConcatWithOutStrBuilder(iterations int) string {
    var result string

    for i := 0; i < iterations; i++ {
        result += "Gopher" + " " + "run" + "\n"
    }

    return result
}
```

You can evaluate their execution speed and allocation metrics by writing standard Go benchmark tests.

```go
func BenchmarkConcatWithSB(b *testing.B) {
    for b.Loop() {
        ConcatWithStrBuilder(10)
    }
}

func BenchmarkConcatWithOutStrBuilder(b *testing.B) {
    for b.Loop() {
        ConcatWithOutStrBuilder(10)
    }
}
```

### Benchmark Results

Running these tests on a modern architecture yields the following performance profiles.

```shell
goos: darwin
goarch: arm64
BenchmarkConcatWithSB-10                 7434296               146.8 ns/op           248 B/op          5 allocs/op
BenchmarkConcatWithOutStrBuilder-10      4913868               255.7 ns/op           664 B/op          9 allocs/op
```
