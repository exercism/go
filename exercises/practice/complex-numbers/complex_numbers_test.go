package complex

import (
	"math"
	"testing"
)

func TestNumber_Real(t *testing.T) {
	for _, tt := range realTestCases {
		t.Run(tt.description, func(t *testing.T) {
			n := Number{tt.in.a, tt.in.b}
			if got := n.Real(); got != tt.want {
				t.Errorf("Number%+v.Real() = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestNumber_Imaginary(t *testing.T) {
	for _, tt := range imaginaryTestCases {
		t.Run(tt.description, func(t *testing.T) {
			n := Number{tt.in.a, tt.in.b}
			if got := n.Imaginary(); got != tt.want {
				t.Errorf("Number%+v.Imaginary() = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestNumber_Add(t *testing.T) {
	for _, tt := range addTestCases {
		t.Run(tt.description, func(t *testing.T) {
			n1 := Number{tt.n1.a, tt.n1.b}
			n2 := Number{tt.n2.a, tt.n2.b}
			if got := n1.Add(n2); got.Real() != tt.want.a || got.Imaginary() != tt.want.b {
				t.Errorf("Number%+v.Add%+v\n got: a=%v, b=%v\nwant: a=%v, b=%v", tt.n1, tt.n2, got.Real(), got.Imaginary(), tt.want.a, tt.want.b)
			}
		})
	}
}

func TestNumber_Subtract(t *testing.T) {
	for _, tt := range subtractTestCases {
		t.Run(tt.description, func(t *testing.T) {
			n1 := Number{tt.n1.a, tt.n1.b}
			n2 := Number{tt.n2.a, tt.n2.b}
			if got := n1.Subtract(n2); got.Real() != tt.want.a || got.Imaginary() != tt.want.b {
				t.Errorf("Number%+v.Subtract%+v\n got: a=%v, b=%v\nwant: a=%v, b=%v", tt.n1, tt.n2, got.Real(), got.Imaginary(), tt.want.a, tt.want.b)
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
			if got := n1.Multiply(n2); got.Real() != tt.want.a || got.Imaginary() != tt.want.b {
				t.Errorf("Number%+v.Multiply%+v\n got: a=%v, b=%v\nwant: a=%v, b=%v", tt.n1, tt.n2, got.Real(), got.Imaginary(), tt.want.a, tt.want.b)
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
			if got := n.Times(tt.factor); got.Real() != tt.want.a || got.Imaginary() != tt.want.b {
				t.Errorf("Number%+v.Times(%v)\n got: a=%v, b=%v\nwant: a=%v, b=%v", tt.n1, tt.factor, got.Real(), got.Imaginary(), tt.want.a, tt.want.b)
			}
		})
	}
}

func TestNumber_Divide(t *testing.T) {
	for _, tt := range divideTestCases {
		t.Run(tt.description, func(t *testing.T) {
			n1 := Number{tt.n1.a, tt.n1.b}
			n2 := Number{tt.n2.a, tt.n2.b}
			if got := n1.Divide(n2); got.Real() != tt.want.a || got.Imaginary() != tt.want.b {
				t.Errorf("Number%+v.Divide%+v\n got: a=%v, b=%v\nwant: a=%v, b=%v", tt.n1, tt.n2, got.Real(), got.Imaginary(), tt.want.a, tt.want.b)
			}
		})
	}
}

func TestNumber_Abs(t *testing.T) {
	for _, tt := range absTestCases {
		t.Run(tt.description, func(t *testing.T) {
			n := Number{tt.in.a, tt.in.b}
			if got := n.Abs(); got != tt.want {
				t.Errorf("Number.Abs%+v = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestNumber_Conjugate(t *testing.T) {
	for _, tt := range conjugateTestCases {
		t.Run(tt.description, func(t *testing.T) {
			n := Number{tt.in.a, tt.in.b}
			if got := n.Conjugate(); got.Real() != tt.want.a || got.Imaginary() != tt.want.b {
				t.Errorf("Number%+v.Conjugate()\n got: a=%v, b=%v\nwant: a=%v, b=%v", tt.in, got.Real(), got.Imaginary(), tt.want.a, tt.want.b)
			}
		})
	}
}

func TestNumber_Exp(t *testing.T) {
	const precision = 0.000001
	for _, tt := range expTestCases {
		t.Run(tt.description, func(t *testing.T) {
			n := Number{tt.in.a, tt.in.b}
			if got := n.Exp(); math.Abs(got.Real()-tt.want.a) > precision || math.Abs(got.Imaginary()-tt.want.b) > precision {
				t.Errorf("Number%+v.Exp()\n got: a=%v, b=%v\nwant: a=%v, b=%v", tt.in, got.Real(), got.Imaginary(), tt.want.a, tt.want.b)
			}
		})
	}
}
