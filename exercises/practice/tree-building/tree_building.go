package tree

import "errors"

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

func Build(records []Record) (*Node, error) {
	if len(records)==0{
        return nil,nil
    }

    positions := make([]int, len(records))
    nodes := make([]Node,len(records))
    for i, r := range records {
        if r.ID < 0 || r.ID >= len(records) {
			return nil, errors.New("records ID is wrong1.")
		}
		positions[r.ID] = i
	}
    for i := range positions{
        d := records[positions[i]]
        if d.ID != i || d.ID<d.Parent || (d.ID==d.Parent && d.ID!=0){
            return nil,errors.New("records ID is wrong2.")
        }

        nodes[i].ID=d.ID
        if i!=0{
            p := &nodes[d.Parent]
            p.Children = append(p.Children,&nodes[i])
        }
    
    }

    return &nodes[0],nil
}
