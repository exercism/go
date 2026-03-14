package zipper

import "errors"

type Node struct {
	value int
	left  *Node
	right *Node
}

// Clone this tree, but replace the focus node with a custom func version.
func (n *Node) cloneWithChange(focus *Node, change func(*Node) *Node) (*Node, *Node) {
	if n == nil {
		return nil, nil
	}
	if n == focus {
		changed := change(n)
		return changed, changed
	}
	left, lFocus := n.left.cloneWithChange(focus, change)
	if lFocus != nil {
		focus = lFocus
	}
	right, rFocus := n.right.cloneWithChange(focus, change)
	if rFocus != nil {
		focus = rFocus
	}
	return &Node{value: n.value, left: left, right: right}, focus
}

type Zipper struct {
	root    *Node
	parents map[*Node]*Node
	focus   *Node
}

// Recursively populate the parents map.
func setParents(parents map[*Node]*Node, parent *Node, child *Node) {
	if child == nil {
		return
	}
	parents[child] = parent
	setParents(parents, child, child.left)
	setParents(parents, child, child.right)
}

func NewZipper(tree *Node) Zipper {
	return newZipperWithFocus(tree, tree)
}

func newZipperWithFocus(tree, focus *Node) Zipper {
	parents := make(map[*Node]*Node)
	setParents(parents, nil, tree)
	return Zipper{root: tree, focus: focus, parents: parents}
}

func (z Zipper) Value() int {
	return z.focus.value
}

func (z Zipper) ToTree() *Node {
	return z.root
}

// Return a new Zipper with the focus changed.
func (z Zipper) withFocus(focus *Node) (Zipper, error) {
	if focus == nil {
		return Zipper{}, errors.New("no node")
	}
	return Zipper{root: z.root, parents: z.parents, focus: focus}, nil
}

func (z Zipper) Left() (Zipper, error) {
	return z.withFocus(z.focus.left)
}

func (z Zipper) Right() (Zipper, error) {
	return z.withFocus(z.focus.right)
}

func (z Zipper) Up() (Zipper, error) {
	return z.withFocus(z.parents[z.focus])
}

func (z Zipper) SetValue(v int) Zipper {
	tree, focus := z.root.cloneWithChange(z.focus, func(n *Node) *Node {
		l, _ := n.left.cloneWithChange(nil, nil)
		r, _ := n.right.cloneWithChange(nil, nil)
		return &Node{value: v, left: l, right: r}
	})
	return newZipperWithFocus(tree, focus)
}

func (z Zipper) SetLeft(v *Node) Zipper {
	tree, focus := z.root.cloneWithChange(z.focus, func(n *Node) *Node {
		r, _ := n.right.cloneWithChange(nil, nil)
		return &Node{value: n.value, left: v, right: r}
	})
	return newZipperWithFocus(tree, focus)
}

func (z Zipper) SetRight(v *Node) Zipper {
	tree, focus := z.root.cloneWithChange(z.focus, func(n *Node) *Node {
		l, _ := n.left.cloneWithChange(nil, nil)
		return &Node{value: n.value, left: l, right: v}
	})
	return newZipperWithFocus(tree, focus)
}
