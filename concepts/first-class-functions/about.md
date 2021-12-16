# About

Go supports functions as first-class values. It means a function can be an argument to other functions, returned by another function, and assigned to variables.

For example:
```go
import "fmt"

func englishGreeting(name string) string {
  return fmt.Sprintf("Hello %s, nice to meet you!", name)
}

func spanishGreeting(name string) string {
  return fmt.Sprintf("Hola %s, mucho gusto!", name)
}

greeting := englishGreeting
fmt.Println(greeting("Sasha")) // "Hello Sasha, nice to meet you!"

greeting = spanishGreeting
fmt.Println(greeting("Sasha")) // "Hola Sasha, mucho gusto!"
```

## Function types
Like other values, function values have types. A function type denotes the set of all functions with the same sequence of parameter types and the same sequence of result types. User-defined types can be declared on top of function types.
The value of an uninitialized variable of function type is `nil`. Therefore, calling a `nil` function value cases a panic.
Function values can be compared with `nil`, but functional values are not comparable against each other.

```go
var dutchGreeting func(string) string
dutchGreeting("Sasha") // panic: call of nil function

type greetingFunc func(string) string

func greet(name string, f greetingFunc) string {
  if f != nil {
    return f(name)
  }

  return ""
}

func GreetIn(language, name string) string {
  var gf greetingFunc
  switch language: {
    case "eng", "ENG":
      gf = englishGreeting
    case "esp", "ESP":
      gf = spanishGreeting
  }

  return Greeting(name, gf)
}
```