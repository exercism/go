## time

In Go, functionality for working with times is provided by the `time` package. The types and methods in this package allow us to manipulate times, get the current time, determine elapsed time, parse times from strings, and more.

To work with time, you will usually call a method on a `Time` instance, but there are also some functions called on the `time` package itself.

Parsing strings works a little differently in Go. Many languages use a format string with various codes like YYYY for four-digit year, MM for two-digit month, etc. Go instead uses an exact date, `Mon Jan 2 15:04:05 -0700 MST 2006`, to parse strings. You can rearrange these components or omit them as necessary, and you can spell out month and day names, or use their number values, but in order for Go to understand the parts of the string you are parsing, the corresponding parts of the format string must be from this exact date.
