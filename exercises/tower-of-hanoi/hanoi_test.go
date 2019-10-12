package towerofhanoi

import (
	"fmt"
	"testing"
)

// board defined only for testing.
// student solution is not expected to use it.
type board struct {
	from []int
	to   []int
	via  []int
}

// constants defined only for testing.
// student solution is not expected to use them.
const (
	rodFrom = 'A'
	rodTo   = 'B'
	rodVia  = 'C'
)

func TestTowerOfHanoi(t *testing.T) {

	for i := 0; i < 8; i++ {
		if err := testDisks(t, i); err != nil {
			t.Errorf("FAIL totalDisks=%d %v", i, err)
			continue
		}
		t.Logf("PASS totalDisks=%d", i)
	}
}

func testDisks(t *testing.T, disks int) error {
	// create board
	b := board{}
	for i := disks; i > 0; i-- {
		b.from = append(b.from, i) // push disks into rod FROM
	}

	// get moves from solution
	moves := Solve(disks, rodFrom, rodTo, rodVia)

	var rodLetter bool
	if (rodFrom >= 'A' && rodFrom <= 'Z') && (rodTo >= 'A' && rodTo <= 'Z') && (rodVia >= 'A' && rodVia <= 'Z') {
		rodLetter = true
	}

	// apply moves to board
	err := apply(t, disks, &b, moves, rodLetter)
	if err != nil {
		return err
	}

	// check board
	if len(b.to) != disks {
		return fmt.Errorf("rod TO=%v has %d disks, should have %d: %v", rod(rodTo, rodLetter), len(b.to), disks, b.to)
	}

	return nil
}

func rod(r int, rodLetter bool) string {
	if rodLetter {
		return fmt.Sprintf("%c", r)
	}
	return fmt.Sprintf("%v", r)
}

func move(m Move, rodLetter bool) string {
	if rodLetter {
		return fmt.Sprintf("{%c %c}", m.from, m.to)
	}
	return fmt.Sprintf("%v", m)
}

func apply(t *testing.T, disks int, b *board, moves []Move, rodLetter bool) error {
	for i, m := range moves {
		t.Logf("disks=%d move=%v before: %v", disks, move(m, rodLetter), *b)
		rod1 := m.from
		rod2 := m.to
		if rod1 == rod2 {
			return fmt.Errorf("invalid null move %d: %v", i, move(m, rodLetter))
		}
		var disk int
		switch rod1 {
		case rodFrom:
			if len(b.from) == 0 {
				return fmt.Errorf("invalid move from empty rod FROM %d: %v", i, move(m, rodLetter))
			}
			last := len(b.from) - 1
			disk, b.from = b.from[last], b.from[:last]
		case rodTo:
			if len(b.to) == 0 {
				return fmt.Errorf("invalid move from empty rod TO %d: %v", i, move(m, rodLetter))
			}
			last := len(b.to) - 1
			disk, b.to = b.to[last], b.to[:last]
		case rodVia:
			if len(b.via) == 0 {
				return fmt.Errorf("invalid move from empty rod VIA %d: %v", i, move(m, rodLetter))
			}
			last := len(b.via) - 1
			disk, b.via = b.via[last], b.via[:last]
		default:
			return fmt.Errorf("invalid source rod %d: %v", i, move(m, rodLetter))
		}
		switch rod2 {
		case rodFrom:
			if len(b.from) > 0 {
				last := len(b.from) - 1
				top := b.from[last]
				if disk > top {
					return fmt.Errorf("invalid move disk=%d is greater than top=%d of rod FROM %d: %v", disk, top, i, move(m, rodLetter))
				}
			}
			b.from = append(b.from, disk)
		case rodTo:
			if len(b.to) > 0 {
				last := len(b.to) - 1
				top := b.to[last]
				if disk > top {
					return fmt.Errorf("invalid move disk=%d is greater than top=%d of rod TO %d: %v", disk, top, i, move(m, rodLetter))
				}
			}
			b.to = append(b.to, disk)
		case rodVia:
			if len(b.via) > 0 {
				last := len(b.via) - 1
				top := b.via[last]
				if disk > top {
					return fmt.Errorf("invalid move disk=%d is greater than top=%d of rod VIA %d: %v", disk, top, i, move(m, rodLetter))
				}
			}
			b.via = append(b.via, disk)
		default:
			return fmt.Errorf("invalid destination rod %d: %v", i, move(m, rodLetter))
		}
		t.Logf("disks=%d move=%v after: %v", disks, move(m, rodLetter), *b)
	}
	return nil
}
