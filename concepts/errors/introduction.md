An `error` is a built-in interface type in the Go language.

In other languages, whether or not a function will throw an exception is unknown.
Errors in Go are not treated as exceptions: they are _values_ that are returned together with the function's result.

```go
func doStuff(a, b int) (int, error) {
    ...
}
```

If a function can fail for some reason, returning the error to the caller indicates there was a problem,
otherwise returning an error's zero value `nil` represents no error.

```go
result, err := doStuff(a, b)
if err != nil {
    // handle the error
}
...
```

The above is a very common snippet to come across.
Returning the error together with the results makes it harder to introduce programming issues in your code.

TODO: finish up
