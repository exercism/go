# About

In Go functions are first-class values. This means that you can do with functions the same things you can do with all other values - assign functions to variables, pass them as arguments to other functions, return functions from other functions, etc.
For example:

```go
import "fmt"

func engGreeting(name string) string { return fmt.Sprintf("Hello %s, nice to meet you!", name) }
func espGreeting(name string) string { return fmt.Sprintf("¡Hola %s, mucho gusto!", name) }

greeting := engGreeting
fmt.Println(greeting("Alice")) // Hello Alice, nice to meet you!

greeting = espGreeting
fmt.Println(greeting("Alice")) // ¡Hola Alice, mucho gusto!
```

Function values provide an opportunity to parametrize functions not with data only, but behaviour too. For example:

```go
func dialog(name string, greetingFunc func(string) string) {
	fmt.Println(greetingFunc(name))
	fmt.Println("I'm a dialog bot.")
}

dialog("Alice", greeting)
// Output:
// ¡Hola Alice, mucho gusto!
// I'm a dialog bot.
```

The value of an uninitialized variable of function type is `nil`. Therefore, calling a `nil` function value causes a panic. Function values can be compared with `nil` and it is used to avoid unnecessary program panics. But functional values are not comparable against each other.

```go
var dutchGreeting func(string) string
dutchGreeting("Alice") // panic: call of nil function
```

Using function values is possible thanks to the function types in Go. A function type denotes the set of all functions with the same sequence of parameter types and the same sequence of result types. User-defined types can be declared on top of function types. For instance, the `dialog` function from one of the previous examples can be updated as following:

```go
type greetingFunc func(string) string

func dialog(name string, f greetingFunc) {
	fmt.Println(f(name))
	fmt.Println("I'm a dialog bot.")
}
```

Another powerful tool that is available thanks to first-class functions support is anonymous functions. Anonymous functions are defined at their point of use, without a name following the `func` keyword. Such functions have access to the variables of the enclosing function. For example:

```go
func fib() func() int {
	var n1, n2 int

	return func() int {
		if n1 == 0 && n2 == 0 {
			n1 = 1
		} else {
			n1, n2 = n2, n1 + n2
		}
		return n2
	}
}

next := fib()
for i := 0; i < N; i++ {
  fmt.Printf("F%d\t= %4d\n", i, next())
}
```

A call to fib initialises local variables n1, n2, and returns an anonymous function that, each time it is called, changes variables of the enclosing function. Nth call of the anonymous function returns the Nth number of the Fibonacci sequence starting from 0. It demonstrates that function values can have state. This feature is broadly used in JavaScript and is known as Immediately-Invoked Function Expression (IIFE).

## Fist class function usage
* Event handlers - one of the most used is HTTP handler
* Middlewares
* Collection operations such as map, filter, reduce

The standard package library in Go is full of great examples of first-class functions usage. For instance, package [strings][strings-pkg] - all package functions with the suffix `Func` accept a custom function as parameter to support the operation logic.

[strings-pkg]: https://pkg.go.dev/strings