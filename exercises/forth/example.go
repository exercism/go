// Package forth implements a tiny subset of the Forth language.
package forth

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

type operatorFn func(stack *[]int) error

type operatorID byte

const (
	opAdd operatorID = iota
	opSub
	opMul
	opDiv
	opDrop
	opDup
	opSwap
	opOver
	opConst
	opUserDef
	opEndDef
)

type operatorTyp struct {
	fn operatorFn
	id operatorID
}

func Forth(input []string) (result []int, err error) {
	if len(input) == 0 {
		return []int{}, nil
	}

	// Allocate an initially empty stack, with arbitrary starting capacity of 8.
	stack := make([]int, 0, 8)
	// Allocate a map for user defined words.
	userDefs := make(map[string][]operatorTyp, 8)
	for _, phrase := range input {
		// Parse one phrase of input, building up an operator list.
		opList, err := parse(phrase, userDefs)
		if err != nil {
			return nil, err
		}
		// Perform any operators from that phrase, updating stack.
		for _, opr := range opList {
			err = opr.fn(&stack)
			if err != nil {
				return nil, err
			}
		}
	}

	return stack, nil
}

// parse given phrase, returning an operator list, and updating
// a userDefs map for any user definifion of words in phrase.
func parse(phrase string, userDefs map[string][]operatorTyp) (oplist []operatorTyp, err error) {
	words := strings.FieldsFunc(phrase,
		func(r rune) bool {
			return unicode.IsSpace(r) || unicode.IsControl(r)
		})

	// t is token index into words[]
	for t := 0; t < len(words); t++ {
		w := strings.ToUpper(words[t])
		// Handle reference to user defined word.
		if udef, ok := userDefs[w]; ok {
			oplist = append(oplist, udef...)
		} else if op, ok := builtinOps[w]; ok {
			if op.id == opUserDef {
				// Handle user defined word definition.
				t++
				if t >= len(words)-2 {
					return nil, errEmptyUserDef
				}
				userword := strings.ToUpper(words[t])
				if _, numerr := strconv.Atoi(userword); numerr == nil {
					return nil, errInvalidUserDef
				}
				t++
				var userops []operatorTyp
				for t < len(words) {
					oneOp, err := parse(words[t], userDefs)
					if err != nil {
						return nil, err
					}
					if oneOp[0].id == opEndDef {
						break
					}
					userops = append(userops, oneOp...)
					t++
				}
				if len(userops) == 0 {
					return nil, errEmptyUserDef
				}
				userDefs[userword] = userops
			} else {
				// Normal builtin operator.
				oplist = append(oplist, op)
			}
		} else {
			// Handle constant literal.
			var x int
			x, err = strconv.Atoi(w)
			if err != nil {
				return nil, err
			}
			oplist = append(oplist,
				operatorTyp{id: opConst,
					fn: func(stack *[]int) error {
						push(stack, x)
						return nil
					},
				})
		}
	}
	return oplist, nil
}

// builtinOps are the pre-defined operators to support.
var builtinOps = map[string]operatorTyp{
	"+":    {add, opAdd},
	"-":    {subtract, opSub},
	"*":    {multiply, opMul},
	"/":    {divide, opDiv},
	"DUP":  {dup, opDrop},
	"DROP": {drop, opDup},
	"SWAP": {swap, opSwap},
	"OVER": {over, opOver},
	":":    {nil, opUserDef},
	";":    {nil, opEndDef},
}

func pop(stack *[]int) (v int, err error) {
	slen := len(*stack)
	if slen >= 1 {
		v = (*stack)[slen-1]
		*stack = (*stack)[:slen-1]
		return v, nil
	}
	return 0, errNotEnoughOperands
}

func pop2(stack *[]int) (v1, v2 int, err error) {
	v1, err = pop(stack)
	if err != nil {
		return 0, 0, err
	}
	v2, err = pop(stack)
	return v1, v2, err
}

func push(stack *[]int, v int) {
	*stack = append(*stack, v)
}

func binaryOp(stack *[]int, op func(a, b int) int) error {
	v1, v2, err := pop2(stack)
	if err != nil {
		return err
	}
	push(stack, op(v2, v1))
	return nil
}

func add(stack *[]int) (err error) {
	return binaryOp(stack, func(a, b int) int { return a + b })
}

func subtract(stack *[]int) (err error) {
	return binaryOp(stack, func(a, b int) int { return a - b })
}

func multiply(stack *[]int) error {
	return binaryOp(stack, func(a, b int) int { return a * b })
}

func divide(stack *[]int) error {
	v1, v2, err := pop2(stack)
	if err != nil {
		return err
	}
	if v1 == 0 {
		return errDivideByZero
	}
	push(stack, v2/v1)
	return nil
}

func dup(stack *[]int) error {
	v1, err := pop(stack)
	if err != nil {
		return err
	}
	push(stack, v1)
	push(stack, v1)
	return nil
}

func drop(stack *[]int) error {
	_, err := pop(stack)
	return err
}

func over(stack *[]int) error {
	v1, v2, err := pop2(stack)
	if err != nil {
		return err
	}
	push(stack, v2)
	push(stack, v1)
	push(stack, v2)
	return nil
}

func swap(stack *[]int) error {
	v1, v2, err := pop2(stack)
	if err != nil {
		return err
	}
	push(stack, v1)
	push(stack, v2)
	return nil
}

var errNotEnoughOperands = errors.New("not enough operands")
var errDivideByZero = errors.New("attempt to divide by zero")
var errEmptyUserDef = errors.New("empty user definition")
var errInvalidUserDef = errors.New("invalid user def word")
