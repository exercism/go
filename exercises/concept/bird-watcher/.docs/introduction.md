# Introduction

## General syntax

The for loop is one of the most commonly used statements to repeatedly execute some logic. In Go it consists of the `for` keyword, a header and a code block that contains the body of the loop wrapped in curly brackets. The header consists of 3 components separated by semicolons `;`: initilization, condition and step.

```go
for initialization; condition; step {
  // loop body - code that is executed repeatedly as long as the condition is true
}
```

- The **initialization** component is some code that runs only once before the loop starts.
- The **condition** component must be some expression that evaluates to a boolean and controls when the loop should stop. The code inside the loop will run as long as this condition evaluates to true. As soon as this expression evaluates to false, no more iterations of the loop will run.
- The **step** component is some code that will run at the end of each iteration.

**Note:** Unlike other languages, there are no parentheses `()` surrounding the three components of the header. In fact, inserting such parenthesis is a compilation error. However, the braces `{ }` surrounding the loop body are always required.

## For Loops - An example

The initialization component usually sets up a counter variable, the condition checks whether the loop should be continued or stopped and the step increments the counter at the end of each repetition.

```go
for i := 1; i < 10; i++ {
  fmt.Println(i)
}
```

This loop will print the numbers from `1` to `9` inclusivé. 
Defining the step is often done using an increment or decrement statement, as shown in the example above.

