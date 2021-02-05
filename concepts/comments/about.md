Go supports two types of [comments][comments]. Single line comments are preceded by `//` and multiline comments are inserted between `/*` and `*/`.

## Documentation comments

In Go, comments play an important role in documenting code. They are used by the [godoc][godoc] command, which extracts these comments to create documentation about Go packages (see [pkg.go.dev][go packages] for examples). A documentation comment should be a complete sentence that starts with the name of the thing being described and ends with a period.

Documentation comments are about `what` the thing does and contains, whereas other types of comments such as code comments are about `why` something is done. The `why` helps others and future you to understand the reason behind a particular piece of code. However, if the `why` feels obvious, the code comment is not necessary. A [good rule of thumb][less comments] and more sustainable solution is to write code that is easier to understand so that explanatory comments are hopefully not needed.

## Comments for exported identifiers

Documentation comments should precede packages as well as [exported identifier][exported identifiers]s, for example exported functions, methods, package variables, constants, and structs, which you will learn more about in the next exercises. Comments written for packages and exported identifiers are useful for the users who import and use these packages.

Note, however, that identifiers (such as variables) that are declared inside of functions and methods are private and do not necessarily require comments for the users of the packages.

A package-level variable can look like this:

```go
// TemperatureFahrenheit represents a certain temperature in degrees Fahrenheit.
var TemperatureFahrenheit float64
```

Note that `TemperatureFahrenheit` is capitalized, which makes this identifier exported, and thus usable in any code which imports this package.

## Package comments

Package comments should be written directly before a package clause (`package x`) and begin with `Package x ...` like this:

```go
// Package kelvin provides tools to convert temperatures to and from Kelvin.
package kelvin
```

## Function comments

A function comment should be written directly before the function declaration. It should be a full sentence that starts with the function name. For example, an exported comment for the function `Test` should take the form `Test ...` and end with a period. It could also explain what arguments the function takes, what it does with them, and what its return values mean:

```go
// CelsiusFreezingTemp returns an integer value equal to the temperature at which water freezes in degrees Celsius.
func CelsiusFreezingTemp() int {
	return 0
}
```

## Golint

[Golint][golint] is a great tool to check for missing comments and other common stylistic issues, which you can read more about at [Effective Go][effective go].

You can install `golint` on your machine with the following command:

```
go get -u golang.org/x/lint/golint
```

It's a good idea to configure your editor to run `golint` for you, otherwise you can invoke it like this:

```
golint weather.go
```

To use the `golint` command globally, make sure that it is in your `$PATH`.

[godoc]: https://golang.org/cmd/go/#hdr-Show_documentation_for_package_or_symbol
[go packages]: https://pkg.go.dev/
[less comments]: https://dave.cheney.net/practical-go/presentations/qcon-china.html#_dont_comment_bad_code_rewrite_it
[exported identifiers]: https://www.ardanlabs.com/blog/2014/03/exportedunexported-identifiers-in-go.html
[golint]: https://github.com/golang/lint
[effective go]: https://golang.org/doc/effective_go.html
[comments]: https://golang.org/ref/spec#Comments
