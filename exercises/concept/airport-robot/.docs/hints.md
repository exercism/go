# Hints

## General

- Maybe have a look at [Interfaces in the Tour of Go][interfaces-tour-of-go] to see more examples if you struggle with the exercise.

## 1. Create the abstract greeting functionality

- Look back at `Counter` example the introduction to find out how to define an interface.
- Revisit the [functions concepts][concept-functions] to recap how to write function signatures in Go.
- The abstract `Greeter` type can be used in a function signature the same way as a normal concrete type like `string`.
- To implement `SayHello`, call the methods on the argument of type `Greeter`.
  Then use string formatting or string concatenation to construct the final result.

## 2. Implement Italian

- Revisit the [structs concept][concept-structs] to see how to define a struct type.
- To solve the task, the struct does not need any fields at all.
- Once you defined the struct, you want to add the methods `LanguageName` and `Greet` to it.
  Revisit the [methods concept][concept-methods] to find out how to add a method for a type.

## 3. Implement Portuguese

- See hints for task 2.

[interfaces-tour-of-go]: https://go.dev/tour/methods/9
[concept-functions]: /tracks/go/concepts/functions
[concept-structs]: /tracks/go/concepts/structs
[concept-methods]: /tracks/go/concepts/methods
