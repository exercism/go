package kindergarten

// Define the Garden type here.

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`
//
// If the children argument is empty, use the list of children defined in the instructions.
// If it is not empty, use the given value.

func NewGarden(diagram string, children []string) (*Garden, error) {
	panic("Please implement the NewGarden function")
}

func (g *Garden) Plants(child string) ([]string, bool) {
	panic("Please implement the Plants function")
}
