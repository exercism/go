# Introduction

A `string` in Go is an immutable sequence of bytes, which don't necessarily have to represent characters.

A string literal is defined between double quotes:


```go
const name = "Jane"
```
Strings can be concatenated via the `+` operator:

```go

fullName := "Jane" + " " + "Austen" // "Jane Austen"
```

Some special characters need to be escaped with a leading backslash, such as `\t` for a tab and `\n` for a new line in strings.

```go
Dialogue := "How is the weather today?\nIt's sunny"  
// "How is the weather today?
// It's sunny"  
```

The `strings` package contains many useful functions to work on strings, such as `Repeat`, `ToUpper`, `ToLower`, `ReplaceAll`, `TrimSpace`, and many more. For getting more information about strings functions check out the [strings package documentation](https://pkg.go.dev/strings).

```go
import "strings"

// strings.ToLower returns the string given as argument with all its characters lowercased
fmt.Println(strings.ToLower("MaKEmeLoweRCase"))
// Output: "makemelowercase"

// strings.ToUpper(s string) returns the string given as argument with all its characters uppercased
fmt.Println(strings.ToUpper("MaKemeUpPercaSe"))
// Output: "MAKEMEUPPERCASE"

// strings.Repeat returns a string with a substring given as argument repeated many times
fmt.Println(strings.Repeat("Go", 3))
// Output: "GoGoGo" 

// strings.ReplaceAll returns a copy of a string with all non-overlapping instances of a substring replaced by another substring
fmt.Println(strings.ReplaceAll("your cat is playing with your pillow", "your", "my"))
// Output: "my cat is playing with my pillow"

// strings.TrimSpace removes all leading and trailing whitespace from a string
fmt.Println(strings.TrimSpace(" \t\n Hello, Gophers \n\t\r\n"))
// Output: "Hello, Gophers"

```
