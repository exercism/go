# Introduction

A `string` in Go is an immutable sequence of bytes, which don't necessarily have to represent characters.

A `string` literal is a sequence of bytes enclosed in double quotes:

```go
firstName := "Jane"
```

Strings can be concatenated via the `+` operator:

```go
fullName := "Jane" + " " + "Austen" // fullName = "Jane Austen"
```

Some special characters need to be escaped with a leading backslash, such as "/n" for new lines in strings.
```go
Dialogue := "How is the weather today?\nIt's sunny"  


// Dialogue = "How is the weather today?
// It's sunny"  
```


The `strings` package contains many useful functions to work on strings, such as Repeat, ToUpper, ToLower, ReplaceAll, TrimSpace and many more. For getting more information about strings functions check out [strings page](https://pkg.go.dev/strings).

```go
import "strings"

fmt.Println(strings.ToLower("MaKEmeLoweRCase")) // Output: "makemelowercase"

fmt.Println(strings.ToUpper("MaKemeUpPercaSe")) // Output: "MAKEMEUPPERCASE"

fmt.Println(strings.Repeat("Go", 3)) // Output: "GoGoGo" 

fmt.Println(strings.ReplaceAll("your cat is playing with your pillow", "your", "my")) // Output: "my cat is playing with my pillow

fmt.Println(strings.TrimSpace(" \t\n Hello, Gophers \n\t\r\n")) // Output: "Hello, Gophers"

```
