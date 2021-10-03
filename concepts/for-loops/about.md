# Introduction

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

Also, the init and post statements are optional.

```go
i := 1

for i <= 3 {

  fmt.Println(i)
  i = i + 1

}
```

## For is Go's "while"

At that point you can drop the semicolons: C's `while` is spelled `for` in Go.

```go
var sum int = 1

for sum < 1000 {

	sum += sum

}
fmt.Println(sum)
```

## Break, Continue and Labels

Inside a loop body you can use the `break` keyword to stop the execution of the loop entirely.

```go
for n := 0; n <= 5; n++ {

  if n == 3 {
    break
  }

  fmt.Println(n)
}
```

In contrast, the keyword `continue` only stops the execution of the current iteration and continues with the next one.

```go
for n := 0; n <= 5; n++ {

  if n%2 == 0 {
    continue
  }

  fmt.Println(n)
}
```

## Infinite for loop

If you forget or omit the loop condition it loops forever.

```go
for {

  // code

}
```
