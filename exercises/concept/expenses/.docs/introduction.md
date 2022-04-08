# Introduction

In Go, functions are first-class values. This means that you can do with functions the same things you can do with all other values - assign functions to variables, pass them as arguments to other functions or even return functions from other functions. 

Below we are creating two functions, `engGreeting` and `espGreeting` and we are assigning them to the variable `greeting`:

```go
import "fmt"

func engGreeting(name string) string {
	return fmt.Sprintf("Hello %s, nice to meet you!", name)
}

func espGreeting(name string) string {
	return fmt.Sprintf("¡Hola %s, mucho gusto!", name)
}

greeting := engGreeting			// greeting is a variable of type func(string) string
fmt.Println(greeting("Alice"))	// Hello Alice, nice to meet you!

greeting = espGreeting
fmt.Println(greeting("Alice")) 	// ¡Hola Alice, mucho gusto!
```

Function values provide an opportunity to parametrize functions not only with data but with behavior too.
In the following example, we are passing behaviour to the `dialog` function via the `greetingFunc` parameter:

```go
func dialog(name string, greetingFunc func(string) string) {
	fmt.Println(greetingFunc(name))
	fmt.Println("I'm a dialog bot.")
}

func espGreeting(name string) string {
	return fmt.Sprintf("¡Hola %s, mucho gusto!", name)
}

greeting := espGreeting
dialog("Alice", greeting)
// Output:
// ¡Hola Alice, mucho gusto!
// I'm a dialog bot.
```

The value of an uninitialized variable of function type is `nil`.
Therefore, calling a `nil` function value causes a panic.

```go
var dutchGreeting func(string) string
dutchGreeting("Alice") // panic: call of nil function
```

Function values can be compared with `nil`. This can be useful to avoid unnecessary program panics.

```go
var dutchGreeting func(string) string
if dutchGreeting != nil {
	dutchGreeting("Alice") // safe to call dutchGreeting
}
```

## Function types

Using function values is possible thanks to the function types in Go. A function type denotes the set of all functions with the same sequence of parameter types and the same sequence of result types. User-defined types can be declared on top of function types. For instance, the `dialog` function from the previous examples can be updated as following:

```go
type greetingFunc func(string) string

func dialog(name string, f greetingFunc) {
	fmt.Println(f(name))
	fmt.Println("I'm a dialog bot.")
}
```

## Anonymous functions

Another powerful tool that is available thanks to first-class functions support is anonymous functions. Anonymous functions are defined at their point of use, without a name following the `func` keyword. Such functions have access to the variables of the enclosing function.

For example:

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

A call to `fib` declares the variables `n1` and `n2` and returns an anonymous function that, in turn, changes the values of these variables each time the function is called. Nth calls of the anonymous function return the Nth number of the Fibonacci sequence starting from 0. The anonymous inner function has access to the local variables (`n1` and `n2`) of the enclosing function `fib`. This is a great way to have function values keep state between calls. We say that the anonymous function is a closure of the variables `n1` and `n2`. [Closures][closure] are widely used in programming and you might see other languages supporting them.

[closure]: https://en.wikipedia.org/wiki/Closure_(computer_programming)
