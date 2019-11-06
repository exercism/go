package tree

import "fmt"

const rootID = 0

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

// Build converts an unordered slice of Records into a hierarchical tree of Nodes,
// after validating that the tree is not a cyclic graph and that the records have
// a contiguous range of IDs.
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	//sort so we can check contiguity later
	for i, r := range records {
		if r.ID < rootID || r.ID >= len(records) {
			return nil, fmt.Errorf("out of bounds record id %d", r.ID)
		}
		records[r.ID], records[i] = records[i], records[r.ID]
	}

	//build and validate
	nodes := make([]Node, len(records))
	for i, r := range records {
		nodes[i].ID = r.ID

		if r.ID == rootID && r.Parent == 0 && i == 0 {
			continue
		}

		if err := validate(i, r); err != nil {
			return nil, err
		}

		p := &nodes[r.Parent]
		p.Children = append(p.Children, &nodes[i])
	}

	return &nodes[0], nil
}

func validate(i int, r Record) error {
	if i != r.ID {
		return fmt.Errorf("non-contiguous node %d (want %d)", r.ID, i)
	}

	if r.ID == rootID {
		return fmt.Errorf("root node has parent %d", r.Parent)
	}

	if r.ID <= r.Parent {
		return fmt.Errorf("node %d has impossible parent %d", r.ID, r.Parent)
	}
	return nil
}
