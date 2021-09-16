# Introduction

A method is a function with a special _receiver_ argument. The receiver appears in its own argument list between `func` keyword and the name of the method.

```go
func (receiver type) MethodName(parameters) (returnTypes) {

}
```

You can only define a method with a receiver whose type is defined in the same package as the method. The receiver can be either a struct type or a non-struct type as in the example below.

```go
type Name string

func (s Name) Greetings() string {
	return fmt.Sprintf("Welcome %s !", s)
}


s := Name("Bronson")
fmt.Println(s.Greetings())
// Output: Welcome Bronson !
```

Notice the way we called the method Greetings() on the Name instance s. It’s exactly like the way you call methods in an object-oriented programming language.

Remember: a method is just a function with a receiver argument. Methods help to avoid naming conflicts - since a method is tied to a particular reciever type, you can have the same method name on different types.

```go
import "math"

type rect struct {
	width, hight int
}
func (r rect) area() int {
	return r.width * r.height
}

type circle struct {
	radius int
}
func (c circle) area() float64 {
	return 2*c.radius*math.Pi
}
```

All the methods we have seen so far have a value receiver which means they on a copy of the value passed to the method, meaning that any modification done to the receiver inside the method is not visible to the caller.

You can declare methods with pointer receivers in order to modify the value to which the receiver points. This is done by prefixing the type name with a `*`, for example with the `rect` type, a pointer reciever would be declared as `*rect`. Such modifications are visible to the caller or the method as well.

```go
type rect struct {
	width, height int
}
func (r *rect) squareIt() {
	r.height = r.width
}

r := rect{width: 10, height: 20}
fmt.Printf("Width: %d, Height: %d\n", r.width, r.height)
// Output: Width: 10, Height: 20

r.squareIt()
fmt.Printf("Width: %d, Height: %d\n", r.width, r.height)
// Output: Width: 10, Height: 10
```
