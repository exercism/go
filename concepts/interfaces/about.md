# About

## Interface as a set of methods

In its simplest form, an **interface type** is a set of method signatures.
Here is an example of an interface definition that includes two methods `Add` and `Value`:

```go
type Counter interface {
    Add(increment int)
    Value() int
}
```

The parameter names like `increment` can be omitted from the interface definition but they often increase readability.

Interface names in Go do not contain the word `Interface` or `I`.
Instead, they often end with `er`, e.g. `Reader`, `Stringer`.

## Implementing an interface

Any type that defines the methods of the interface automatically implicitly "implements" the interface.
There is no `implements` keyword in Go.

The following type implements the `Counter` interface we saw above.

```go
type Stats struct {
    value int
    // ...
}

func (s Stats) Add(v int) {
    s.value += v
}

func (s Stats) Value() int {
    return s.value
}

func (s Stats) SomeOtherMethod() {
    // The type can have additional methods not mentioned in the interface.
}
```

For implementing the interface, it does not matter whether the method has a value or pointer receiver.
(Revisit the [methods concepts][concept-methods] if you are unsure about those.)

> A value of interface type can hold any value that implements those methods. [^1]

That means `Stats` can now be used in all the places that expect the `Counter` interface.

```go
func SetUpAnalytics(counter Counter) {
    // ...
}

stats := Stats{}
SetUpAnalytics(stats)
// works because Stats implements Counter
```

If all function parameters and return values needed to use concrete (ie. non-interface) types, that would strongly limit their re-usability.
By using interface types for the parameters, only the behavior that is needed for the function to do its job is defined.
Now the function can be used with parameters of different concrete types, as long as those concrete types implement the interface.
The same logic applies when using interface types as return values.

Because interfaces are implemented implicitly, a type can easily implement multiple interfaces.
It only needs to have all the necessary methods defined.

## Combining interfaces

Interface definitions can include other interfaces.
Here is an example from the standard library.

```go
type ReadCloser interface {
	Reader
	Closer
}

// Based on these interfaces:

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}
```

Instead of repeating the method signatures for `Read` and `Close`, to define `ReadCloser`, the interface definition makes use of the existing definitions for `Reader` and `Closer`.

## Empty interface

There is one very special interface type in Go, the **empty interface** type that contains zero methods.
The empty interface is written like this: `interface{}`.
In Go 1.18 or higher, `any` can be used as well. It was defined as an alias.

Since the empty interface has no methods, every type implements it implicitly.
This is helpful for defining a function that can generically accept any value.
In that case, the function parameter uses the empty interface type.

For example:

```go
func AnythingGoes(i interface{}) string {
    if i == 0 {
        return "zero"
    }
    if i == "" {
        return "empty string"
    }
    return "something else"
}
```

[^1]: [Tour of Go: Interfaces][tour-of-go-interfaces]

[concept-methods]: /tracks/go/concepts/methods
[tour-of-go-interfaces]: https://go.dev/tour/methods/9
