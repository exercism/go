//nolint:gosec // In the context of this exercise, it is fine to use math.Rand instead of crypto.Rand.
package treebuilding

import (
	"fmt"
	"math/rand"
	"slices"
	"testing"
)

// Define a function Build(records []Record) (*Node, error)
// where Record is a struct containing int fields ID and Parent
// and Node is a struct containing int field ID and []*Node field Children.

func (n Node) String() string {
	return fmt.Sprintf("%d:%s", n.ID, n.Children)
}

func TestBuildTree(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual, err := Build(tc.input)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("Build returned %s but was expected to fail.", actual)
				}
			} else if !nodeEqual(actual, tc.expected) {
				t.Fatalf("Build for test case %q returned %s but was expected to return %s.",
					tc.description, actual, tc.expected)
			}
		})
	}
}

func shuffleRecords(records []Record) []Record {
	gen := rand.New(rand.NewSource(42))
	newRecords := make([]Record, len(records))
	for i, idx := range gen.Perm(len(records)) {
		newRecords[i] = records[idx]
	}
	return newRecords
}

// Binary tree
func makeTwoTreeRecords() []Record {
	records := make([]Record, 1<<16)
	for i := range records {
		if i == 0 {
			records[i] = Record{ID: 0}
		} else {
			records[i] = Record{ID: i, Parent: i >> 1}
		}
	}
	return shuffleRecords(records)
}

var twoTreeRecords = makeTwoTreeRecords()

func BenchmarkTwoTree(b *testing.B) {
	for range b.N {
		Build(twoTreeRecords)
	}
}

// Each node but the root node and leaf nodes has ten children.
func makeTenTreeRecords() []Record {
	records := make([]Record, 10000)
	for i := range records {
		if i == 0 {
			records[i] = Record{ID: 0}
		} else {
			records[i] = Record{ID: i, Parent: i / 10}
		}
	}
	return shuffleRecords(records)
}

var tenTreeRecords = makeTenTreeRecords()

func BenchmarkTenTree(b *testing.B) {
	for range b.N {
		Build(tenTreeRecords)
	}
}

func makeShallowRecords() []Record {
	records := make([]Record, 10000)
	for i := range records {
		records[i] = Record{ID: i, Parent: 0}
	}
	return shuffleRecords(records)
}

var shallowRecords = makeShallowRecords()

func BenchmarkShallowTree(b *testing.B) {
	for range b.N {
		Build(shallowRecords)
	}
}

func nodeEqual(node1, node2 *Node) bool {
	switch {
	case node1 == nil && node2 == nil:
		return true
	case node1 == nil || node2 == nil:
		return false
	default:
		return node1.ID == node2.ID && slices.EqualFunc(node1.Children, node2.Children, nodeEqual)
	}
}
