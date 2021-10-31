# About

## General syntax

The for loop is one of the most commonly used statements to repeatedly execute some logic. In Go it consists of the `for` keyword, a header and a code block that contains the body of the loop wrapped in curly brackets. The header consists of 3 components separated by semicolons `;`: init, condition and post.

```go
for init; condition; post {
  // loop body - code that is executed repeatedly as long as the condition is true
}
```

- The **init** component is some code that runs only once before the loop starts.
- The **condition** component must be some expression that evaluates to a boolean and controls when the loop should stop. The code inside the loop will run as long as this condition evaluates to true. As soon as this expression evaluates to false, no more iterations of the loop will run.
- The **post** component is some code that will run at the end of each iteration.

**Note:** Unlike other languages, there are no parentheses `()` surrounding the three components of the header. In fact, inserting such parenthesis is a compilation error. However, the braces `{ }` surrounding the loop body are always required.

## For Loops - An example

The init component usually sets up a counter variable, the condition checks whether the loop should be continued or stopped and the post component usually increments the counter at the end of each repetition.

```go
for i := 1; i < 10; i++ {
  fmt.Println(i)
}
```

This loop will print the numbers from `1` to `9` inclusivé. 
Defining the step is often done using an increment or decrement statement, as shown in the example above.

## Optional components of the header

The init and post components of the header are optional:

```go
var sum = 1
for sum < 1000 {
	sum += sum
}
fmt.Println(sum)
// Output: 1024
```

This is similar to a `while` loop in other languages.

In Go, there is no `while` keyword. A `while` loop can be expressed in Go by writing a `for` loop without the init and post parts of the header, like in the example above.

## Break and Continue

Inside a loop body you can use the `break` keyword to stop the execution of the loop entirely:

```go
for n := 0; n <= 5; n++ {
  if n == 3 {
    break
  }
  fmt.Println(n)
}
// Output:
// 0
// 1
// 2
```

In contrast, the keyword `continue` only stops the execution of the current iteration and continues with the next one:

```go
for n := 0; n <= 5; n++ {
  if n%2 == 0 {
    continue
  }
  fmt.Println(n)
}
// Output:
// 1
// 3
// 5
```
## Infinite for loop

The condition part of the loop header is also optional. In fact, you can write a loop with no header: 

```go
for {
  // Endless loop...
}
```

This loop will never end and will only ever finish if the program exits or has a `break` in its body.

## Labels and goto

When we use `break`, Go will stop running the most inner loop. Similarly, when we use `continue`, Go will run the next iteration of the most inner loop.

However, this is not always desirable. We can use labels together with `break` and `continue` to specifcy exactly from which loop we want to exit or continue, respectively.

In this example we are creating a label `OuterLoop`, that will refer to the most outter loop. In the most inner loop, to say that we want exit from the most outter loop, we then use `break` followed by the name of the label of the most outter loop:

```go
OuterLoop:
    for i := 0; i < 10; i++ {
        for j := 0; j < 10; j++ {
            // ...
            break OuterLoop
        }
    }
```

Using labels with `continue` would also work, in which case, Go would continue in the next iteration of the loop referenced by the label.

Go also has a `goto` keyword that works in a similar way and allows us to jump to from a piece of code to another labeled piece of code.

**Warning:** Even though Go allows to jump to a piece of code marked with a label with `continue`, `break` or `goto`, using this feature of the language can easily make the code very hard to read. For this reason, using labels is often not recommended.