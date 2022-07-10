# Introduction

### Panic

Panic is a built-in function that stops the ordinary flow of control and begins panicking. 
When the function F calls panic, execution of F stops, any deferred functions in F are executed normally, 
and then F returns to its caller. To the caller, F then behaves like a call to panic. 
The process continues up the stack until all functions in the current goroutine have returned, 
at which point the program crashes. Panics can be initiated by invoking panic directly. 
They can also be caused by runtime errors, such as out-of-bounds array accesses.

```
func f() {
    testList := make([]string, 2)
    fmt.Println(testList[3])
}
```

Above function will Panic because we are accessing an element that is out of bounds.

### Defer

Your program may have resources that it must clean up properly, even while a panic is being processed by the runtime. Go allows you to defer the execution of a function call until its calling function has completed execution. Deferred functions run even in the presence of a panic, and are used as a safety mechanism to guard against the chaotic nature of panics. Functions are deferred by calling them as usual, then prefixing the entire statement with the defer keyword, as in below example. Run this example to see how a message can be printed even though a panic was produced:

```
func main() {
	defer func() {
		fmt.Println("hello from the deferred function!")
	}()

	panic("oh no!")
}
```

### Recover

Recover is a built-in function that regains control of a panicking goroutine. 
Recover is only useful inside deferred functions. During normal execution, a call to recover will return nil and have no other effect. If the current goroutine is panicking, a call to recover will capture the value given to panic and resume normal execution.


```
func f() {
    testList := make([]string, 2)
    fmt.Println(testList[3])
}

func Rec(){
    if rec := Recover(); rec != nil {
        fmt.Println("Recovered in Rec()")
    }
}
```

Here in the above `Rec()` function we are recoving the the panic created in `f()` function.