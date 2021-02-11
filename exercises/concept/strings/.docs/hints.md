# General

- The `fmt` package of the standard library has some [formatting functionality for strings][fmt-package].

- To add one string to another, you can use the `"+="` assignment.

```go
s := "Hello"
s += " word."
fmt.Println(s)
// Output: Hello world.
```

- The [Go 101 string section][go101] has more details about this concept.

[fmt-package]: https://golang.org/pkg/fmt/
[go101]: https://go101.org/article/string.html
