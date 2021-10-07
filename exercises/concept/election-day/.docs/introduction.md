# Introduction

Like many other languages, Go has pointers. 
If you're new to pointers, they can feel a little mysterious but once you get used to them, they're quite straight-forward.
They're a crucial part of Go, so take some time to really understand them.

Before digging into the details, it's worth understanding the use of pointers. Pointers are a way to share memory with other parts of our program, which is useful for two major reasons:
1. When we have large amounts of data, making copies to pass between functions is very inefficient. 
  By passing the memory location of where the data is stored instead, we can dramatically reduce the resource-footprint of our programs.
2. By passing pointers between functions, we can access and modify the single copy of the data directly, meaning that any changes made by one function are immediately visible to other parts of the program when the function ends.

## Variables and Memory

Let's say we have a regular integer variable `a`:

```go
var a int
```

When we declare a variable, Go has to find a place in memory to store its value. This is largely abstracted from us — when we need to fetch the value stored in that piece of memory, we can just refer to it by the variable name.

For instance, when we write `a + 2`, we are effectively fetching the value stored in the memory associated with the variable `a` and adding 2 to it.

Similarly, when we need to change the value in the piece of memory of `a`, we can use the variable name to do an assignment:

```go
a = 3
```

The piece of memory that is associated with `a` will now be storing the value `3`.

## Pointers

While variables allow us to refer to values in memory, sometimes it's useful to know the **memory address** to which the variable is pointing. **Pointers** hold the memory addresses of those values. You declare a variable with a pointer type by prefixing the underlying type with an asterisk: 

```go
var p *int // 'p' contains the memory address of an integer
```

Here we declaring a variable `p` of type "pointer to int" (`*int`). This means that `p` will hold the memory address of an integer. The zero value of pointers is `nil` because a `nil` pointer holds no memory address.

### Getting a pointer to a variable

To find the memory address of the value of a variable, we can use the `&` operator. 
For example, if we want to find and store the memory address of variable `a` in the pointer `p`, we can do the following:

```go
var a int
a = 2

var p *int
p = &a // the variable 'p' contains the memory address of 'a'
```

### Accessing the value via a pointer (dereferencing)

When we have a pointer, we might want to know the value stored in the memory address the pointer represents. We can do this using the `*` operator:

```go
var a int 
a = 2

var p *int 
p = &a // the variable 'p' contains the memory address of 'a'

var b int
b = *p // b == 2
```

The operation `*p` fetches the value stored at the memory address stored in `p`. This operation is often called "dereferencing".

We can also use the derefering operator to assign a new value to the memory address referenced by the pointer:

```go
var a int
a = 2   // declare int variable 'a' and assign it the value of 2

var pa *int
pa = &a          // 'pa' now contains to the memory address of 'a'  
*pa = *pa + 2    // increment by 2 the value at memory address 'pa'

fmt.Println(a)   // Output: 4
                 // 'a' will have the new value that was changed via the pointer!
```

Assigning to `*pa` will change the value stored at the memory address `pa` holds. Since `pa` holds the memory address of `a`, by assigning to `*pa` we are effectively changing the value of `a`! 

A note of caution however: always check if a pointer is not `nil` before dereferencing. Dereferencing a `nil` pointer will make the program crash at runtime!

```go
var p *int // p is nil initially
fmt.Println(*p)
// panic: runtime error: invalid memory address or nil pointer dereference
```

### Pointers to structs

So far we've only seen pointers to primitive values. We can also create pointers for structs:

```go
type Person struct {
    Name string
    Age  int
}

var peter Person
peter = Person{Name: "Peter", Age: 22}

var p *Person
p = &peter
```

We could have also created a new `Person` and immediately stored a pointer to it:

```go
var p *Person
p = &Person{Name: "Peter", Age: 22}
```

When we have a pointer to a struct, we don't need to dereference the pointer before accessing one of the fields:

```go
var p *Person
p = &Person{Name: "Peter", Age: 22}

fmt.Println(p.Name) // Output: "Peter" 
                    // Go automatically dereferences 'p' to allow
                    // access to the 'Name' field
```

## Slices and maps are already pointers

Slices and maps are special types because they already have pointers in their implementation. This means that more often that not, we don't need to create pointers for these types to share the memory address for their values. Imagine we have a function that increments the value of a key in a map:


```go
func incrementPeterAge(m map[string]int) {
	m["Peter"] += 1
}
```

If we create a map and call this function, the changes the function made to the map persist after the function ended. This is a similar behavior we get if we were using a pointer, but note how on this example we are not using any referencing/dereferencing or any of the pointer syntax:

```go
ages := map[string]int{
  "Peter": 21
}
incrementPeterAge(ages)
fmt.Println(ages)
// Output: map[Peter:22]
// The changes the function 'addPeterAge' made to the map are visible after the function ends!
```

The same applies when changing an existing item in a slice.

However, actions that return a new slice like `append` are a special case and **might not** modify the slice outside of the function. This is due to the way slices work internally, but we won't cover this in detail in this exercise, as this is a more advanced topic. If you are really curious you can read more about this in [Go Blog: Mechanics of 'append'][mechanics-of-append]

[mechanics-of-append]: https://go.dev/blog/slices
