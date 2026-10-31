# About

In Go 1.18, Go added support for **generics**.
Generics allow you to write functions and data structures that work with different types without duplicating code.

## Why generics?

In Go, functions and types normally work with specific types, such as `int` or `string`.
Before generics, you had two choices to support multiple types.
First, you could write multiple functions with different names, such as `SumInts` and `SumFloats`.
This approach creates duplicate code that is difficult to maintain.
Second, you could use the empty interface `any` (`interface{}`).
However, the empty interface removes compile-time type checks and requires type assertions at runtime.
Generics solve this problem by providing type safety while avoiding duplicate code.

## Generic functions

You define a generic function by adding **type parameters** in square brackets `[...]` after the function name.
Here is an example of a generic function that returns a reversed slice:

```go
func Reverse[T any](s []T) []T {
    result := make([]T, len(s))
    for i, v := range s {
        result[len(s)-1-i] = v
    }
    return result
}
```

In this example, `T` is a type parameter.
You can think of `T` as a placeholder for a real type.
The `any` keyword inside the square brackets is a **type constraint**.
The constraint tells the compiler what types are permitted for `T`.
Here, `any` means `T` can be any type.

### Calling generic functions

You can call a generic function by passing the actual type in square brackets.
This actual type is called a **type argument**.

```go
numbers := []int{1, 2, 3}
reversedNumbers := Reverse[int](numbers)
```

In most cases, the Go compiler can automatically detect the type from the function arguments.
This feature is called **type inference**.
With type inference, you do not need to write the type argument in square brackets.

```go
words := []string{"apple", "banana"}
reversedWords := Reverse(words)
```

In this example, the compiler sees that `words` is a slice of strings, so it infers that `T` is `string`.

## Type constraints

Every type parameter must have a constraint.
A constraint specifies which types you can use as type arguments.
A constraint also determines the operations you are allowed to perform on values of that type.

### The `any` constraint

The `any` constraint allows any type.
In Go, `any` is an alias for the empty interface `interface{}`.
Because `any` allows every type, the function can only perform operations that work on all Go types.
For example, you cannot compare two values of type `any` using `==`, because functions and slices cannot be compared with `==`.

### The `comparable` constraint

Go provides a built-in constraint called `comparable`.
The `comparable` constraint allows any type that supports the `==` and `!=` operators.
Comparable types include booleans, numbers, strings, pointers, channels, and structs or arrays made of comparable types.

Here is an example of a function that checks whether a slice contains a specific item:

```go
func Contains[T comparable](elems []T, target T) bool {
    for _, elem := range elems {
        if elem == target {
            return true
        }
    }
    return false
}
```

Because `T` uses the `comparable` constraint, using `elem == target` is valid.
If you replace `comparable` with `any`, the compiler will report an error.

### Custom constraints using interface type sets

Starting with Go 1.18, interfaces can define sets of types in addition to methods.
You can use the pipe character `|` to list the allowed types in an interface.

```go
type Number interface {
    int | int64 | float64
}
```

You can now use `Number` as a constraint for a function:

```go
func Double[T Number](value T) T {
    return value * 2
}
```

Because all types in `Number` support the `*` operator, Go allows `value * 2` inside the function.

### The underlying type approximation `~`

Sometimes you create a custom type definition based on a primitive type:

```go
type MyInt int
```

The `Number` constraint defined above only permits the exact types `int`, `int64`, and `float64`.
If you pass a `MyInt` value to `Double`, the compiler produces an error because `MyInt` does not match `int`.
To allow any type whose underlying type is `int`, add a tilde `~` before the type:

```go
type Number interface {
    ~int | ~int64 | ~float64
}
```

Now you can use `Double` with values of type `MyInt`.

The standard library provides the `cmp.Ordered` constraint in the [`cmp` package][doc-cmp].
The `cmp.Ordered` constraint allows any type that supports the comparison operators `<`, `<=`, `>`, and `>=`.

## Generic types

You can also use type parameters to define generic types, such as structs.
Here is an example of a generic `Stack` data structure:

```go
type Stack[T any] struct {
    elements []T
}

func (s *Stack[T]) Push(v T) {
    s.elements = append(s.elements, v)
}

func (s *Stack[T]) Pop() (T, bool) {
    if len(s.elements) == 0 {
        var zero T
        return zero, false
    }
    lastIndex := len(s.elements) - 1
    top := s.elements[lastIndex]
    s.elements = s.elements[:lastIndex]
    return top, true
}
```

When you declare a variable of a generic type, you must provide the type argument in square brackets:

```go
intStack := Stack[int]{}
intStack.Push(42)

stringStack := Stack[string]{}
stringStack.Push("hello")
```

Type inference only works for generic functions.
You must always specify the type argument when creating values of a generic type.

## When to use generics

Generics are a powerful feature, but you do not need to use them for every problem.

Use generics in situations like these:
- You are building data structures, such as trees, linked lists, or caches, that work with any element type.
- You are writing utility functions that operate on slices or maps of any type, such as filtering, sorting, or searching.
- You need the exact same implementation for several types, and you want compile-time type safety.

Use interfaces instead in situations like these:
- You care about behavior and methods, rather than the underlying data type.
- Different types require different method implementations to accomplish their work.

For more information, see the [Go Generics Tutorial][doc-tutorial-generics], the [Go Blog introduction to generics][blog-intro-generics], and the [Go Language Specification on Type Parameters][ref-spec-type-params].

[doc-tutorial-generics]: https://go.dev/doc/tutorial/generics
[blog-intro-generics]: https://go.dev/blog/intro-generics
[ref-spec-type-params]: https://go.dev/ref/spec#Type_parameter_declarations
[doc-cmp]: https://pkg.go.dev/cmp
