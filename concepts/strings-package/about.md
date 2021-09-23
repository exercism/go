# About

The [`strings` package](https://pkg.go.dev/strings) contains many useful functions to work on strings:

| Function                                            | Purpose                                                             |
| --------------------------------------------------- | ------------------------------------------------------------------- |
| [ToLower](https://pkg.go.dev/strings#ToLower)       | Convert the string to lower case                                    |
| [ToUpper](https://pkg.go.dev/strings#ToUpper)       | Convert the string to upper case                                    |
| [TrimSpace](https://pkg.go.dev/strings#TrimSpace)   | Trim leading and trailing spaces                                    |
| [Index](https://pkg.go.dev/strings#Index)           | Find the index of the first instance of a substring within a string |
| [Replace](https://pkg.go.dev/strings#Replace)       | Replace one occurence of a substring in a string                    |
| [ReplaceAll](https://pkg.go.dev/strings#ReplaceAll) | Replace all occurences of a substring in a string                   |
| [Split](https://pkg.go.dev/strings#Split)           | Split a string into parts based on a separator                      |
| [HasSuffix](https://pkg.go.dev/strings#HasSuffix)   | Check if a string ends with a specific substring                    |
| [Count](https://pkg.go.dev/strings#Count)           | Count the number of occurrences of a substring within a string      |

```go
import "strings"

strings.ToLower("Gopher")    // Output: "gopher"
strings.Index("Apple", "le") // Output: 3
strings.Count("test", "t")   // Output: 2
```
