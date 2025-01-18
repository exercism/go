package binarysearchtree

import (
	"errors"
	"strings"
	"testing"
)

type DataTestCase struct {
	description string
	input       []int
	paths       []string
	data        []int
}

var newBstTestCase DataTestCase = DataTestCase{
	description: "data is retained",
	input:       []int{4},
	paths:       []string{""},
	data:        []int{4},
}

var insertTestCases = []DataTestCase{
	{
		description: "smaller number at left node",
		input:       []int{4, 2},
		paths:       []string{"", "l"},
		data:        []int{4, 2},
	},
	{
		description: "same number at left node",
		input:       []int{4, 4},
		paths:       []string{"", "l"},
		data:        []int{4, 4},
	},
	{
		description: "greater number at right node",
		input:       []int{4, 5},
		paths:       []string{"", "r"},
		data:        []int{4, 5},
	},
	{
		description: "can create complex tree",
		input:       []int{4, 2, 6, 1, 3, 5, 7},
		paths:       []string{"", "l", "ll", "lr", "r", "rl", "rr"},
		data:        []int{4, 2, 1, 3, 6, 5, 7},
	},
}

type SortedDataTestCase struct {
	description string
	input       []int
	expected    []int
}

var sortedDataTestCases = []SortedDataTestCase{
	{
		description: "can sort single number",
		input:       []int{2},
		expected:    []int{2},
	},
	{
		description: "can sort if second number is smaller than first",
		input:       []int{2, 1},
		expected:    []int{1, 2},
	},
	{
		description: "can sort if second number is same as first",
		input:       []int{2, 2},
		expected:    []int{2, 2},
	},
	{
		description: "can sort if second number is greater than first",
		input:       []int{2, 3},
		expected:    []int{2, 3},
	},
	{
		description: "can sort complex tree",
		input:       []int{2, 1, 3, 6, 7, 5},
		expected:    []int{1, 2, 3, 5, 6, 7},
	},
}

func (bst *BinarySearchTree) Size() int {
	if bst == nil {
		return 0
	} else {
		return 1 + bst.left.Size() + bst.right.Size()
	}
}

func (bst *BinarySearchTree) Data(path []rune) (int, error) {
	if bst == nil {
		return 0, errors.New("nil")
	}
	if len(path) == 0 {
		return bst.data, nil
	}
	switch path[0] {
	case 'l':
		return bst.left.Data(path[1:])
	case 'r':
		return bst.right.Data(path[1:])
	default:
		return 0, errors.New("Invalid path element:" + string(path[0]))
	}
}

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 {
		return true
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// makeBst builds tree by calling user defined functions NewBst and Insert
func makeBst(input []int) *BinarySearchTree {
	if len(input) == 0 {
		return nil
	}
	result := NewBst(input[0])
	for i := 1; i < len(input); i++ {
		result.Insert(input[i])
	}
	return result
}

func expandPath(p string) string {
	parts := make([]string, len(p))
	for i, c := range p {
		if c == 'l' {
			parts[i] = "left"
		} else {
			parts[i] = "right"
		}
	}
	return strings.Join(append([]string{"bst"}, parts...), ".")
}

func TestNewBst(t *testing.T) {
	td := newBstTestCase
	t.Run(td.description, func(t *testing.T) {
		tree := NewBst(td.input[0])
		if tree == nil {
			t.Fatalf("bst should not be nil")
		}
		if tree.Size() != len(td.input) {
			t.Fatalf("bst should have same number of elements as input, expected: %v, got: %v",
				len(td.input), tree.Size())
		}
		expected := td.data[0]
		got := tree.data
		if got != expected {
			t.Fatalf("expected %d, got %d", expected, got)
		}
	})
}

func TestInsert(t *testing.T) {
	for _, td := range insertTestCases {
		t.Run(td.description, func(t *testing.T) {
			tree := makeBst(td.input)
			if tree == nil {
				t.Fatalf("bst should not be nil")
			}
			if tree.Size() != len(td.input) {
				t.Fatalf("bst should have same number of elements as input, expected: %v, got: %v",
					len(td.input), tree.Size())
			}
			for i, path := range td.paths {
				expected := td.data[i]
				expPath := expandPath(path)
				got, err := tree.Data([]rune(path))
				if err != nil {
					t.Fatalf("%v should not be nil:", expPath)
				}
				if got != expected {
					t.Fatalf("%v: expected %d, got %d", expPath+".data", expected, got)
				}
			}
		})
	}
}

func TestSortedData(t *testing.T) {
	for _, td := range sortedDataTestCases {
		t.Run(td.description, func(t *testing.T) {
			tree := makeBst(td.input)
			if tree == nil {
				t.Fatalf("bst should not be nil")
			}
			if tree.Size() != len(td.input) {
				t.Fatalf("bst should have same number of elements as input, expected: %v, got: %v",
					len(td.input), tree.Size())
			}
			got := tree.SortedData()
			if !slicesEqual(got, td.expected) {
				t.Fatalf("expected %d, got %d", td.expected, got)
			}
		})
	}
}

var benchmarkResult []int

func BenchmarkSortedData(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	input := []int{2, 1, 3, 6, 7, 5, 2, 1, 3, 6, 7, 5}
	tree := makeBst(input)
	var result []int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result = tree.SortedData()
	}
	benchmarkResult = result
}
