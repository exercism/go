# About

## Adding context to errors

We explored basic error handling in Go in the [errors concept][concept-errors].
As you learned there, by default errors do not carry around stack traces.
That makes it crucial to ensure the error itself contains enough information to identify the problem.

If we wanted to add information to an existing error with the tools we already know, we could write something like this to create a new error with the combined text of the original error and the additional information:

```go
err := errors.New(fmt.Sprintf("parsing age for user %d: %v", userID, originalError))
```

Luckily, the `fmt` package from the standard library contains a short-hand version for this in form of the `Errorf` function.
That allows you, for example, to easily add information about the context in which the error occurred.

```go
originalError := errors.New("unexpected negative number")
userID := 123
err := fmt.Errorf("parsing age for user %d: %v", userID, originalError)
err.Error()
// => "parsing age for user 123: unexpected negative number"
```

Unless you are sure that the error already has enough information included, you should try to get into the habit of always adding context to an error before returning it from a function.
That means your standard error handling pattern should look like this:

```go
result, err := calculateSomething()
if err != nil {
  return 0, fmt.Errorf("calculating something: %v", err)
}
```

As shown in the examples, the convention is to separate the error from the added message with a colon.
Just like the message of the original error, also the context message should not be capitalized and no period should be added at the end.

~~~~exercism/note
When adding context to an error, it is tempting to start the message with "failed to ..." or similar.
However, as an error can accumulate a lot of additional context as it travels up the function chain, this can make the final message that will be logged or the user will see harder to read/ full of duplication.

Here is an example to illustrate:

```text
failed to validate request: failed to parse path parameter: failed to convert number: invalid input "abc"
vs.
validating request: parsing path parameter: converting to number: invalid input "abc"
```

So keep this end result in mind when writing the context messages and refrain from adding text about the fact that something did not go as expected.
This is usually clear from the context (e.g. log level "ERROR") and the original error message.
~~~~

Often this way of adding information for an error is good enough but there are cases where you want to allow the consumer of your error to check for or retrieve the original underlying error.
Adding context in way that allows this is called "wrapping an error" in Go.

## Wrapping errors and checking for sentinel errors

Error wrapping can be achieved with a very minor change to the code we saw above.
To wrap an error, you need to use the formatting verb `%w` instead of `%v`.
Behind the scenes, this will make sure that the resulting error implements an `Unwrap` method which returns the original error.
Because of that, then `errors.Is` can be used to check whether a specific sentinel error appears somewhere along the "error chain".
It does that by secretly calling `Unwrap` repeatedly until the error in question was found or the chain ended.

```go
originalError := errors.New("unexpected negative number")
err := fmt.Errorf("parsing age: %w", originalError)
errors.Is(err, originalError)
// => true
```

Checking for the original error with `errors.Is` would not work with the regular `%v` formatting verb.

```go
originalError := errors.New("unexpected negative number")
err := fmt.Errorf("parsing age: %v", originalError)
errors.Is(err, originalError)
// => false
```

As a result, it is good practice to use `%v` by default and only use `%w` if you explicitly want to allow your consumer to access an underlying error ([Google Go Styleguide][google-go-styleguide]).

If you find ourself in a situation where you want to check for a sentinel error but explicitly don't want to unwrap, you can use `==` instead of `errors.Is`.

```go
var originalError = errors.New("unexpected negative number")
func someFunc() error {
  return originalError
}

err := someFunc()
err == originalError 
// => true
```

In contrast, `==` would not identify the error if it was wrapped.

```go
var originalError = errors.New("unexpected negative number")
func someFunc() error {
  return fmt.Errorf("parsing age: %v", originalError)
}

err := someFunc()
err == originalError 
// => false
```

It is fine to work with `errors.Is` by default, but you should always be aware that this means the whole error chain will be searched.

## Checking for (custom) error types

There is an equivalent to `errors.Is` that allows to check for and retrieve an error of a specific type from the error chain.
The function for that is `errors.As` and just like `errors.Is` it will search for the given error type along the whole chain.
`errors.As` does only then extract the error it found that matches the type so you can further work with it, e.g. retrieve specific fields.

```go
type MyCustomError struct {
  Message string
  Details string
}

func (e *MyCustomError) Error() string {
  return fmt.Sprintf("%s, details: %s", e.Message, e.Details)
}

func someFunc() error {
  originalError := &MyCustomError{
    Message: "some message",
    Details: "some details",
  }

  return fmt.Errorf("doing something: %w", originalError)
}

err := someFunc()
var customError *MyCustomError
errors.As(err, &customError)
// => true

// customError now contains the error that was found in the error chain.
customError.Details
// => "some details"
```

~~~~exercism/caution
Be careful with the syntax regarding the pointers above.
The code will only work and compile correctly if `customError` has the exact type that implements the error interface.
In our case, that is `*MyCustomError` (a [pointer][concept-pointers] to `MyCustomError`), not `MyCustomError` itself.

On top of that, the second argument needs to be a pointer to the error variable.
Since our error is already a pointer, what we are passing to `errors.As` is a pointer to a pointer (to MyCustomError).
Only with this set up correctly, Go can then fill the variable with the error it found.
~~~~

As before, `errors.As` would not have found the error type if `%v` would have been used when calling `Errorf`.

If you don't want to unwrap for some reason, type assertion can be used instead (equivalent to `==` above).

```go
// MyCustomError defined as above.

func someFunc() error {
  return &MyCustomError{
    Message: "some message",
    Details: "some details",
  }
}

err := someFunc()
customError, ok := err.(*CustomError)
// "ok" is now true
customError.Details
// => "some details"
```

Type assertion will not be able to identify the error type if the error would have been wrapped.

## Allowing errors of custom types to be unwrapped

Sometimes just wrapping an error with some additional text is not enough.
You can create a custom error type instead that holds the original error and the additional structured data that you want to add.
If you want to allow unwrapping for your error type, the only thing you have to do is to manually add an `Unwrap() error` method so the `Unwrap` interface is satisfied.

```go
type SpecialError struct {
  originalError error
  metadata string
}

func (e *SpecialError) Error() string {
  // The usual serialization code goes here.
}

func (e *SpecialError) Unwrap() error {
  return e.originalError
}
```

## Combining multiple errors

In Go 1.20, a new `Join` function was added to the built-in `errors` package.
`Join` allows it to combine multiple errors together in a way that still allows unwrapping, i.e. checking for a specific error or error type along the chain.
With that, the error chain can actually be an error tree in reality.
More information can be found in the [release notes][release-notes] and the [documentation][doc-join].

[concept-errors]: /tracks/go/concepts/errors
[concept-pointers]: /tracks/go/concepts/pointers
[google-go-styleguide]: https://google.github.io/styleguide/go/best-practices#adding-information-to-errors
[release-notes]: https://tip.golang.org/doc/go1.20#errors
[doc-join]: https://pkg.go.dev/errors#Join