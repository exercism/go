package constants

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"strconv"
	"testing"
)

func TestGetFixedInterestRate(t *testing.T) {
	tests := map[string]struct {
		want float32
	}{
		"GetFixedInterestRate 1": {want: 0.05},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if got := GetFixedInterestRate(); got != tc.want {
				t.Errorf("GetFixedInterestRate() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestGetDaysPerYear(t *testing.T) {
	tests := map[string]struct {
		want int
	}{
		"GetDaysPerYear 1": {want: 365},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if got := GetDaysPerYear(); got != tc.want {
				t.Errorf("GetDaysPerYear() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestGetMonth(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "./constants.go", nil, 0)
	if err != nil {
		t.Fatal(err)
	}

	var conf types.Config
	pkg, err := conf.Check("P", fset, []*ast.File{f}, nil)
	if err != nil {
		t.Fatal(err)
	}

	tests := map[string]struct {
		arg  string
		want int
	}{
		"GetMonth(Jan)": {arg: "Jan", want: 1},
		"GetMonth(Mar)": {arg: "Mar", want: 3},
		"GetMonth(Oct)": {arg: "Oct", want: 10},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tv, err := types.Eval(fset, pkg, 0, tc.arg)
			if err != nil {
				t.Errorf("Eval(%s) failed: %v", tc.arg, err)
			}
			if tv.Type.String() != "untyped int" {
				t.Error("Invalid const type:", tv.Type.String())
			}
			value, err := strconv.Atoi(tv.Value.String())
			if err != nil {
				t.Error("Invalid const type 2")
			}

			if got := GetMonth(value); got != tc.want {
				t.Errorf("GetMonth(%v) = %v, want %v", tc.arg, got, tc.want)
			}
		})
	}
}

func TestGetAccountNumber(t *testing.T) {
	tests := map[string]struct {
		want AccNo
	}{
		"GetAccountNumber 1": {want: "XF348IJ"},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if got := GetAccountNumber(); got != tc.want {
				t.Errorf("GetAccountNumber() = %v, want %v", got, tc.want)
			}
		})
	}
}
