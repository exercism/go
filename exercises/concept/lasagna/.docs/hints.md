# Hints

## General

- An [integer value][integers] can be defined as one or more consecutive digits.

If you see a `panic:` error when running the tests, this is as you haven't implemented one of the functions (it should say which one) or you have left the boilerplate in place, you need to remove the `panic(...)` line from the supplied code.

## 1. Define the expected oven time in minutes

- You need to define a [constant][constants] and assign it the expected oven time in minutes.

If you see an `undefined: OvenTime` error then double check that you have the constant defined.

If you see an `invalid operation: got != tt.expected (mismatched types float64 and int)` error then you've put a decimal point into the `OvenTime` - remove the decimal so that it'll be an integer.

If you see a `syntax error: non-declaration statement outside function body` error then this is as you forgot the `const` keyword.

If you see an `syntax error: unexpected :=, expecting =` error then this is as you are trying to assign to the constant like a variable; contants are assigned using `=` not `:=`.

## 2. Calculate the remaining oven time in minutes

- You need to define a [function][functions] with a single parameter.
- You have to [explicitly return an integer][return] from a function.
- The function's parameter is an [integer][integers].
- You can [call][calls] one of the other functions you've defined previously.
- You can use the [mathematical operator for subtraction][operators] to subtract values.

## 3. Calculate the preparation time in minutes

- You need to define a [function][functions] with a single parameter.
- You have to [explicitly return an integer][return] from a function.
- The function's parameter is an [integer][integers].
- You can use the [mathematical operator for multiplication][operators] to multiply values.

## 4. Calculate the elapsed working time in minutes

- You need to define a [function][functions] with two parameters.
- You have to [explicitly return an integer][return] from a function.
- The function's parameter is an [integer][integers].
- You can [call][calls] one of the other functions you've defined previously.
- You can use the [mathematical operator for addition][operators] to add values.

[functions]: https://tour.golang.org/basics/4
[return]: https://golang.org/ref/spec#Return_statements
[operators]: https://golang.org/ref/spec#Operators
[integers]: https://golang.org/ref/spec#Integer_literals
[calls]: https://golang.org/ref/spec#Calls
[constants]: https://tour.golang.org/basics/15
