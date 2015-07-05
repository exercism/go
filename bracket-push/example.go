package bracket_push

import (
	"fmt"
)

type bracketKind int

const (
	PAREN   bracketKind = iota
	BRACKET bracketKind = iota
	BRACE   bracketKind = iota
)

// Whether a bracket is an opening or closing bracket
type bracketForm bool

const (
	OPEN  bracketForm = true
	CLOSE bracketForm = false
)

type charInfo struct {
	kind bracketKind
	form bracketForm
}

// Bracket returns whether all the brackets in the input string are
// correctly nested.
func Bracket(input string) (bool, error) {
	var stack []bracketKind
	for _, char := range input {
		ci, err := info(char)
		if err != nil {
			return false, err
		}
		if ci.form == OPEN {
			stack = append(stack, ci.kind)
		} else {
			if len(stack) == 0 {
				return false, nil
			} else if stack[len(stack)-1] == ci.kind {
				stack = stack[:len(stack)-1]
			} else {
				return false, nil
			}
		}
	}
	return len(stack) == 0, nil
}

func info(char rune) (charInfo, error) {
	switch char {
	case '(':
		return charInfo{kind: PAREN, form: OPEN}, nil
	case '[':
		return charInfo{kind: BRACKET, form: OPEN}, nil
	case '{':
		return charInfo{kind: BRACE, form: OPEN}, nil
	case ')':
		return charInfo{kind: PAREN, form: CLOSE}, nil
	case ']':
		return charInfo{kind: BRACKET, form: CLOSE}, nil
	case '}':
		return charInfo{kind: BRACE, form: CLOSE}, nil
	default:
		return charInfo{}, fmt.Errorf("Unknown bracket %v", char)
	}
}
