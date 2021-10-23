# Debug

When a test fails, a message is displayed describing what went wrong and for which input. You can also use the fact that [console output][fmt-println] will be shown too. You can write to the console using:

```go
import "fmt"
fmt.Println("Debug message")
```

[fmt-println]: https://pkg.go.dev/fmt#Println

Note: When debugging with the in-browser editor, using e.g. `fmt.Print` will not add a newline after the output which can provoke a bug in `go test --json` and result in messed up test output.
