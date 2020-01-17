# Comments

> The Go project takes documentation seriously. Documentation is a huge part of making software accessible and maintainable.
> Of course it must be well-written and accurate, but it also must be easy to write and to maintain.
> Ideally, it should be coupled to the code itself so the documentation evolves along with the code.
> The easier it is for programmers to produce good documentation, the better for everyone.
> [Godoc: documenting Go code](https://blog.golang.org/godoc-documenting-go-code)

For this reason documentation in Go is created from the comments and there are some rules around comments to help ensure
high quality documentation.
The most visible rules are:

- Anything exported has a comment.
- A comment for an exported function has the format `// [FunctionName] does x, y, z.`
- Every package should have a comment and it starts with `// Package [PackageName]`

Comment rules are enforced by the [golint](https://github.com/golang/lint) tool which is _not_ installed with Go by default.

Further material on Comments in Go:
- [Effective Go](https://golang.org/doc/effective_go.html#commentary)
- [Practical Go (Dave Chaney)](https://dave.cheney.net/practical-go/presentations/qcon-china.html#_comments)

# Documentation

The Go tool chain provides the [godoc](https://godoc.org/golang.org/x/tools/cmd/godoc) tool for producing documentation
as HTML or plain text from the Go source code. The generated documentation for many open source packages is available 
at [godoc.org](https://godoc.org/).
