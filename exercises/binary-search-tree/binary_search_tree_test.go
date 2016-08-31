// API:
//
// type SearchTreeData struct {
//	left  *SearchTreeData
//	data  int
//	right *SearchTreeData
// }
//
// func Bst(int) *SearchTreeData
// func (*SearchTreeData) Insert(int)
// func (*SearchTreeData) MapString(func(int) string) []string
// func (*SearchTreeData) MapInt(func(int) int) []int

package binarysearchtree

import (
	"reflect"
	"strconv"
	"testing"
)

func TestDataIsRetained(t *testing.T) {
	actual := Bst(4).data
	expected := 4
	if actual != expected {
		t.Errorf("Bst(4).data: %d, want %d.", actual, expected)
	}
}

func TestInsertingLess(t *testing.T) {
	bst := SearchTreeData{data: 4}
	bst.Insert(2)

	actual := bst.data
	expected := 4
	if actual != expected {
		t.Errorf("bst.data: %d, want %d.", actual, expected)
	}

	actual = bst.left.data
	expected = 2
	if actual != expected {
		t.Errorf("bst.left.data: %d, want %d.", actual, expected)
	}
}

func TestInsertingSame(t *testing.T) {
	bst := SearchTreeData{data: 4}
	bst.Insert(4)

	actual := bst.data
	expected := 4
	if actual != expected {
		t.Errorf("bst.data: %d, want %d.", actual, expected)
	}

	actual = bst.left.data
	expected = 4
	if actual != expected {
		t.Errorf("bst.left.data: %d, want %d.", actual, expected)
	}
}

func TestInsertingMore(t *testing.T) {
	bst := SearchTreeData{data: 4}
	bst.Insert(5)

	actual := bst.data
	expected := 4
	if actual != expected {
		t.Errorf("bst.data: %d, want %d.", actual, expected)
	}

	actual = bst.right.data
	expected = 5
	if actual != expected {
		t.Errorf("bst.data: %d, want %d.", actual, expected)
	}
}

func TestComplexTree(t *testing.T) {
	bst := SearchTreeData{data: 4}
	bst.Insert(2)
	bst.Insert(6)
	bst.Insert(1)
	bst.Insert(3)
	bst.Insert(7)
	bst.Insert(5)

	actual := bst.data
	expected := 4
	if actual != expected {
		t.Errorf("bst.data: %d, want %d.", actual, expected)
	}

	actual = bst.left.data
	expected = 2
	if actual != expected {
		t.Errorf("bst.left.data: %d, want %d.", actual, expected)
	}

	actual = bst.left.left.data
	expected = 1
	if actual != expected {
		t.Errorf("bst.left.left.data: %d, want %d.", actual, expected)
	}

	actual = bst.left.right.data
	expected = 3
	if actual != expected {
		t.Errorf("bst.left.right.data: %d, want %d.", actual, expected)
	}

	actual = bst.right.data
	expected = 6
	if actual != expected {
		t.Errorf("bst.right.data: %d, want %d", actual, expected)
	}

	actual = bst.right.left.data
	expected = 5
	if actual != expected {
		t.Errorf("bst.right.left.data: %d, want %d", actual, expected)
	}

	actual = bst.right.right.data
	expected = 7
	if actual != expected {
		t.Errorf("bst.right.right.data: %d, want %d", actual, expected)
	}
}

func TestMapStringWithOneElement(t *testing.T) {
	bst := SearchTreeData{data: 4}

	actual := bst.MapString(strconv.Itoa)
	expected := []string{"4"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("bst.MapString(): %q, want %q.", actual, expected)
	}
}

func TestMapStringWithSmallElement(t *testing.T) {
	bst := SearchTreeData{data: 4}
	bst.Insert(2)

	actual := bst.MapString(strconv.Itoa)
	expected := []string{"2", "4"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("bst.MapString(): %q, want %q.", actual, expected)
	}
}

func TestMapStringWithLargeElement(t *testing.T) {
	bst := SearchTreeData{data: 4}
	bst.Insert(5)

	actual := bst.MapString(strconv.Itoa)
	expected := []string{"4", "5"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("bst.MapString(): %q, want %q.", actual, expected)
	}
}

func TestMapStringWithComplexStructure(t *testing.T) {
	bst := SearchTreeData{data: 4}
	bst.Insert(2)
	bst.Insert(1)
	bst.Insert(3)
	bst.Insert(6)
	bst.Insert(7)
	bst.Insert(5)

	actual := bst.MapString(strconv.Itoa)
	expected := []string{"1", "2", "3", "4", "5", "6", "7"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("bst.MapString(): %q, want %q.", actual, expected)
	}
}

func TestMapIntWithComplexStructure(t *testing.T) {
	bst := SearchTreeData{data: 4}
	bst.Insert(2)
	bst.Insert(1)
	bst.Insert(3)
	bst.Insert(6)
	bst.Insert(7)
	bst.Insert(5)

	f := func(i int) int {
		return i
	}

	actual := bst.MapInt(f)
	expected := []int{1, 2, 3, 4, 5, 6, 7}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("bst.MapString(): %v, want %v.", actual, expected)
	}
}
