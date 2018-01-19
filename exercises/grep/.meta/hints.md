## Implementation

In package grep, Define a single Go func, Search, which accepts a pattern string,
a slice of flags which are strings, and a slice of filename strings.
Search should return a slice of strings of the output for
the given flags and filenames.

Use the following signature for func Search:

```
func Search(pattern string, flags, files []string) []string {
```

