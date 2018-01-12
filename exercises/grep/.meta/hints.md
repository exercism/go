## Implementation

Define a single Go func, GrepFiles, which accepts a pattern string,
a slice of flags which are strings, and a slice of filename strings.
GrepFiles should return a slice of strings of the output for
the given flags and filenames.

Use the following signature for func GrepFiles:

```
func GrepFiles(pattern string, flags []string, files []string) []string {
```

