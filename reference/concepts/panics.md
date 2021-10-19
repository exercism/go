# Panics

[Panics](https://blog.golang.org/defer-panic-and-recover) can be caused by runtime errors like out-of-bounds array accesses or nil pointer dereferences.
In this case they should be considered bugs and be fixed by the developer.

Panics can also be invoked by the developer with the `panic` function. Since Go handles errors differently, `panic` should
be used only in exceptional cases to crash the program. It should **not** be used to signal an error to the caller.
One valid example is to crash a server _on startup_ if the database is not reachable.

> Panic is a built-in function that stops the ordinary flow of control and begins _panicking_.
> When the function `F` calls `panic`, execution of `F` stops, any deferred functions in `F` are executed normally, and then `F` returns to its caller.
> To the caller, `F` then behaves like a call to `panic`.
> The process continues up the stack until all functions in the current goroutine have returned, at which point the program crashes.

A panic crashes the program unless it is [recovered](https://blog.golang.org/defer-panic-and-recover).
