package react

import (
	"fmt"
	"maps"
	"strings"
	"testing"
)

type ReactTest struct {
	r         Reactor
	cells     map[string]Cell
	cancelers map[string]Canceler
	callbacks map[string]int
}

func NewReactTest() ReactTest {
	return ReactTest{
		r:         New(),
		cells:     make(map[string]Cell),
		cancelers: make(map[string]Canceler),
		callbacks: make(map[string]int),
	}
}

func (rt ReactTest) Create(cc CreateCell) {
	switch cc.cType {
	case "input":
		rt.cells[cc.name] = rt.r.CreateInput(cc.value)
	case "compute1":
		rt.cells[cc.name] = rt.r.CreateCompute1(rt.cells[cc.inputs[0]], cc.fn1)
	case "compute2":
		rt.cells[cc.name] = rt.r.CreateCompute2(rt.cells[cc.inputs[0]], rt.cells[cc.inputs[1]], cc.fn2)
	default:
		panic("Unknown cell type")
	}
}

func (cc CreateCell) String() string {
	switch cc.cType {
	case "input":
		return fmt.Sprintf("%s = r.CreateInput(%d)", cc.name, cc.value)
	case "compute1":
		return fmt.Sprintf("%s = r.CreateCompute1(%s, %s)", cc.name, cc.inputs[0], cc.fnString)
	case "compute2":
		return fmt.Sprintf("%s = r.CreateCompute2(%s, %s, %s)", cc.name, cc.inputs[0], cc.inputs[1], cc.fnString)
	default:
		panic("Unknown cell type")
	}
}

func (rt ReactTest) Operate(op Operation) error {
	switch op.oType {
	case "SetValue":
		for key := range rt.callbacks {
			delete(rt.callbacks, key)
		}
		inputCell := rt.cells[op.cell].(InputCell)
		inputCell.SetValue(op.value)
		if len(op.wantCallbacks) != 0 {
			if !maps.Equal(rt.callbacks, op.wantCallbacks) {
				return fmt.Errorf("incorrect callbacks\nactual callbacks: %v\nexpected callbacks: %v", rt.callbacks, op.wantCallbacks)
			}
		}
		for _, name := range op.wantNoCallbacks {
			if _, ok := rt.callbacks[name]; ok {
				return fmt.Errorf("%s.SetValue() called callback %s unexpectedly", op.cell, name)
			}
		}
	case "AddCallback":
		computeCell := rt.cells[op.cell].(ComputeCell)
		rt.cancelers[op.name] = computeCell.AddCallback(func(x int) {rt.callbacks[op.name] = x})
	case "Cancel":
		rt.cancelers[op.name].Cancel()
	case "Value":
		if got := rt.cells[op.cell].Value(); got != op.value {
			return fmt.Errorf("%s.Value() = %d, want %d", op.cell, got, op.value)
		}
	default:
		panic("Unknown operation type")
	}
	return nil
}

func (op Operation) String() string {
	switch op.oType {
	case "SetValue":
		return fmt.Sprintf("%s.SetValue(%d)", op.cell, op.value)
	case "AddCallback":
		return fmt.Sprintf("%sCanceler = %s.AddCallback(callbackFunc[%q])", op.name, op.cell, op.name)
	case "Cancel":
		return fmt.Sprintf("%sCanceler.Cancel()", op.name)
	case "Value":
		return fmt.Sprintf("%s.Value()", op.cell)
	default:
		panic("Unknown operation type")
	}
}

func (tc testCase) CallStack(i int) string {
	calls := []string{"r = New()"}
	for _, cc := range tc.cells {
		calls = append(calls, cc.String())
	}
	for i := range i + 1 {
		calls = append(calls, tc.operations[i].String())
	}
	return strings.Join(calls, "\n")

}

func TestReact(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			rt := NewReactTest()
			for _, cc := range tc.cells {
				rt.Create(cc)
			}
			for i, op := range tc.operations {
				if err := rt.Operate(op); err != nil {
					t.Fatalf("Call stack:\n%s\nError: %s", tc.CallStack(i), err)
				}
			}
		})
	}
}

func BenchmarkReact(b *testing.B) {
	for range b.N {
		for _, tc := range testCases {
			rt := NewReactTest()
			for _, cc := range tc.cells {
				rt.Create(cc)
			}
			for _, op := range tc.operations {
				rt.Operate(op)
			}
		}
	}
}
