# About

The [`strings` package](https://pkg.go.dev/strings) contains many useful functions to work on strings:

| Function                                      | Purpose                                                             |
| --------------------------------------------- | ------------------------------------------------------------------- |
| [ToLower](https://pkg.go.dev/strings#ToLower) | Convert the string to lower case                                    |
| [Index](https://pkg.go.dev/strings#Index)     | Find the index of the first instance of a substring within a string |
| [Count](https://pkg.go.dev/strings#Count)     | Count the number of occurrences of a substring within a string      |

```go
import "strings"

strings.ToLower("Gopher")    // Output: "gopher"
strings.Index("Apple", "le") // Output: 3
strings.Count("test", "t")   // Output: 2
```
