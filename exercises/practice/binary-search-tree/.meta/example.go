package binarysearchtree

import "fmt"

type BinarySearchTree struct {
	left  *BinarySearchTree
	data  int
	right *BinarySearchTree
}

func NewBst(i int) *BinarySearchTree {
	return &BinarySearchTree{data: i}
}

func (bst *BinarySearchTree) Insert(i int) {
	if i <= bst.data {
		if bst.left != nil {
			bst.left.Insert(i)
		} else {
			bst.left = &BinarySearchTree{data: i}
		}
	} else {
		if bst.right != nil {
			bst.right.Insert(i)
		} else {
			bst.right = &BinarySearchTree{data: i}
		}
	}
}

func (bst *BinarySearchTree) SortedData() (result []int) {
	if bst.left != nil {
		result = append(bst.left.SortedData(), result...)
	}
	result = append(result, bst.data)
	if bst.right != nil {
		result = append(result, bst.right.SortedData()...)
	}
	return result
}

// useful for debugging
func (bst *BinarySearchTree) String() string {
	if bst == nil {
		return "nil"
	}
	return fmt.Sprintf("(%d %v %v)", bst.data, bst.left, bst.right)
}
