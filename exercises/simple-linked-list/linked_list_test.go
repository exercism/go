// API:
//
// type Element struct {
//  data int
//  next *Element
// }
//
// type List struct {
//  head *Element
//  size int
// }
//
// func New([]int) *List
// func (*List) Size() int
// func (*List) Push(int)
// func (*List) Pop() (int, error)
// func (*List) Array() []int
// func (*List) Reverse() *List

package linkedlist

import (
	"reflect"
	"testing"
)

var array1To10 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func TestEmptyListHasSizeZero(t *testing.T) {
	list := New([]int{})
	if size := list.Size(); size != 0 {
		t.Fatalf("Size of empty list: %d, expected: %d", size, 0)
	}
	list = New(nil)
	if size := list.Size(); size != 0 {
		t.Fatalf("Size of empty list: %d, expected: %d", size, 0)
	}
}

func TestSingletonListHasSizeOne(t *testing.T) {
	list := New([]int{1})
	if size := list.Size(); size != 1 {
		t.Fatalf("Size of singleton list: %d, expected: %d", size, 1)
	}
}

func TestNonEmptyListHasCorrectSize(t *testing.T) {
	list := New([]int{1, 2, 3})
	if size := list.Size(); size != 3 {
		t.Fatalf("Size of list from [1, 2, 3]: %d, expected: %d", size, 3)
	}
}

func TestEmptyListToEmptyArray(t *testing.T) {
	list := New([]int{})
	if array := list.Array(); len(array) != 0 {
		t.Fatalf("Test empty list to array: %v, want empty array", array)
	}
	list = New(nil)
	if array := list.Array(); len(array) != 0 {
		t.Fatalf("Test empty list to array: %v, want empty array", array)
	}
}

func TestNonEmptyListToArray(t *testing.T) {
	expected := []int{1, 2, 3}
	list := New(expected)
	array := list.Array()
	if !reflect.DeepEqual(array, expected) {
		t.Fatalf("Test non empty list to array: %v, want %v", array, expected)
	}
}

func TestPopFromEmptyList(t *testing.T) {
	list := New([]int{})
	if _, err := list.Pop(); err == nil {
		t.Fatalf("Pop from empty list: expected error but there was not")
	}
	list = New(nil)
	if _, err := list.Pop(); err == nil {
		t.Fatalf("Pop from empty list: expected error but there was not")
	}
}

func TestPopFromNonEmptyList(t *testing.T) {
	list := New([]int{1, 2, 3})
	elem, err := list.Pop()
	if err != nil {
		t.Fatalf("Pop from non empty list: unexpected error %v", err)
	}
	if elem != 3 {
		t.Fatalf("Pop from non empty list: %d, want %d", elem, 3)
	}
	actual := list.Array()
	expected := []int{1, 2}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Pop from non empty list: %v, want %v", actual, expected)
	}
}

func TestPushToEmptyList(t *testing.T) {
	list := New([]int{})
	list.Push(1)
	actual := list.Array()
	expected := []int{1}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Push to empty list: %v, want %v", actual, expected)
	}
	list = New(nil)
	list.Push(1)
	actual = list.Array()
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Push to empty list: %v, want %v", actual, expected)
	}
}

func TestPushToNonEmptyList(t *testing.T) {
	list := New([]int{1, 2, 3})
	list.Push(4)
	actual := list.Array()
	expected := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Push to non empty list: %v, want %v", actual, expected)
	}
}

func TestPushAndPop(t *testing.T) {
	list := New([]int{1, 2, 3})
	list.Pop()
	list.Push(4)
	list.Push(5)
	list.Pop()
	list.Push(6)
	actual := list.Array()
	expected := []int{1, 2, 4, 6}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Test push and pop: %v, want %v", actual, expected)
	}
}

func TestReverseEmptyList(t *testing.T) {
	list := New([]int{})
	if reversed := list.Reverse().Array(); len(reversed) != 0 {
		t.Fatalf("Reverse empty list: %v, want empty list", reversed)
	}
	list = New(nil)
	if reversed := list.Reverse().Array(); len(reversed) != 0 {
		t.Fatalf("Reverse empty list: %v, want empty list", reversed)
	}
}

func TestReverseNonEmptyList(t *testing.T) {
	list := New([]int{1, 2, 3})
	actual := list.Reverse().Array()
	expected := []int{3, 2, 1}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Reverse non empty list: %v, want %v", actual, expected)
	}
}

func BenchmarkNewList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = New(array1To10)
	}
}

func BenchmarkListSize(b *testing.B) {
	list := New(array1To10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = list.Size()
	}
}

func BenchmarkListPush(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		list := New([]int{})
		b.StartTimer()
		for k := 0; k < 1000; k++ {
			list.Push(k)
		}
	}
}

func BenchmarkListPop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		list := New([]int{})
		for k := 0; k < 1000; k++ {
			list.Push(k)
		}
		b.StartTimer()
		for k := 0; k < 1000; k++ {
			_, _ = list.Pop()
		}
	}
}

func BenchmarkListToArray(b *testing.B) {
	list := New(array1To10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = list.Array()
	}
}

func BenchmarkListReverse(b *testing.B) {
	list := New(array1To10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = list.Reverse()
	}
}
