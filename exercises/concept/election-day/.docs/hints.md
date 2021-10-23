# Hints

## General

- `*T` can be used to declared variables that are pointers to some type `T`, e.g `var i *int` declares a variable `i` that is a pointer to an `int`
- You can get a pointer for a variable (its memory address) by using the `&` operator, e.g `mypointer := &anIntVariable`.
- You can get the value stored in a pointer by using the `*` operator on a pointer, eg. `var i int = *aPointerToInt`. This is called dereferencing the pointer.
- You check if a pointer is not `nil` before dereferencing it. Attempting to dereference a `nil` pointer will give you a runtime error.
- If you are unsure how pointers work, try reading [Tour of Go: Pointers][go-tour-pointers] or [Go by Example: Pointers][go-by-example-pointers]

## 1. Create a vote counter

- You need to create a pointer to an `int`, in other words, a `*int`.
- You can use the `&` operator on a variable to create a pointer to it, e.g `&myInt`
- You can create a pointer to a new variable defined by you or you can use the variable of the function argument

## 2. Get number of votes from a counter

- You can use the `*` operator on a pointer to dereference it and get its value, e.g `*myPointer`
- Dereferencing `nil` pointers will give you a runtime error. Always make sure a pointer is not `nil` before dereferencing it.

## 3. Increment the votes of a counter

- If you have a pointer `var myPointer *int`, you can assign to `*myPointer` to change the value pointed by `myPointer`
- To get the current value of the pointer, you need to dereference it using the `*` operator, or call the function you made in the previous task.

## 4. Create the election results

- Create a new `ElectionResult` literal or variable with the fields `Name` and `Votes` filled with the values in the arguments of the function.
- You can create a pointer from a variable or literal by using the `&` operator before the variable name/literal declaration, e.g `&myVariable` or `&ElectionResult{Name: "John", Votes: 1}`

## 5. Announce the results

- Although you are receiving a pointer to an `ElectionResult`, you can access its fields with the dot `.` notation, like if it wasn't a pointer!
- Build the message by accessing the `Name` and `Value` fields on the struct.
- Even though you are accessing fields from a pointer to a struct, you don't need to do any dereferencing. Go will automatically dereference the pointer for you, like in this example:

```go
result := &ElectionResult{
    Name: "John",
    Votes: 32
}

result.Name // "John" - Go will automatically dereference the pointer
            //          and access the 'Name' field of the dereferenced struct 
```

## 6. Vote recounting

- You can think of maps as being pointers already. This means that changes you make to the map inside the function will be visible outside the function.
- To increment the value of a key in a `var m map[string]int`, you have several options: `m["mykey"] = m["mykey"] + 1 `, `m["mykey"] += 1 ` or `m["mykey"]++ `

[go-tour-pointers]: https://tour.golang.org/moretypes/1
[go-by-example-pointers]: https://gobyexample.com/pointers

