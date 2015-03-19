package tree

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

// Define a function Build(records []Record) (*Node, error)
// where Record is a struct containing int fields Id and Parent
// and Node is a struct containing int field Id and []*Node field Children.
//
// Also define an exported TestVersion with a value that matches
// the internal testVersion here.

const testVersion = 1

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
			{Id: 0},
		},
		expected: &Node{
			Id: 0,
		},
	},
	{
		name: "three nodes in order",
		input: []Record{
			{Id: 0},
			{Id: 1, Parent: 0},
			{Id: 2, Parent: 0},
		},
		expected: &Node{
			Id: 0,
			Children: []*Node{
				{Id: 1},
				{Id: 2},
			},
		},
	},
	{
		name: "three nodes in reverse order",
		input: []Record{
			{Id: 2, Parent: 0},
			{Id: 1, Parent: 0},
			{Id: 0},
		},
		expected: &Node{
			Id: 0,
			Children: []*Node{
				{Id: 1},
				{Id: 2},
			},
		},
	},
	{
		name: "more than two children",
		input: []Record{
			{Id: 3, Parent: 0},
			{Id: 2, Parent: 0},
			{Id: 1, Parent: 0},
			{Id: 0},
		},
		expected: &Node{
			Id: 0,
			Children: []*Node{
				{Id: 1},
				{Id: 2},
				{Id: 3},
			},
		},
	},
	{
		name: "binary tree",
		input: []Record{
			{Id: 5, Parent: 1},
			{Id: 3, Parent: 2},
			{Id: 2, Parent: 0},
			{Id: 4, Parent: 1},
			{Id: 1, Parent: 0},
			{Id: 0},
			{Id: 6, Parent: 2},
		},
		expected: &Node{
			Id: 0,
			Children: []*Node{
				{
					Id: 1,
					Children: []*Node{
						{Id: 4},
						{Id: 5},
					},
				},
				{
					Id: 2,
					Children: []*Node{
						{Id: 3},
						{Id: 6},
					},
				},
			},
		},
	},
	{
		name: "unbalanced tree",
		input: []Record{
			{Id: 5, Parent: 2},
			{Id: 3, Parent: 2},
			{Id: 2, Parent: 0},
			{Id: 4, Parent: 1},
			{Id: 1, Parent: 0},
			{Id: 0},
			{Id: 6, Parent: 2},
		},
		expected: &Node{
			Id: 0,
			Children: []*Node{
				{
					Id: 1,
					Children: []*Node{
						{Id: 4},
					},
				},
				{
					Id: 2,
					Children: []*Node{
						{Id: 3},
						{Id: 5},
						{Id: 6},
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
			{Id: 0, Parent: 1},
			{Id: 1, Parent: 0},
		},
	},
	{
		name: "no root node",
		input: []Record{
			{Id: 1, Parent: 0},
		},
	},
	{
		name: "non-continuous",
		input: []Record{
			{Id: 2, Parent: 0},
			{Id: 4, Parent: 2},
			{Id: 1, Parent: 0},
			{Id: 0},
		},
	},
	{
		name: "cycle directly",
		input: []Record{
			{Id: 5, Parent: 2},
			{Id: 3, Parent: 2},
			{Id: 2, Parent: 2},
			{Id: 4, Parent: 1},
			{Id: 1, Parent: 0},
			{Id: 0},
			{Id: 6, Parent: 3},
		},
	},
	{
		name: "cycle indirectly",
		input: []Record{
			{Id: 5, Parent: 2},
			{Id: 3, Parent: 2},
			{Id: 2, Parent: 6},
			{Id: 4, Parent: 1},
			{Id: 1, Parent: 0},
			{Id: 0},
			{Id: 6, Parent: 3},
		},
	},
	{
		name: "higher id parent of lower id",
		input: []Record{
			{Id: 0},
			{Id: 2, Parent: 0},
			{Id: 1, Parent: 2},
		},
	},
}

func (n Node) String() string {
	return fmt.Sprintf("%d:%s", n.Id, n.Children)
}

func TestMakeTreeSuccess(t *testing.T) {
	if TestVersion != testVersion {
		t.Fatalf("Found TestVersion = %v, want %v", TestVersion, testVersion)
	}
	for _, tt := range successTestCases {
		actual, err := Build(tt.input)
		if err != nil {
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
			records[i] = Record{Id: 0}
		} else {
			records[i] = Record{Id: i, Parent: i >> 1}
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
			records[i] = Record{Id: 0}
		} else {
			records[i] = Record{Id: i, Parent: i / 10}
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
		records[i] = Record{Id: i, Parent: 0}
	}
	return shuffleRecords(records)
}

var shallowRecords = makeShallowRecords()

func BenchmarkShallowTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Build(shallowRecords)
	}
}
