package prism

import "math"

type Position struct {
	x     float64
	y     float64
	angle float64
}

type Prism struct {
	id    int
	x     float64
	y     float64
	angle float64
}

func hits(ax, ay, bx, by, angle float64) bool {
	dx := bx - ax
	dy := by - ay
	r := math.Hypot(dx, dy)
	rad := math.Pi * angle / 180
	return math.Abs(math.Sin(rad)*r-dy) < 0.1 && math.Abs(math.Cos(rad)*r-dx) < 0.1
}

func nextHit(x, y, angle float64, prisms []Prism) (Prism, bool) {
	next := Prism{id: -1}
	var distance float64
	for _, prism := range prisms {
		pDist := math.Hypot(prism.x-x, prism.y-y)
		if pDist != 0 && hits(x, y, prism.x, prism.y, angle) && (next.id == -1 || pDist < distance) {
			next = prism
			distance = pDist
		}
	}
	return next, next.id != -1
}

func FindSequence(start Position, prisms []Prism) []int {
	x := start.x
	y := start.y
	angle := start.angle
	var seen []int
	next, ok := nextHit(x, y, angle, prisms)
	for ok {
		x = next.x
		y = next.y
		angle += next.angle
		seen = append(seen, next.id)
		next, ok = nextHit(x, y, angle, prisms)
	}
	return seen
}
