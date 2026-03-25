package simplelinkedlist

import (
	"fmt"
	"slices"
	"strings"
	"testing"
)

func (op Operation) Execute(l *List) (*List, int, []int, error) {
	switch op.operation {
	case "Push":
		l.Push(op.value)
		return l, 0, nil, nil
	case "Pop":
		gotInt, err := l.Pop()
		return l, gotInt, nil, err
	case "Peek":
		gotInt, err := l.Peek()
		return l, gotInt, nil, err
	case "Size":
		gotInt := l.Size()
		return l, gotInt, nil, nil
	case "Reverse":
		l := l.Reverse()
		return l, 0, nil, nil
	case "Array":
		gotInts := l.Array()
		return l, 0, gotInts, nil
	default:
		panic(fmt.Sprintf("Invalid op %q", op.operation))
	}
}

func (op Operation) String() string {
	switch op.operation {
	case "Push":
		return fmt.Sprintf("l.%s(%d)", op.operation, op.value)
	case "Pop", "Peek", "Size":
		return fmt.Sprintf("l.%s()", op.operation)
		return fmt.Sprintf("l.%s()", op.operation)
	case "Reverse":
		return fmt.Sprintf("l = l.%s()", op.operation)
	case "Array":
		return fmt.Sprintf("l.%s()", op.operation)
	default:
		panic(fmt.Sprintf("Invalid op %q", op.operation))
	}
}

func (tc testCase) callStack(i int) string {
	calls := []string{fmt.Sprintf("l := New(%#v)", tc.initialValues)}
	for i := range i + 1 {
		calls = append(calls, tc.operations[i].String())
	}
	return strings.Join(calls, "; ")
}

func TestList(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			l := New(tc.initialValues)

			for i, op := range tc.operations {
				newList, gotInt, gotInts, err := op.Execute(l)
				// Preserve across loops
				l = newList

				// Verify
				switch op.operation {
				case "Pop", "Peek", "Size":
					if op.expectedErr != "" {
						if err == nil {
							t.Fatalf("%s = %d, expected error %s", tc.callStack(i), gotInt, op.expectedErr)
						} else if err.Error() != op.expectedErr {
							t.Fatalf("%s = %v, expected error %s", tc.callStack(i), err, op.expectedErr)
						}
					} else if err != nil {
						t.Fatalf("%s\ngot unexpected error %v", tc.callStack(i), err)
					} else if gotInt != op.expectedInt {
						t.Fatalf("%s = %d, want %d", tc.callStack(i), gotInt, op.expectedInt)
					}
				case "Array":
					if !slices.Equal(gotInts, op.expectedInts) {
						t.Fatalf("%s = %#v, want %#v", tc.callStack(i), gotInts, op.expectedInts)
					}
				}
			}
		})
	}
}

func BenchmarkListOps(b *testing.B) {
	for range b.N {
		for _, tc := range testCases {
			l := New(tc.initialValues)

			for _, op := range tc.operations {
				l, _, _, _ = op.Execute(l)
			}
		}
	}
}

var array1To10 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func BenchmarkNewList(b *testing.B) {
	for range b.N {
		New(array1To10)
	}
}

func BenchmarkListSize(b *testing.B) {
	list := New(array1To10)
	b.ResetTimer()
	for range b.N {
		list.Size()
	}
}

func BenchmarkListPush(b *testing.B) {
	for range b.N {
		b.StopTimer()
		list := New([]int{})
		b.StartTimer()
		for k := range 1000 {
			list.Push(k)
		}
	}
}

func BenchmarkListPop(b *testing.B) {
	for range b.N {
		b.StopTimer()
		list := New([]int{})
		for k := range 1000 {
			list.Push(k)
		}
		for range 1000 {
			list.Pop()
		}
	}
}

func BenchmarkListToArray(b *testing.B) {
	list := New(array1To10)
	b.ResetTimer()
	for range b.N {
		list.Array()
	}
}

func BenchmarkListReverse(b *testing.B) {
	list := New(array1To10)
	b.ResetTimer()
	for range b.N {
		list.Reverse()
	}
}
