package robot

import "fmt"

// definitions used in step 1

var Step1Robot struct {
	X, Y int
	Dir
}

type Dir int

var _ fmt.Stringer = Dir(1729)

// additional definitions used in step 2

type Command byte // valid values are 'R', 'L', 'A'
type RU int
type Pos struct{ Easting, Northing RU }
type Rect struct{ Min, Max Pos }
type Step2Robot struct {
	Dir
	Pos
}

// additional definition used in step 3

type Step3Robot struct {
	Name string
	Step2Robot
}
