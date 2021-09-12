// This code is largely taken from the Go strconv package.
// The original code can be found at http://golang.org/src/pkg/strconv/atoi.go

package hexadecimal

import (
	"errors"
	"math"
)

// ErrRange indicates that a value is out of range for the target type.
var ErrRange = errors.New("value out of range")

// ErrSyntax indicates that a value does not have the right syntax for the target type.
var ErrSyntax = errors.New("invalid syntax")

// A ParseError records a failed conversion.
type ParseError struct {
	Num string // the input
	Err error  // the reason the conversion failed (ErrRange, ErrSyntax)
}

func (e *ParseError) Error() string {
	return "hexadecimal.ParseHex: parsing \"" + e.Num + "\": " + e.Err.Error()
}

// ParseHex interprets a string s in base 16 and returns the corresponding
// value n.
//
// The errors that ParseHex returns have concrete type *ParseError
// and include err.Num = s.  If s is empty or contains invalid
// digits, ParseError.Err = ErrSyntax; if the value corresponding
// to s cannot be represented by an int64, ParseError.Err = ErrRange.
func ParseHex(s string) (n int64, err error) {
	if len(s) < 1 {
		err = ErrSyntax
		goto Error
	}

	for i := 0; i < len(s); i++ {
		d := s[i]
		var v byte
		switch {
		case '0' <= d && d <= '9':
			v = d - '0'
		case 'a' <= d && d <= 'f':
			v = d - 'a' + 10
		case 'A' <= d && d <= 'F':
			v = d - 'A' + 10
		default:
			n = 0
			err = ErrSyntax
			goto Error
		}

		if n >= math.MaxInt64/16+1 {
			// n*16 overflows
			n = math.MaxInt64
			err = ErrRange
			goto Error
		}

		n *= 16
		n1 := n + int64(v)

		if n1 < n {
			// n+v overflows
			n = math.MaxInt64
			err = ErrRange
			goto Error
		}
		n = n1
	}
	return n, nil

Error:
	return n, &ParseError{s, err}
}

func HandleErrors(tests []string) []string {
	e := make([]string, len(tests))
	for i, s := range tests {
		_, err := ParseHex(s)
		switch pe, ok := err.(*ParseError); {
		case err == nil:
			e[i] = "none"
		case ok && pe.Err == ErrSyntax:
			e[i] = "syntax"
		case ok && pe.Err == ErrRange:
			e[i] = "range"
		}
	}
	return e
}
