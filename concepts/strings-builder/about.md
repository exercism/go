# About String Builder

The [`strings` package](https://pkg.go.dev/strings) provides a specialized type called [`Builder`](https://go.dev) to efficiently build strings using write methods.

In Go, strings are immutable. Any standard modification or concatenation (like `s += "text"` in a loop) forces Go to allocate a completely new string in memory and copy the old data over. The `strings.Builder` type solves this by accumulating text inside an internal, mutable byte buffer, minimizing memory allocation and copying overhead.

## Core Methods

Below are the essential methods provided by `strings.Builder` for constructing strings:

| Role                  | Method                                                 | Purpose                                                             |
| --------------------- | ------------------------------------------------------ | ------------------------------------------------------------------- |
| Writing Data          | [WriteString](https://go.dev.WriteString) | Appends the contents of a string to the buffer                      |
| Writing Data          | [WriteRune](https://go.dev.WriteRune)   | Appends the UTF-8 encoding of a Unicode rune to the buffer          |
| Writing Data          | [WriteByte](https://go.dev.WriteByte)   | Appends a single raw byte to the buffer                             |
| Performance Tuning    | [Grow](https://go.dev.Grow)             | Pre-allocates memory for another `n` bytes to avoid re-allocations  |
| Performance Tuning    | [Cap](https://go.dev.Cap)               | Returns the total capacity of the underlying allocated byte slice   |
| Performance Tuning    | [Len](https://go.dev.Len)               | Returns the number of accumulated bytes                             |
| Control & Retrieval   | [String](https://go.dev.String)         | Returns the accumulated text as a final string                      |
| Control & Retrieval   | [Reset](https://go.dev.Reset)           | Resets the builder to be empty, making it ready for reuse           |

## Code Examples

### Basic String Construction

The zero-value of a `strings.Builder` is immediately ready to use. You can write strings and individual runes seamlessly:

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    var b strings.Builder

    // Appending standard strings
    b.WriteString("Gopher")

    // Appending individual Unicode characters (runes)
    b.WriteRune(' ')
    b.WriteRune('🚀')

    // Retrieving the final result
    fmt.Println(b.String()) // Output: Gopher 🚀
}
```

### Resource Reuse with Reset()

If you need to process a batch of data or build multiple strings in a sequence, you don't need to redeclare a new builder variable. Calling `Reset()` clears the internal buffer and brings the builder back to its pristine initial state:

```go
var b strings.Builder

// First task
b.WriteString("Task One")
fmt.Println(b.String()) // "Task One"

// Clear and reuse the exact same builder instance
b.Reset() 

b.WriteString("Task Two")
fmt.Println(b.String()) // "Task Two"
```

### Performance Optimization with Grow()

If you happen to know or can approximate the final size of the string you are building, you can use `Grow()` to explicitly reserve memory upfront. This eliminates the CPU overhead of dynamically resizing the buffer multiple times:

```go
var b strings.Builder

// Pre-allocate space for 100 bytes ahead of time
b.Grow(100) 

for i := 0; i < 10; i++ {
    b.WriteString("item...") 
}
```

## Important Safety Restriction

Once a `strings.Builder` has had data written to it, **it must not be copied**. Because the builder internally holds a pointer to its own growable buffer, copying the builder structure will result in multiple instances pointing to the same memory.

If you attempt to modify or use a copied instance of a non-zero builder, Go will immediately trigger a **runtime panic** to protect against memory corruption.
