# Keep track of the size

```go
package linkedlist

import "errors"

type Element struct {
	data int
	next *Element
}

type List struct {
	head *Element
	size int
}

func New(numbers []int) *List {
	output := &List{}

	for _, number := range numbers {
		output.Push(number)
	}

	return output
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Push(element int) {
	l.head = &Element{element, l.head}
	l.size++
}

func (l *List) Pop() (int, error) {
	if l.size < 1 {
		return 0, errors.New("no elements")
	}
	deadHead := l.head
	l.head = deadHead.next
	deadHead.next = nil
	l.size--
	return deadHead.data, nil
}

func (l *List) Array() []int {
	output := make([]int, l.size)
	for i, head := l.size-1, l.head; i > -1; i, head = i-1, head.next {
		output[i] = head.data
	}
	return output
}

func (l *List) Reverse() *List {
	output := &List{}
	for head := l.head; head != nil; head = head.next {
		output.Push(head.data)
	}
	return output
}
```

This approach starts by importing from packages for what is needed.

The `Element` struct is defined, and then the `List` struct is defined with a `size` field.

The `New` function uses the `Push` function to create a new `List` from a slice of integers.

The `Size` [pointer receiver method][pointer-receiver] returns the value of the `size` field.

The `Push` method increments the `size` field for each element that is added.

The `Pop` method decrements the `size` field for each element that is removed.

The `Array` method makes use of the `size` field without having to calculate the size dynamically.

The `Reverse` method does not need to read or change the size of the `List`.

To keep track of the size as you go may seem wasteful if the size is never requested.
On the other hand, if the the size is requested more than once on an unchanged list, it may seem wasteful to calculate the same size
multiple times.

[pointer-receiver]: https://go.dev/tour/methods/4
