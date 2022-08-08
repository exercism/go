package complex

import (
	"fmt"
	"math"
	"testing"
)

func TestNumber_Real(t *testing.T) {
	for _, tt := range realTestCases {
		t.Run(tt.description, func(t *testing.T) {
			n := Number{tt.in.a, tt.in.b}
			if got := n.Real(); got != tt.want {
				t.Errorf("Number({%v, %v}).Real() = %v, want %v", tt.in.a, tt.in.b, got, tt.want)
			}
		})
	}
}

func TestNumber_Imaginary(t *testing.T) {
	for _, tt := range imaginaryTestCases {
		t.Run(tt.description, func(t *testing.T) {
			n := Number{tt.in.a, tt.in.b}
			if got := n.Imaginary(); got != tt.want {
				t.Errorf("Number({%v, %v}).Imaginary() = %v, want %v", tt.in.a, tt.in.b, got, tt.want)
			}
		})
	}
}

func TestNumber_Add(t *testing.T) {
	for _, tt := range addTestCases {
		t.Run(tt.description, func(t *testing.T) {
			n1 := Number{tt.n1.a, tt.n1.b}
			n2 := Number{tt.n2.a, tt.n2.b}
			if got := n1.Add(n2); fmt.Sprintf("%v", got) != fmt.Sprintf("%v", tt.want) {
				t.Errorf("Number({%v, %v}).Add({%v, %v}) = %v, want %v", tt.n1.a, tt.n1.b, tt.n2.a, tt.n2.b, got, tt.want)
			}
		})
	}
}

func TestNumber_Subtract(t *testing.T) {
	for _, tt := range subtractTestCases {
		t.Run(tt.description, func(t *testing.T) {
			n1 := Number{tt.n1.a, tt.n1.b}
			n2 := Number{tt.n2.a, tt.n2.b}
			if got := n1.Subtract(n2); fmt.Sprintf("%v", got) != fmt.Sprintf("%v", tt.want) {
				t.Errorf("Number{%v, %v}.Subtract({%v, %v}) = %v, want %v", tt.n1.a, tt.n1.b, tt.n2.a, tt.n2.b, got, tt.want)
			}
		})
	}
}

func TestNumber_Multiply(t *testing.T) {
	for _, tt := range multiplyTestCases {
		t.Run(tt.description, func(t *testing.T) {
			if tt.n2 == nil {
				t.Skip("skipping tests with factor used withNumber.Times()")
			}
			n1 := Number{tt.n1.a, tt.n1.b}
			n2 := Number{tt.n2.a, tt.n2.b}
			if got := n1.Multiply(n2); fmt.Sprintf("%v", got) != fmt.Sprintf("%v", tt.want) {
				t.Errorf("Number({%v, %v}).Multiply({%v, %v}) = %v, want %v", tt.n1.a, tt.n1.b, tt.n2.a, tt.n2.b, got, tt.want)
			}
		})
	}
}

func TestNumber_Times(t *testing.T) {
	for _, tt := range multiplyTestCases {
		t.Run(tt.description, func(t *testing.T) {
			if tt.n2 != nil {
				t.Skip("skipping tests with complex multiplier used withNumber.Multiply()")
			}
			n := Number{tt.n1.a, tt.n1.b}
			if got := n.Times(tt.factor); fmt.Sprintf("%v", got) != fmt.Sprintf("%v", tt.want) {
				t.Errorf("Number({%v, %v}).Times(%v) = %v, want %v", tt.n1.a, tt.n1.b, tt.factor, got, tt.want)
			}
		})
	}
}

func TestNumber_Divide(t *testing.T) {
	for _, tt := range divideTestCases {
		t.Run(tt.description, func(t *testing.T) {
			n1 := Number{tt.n1.a, tt.n1.b}
			n2 := Number{tt.n2.a, tt.n2.b}
			if got := n1.Divide(n2); fmt.Sprintf("%v", got) != fmt.Sprintf("%v", tt.want) {
				t.Errorf("Number({%v, %v}).Divide({%v, %v}) = %v, want %v", tt.n1.a, tt.n1.b, tt.n2.a, tt.n2.b, got, tt.want)
			}
		})
	}
}

func TestNumber_Abs(t *testing.T) {
	for _, tt := range absTestCases {
		t.Run(tt.description, func(t *testing.T) {
			n := Number{tt.in.a, tt.in.b}
			if got := n.Abs(); got != tt.want {
				t.Errorf("Number.Abs({%v, %v}) = %v, want %v", tt.in.a, tt.in.b, got, tt.want)
			}
		})
	}
}

func TestNumber_String(t *testing.T) {
	tests := []struct {
		description string
		input       complexNumber
		want        string
	}{
		{
			description: "real number",
			input: complexNumber{
				a: 17.75,
			},
			want: "17.750 + 0.000 * i",
		},
		{
			description: "complex number",
			input: complexNumber{
				b: 170.7547,
			},
			want: "0.000 + 170.755 * i",
		},
		{
			description: "negative complex number",
			input: complexNumber{
				b: -170.7547,
			},
			want: "0.000 - 170.755 * i",
		},
		{
			description: "negative real, positive complex",
			input: complexNumber{
				a: -15.75,
				b: 0.37,
			},
			want: "-15.750 + 0.370 * i",
		},
		{
			description: "positive real, negative complex",
			input: complexNumber{
				a: 7,
				b: -4.75,
			},
			want: "7.000 - 4.750 * i",
		},
		{
			description: "zero",
			input: complexNumber{
				a: 0,
				b: 0,
			},
			want: "0.000 + 0.000 * i",
		},
		{
			description: "negative zero",
			input: complexNumber{
				a: -1 * 0,
				b: -1 * 0,
			},
			want: "0.000 + 0.000 * i",
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			n := Number{tt.input.a, tt.input.b}
			if got := n.Format(); got != tt.want {
				t.Errorf("Number({%v, %v}).String() = %v, want %v", tt.input.a, tt.input.b, got, tt.want)
			}
		})
	}
}

func TestNumber_Conjugate(t *testing.T) {
	for _, tt := range conjugateTestCases {
		t.Run(tt.description, func(t *testing.T) {
			n := Number{tt.in.a, tt.in.b}
			if got := n.Conjugate(); got.Format() != getFormatted(tt.want.a, tt.want.b) {
				t.Errorf("Number({%v, %v}).Conjugate() = %v, want %v", tt.in.a, tt.in.b, got, tt.want)
			}
		})
	}
}

func TestNumber_Exp(t *testing.T) {
	for _, tt := range expTestCases {
		t.Run(tt.description, func(t *testing.T) {
			n := Number{tt.in.a, tt.in.b}
			want := getFormatted(tt.want.a, tt.want.b)
			if got := n.Exp(); got.Format() != want {
				t.Errorf("Number({%v, %v}).Exp() = %q, want %q", tt.in.a, tt.in.b, got.Format(), want)
			}
		})
	}
}

func getFormatted(a, b float64) string {
	sign := "+"
	unpreciseB := float64(int(b*10000)) / 10000
	if unpreciseB < 0 {
		sign = "-"
	}
	return fmt.Sprintf("%.3f %s %.3f * i", a, sign, math.Abs(unpreciseB))
}
