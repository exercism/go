# Introduction

### Panic

Panic is a built-in function that stops the ordinary flow of control and begins panicking. 
When the function F calls panic, execution of F stops, any deferred functions in F are executed normally, 
and then F returns to its caller. To the caller, F then behaves like a call to panic. 
The process continues up the stack until all functions in the current goroutine have returned, 
at which point the program crashes. Panics can be initiated by invoking panic directly. 
They can also be caused by runtime errors, such as out-of-bounds array accesses.

### Recover

Recover is a built-in function that regains control of a panicking goroutine. 
Recover is only useful inside deferred functions. During normal execution, a call to recover will return nil and have no other effect. If the current goroutine is panicking, a call to recover will capture the value given to panic and resume normal execution.
