package twobucket

import (
	"errors"
)

type bucket int

const (
	bOne bucket = iota
	bTwo
)

type step int

const (
	emptyOne step = iota
	emptyTwo
	fillOne
	fillTwo
	pourOneToTwo
	pourTwoToOne
)

type problem struct {
	capacity [2]int
	goal     int
	start    bucket
}

type state struct {
	level        [2]int
	previousStep step
	numSteps     int
}

const (
	FirstBucketName  = "one"
	SecondBucketName = "two"
)

// Solve uses given bucket sizes, the goal amount, and starting bucket to
// solve the two-bucket problem to measure exactly the goal mount,
// returning the goal bucket name "one" or "two",
// the required number of steps, and the fill-level of the other bucket.
func Solve(sizeBucketOne,
	sizeBucketTwo,
	goalAmount int,
	startBucket string) (goalBucket string, numSteps, otherBucketLevel int, err error) {

	if err := validateParameters(sizeBucketOne, sizeBucketTwo, goalAmount, startBucket); err != nil {
		return "", 0, 0, err
	}

	p := problem{
		capacity: [2]int{sizeBucketOne, sizeBucketTwo},
		goal:     goalAmount,
	}
	var s state
	if startBucket == FirstBucketName {
		p.start = bOne
		performStep(p, &s, fillOne)
	} else {
		p.start = bTwo
		performStep(p, &s, fillTwo)
	}
	// Initial step might be a solution.
	if !isSolution(p, s) {
		s = findGoal(p, s)
	}
	switch {
	case s.level[bOne] == p.goal:
		return FirstBucketName, s.numSteps, s.level[bTwo], nil
	case s.level[bTwo] == p.goal:
		return SecondBucketName, s.numSteps, s.level[bOne], nil
	}
	return "", 0, 0, errors.New("no solution")
}

func validateParameters(sizeBucketOne, sizeBucketTwo, goalAmount int, startBucket string) error {
	if sizeBucketOne <= 0 {
		return errors.New("sizeBucketOne invalid")
	}
	if sizeBucketTwo <= 0 {
		return errors.New("sizeBucketTwo invalid")
	}
	if goalAmount <= 0 {
		return errors.New("goalAmount invalid")
	}
	if startBucket != FirstBucketName && startBucket != SecondBucketName {
		return errors.New("startBucket invalid")
	}
	return nil
}

func isSolution(p problem, s state) bool {
	return s.level[bOne] == p.goal || s.level[bTwo] == p.goal
}

func findGoal(p problem, s state) (g state) {
	searchList := make([]state, 1)
	searchList[0] = s

	// Use breadth-first search to find the goal, tracking any previously visited states.
	visited := map[[2]int]bool{}

	// Mark as already visited two invalid bucket levels: 0,0 and the reverse starting position.
	visited[[2]int{0, 0}] = true
	if p.start == bOne {
		visited[[2]int{0, p.capacity[bTwo]}] = true
	} else {
		visited[[2]int{p.capacity[bOne], 0}] = true
	}

	for len(searchList) != 0 {
		// Pop one item from the searchList each pass.
		current := searchList[0]
		searchList = searchList[1:]

		for _, x := range getPossibleSteps(p, current) {
			next := current
			performStep(p, &next, x)
			if isSolution(p, next) {
				return next
			}
			if !visited[next.level] {
				searchList = append(searchList, next)
				visited[next.level] = true
			}
		}
	}
	return state{}
}

func performStep(p problem, s *state, x step) {
	switch x {
	case emptyOne:
		s.level[bOne] = 0
	case emptyTwo:
		s.level[bTwo] = 0
	case fillOne:
		s.level[bOne] = p.capacity[bOne]
	case fillTwo:
		s.level[bTwo] = p.capacity[bTwo]
	case pourOneToTwo:
		pour(p, s, bOne, bTwo)
	case pourTwoToOne:
		pour(p, s, bTwo, bOne)
	}
	s.numSteps++
	s.previousStep = x
}

// pour from bucket a to b.
func pour(p problem, s *state, a, b bucket) {
	amount := p.capacity[b] - s.level[b]
	if amount > s.level[a] {
		amount = s.level[a]
	}
	s.level[b] += amount
	s.level[a] -= amount
}

func getPossibleSteps(p problem, s state) (list []step) {
	for x := emptyOne; x <= pourTwoToOne; x++ {
		if canPerformStep(p, s, x) {
			list = append(list, x)
		}
	}
	return list
}

func canPerformStep(p problem, s state, x step) bool {
	switch x {
	case emptyOne:
		if s.previousStep == fillOne || s.previousStep == pourOneToTwo {
			return false
		}
		return s.level[bOne] != 0
	case emptyTwo:
		if s.previousStep == fillTwo || s.previousStep == pourTwoToOne {
			return false
		}
		return s.level[bTwo] != 0
	case fillOne:
		if s.previousStep == emptyOne || s.level[bOne] == p.capacity[bOne] {
			return false
		}
		return true
	case fillTwo:
		if s.previousStep == emptyTwo || s.level[bTwo] == p.capacity[bTwo] {
			return false
		}
		return true
	case pourOneToTwo:
		return s.level[bOne] != 0 && s.level[bTwo] < p.capacity[bTwo]
	case pourTwoToOne:
		return s.level[bTwo] != 0 && s.level[bOne] < p.capacity[bOne]
	}
	return false
}
