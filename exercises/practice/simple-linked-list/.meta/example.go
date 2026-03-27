package simplelinkedlist

import (
	"errors"
	"slices"
)

type element struct {
	data int
	next *element
}

type List struct {
	top  *element
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
	list.top = &element{elem, list.top}
	list.size++
}

func (list *List) Pop() (int, error) {
	if list.size == 0 {
		return 0, errors.New("list is empty")
	}
	elem := list.top.data
	list.top = list.top.next
	list.size--
	return elem, nil
}

func (list *List) Peek() (int, error) {
	if list.size == 0 {
		return 0, errors.New("list is empty")
	}
	return list.top.data, nil
}

func (list *List) values() []int {
	res := []int{}
	for node := list.top; node != nil; node = node.next {
		res = append(res, node.data)
	}
	return res
}

func (list *List) Array() []int {
	res := list.values()
	slices.Reverse(res)
	return res
}

func (list *List) Reverse() *List {
	return New(list.values())
}
