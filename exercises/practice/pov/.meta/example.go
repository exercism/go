package pov

type RoseTree struct {
	value    string
	children []*RoseTree
	parent   *RoseTree
}

func New(value string, children ...*RoseTree) *RoseTree {
	if value == "" {
		return nil
	}
	result := &RoseTree{value, children, nil}
	for _, child := range children {
		if child != nil {
			child.parent = result
		}
	}
	return result
}

func (tr *RoseTree) Value() string {
	if tr == nil {
		return ""
	} else {
		return tr.value
	}
}

func (tr *RoseTree) Children() []*RoseTree {
	if tr == nil {
		return nil
	} else {
		return tr.children
	}
}

func (tr *RoseTree) FindNode(value string) *RoseTree {
	if tr == nil || tr.value == value {
		return tr
	}
	for _, child := range tr.children {
		res := child.FindNode(value)
		if res != nil {
			return res
		}
	}
	return nil
}

func (tree *RoseTree) FromPov(from string) *RoseTree {
	node := tree.FindNode(from)
	seen := make(map[string]bool)
	var f func(*RoseTree) *RoseTree
	f = func(tr *RoseTree) *RoseTree {
		if tr == nil || seen[tr.value] {
			return nil
		}
		seen[tr.value] = true
		children := make([]*RoseTree, 0, len(tr.children)+1)
		for _, child := range tr.children {
			fChild := f(child)
			if fChild != nil {
				children = append(children, fChild)
			}
		}
		if tr.parent != nil {
			fParent := f(tr.parent)
			if fParent != nil {
				children = append(children, fParent)
			}
		}
		return New(tr.value, children...)
	}
	return f(node)
}

func (tr *RoseTree) PathFromRoot(value string) []string {
	node := tr.FindNode(value)
	if node == nil {
		return nil
	}
	result := make([]string, 0)
	for node != tr {
		result = append(result, node.value)
		node = node.parent
	}
	result = append(result, tr.value)
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}

func (tr *RoseTree) PathTo(from, to string) []string {
	if tr == nil {
		return nil
	}
	if from == to {
		return []string{to}
	}
	tr1 := tr.FromPov(from)
	if tr1 == nil {
		return nil
	}
	return tr1.PathFromRoot(to)
}
