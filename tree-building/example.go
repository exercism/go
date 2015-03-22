// +build example

package tree

import (
	"fmt"
	"sort"
)

const TestVersion = 1

type Record struct {
	Id, Parent int
}

type Node struct {
	Id       int
	Children []*Node
}

type NodeSlice []*Node

func (n NodeSlice) Len() int           { return len(n) }
func (n NodeSlice) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n NodeSlice) Less(i, j int) bool { return n[i].Id < n[j].Id }

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	// At the end of this function this will hold: nodes[i].Id == i
	nodes := make([]Node, len(records))
	parents := make([]*Node, len(records))
	seen := make([]bool, len(records))

	for _, record := range records {
		if record.Id >= len(records) {
			return nil, fmt.Errorf("Too high id %d", record.Id)
		}
		if record.Id != 0 && record.Id <= record.Parent {
			return nil, fmt.Errorf("Record %d has self or later parent %d", record.Id, record.Parent)
		}
		if seen[record.Id] {
			return nil, fmt.Errorf("Record with id %d occurs multiple times", record.Id)
		}
		seen[record.Id] = true
		if record.Id != 0 {
			parents[record.Id] = &nodes[record.Parent]
		} else if record.Parent != 0 {
			return nil, fmt.Errorf("Root node has non-0 parent %d", record.Parent)
		}
	}

	for i := 1; i < len(nodes); i++ {
		parents[i].Children = append(parents[i].Children, &nodes[i])
	}

	for i, node := range nodes {
		// The Id field isn't actually used in this function, so we can delay
		// setting it to an opportune moment.
		nodes[i].Id = i
		sort.Sort(NodeSlice(node.Children))
	}
	return &nodes[0], nil
}
