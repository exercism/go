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
	rodFrom = iota
	rodTo
	rodVia
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

	// apply moves to board
	err := apply(t, disks, &b, moves)
	if err != nil {
		return err
	}

	// check board
	if len(b.from) != 0 {
		return fmt.Errorf("rod FROM has %d disks: %v", len(b.from), b.from)
	}
	if len(b.to) != disks {
		return fmt.Errorf("rod TO has %d disks: %v", len(b.to), b.to)
	}
	if len(b.via) != 0 {
		return fmt.Errorf("rod VIA has %d disks: %v", len(b.via), b.via)
	}

	return nil
}

func apply(t *testing.T, disks int, b *board, moves [][2]int) error {
	for i, m := range moves {
		t.Logf("disks=%d move=%v before: %v", disks, m, *b)
		rod1 := m[0]
		rod2 := m[1]
		if rod1 == rod2 {
			return fmt.Errorf("invalid null move %d: %v", i, m)
		}
		var disk int
		switch rod1 {
		case rodFrom:
			if len(b.from) == 0 {
				return fmt.Errorf("invalid move from empty rod FROM %d: %v", i, m)
			}
			last := len(b.from) - 1
			disk, b.from = b.from[last], b.from[:last]
		case rodTo:
			if len(b.to) == 0 {
				return fmt.Errorf("invalid move from empty rod TO %d: %v", i, m)
			}
			last := len(b.to) - 1
			disk, b.to = b.to[last], b.to[:last]
		case rodVia:
			if len(b.via) == 0 {
				return fmt.Errorf("invalid move from empty rod VIA %d: %v", i, m)
			}
			last := len(b.via) - 1
			disk, b.via = b.via[last], b.via[:last]
		default:
			return fmt.Errorf("invalid source rod %d: %v", i, m)
		}
		switch rod2 {
		case rodFrom:
			if len(b.from) > 0 {
				last := len(b.from) - 1
				top := b.from[last]
				if disk > top {
					return fmt.Errorf("invalid move disk=%d is greater than top=%d of rod FROM %d: %v", disk, top, i, m)
				}
			}
			b.from = append(b.from, disk)
		case rodTo:
			if len(b.to) > 0 {
				last := len(b.to) - 1
				top := b.to[last]
				if disk > top {
					return fmt.Errorf("invalid move disk=%d is greater than top=%d of rod TO %d: %v", disk, top, i, m)
				}
			}
			b.to = append(b.to, disk)
		case rodVia:
			if len(b.via) > 0 {
				last := len(b.via) - 1
				top := b.via[last]
				if disk > top {
					return fmt.Errorf("invalid move disk=%d is greater than top=%d of rod VIA %d: %v", disk, top, i, m)
				}
			}
			b.via = append(b.via, disk)
		default:
			return fmt.Errorf("invalid destination rod %d: %v", i, m)
		}
		t.Logf("disks=%d move=%v after: %v", disks, m, *b)
	}
	return nil
}
