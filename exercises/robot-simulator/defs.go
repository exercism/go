package robot

import "fmt"

// definitions used in step 1

var Robot1 struct {
	X, Y   int
	Dir
}

type Dir int

var _ fmt.Stringer = Dir(1729)

// additional definitions used in step 2

type Command byte // valid values are 'R', 'L', 'A'
type RU int
type Pos struct{ Easting, Northing RU }
type Rect struct{ Min, Max Pos }
type Robot2 struct {
	Dir
	Pos
}

// additional definition used in step 3

type Robot3 struct {
	Name string
	Robot2
}
