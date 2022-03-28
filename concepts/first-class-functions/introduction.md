# Introduction

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
