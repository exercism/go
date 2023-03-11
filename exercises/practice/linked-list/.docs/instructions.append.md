# Instructions append

You will write an implementation of a doubly linked list. Implement a
`Node` to hold a value and pointers to the next and previous nodes. Then
implement a `List` which holds references to the first and last node and
offers functions for adding and removing items.

Your `Node` should have the following fields and methods:

* `Value`: the node's value (we will use `interface{}`).
* `Next() *Node`: pointer to the next node.
* `Prev() *Node`: pointer to the previous node.

You should have a function `NewList()` that creates and returns a `List`:

* `NewList(args ...interface{}) *List`: creates a new linked list preserving the order of the values.

Your `List` should have the following methods:

* `First() *Node`: returns a pointer to the first node (head).
* `Last() *Node`: returns a pointer to the last node (tail).
* `Push(v interface{})`: insert value at the back of the list.
* `Pop() (interface{}, error)`: remove value from the back of the list. function returns an error "list is empty" if Pop called on empty list.
* `Unshift(v interface{}) `: insert value at the front of the list.
* `Shift() (interface{}, error)`: remove value from the front of the list. function returns an error "list is empty" if Shift called on empty list.
* `Reverse()`: reverse the linked list.
