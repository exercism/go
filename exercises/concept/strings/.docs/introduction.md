A `string` in Go is an immutable sequence of bytes. Strings may contain arbitrary bytes but usually they contain human-readable text.
Text strings are conventionally interpreted as UTF-8 encoded sequence of Unicode code points (runes) which will be explained in a future exercise.
A `string` value can be written as a string literal, which is a sequence of bytes enclosed in double quotes:

```go
s := "Hello World!"
```

Go provides an in-built package called `fmt` (format package) which offers a variaty of functions to manipulate the format of input and output. The most common used functions to print output are `Println` and `Printf`.
`Println` simply prints the input received on the console screen while `Printf` formats the input, using verbs like `%s` for strings, before printing it on the console.

In Go floating point values are conveniently printed with Printf's verbs: `%g` (compact representation), `%e` (exponent) or `%f` (non exponent). All three verbs allow the field's width and numeric position to be controlled.

```go
f := 4.3242
fmt.Printf("%.4", f)
// Output: 4.32
```
