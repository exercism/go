# Instructions append

## Implementation

You will write an implementation of a generic doubly linked list.
Implement a `Node` to hold a value and pointers to the next and previous nodes.
Then implement a `List` which holds references to the first and last node and offers functions for adding and removing items.

Using a `comparable` generic allows the list to hold any comparable data type.

Your `Node` should have the following fields and methods:

- `Value`: the node's value.
- `Next() *Node`: pointer to the next node.
- `Prev() *Node`: pointer to the previous node.

You should have a function `NewList()` that creates and returns a `List`:

- `NewList[T comparable](args ...any) *List[T]`: creates a new linked list preserving the order of the values.

Your `List[T]` should have the following methods:

- `First() *Node[T]`: returns a pointer to the first node (head).
- `Last() *Node[T]`: returns a pointer to the last node (tail).
- `Push(v T)`: insert value at the back of the list.
- `Pop() (T, error)`: remove value from the back of the list.
- `Unshift(v T)`: insert value at the front of the list.
- `Shift() (T, error)`: remove value from the front of the list.
- `Reverse()`: reverse the linked list.
