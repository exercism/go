# About

The [`strings` package](https://pkg.go.dev/strings) provides functions for working with strings.
Below are some of the functions in the package; check out the documentation to see the rest!

| Role                  | Function                                            | Purpose                                                             |
| --------------------- | --------------------------------------------------- | ------------------------------------------------------------------- |
| Case Conversion       | [ToLower](https://pkg.go.dev/strings#ToLower)       | Convert a string to lower case                                      |
| Case Conversion       | [ToUpper](https://pkg.go.dev/strings#ToUpper)       | Convert a string to upper case                                      |
| Searching             | [Contains](https://pkg.go.dev/strings#Contains)     | Check if a string contains a specific substring                     |
| Searching             | [Count](https://pkg.go.dev/strings#Count)           | Count non-overlapping occurrences of a substring                    |
| Searching             | [HasPrefix](https://pkg.go.dev/strings#HasPrefix)   | Check if a string starts with a specific substring                  |
| Searching             | [HasSuffix](https://pkg.go.dev/strings#HasSuffix)   | Check if a string ends with a specific substring                    |
| Searching             | [Index](https://pkg.go.dev/strings#Index)           | Return the index of the first occurrence of a substring, or -1      |
| Modifying             | [Replace](https://pkg.go.dev/strings#Replace)       | Replace up to n occurrences of a substring                          |
| Modifying             | [ReplaceAll](https://pkg.go.dev/strings#ReplaceAll) | Replace all occurrences of a substring                              |
| Modifying             | [TrimSpace](https://pkg.go.dev/strings#TrimSpace)   | Remove leading and trailing whitespace                              |
| Splitting or Joining  | [Fields](https://pkg.go.dev/strings#Fields)         | Split a string into words by whitespace                             |
| Splitting or Joining  | [FieldsSeq](https://pkg.go.dev/strings#FieldsSeq)   | Get each word one at a time                                         |
| Splitting or Joining  | [Join](https://pkg.go.dev/strings#Join)             | Join a slice of strings with a separator                            |
| Splitting or Joining  | [Split](https://pkg.go.dev/strings#Split)           | Split a string into parts on a separator                            |
| Splitting or Joining  | [SplitSeq](https://pkg.go.dev/strings#SplitSeq)     | Get each part one at a time                                         |

```go
import (
    "fmt"
    "strings"
)

// Case conversion
strings.ToLower("Gopher")              // "gopher"
strings.ToUpper("Gopher")              // "GOPHER"

// Searching
strings.Contains("Apple", "le")        // true
strings.Index("Apple", "le")           // 3
strings.Index("Apple", "xyz")          // -1
strings.Count("test", "t")             // 2
strings.HasPrefix("Gopher", "Go")      // true
strings.HasSuffix("Gopher", "er")      // true

// Modifying
strings.TrimSpace("  hello  ")         // "hello"
strings.Replace("aabbcc", "b", "x", 1) // "aaxbcc"
strings.ReplaceAll("aabbcc", "b", "x") // "aaxxcc"

// Splitting or Joining
strings.Fields("  a b\tc  ")               // []string{"a", "b", "c"}
strings.Split("a,b,c", ",")                // []string{"a", "b", "c"}
strings.Join([]string{"a", "b", "c"}, ",") // "a,b,c"

// You do not need to worry about this range syntax yet.
// It is covered later in the range iteration concept.
for word := range strings.FieldsSeq("  a b\tc  ") {
    fmt.Println(word)
}
// Output:
// a
// b
// c

for part := range strings.SplitSeq("a,b,c", ",") {
    fmt.Println(part)
}
// Output:
// a
// b
// c
```
