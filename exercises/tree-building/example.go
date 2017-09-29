package tree

import (
	"fmt"
	"sort"
)

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

type NodeSlice []*Node

func (n NodeSlice) Len() int           { return len(n) }
func (n NodeSlice) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n NodeSlice) Less(i, j int) bool { return n[i].ID < n[j].ID }

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	// At the end of this function this will hold: nodes[i].ID == i
	nodes := make([]Node, len(records))
	parents := make([]*Node, len(records))
	seen := make([]bool, len(records))

	for _, record := range records {
		if record.ID >= len(records) {
			return nil, fmt.Errorf("Too high id %d", record.ID)
		}
		if record.ID != 0 && record.ID <= record.Parent {
			return nil, fmt.Errorf("Record %d has self or later parent %d", record.ID, record.Parent)
		}
		if seen[record.ID] {
			return nil, fmt.Errorf("Record with id %d occurs multiple times", record.ID)
		}
		seen[record.ID] = true
		if record.ID != 0 {
			parents[record.ID] = &nodes[record.Parent]
		} else if record.Parent != 0 {
			return nil, fmt.Errorf("Root node has non-0 parent %d", record.Parent)
		}
	}

	for i := 1; i < len(nodes); i++ {
		parents[i].Children = append(parents[i].Children, &nodes[i])
	}

	for i, node := range nodes {
		// The ID field isn't actually used in this function, so we can delay
		// setting it to an opportune moment.
		nodes[i].ID = i
		sort.Sort(NodeSlice(node.Children))
	}
	return &nodes[0], nil
}
