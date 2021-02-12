# Introduction

A [method][methods] is a function with a special _receiver_ argument. The receiver appears in its own argument list between `func` keyword and the name of the method.

```go
func (receiver type) MethodName(parameters) (returnTypes){

}
```

You can only define a method with a receiver whose type is defined in the same package as the method. The receiver can be either a struct type or a non-struct type as in the example below.

```go
type Name string

func (s Name) Greetings() string {
	return fmt.Sprintf("Welcome %s !", s)
}

func main() {
	s := Name("Bronson")
	fmt.Println(s.Greetings())
}
// Output: Welcome Bronson !
```

Methods with a value receiver operate on a copy of the value passed to it, meaning that any modification done to the receiver inside the method is not visible to the caller.

You can declare methods with pointer receivers in order to modify the value to which the receiver points. Such modifications are visible to the caller or the method as well.

```go
type rect struct {
	width, height int
}
func (r *rect) squareIt() {
	r.height = r.width
}

func main() {
	r := rect{width: 10, height: 20}
	fmt.Printf("Width: %d, Height: %d\n", r.width, r.height)
	// Output: Width: 10, Height: 20

	r.squareIt()
	fmt.Printf("Width: %d, Height: %d\n", r.width, r.height)
	// Output: Width: 10, Height: 10
}
```

You can find several examples [here][pointers_receivers]. Also checkout this short tutorial about [methods][methods_tutorial].

Remember: a method is just a function with a receiver argument.

[methods]: https://tour.golang.org/methods/1
[pointers_receivers]: https://tour.golang.org/methods/4
[methods_tutorial]: https://www.callicoder.com/golang-methods-tutorial/
