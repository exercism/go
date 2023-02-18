# Introduction

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

Because interfaces are implemented implicitly, a type can easily implement multiple interfaces.
It only needs to have all the necessary methods defined.

## Empty interface

There is one very special interface type in Go, the **empty interface** type that contains zero methods.
The empty interface is written like this: `interface{}`.
In Go 1.18 or higher, `any` can be used as well. It was defined as an alias.

Since the empty interface has no methods, every type implements it implicitly.
This is helpful for defining a function that can generically accept any value.
In that case, the function parameter uses the empty interface type.

[concept-methods]: /tracks/go/concepts/methods
