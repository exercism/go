//nolint:gosec // In the context of this exercise, it is fine to use math.Rand instead of crypto.Rand.
package tree

import (
	"fmt"
	"math/rand"
	"testing"
)

// Define a function Build(records []Record) (*Node, error)
// where Record is a struct containing int fields ID and Parent
// and Node is a struct containing int field ID and []*Node field Children.

var successTestCases = []struct {
	name     string
	input    []Record
	expected *Node
}{
	{
		name:     "empty input",
		input:    []Record{},
		expected: nil,
	},
	{
		name: "one node",
		input: []Record{
			{ID: 0},
		},
		expected: &Node{
			ID: 0,
		},
	},
	{
		name: "three nodes in order",
		input: []Record{
			{ID: 0},
			{ID: 1, Parent: 0},
			{ID: 2, Parent: 0},
		},
		expected: &Node{
			ID: 0,
			Children: []*Node{
				{ID: 1},
				{ID: 2},
			},
		},
	},
	{
		name: "three nodes in reverse order",
		input: []Record{
			{ID: 2, Parent: 0},
			{ID: 1, Parent: 0},
			{ID: 0},
		},
		expected: &Node{
			ID: 0,
			Children: []*Node{
				{ID: 1},
				{ID: 2},
			},
		},
	},
	{
		name: "three levels of nesting",
		input: []Record{
			{ID: 2, Parent: 1},
			{ID: 1, Parent: 0},
			{ID: 3, Parent: 2},
			{ID: 0},
		},
		expected: &Node{
			ID: 0,
			Children: []*Node{
				{
					ID: 1,
					Children: []*Node{
						{
							ID: 2,
							Children: []*Node{
								{ID: 3},
							},
						},
					},
				},
			},
		},
	},
	{
		name: "more than two children",
		input: []Record{
			{ID: 3, Parent: 0},
			{ID: 2, Parent: 0},
			{ID: 1, Parent: 0},
			{ID: 0},
		},
		expected: &Node{
			ID: 0,
			Children: []*Node{
				{ID: 1},
				{ID: 2},
				{ID: 3},
			},
		},
	},
	{
		name: "binary tree",
		input: []Record{
			{ID: 5, Parent: 1},
			{ID: 3, Parent: 2},
			{ID: 2, Parent: 0},
			{ID: 4, Parent: 1},
			{ID: 1, Parent: 0},
			{ID: 0},
			{ID: 6, Parent: 2},
		},
		expected: &Node{
			ID: 0,
			Children: []*Node{
				{
					ID: 1,
					Children: []*Node{
						{ID: 4},
						{ID: 5},
					},
				},
				{
					ID: 2,
					Children: []*Node{
						{ID: 3},
						{ID: 6},
					},
				},
			},
		},
	},
	{
		name: "unbalanced tree",
		input: []Record{
			{ID: 5, Parent: 2},
			{ID: 3, Parent: 2},
			{ID: 2, Parent: 0},
			{ID: 4, Parent: 1},
			{ID: 1, Parent: 0},
			{ID: 0},
			{ID: 6, Parent: 2},
		},
		expected: &Node{
			ID: 0,
			Children: []*Node{
				{
					ID: 1,
					Children: []*Node{
						{ID: 4},
					},
				},
				{
					ID: 2,
					Children: []*Node{
						{ID: 3},
						{ID: 5},
						{ID: 6},
					},
				},
			},
		},
	},
}

var failureTestCases = []struct {
	name  string
	input []Record
}{
	{
		name: "one root node and has parent",
		input: []Record{
			{ID: 0, Parent: 1},
		},
	},
	{
		name: "root node has parent",
		input: []Record{
			{ID: 0, Parent: 1},
			{ID: 1, Parent: 0},
		},
	},
	{
		name: "no root node",
		input: []Record{
			{ID: 1, Parent: 0},
		},
	},
	{
		name: "duplicate node",
		input: []Record{
			{ID: 0, Parent: 0},
			{ID: 1, Parent: 0},
			{ID: 1, Parent: 0},
		},
	},
	{
		name: "duplicate root",
		input: []Record{
			{ID: 0, Parent: 0},
			{ID: 0, Parent: 0},
		},
	},
	{
		name: "non-continuous",
		input: []Record{
			{ID: 2, Parent: 0},
			{ID: 4, Parent: 2},
			{ID: 1, Parent: 0},
			{ID: 0},
		},
	},
	{
		name: "cycle directly",
		input: []Record{
			{ID: 5, Parent: 2},
			{ID: 3, Parent: 2},
			{ID: 2, Parent: 2},
			{ID: 4, Parent: 1},
			{ID: 1, Parent: 0},
			{ID: 0},
			{ID: 6, Parent: 3},
		},
	},
	{
		name: "cycle indirectly",
		input: []Record{
			{ID: 5, Parent: 2},
			{ID: 3, Parent: 2},
			{ID: 2, Parent: 6},
			{ID: 4, Parent: 1},
			{ID: 1, Parent: 0},
			{ID: 0},
			{ID: 6, Parent: 3},
		},
	},
	{
		name: "higher id parent of lower id",
		input: []Record{
			{ID: 0},
			{ID: 2, Parent: 0},
			{ID: 1, Parent: 2},
		},
	},
}

func (n Node) String() string {
	return fmt.Sprintf("%d:%s", n.ID, n.Children)
}

func TestMakeTreeSuccess(t *testing.T) {
	for _, tt := range successTestCases {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := Build(tt.input)
			if err != nil {
				var _ error = err
				t.Fatalf("Build for test case %q returned error %q. Error not expected.",
					tt.name, err)
			}
			if !nodeEqual(actual, tt.expected) {
				t.Fatalf("Build for test case %q returned %s but was expected to return %s.",
					tt.name, actual, tt.expected)
			}
		})
	}
}

func TestMakeTreeFailure(t *testing.T) {
	for _, tt := range failureTestCases {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := Build(tt.input)
			if err == nil {
				t.Fatalf("Build for test case %q returned %s but was expected to fail.",
					tt.name, actual)
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
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
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
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
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
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		Build(shallowRecords)
	}
}

func nodeEqual(node1, node2 *Node) bool {
	switch {
	case node1 == nil && node2 == nil:
		return true
	case node1 == nil && node2 != nil:
		return false
	case node1 != nil && node2 == nil:
		return false
	default:
		return node1.ID == node2.ID && nodeSliceEqual(node1.Children, node2.Children)
	}
}

func nodeSliceEqual(nodes1, nodes2 []*Node) bool {
	if len(nodes1) == 0 && len(nodes2) == 0 {
		return true
	}
	if len(nodes1) != len(nodes2) {
		return false
	}
	for i := range nodes1 {
		if !nodeEqual(nodes1[i], nodes2[i]) {
			return false
		}
	}
	return true
}
