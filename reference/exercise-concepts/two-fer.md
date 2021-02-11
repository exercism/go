# Two fer

## Exercise

Given a name, return a string with the message: "One for X, one for me.", where the X can be provided and defaults to "you" if not given.

## Concepts encountered while solving it

### Code formatting

It's the first time most students work with Go, and many have not yet used `gofmt`.

### Linting

It's the first time most students work with Go, and many have not yet used `golint`. Linting already provides some hints
on what (not) to do, e.g. use of variable names, not using else if the the regular condition ends with return, and missing
documentation on exported function (or in the wrong format).

### Comments

It's the first time most students work with Go, and many have not yet been using proper documentation of the package and
the exported function. Most of the time they leave the stub comments in the file.

### Functions

Everything is done via functions. In this case:

- there is no concept of default values for parameters; it's done in the body of the function
- parameters are passed by value, so in the scope of the function they can be altered

### Strings

The parameter passed is a string. For the default value, a check needs to be made if the string is empty.

- how to do an idiomatic check for empty string
- how to concatenate strings, either using a simple `+`, or the `fmt` package

### Variables / Assignment / Scope

The parameters are passed by value, so in the scope of the function they can be altered.

- scoping

### Comparison / Conditionals

The code needs to check whether the parameter is empty.

No else branch is needed. Keeps the code simpler.

### Return value

Value needs to be returned. No naked return.
