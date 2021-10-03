## General syntax

The for loop is one of the most commonly used statements to repeatedly execute some logic. The basic `for` loop has three components separated by semicolons:

- the init statement: executed before the first iteration

- the condition expression: evaluated before every iteration

- the post statement: executed at the end of every iteration

```go

for initialization; condition; step {

  // code that is executed repeatedly as long as the condition is true

}



```

**Note:** Unlike other languages like C, Java, or JavaScript there are no parentheses surrounding the three components of the `for` statement and the braces `{ }` are always required.

## Header

The initialization usually sets up a counter variable, the condition checks whether the loop should be continued or stopped and the step increments the counter at the end of each repetition.

The individual parts of the header are separated by semicolons.

```go

for i := 1; i < 10; i++ {

  fmt.Println(i)

}

```

Defining the step is often done using the increment or decrement operator as shown in the example above.
