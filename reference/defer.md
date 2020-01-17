# Defer

In Go the `defer` statement defers the execution of a function until the surrounding function returns.
The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
This is a very simple, yet very powerful concept enabling clean code and preventing a lot of bugs from happening in the first place.

### Without Defer

In this example all data is being read from a file with a typical read loop. Then the file needs to be closed.
In this case that means calling `f.Close` in 3 different places.

```go
func readFile(file string) ([]byte, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	var (
		n   int
		b   = make([]byte, 100)
		res []byte
	)
	// read the entire file into res
	for {
		n, err = f.Read(b)
		res = append(res, b[:n]...)
		if err != nil {
			break
		}
	}
	// ignore end-of-file error
	if err != io.EOF {
		f.Close()
		return nil, err
	}
	// create error if file was empty
	if len(res) == 0 {
		f.Close()
		return nil, errors.New("file is empty")
	}
	f.Close()
	return res, nil
}
```

If the function needs extending possibly with another case that returns early,
the developer has to remember to call `f.Close` in that case as well or the file descriptor will not be closed.
This is a typical example where `defer` makes the code cleaner and easier to maintain, thereby preventing future bugs.

### Using Defer

The defer statement is used right after the file was opened successfully. First the error needs to be checked in case
`os.Open` failed and then `defer f.Close()` is used to make sure the file is closed when the function returns -
no matter where it returns from that point on.

```go
func readFile(file string) ([]byte, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var (
		n   int
		b   = make([]byte, 100)
		res []byte
	)
	// read the entire file into res
	for {
		n, err = f.Read(b)
		res = append(res, b[:n]...)
		if err != nil {
			break
		}
	}
	// ignore end-of-file error
	if err != io.EOF {
		return nil, err
	}
	// create error if file was empty
	if len(res) == 0 {
		return nil, errors.New("file is empty")
	}
	return res, nil
}
```

Extending this function with more functionality does not require to think about closing the file any more.
That has been taken care of.

### Pro Tip
The trick of using `defer` successfully lies in scoping functions accordingly.
If defer does not fit in, restructuring should be considered.
