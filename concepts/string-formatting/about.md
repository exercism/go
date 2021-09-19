# About

Go provides an in-built package called `fmt` (format package) which offers a variety of functions to manipulate the format of input and output. The most common used functions to print output are `Println` and `Printf`.
`Println` simply prints the input received on the console screen while `Printf` formats the input, using verbs like `%s` for strings, before printing it on the console.

In Go floating point values are conveniently printed with Printf's verbs: `%g` (compact representation), `%e` (exponent) or `%f` (non exponent). All three verbs allow the field's width and numeric position to be controlled.

```go
f := 4.3242
fmt.Printf("%.4", f)
// Output: 4.32
```

[string concatenation]: https://golang.org/ref/spec
