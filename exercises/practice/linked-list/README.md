# Linked List

Implement a doubly linked list.

Like an array, a linked list is a simple linear data structure. Several
common data types can be implemented using linked lists, like queues,
stacks, and associative arrays.

A linked list is a collection of data elements called *nodes*. In a
*singly linked list* each node holds a value and a link to the next node.
In a *doubly linked list* each node also holds a link to the previous
node.

You will write an implementation of a doubly linked list. Implement a
`Node` to hold a value and pointers to the next and previous nodes. Then
implement a `List` which holds references to the first and last node and
offers functions for adding and removing items.

Your `Node` should have the following fields and functions:

* `Val`: the node's value (we will use `interface{}`).
* `Next()`: pointer to the next node.
* `Prev()`: pointer to the previous node.
* `First()`: pointer to the first node (head).
* `Last()`: pointer to the last node (tail).

Your `List` should have the following fields and functions:

* `NewList(args ...interface{}) *List`: creates a new linked list preserving the order of the values.
* `First()`: pointer to the first node (head).
* `Last()`: pointer to the last node (tail).
* `PushBack(v interface{})`: insert value at back.
* `PopBack() (interface{}, error)`: remove value at back.
* `PushFront(v interface{}) `: insert value at front.
* `PopFront() (interface{}, error)`: remove value at front.
* `Reverse()`: reverse the linked list.

Instead of not covering error conditions, like calling `PopBack` or `PopFront` on an empty list,
we will follow Go's idiomatic style and implement the error checks as well.

If you want to know more about linked lists, check [Wikipedia](https://en.wikipedia.org/wiki/Linked_list).

## Coding the solution

Look for a stub file having the name linked_list.go
and place your solution code in that file.

## Running the tests

To run the tests run the command `go test` from within the exercise directory.

If the test suite contains benchmarks, you can run these with the `--bench` and `--benchmem`
flags:

    go test -v --bench . --benchmem

Keep in mind that each reviewer will run benchmarks on a different machine, with
different specs, so the results from these benchmark tests may vary.

## Further information

For more detailed information about the Go track, including how to get help if
you're having trouble, please visit the exercism.io [Go language page](http://exercism.io/languages/go/resources).

## Source

Classic computer science topic

## Submitting Incomplete Solutions
It's possible to submit an incomplete solution so you can see how others have completed the exercise.
