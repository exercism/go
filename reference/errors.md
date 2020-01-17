# Errors and Error Handling

> In Go, error handling is important. The language's design and conventions encourage you to explicitly check for errors
> where they occur (as distinct from the convention in other languages of throwing exceptions and sometimes catching them).
>[Error handling and Go](https://blog.golang.org/error-handling-and-go)

As the quote already hints at: Error handling in Go is different than in most other languages. It actually has almost
nothing in common with the typical catching of exceptions. The difference becomes clear, when looking at how to deal with errors in Go.

## Simple Example

Errors in Go are explicitly returned by every function that can fail. The error should always be the last return argument.
Here for example the function `os.Open` to open a file:

```go
func Open(name string) (file *File, err error)
```

A usage of this function would look like this:

```go
func readFile(file string) ([]byte, error) {
    if file == "" {
        // create a new error
        return nil, errors.New("missing file name")
    }

    f, err := os.Open(file)
    if err != nil {
        // os.Open failed for some reason
        return nil, err
    }
    defer f.Close()

    // the error of ReadAll is returned directly
    return ioutil.ReadAll(f)
}
```

In this example we decided that we don't know what to do if opening the file fails so we just return the error back to
the caller. This is often done in lower level functions.

Additionally we created a function ourselves that returns an error. If the given `file` string is empty we want to catch
that before trying to open the file and return a newly created error.

At some point up the chain we would then decide to log the error or deal with it for example by creating a new empty file instead.
An error should however not be dealt with and returned up the chain at the same time.

Don't do:
```go
    f, err := os.Open(file)
    if err != nil {
        log.Println(err)
        return nil, err
    }
    defer f.Close()
```

## Advanced Error Handling

There is a lot more to errors than can be covered here. The following links will get you started:
- [Error Handling and Go](https://blog.golang.org/error-handling-and-go)
- [Standard errors Package](https://golang.org/pkg/errors/)
- [Don't just check errors, handle them gracefully (Dave Chaney)](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully)
- [Error Package with Stack Traces (Dave Chaney)](https://github.com/pkg/errors)

Note: Errors in Go are currently being worked on. Go 1.13 (current release at writing) has brought some advancements,
and future versions might bring even more standard features.
