package robot

import "fmt"

// definitions used in step 1

var (
	X, Y   int
	Facing Dir
)

type Dir int

var _ fmt.Stringer = Dir(1729)

// additional definitions used in step 2

type Command byte // valid values are 'R', 'L', 'A'
type RU int
type Pos struct{ Easting, Northing RU }
type Rect struct{ Min, Max Pos }
type DirAt struct {
	Dir
	Pos
}

// additional definition used in step 3

type Place struct {
	Name string
	DirAt
}
