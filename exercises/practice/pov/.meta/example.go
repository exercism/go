package pov

type Tree struct {
	value    string
	children []*Tree
	parent   *Tree
}

func New(value string, children ...*Tree) *Tree {
	if value == "" {
		return nil
	}
	result := &Tree{value, children, nil}
	for _, child := range children {
		if child != nil {
			child.parent = result
		}
	}
	return result
}

func (tr *Tree) Value() string {
	if tr == nil {
		return ""
	} else {
		return tr.value
	}
}

func (tr *Tree) Children() []*Tree {
	if tr == nil {
		return nil
	} else {
		return tr.children
	}
}

func (tr *Tree) FindNode(value string) *Tree {
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

func (tree *Tree) FromPov(from string) *Tree {
	node := tree.FindNode(from)
	seen := make(map[string]bool)
	var f func(*Tree) *Tree
	f = func(tr *Tree) *Tree {
		if tr == nil || seen[tr.value] {
			return nil
		}
		seen[tr.value] = true
		children := make([]*Tree, 0, len(tr.children)+1)
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

func (tr *Tree) PathFromRoot(value string) []string {
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

func (tr *Tree) PathTo(from, to string) []string {
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
