# Code Formatting

The Go programming language is actually the first language ever, that added opinionated tooling to auto-format code
right from the start when it was released in 2009.
Since then many newer languages have followed that example and older languages added opinionated auto-formatting to their tool chain.

## The Tool

The tool to format Go code is [gofmt](https://golang.org/cmd/gofmt/), pronounced as [go-fumpt].

> Gofmt's style is no one's favorite, yet gofmt is everyone's favorite. [Rob Pike (2015)](https://www.youtube.com/watch?v=PAAkCSZUG1c&t=8m43s)

To clarify: there are no options on how Go code is being formatted. There is only one way: the `gofmt` way.

In hindsight this decision by the core team and the overwhelming adoption by the community has proven to have even more advantages than originally anticipated.
Here are some of the advantages -- anticipated and otherwise:

- All Go code looks and feels the same. Code written by others feels familiar and is easy to read.
- There are no discussions about how to format code. (Spaces vs Tabs, end-of-line braces vs newline braces, etc.)
- Very easy to auto-generate code that looks just like human written code.

## Usage of `gofmt`

Most IDE's that support Go or have extensions for Go, enable `gofmt` to run on every save by default or at least provide a way to enable that easily.

To run `gofmt` manually:

```shell script
# show a diff
gofmt -d file.go

# apply changes to the file
gofmt -w file.go
```
