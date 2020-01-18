package react

type reactor struct {
	cells []*cell
}

type cell struct {
	ord       int // Cell ordinal (identifier)
	reac      *reactor
	value     int
	consumers []int
	producers []*cell
	compute   func([]int) int
	nextKey   int
	observers map[int]func(int)
}

type canceler struct {
	cancel func()
}

func (c *canceler) Cancel() {
	c.cancel()
}

func (c *cell) Value() int { return c.value }

func (c *cell) SetValue(v int) {
	oldVal := c.value
	c.value = v
	if v != oldVal {
		c.reac.trigger(c)
	}
}

func (c *cell) AddCallback(cb func(int)) Canceler {
	key := c.nextKey
	c.nextKey++
	c.observers[key] = cb
	return &canceler{
		cancel: func() {
			delete(c.observers, key)
		},
	}
}

func New() Reactor {
	return &reactor{}
}

func (r *reactor) CreateInput(initial int) InputCell {
	newCell := &cell{ord: len(r.cells), reac: r, value: initial}
	r.cells = append(r.cells, newCell)
	return newCell
}

func (r *reactor) CreateCompute1(dep Cell, compute func(int) int) ComputeCell {
	d := dep.(*cell)
	newCell := &cell{
		ord:       len(r.cells),
		reac:      r,
		value:     compute(d.value),
		producers: []*cell{d},
		compute:   func(vs []int) int { return compute(vs[0]) },
		observers: make(map[int]func(int)),
	}
	d.consumers = append(d.consumers, newCell.ord)
	r.cells = append(r.cells, newCell)
	return newCell
}

func (r *reactor) CreateCompute2(dep1, dep2 Cell, compute func(int, int) int) ComputeCell {
	d1 := dep1.(*cell)
	d2 := dep2.(*cell)
	newCell := &cell{
		ord:       len(r.cells),
		reac:      r,
		value:     compute(dep1.Value(), dep2.Value()),
		producers: []*cell{d1, d2},
		compute:   func(vs []int) int { return compute(vs[0], vs[1]) },
		observers: make(map[int]func(int)),
	}
	d1.consumers = append(d1.consumers, newCell.ord)
	d2.consumers = append(d2.consumers, newCell.ord)
	r.cells = append(r.cells, newCell)
	return newCell
}

// Check to see if any cells need updating.
func (r *reactor) trigger(start *cell) {
	// The bitvector is used as a dirty list to indicate which cells need to be re-evaluated.
	// This uses the fact that a cell with a lower ordinal cannot depend on a cell with
	// a higher ordinal.
	bv := newBitv(len(r.cells))
	for _, ord := range start.consumers {
		bv.set(uint(ord))
	}
	for i := start.ord + 1; i < len(r.cells); i++ {
		if !bv.get(uint(i)) {
			continue
		}
		c := r.cells[i]
		vals := make([]int, len(c.producers))
		for i, p := range c.producers {
			vals[i] = p.value
		}
		oldValue := c.value
		c.value = c.compute(vals)
		if oldValue != c.value {
			for _, cb := range c.observers {
				cb(c.value)
			}
			for _, ord := range c.consumers {
				bv.set(uint(ord))
			}
		}
	}
}

// bitv is a simple bit vector with good cache behavior.
type bitv []byte

// newBitv creates a new bit vector with at least 'bits' bits, initialized to false.
func newBitv(bits int) bitv {
	return make([]byte, bits/8+1)
}

func (b bitv) get(pos uint) bool {
	return b[pos/8]&(1<<(pos%8)) != 0
}

func (b bitv) set(pos uint) {
	b[pos/8] |= 1 << (pos % 8)
}
