package spiralmatrix

import (
	"fmt"
)

type coordinate struct {
	x int
	y int
}

func (co coordinate) withinGrid(size int) bool {
	return co.x >= 0 && co.x < size && co.y >= 0 && co.y < size
}

type direction int

const (
	up direction = iota
	right
	down
	left
)

func (d direction) move(co coordinate) coordinate {
	switch d {
	case up:
		return coordinate{
			x: co.x - 1,
			y: co.y,
		}
	case right:
		return coordinate{
			x: co.x,
			y: co.y + 1,
		}
	case down:
		return coordinate{
			x: co.x + 1,
			y: co.y,
		}
	case left:
		return coordinate{
			x: co.x,
			y: co.y - 1,
		}
	default:
		panic(fmt.Sprintf("unexpected direction to move: %v", d))
	}
}

func (d direction) turnRight() direction {
	switch d {
	case up:
		return right
	case right:
		return down
	case down:
		return left
	case left:
		return up
	default:
		panic(fmt.Sprintf("unexpected direction to turn right: %v", d))
	}
}

// SpiralMatrix is an implementation of the spiral matrix exercise.
func SpiralMatrix(size int) [][]int {
	matrix := make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
	}

	co := coordinate{
		x: 0,
		y: 0,
	}
	dir := right
	count := size * size
	for i := 0; i < count; i++ {
		matrix[co.x][co.y] = i + 1

		maybeNext := dir.move(co)
		if maybeNext.withinGrid(size) && matrix[maybeNext.x][maybeNext.y] == 0 {
			co = maybeNext
		} else {
			dir = dir.turnRight()
			co = dir.move(co)
		}
	}

	return matrix
}
