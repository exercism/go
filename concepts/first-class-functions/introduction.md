# Introduction

Go supports functions as first-class values. It means a function can be an argument to other functions, returned by another function, and assigned to variables.

It is possible thanks to the function types in Go. A function type denotes the set of all functions with the same sequence of parameter types and the same sequence of result types. User-defined types can be declared on top of function types.

For example:
```go
import "fmt"

type speakingFunc func(string) string

func engGreeting(name string) string { return fmt.Sprintf("Hello %s, nice to meet you!", name) }
func engFarewell(name string) string { return fmt.Sprintf("Good bye %s :)", name) }
func espGreeting(name string) string { return fmt.Sprintf("¡Hola %s, mucho gusto!", name) }
func espFarewell(name string) string { return fmt.Sprintf("Adios %s :)", name) }

var speakingFuncs = map[string][2]speakingFunc{
	"eng": {engGreeting, engFarewell},
	"esp": {espGreeting, espFarewell},
}

func speakingFuncsLookup(language string) (greeting, farewell speakingFunc) {
  if funcs, ok := speakingFuncs[language]; ok {
    greeting, farewell = funcs[0], funcs[1]
  }
  return
}

func speakTo(name string, f speakingFunc) {
  phrase := f(name)
  fmt.Println(phrase)
}

func Greeting(name, language string) {
  gf, ff := speakingFuncsLookup(language)
  if gf == nil || ff == nil {
		fmt.Printf("unsupported language %s\n", language)
		return
  }
  speakTo(name, gf)
  speakTo(name, ff)
}

Greeting("Alice", "eng")
// Output:
// Hello Alice, nice to meet you!
// Good bye Alice :)

Greeting("Alice", "abc")
// Output: unsupported language abc
```

The value of an uninitialized variable of function type is `nil`. Therefore, calling a `nil` function value causes a panic. Function values can be compared with `nil` and it is used to avoid unnecessary program panics. But functional values are not comparable against each other.
```go
var dutchGreeting func(string) string
dutchGreeting("Alice") // panic: call of nil function
```
