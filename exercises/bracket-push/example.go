package brackets

// testVersion shows the version of the exercise.
const testVersion = 5

type bracketKind int

const (
	kindParen  bracketKind = iota
	kindSquare bracketKind = iota
	kindBrace  bracketKind = iota
	kindOther  bracketKind = iota
)

// Whether a bracket is an opening or closing bracket
type bracketForm bool

const (
	formOpen  bracketForm = true
	formClose bracketForm = false
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
		ci := info(char)
		if ci.kind == kindOther {
			continue
		}
		if ci.form == formOpen {
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

func info(char rune) charInfo {
	switch char {
	case '(':
		return charInfo{kind: kindParen, form: formOpen}
	case '[':
		return charInfo{kind: kindSquare, form: formOpen}
	case '{':
		return charInfo{kind: kindBrace, form: formOpen}
	case ')':
		return charInfo{kind: kindParen, form: formClose}
	case ']':
		return charInfo{kind: kindSquare, form: formClose}
	case '}':
		return charInfo{kind: kindBrace, form: formClose}
	default:
		return charInfo{kind: kindOther}
	}
}
