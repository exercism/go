package tree

type Record struct {
  ID, Parent int
}

type Node struct {
  ID int
  Children []*Node
}

func Build(records []Record) (*Node, error) {
  panic("Please implement the Build function")
}
