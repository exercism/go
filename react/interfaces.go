package react

// A Reactor manages linked cells.
type Reactor interface {
	// CreateInput creates an input cell linked into the reactor
	// with the given initial value.
	CreateInput(int) InputCell

	// CreateCompute1 creates a compute cell which computes its value
	// based on one other cell. The compute function will only be called
	// if the value of the passed cell changes.
	CreateCompute1(Cell, func(int) int) ComputeCell

	// CreateCompute2 is like CreateCompute1, but depending on two cells.
	// The compute function will only be called if the value of any of the
	// passed cells changes.
	CreateCompute2(Cell, Cell, func(int, int) int) ComputeCell
}

// A Cell is conceptually a holder of a value.
type Cell interface {
	// Value returns the current value of the cell.
	Value() int
}

// An InputCell has a changeable value, changing the value triggers updates to
// other cells.
type InputCell interface {
	Cell

	// SetValue sets the value of the cell.
	SetValue(int)
}

// A ComputeCell always computes its value based on other cells and can
// call callbacks upon changes.
type ComputeCell interface {
	Cell

	// AddCallback adds a callback which will be called when the value changes.
	// It returns a callback handle which can be used to remove the callback.
	AddCallback(func(int)) CallbackHandle

	// RemoveCallback removes a previously added callback, if it exists.
	RemoveCallback(CallbackHandle)
}

// A CallbackHandle is used to remove previously added callbacks, see ComputeCell.
type CallbackHandle interface{}
