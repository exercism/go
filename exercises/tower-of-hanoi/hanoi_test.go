package towerofhanoi

import (
	"fmt"
	"strings"
	"testing"
)

type board struct {
	from []int
	to   []int
	via  []int
}

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
		b.from = append(b.from, i) // push disks into rod A
	}

	// get moves from solution
	moves := Solve(disks, "from", "to", "via")

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

func apply(t *testing.T, disks int, b *board, moves []string) error {
	for i, m := range moves {
		t.Logf("disks=%d move=%s before: %v", disks, m, *b)
		s := strings.Split(m, "->")
		if len(s) != 2 {
			return fmt.Errorf("invalid move format %d: [%s]", i, m)
		}
		rod1 := strings.TrimSpace(s[0])
		rod2 := strings.TrimSpace(s[1])
		if rod1 == rod2 {
			return fmt.Errorf("invalid null move %d: [%s]", i, m)
		}
		var disk int
		switch rod1 {
		case "from":
			if len(b.from) == 0 {
				return fmt.Errorf("invalid move from empty rod FROM %d: [%s]", i, m)
			}
			last := len(b.from) - 1
			disk, b.from = b.from[last], b.from[:last]
		case "to":
			if len(b.to) == 0 {
				return fmt.Errorf("invalid move from empty rod TO %d: [%s]", i, m)
			}
			last := len(b.to) - 1
			disk, b.to = b.to[last], b.to[:last]
		case "via":
			if len(b.via) == 0 {
				return fmt.Errorf("invalid move from empty rod VIA %d: [%s]", i, m)
			}
			last := len(b.via) - 1
			disk, b.via = b.via[last], b.via[:last]
		default:
			return fmt.Errorf("invalid source rod %d: [%s]", i, m)
		}
		switch rod2 {
		case "from":
			if len(b.from) > 0 {
				last := len(b.from) - 1
				top := b.from[last]
				if disk > top {
					return fmt.Errorf("invalid move disk=%d is greater than top=%d of rod FROM %d: [%s]", disk, top, i, m)
				}
			}
			b.from = append(b.from, disk)
		case "to":
			if len(b.to) > 0 {
				last := len(b.to) - 1
				top := b.to[last]
				if disk > top {
					return fmt.Errorf("invalid move disk=%d is greater than top=%d of rod TO %d: [%s]", disk, top, i, m)
				}
			}
			b.to = append(b.to, disk)
		case "via":
			if len(b.via) > 0 {
				last := len(b.via) - 1
				top := b.via[last]
				if disk > top {
					return fmt.Errorf("invalid move disk=%d is greater than top=%d of rod VIA %d: [%s]", disk, top, i, m)
				}
			}
			b.via = append(b.via, disk)
		default:
			return fmt.Errorf("invalid destination rod %d: [%s]", i, m)
		}
		t.Logf("disks=%d move=%s after: %v", disks, m, *b)
	}
	return nil
}
