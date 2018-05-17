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

Your `Node` should have the following fields:

* `Val`: the node's value (we will use `interface{}`).
* `Next`: the next node.
* `Prev`: the previous node.

Your `List` should have the following fields and functions:

* `Head`: the first element in the linked list.
* `Tail`: the last element in the linked list.
* `PushBack(v interface{})`: *insert value at back*.
* `PopBack() (interface{}, error)`: *remove value at back*.
* `PushFront(v interface{}) `: *remove value at front*.
* `PopFront() (interface{}, error)`: *insert value at front*.
* `Reverse()`: reverse the linked list (Head, Tail = Tail, Head).

Instead of not covering error conditions, like calling `PopBack` or `PopFront` on an empty list,
we will follow Go's idiomatic style and implement the error checks as well.

If you want to know more about linked lists, check [Wikipedia](https://en.wikipedia.org/wiki/Linked_list).
