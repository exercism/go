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

You should have a function `NewList()` that creates and returns a `List`:

* `NewList(args ...interface{}) *List`: creates a new linked list preserving the order of the values.

Your `List` should have the following methods:

* `First()`: pointer to the first node (head).
* `Last()`: pointer to the last node (tail).
* `PushBack(v interface{})`: insert value at back.
* `PopBack() (interface{}, error)`: remove value at back.
* `PushFront(v interface{}) `: insert value at front.
* `PopFront() (interface{}, error)`: remove value at front.
* `Reverse()`: reverse the linked list.
