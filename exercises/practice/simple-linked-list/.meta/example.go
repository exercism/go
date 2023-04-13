package linkedlist

import "errors"

type element struct {
	data int
	next *element
}

type List struct {
	head *element
	size int
}

func New(elements []int) *List {
	var res List
	for _, elem := range elements {
		res.Push(elem)
	}
	return &res
}

func (list *List) Size() int {
	return list.size
}

func (list *List) Push(elem int) {
	newNode := element{elem, nil}
	if list.head == nil {
		list.head = &newNode
	} else {
		node := list.head
		for node.next != nil {
			node = node.next
		}
		node.next = &newNode
	}
	list.size++
}

func (list *List) Pop() (int, error) {
	var elem int
	switch list.size {
	case 0:
		return 0, errors.New("Cannot call pop on empty list")
	case 1:
		elem = list.head.data
		list.head = nil
	default:
		node := list.head
		for node.next.next != nil {
			node = node.next
		}
		elem = node.next.data
		node.next = nil
	}
	list.size--
	return elem, nil
}

func (list *List) Array() []int {
	res := []int{}
	for node := list.head; node != nil; node = node.next {
		res = append(res, node.data)
	}
	return res
}

func (list *List) Reverse() *List {
	array, size := list.Array(), list.size
	for i := 0; i < size/2; i++ {
		array[i], array[size-1-i] = array[size-1-i], array[i]
	}
	return New(array)
}
