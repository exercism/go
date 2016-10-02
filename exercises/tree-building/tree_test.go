package tree

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

// Define a function Build(records []Record) (*Node, error)
// where Record is a struct containing int fields ID and Parent
// and Node is a struct containing int field ID and []*Node field Children.
//
// Also define a testVersion with a value that matches
// the targetTestVersion here.

const targetTestVersion = 4

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
		actual, err := Build(tt.input)
		if err != nil {
			var _ error = err
			t.Fatalf("Build for test case %q returned error %q. Error not expected.",
				tt.name, err)
		}
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Fatalf("Build for test case %q returned %s but was expected to return %s.",
				tt.name, tt.expected, actual)
		}
	}
}

func TestMakeTreeFailure(t *testing.T) {
	for _, tt := range failureTestCases {
		actual, err := Build(tt.input)
		if err == nil {
			t.Fatalf("Build for test case %q returned %s but was expected to fail.",
				tt.name, actual)
		}
	}
}

func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v", testVersion, targetTestVersion)
	}
}

func shuffleRecords(records []Record) []Record {
	rand := rand.New(rand.NewSource(42))
	newRecords := make([]Record, len(records))
	for i, idx := range rand.Perm(len(records)) {
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
	for i := 0; i < b.N; i++ {
		Build(shallowRecords)
	}
}
