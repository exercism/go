package complex

import (
	"fmt"
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

func TestNumber_Conjugate(t *testing.T) {
	for _, tt := range conjugateTestCases {
		t.Run(tt.description, func(t *testing.T) {
			n := Number{tt.in.a, tt.in.b}
			//TODO: change condition
			if got := n.Conjugate(); got.a != tt.want.a || got.b != tt.want.b {
				t.Errorf("Number({%v, %v}).Conjugate() = %v, want %v", tt.in.a, tt.in.b, got, tt.want)
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

func TestNumber_Exp(t *testing.T) {
	for _, tt := range expTestCases {
		t.Run(tt.description, func(t *testing.T) {
			n := Number{tt.in.a, tt.in.b}
			if got := n.Exp(); fmt.Sprintf("%v", got) != fmt.Sprintf("%v", tt.want) {
				t.Errorf("Number({%v, %v}).Exp() = %v, want %v", tt.in.a, tt.in.b, got, tt.want)
			}
		})
	}
}
