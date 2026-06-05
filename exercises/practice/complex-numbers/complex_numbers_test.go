package complexnumbers

import (
	"math"
	"testing"
)

const floatEqualityThreshold = 1e-5

func floatingPointEquals(got, want float64) bool {
	absoluteDifferenceBelowThreshold := math.Abs(got-want) <= floatEqualityThreshold
	relativeDifferenceBelowThreshold := math.Abs(got-want)/(math.Abs(got)+math.Abs(want)) <= floatEqualityThreshold
	return absoluteDifferenceBelowThreshold || relativeDifferenceBelowThreshold
}

func TestNumber_Real(t *testing.T) {
	for _, tc := range realTestCases {
		t.Run(tc.description, func(t *testing.T) {
			n := Number{tc.in.a, tc.in.b}
			if got := n.Real(); !floatingPointEquals(got, tc.want) {
				t.Errorf("Number%+v.Real() = %v, want %v", tc.in, got, tc.want)
			}
		})
	}
}

func TestNumber_Imaginary(t *testing.T) {
	for _, tc := range imaginaryTestCases {
		t.Run(tc.description, func(t *testing.T) {
			n := Number{tc.in.a, tc.in.b}
			if got := n.Imaginary(); !floatingPointEquals(got, tc.want) {
				t.Errorf("Number%+v.Imaginary() = %v, want %v", tc.in, got, tc.want)
			}
		})
	}
}

func TestNumber_Add(t *testing.T) {
	for _, tc := range addTestCases {
		t.Run(tc.description, func(t *testing.T) {
			n1 := Number{tc.n1.a, tc.n1.b}
			n2 := Number{tc.n2.a, tc.n2.b}
			if got := n1.Add(n2); !floatingPointEquals(got.Real(), tc.want.a) || !floatingPointEquals(got.Imaginary(), tc.want.b) {
				t.Errorf("Number%+v.Add%+v\n got: %+v\nwant: %+v", tc.n1, tc.n2, got, tc.want)
			}
		})
	}
}

func TestNumber_Subtract(t *testing.T) {
	for _, tc := range subtractTestCases {
		t.Run(tc.description, func(t *testing.T) {
			n1 := Number{tc.n1.a, tc.n1.b}
			n2 := Number{tc.n2.a, tc.n2.b}
			if got := n1.Subtract(n2); !floatingPointEquals(got.Real(), tc.want.a) || !floatingPointEquals(got.Imaginary(), tc.want.b) {
				t.Errorf("Number%+v.Subtract%+v\n got: %+v\nwant: %+v", tc.n1, tc.n2, got, tc.want)
			}
		})
	}
}

func TestNumber_Multiply(t *testing.T) {
	for _, tc := range multiplyTestCases {
		t.Run(tc.description, func(t *testing.T) {
			if tc.n2 == nil {
				t.Skip("skipping tests with factor used withNumber.Times()")
			}
			n1 := Number{tc.n1.a, tc.n1.b}
			n2 := Number{tc.n2.a, tc.n2.b}
			if got := n1.Multiply(n2); !floatingPointEquals(got.Real(), tc.want.a) || !floatingPointEquals(got.Imaginary(), tc.want.b) {
				t.Errorf("Number%+v.Multiply%+v\n got: %+v\nwant: %+v", tc.n1, tc.n2, got, tc.want)
			}
		})
	}
}

func TestNumber_Times(t *testing.T) {
	for _, tc := range multiplyTestCases {
		t.Run(tc.description, func(t *testing.T) {
			if tc.n2 != nil {
				t.Skip("skipping tests with complex multiplier used withNumber.Multiply()")
			}
			n := Number{tc.n1.a, tc.n1.b}
			if got := n.Times(tc.factor); !floatingPointEquals(got.Real(), tc.want.a) || !floatingPointEquals(got.Imaginary(), tc.want.b) {
				t.Errorf("Number%+v.Times(%v)\n got: %+v\nwant: %+v", tc.n1, tc.factor, got, tc.want)
			}
		})
	}
}

func TestNumber_Divide(t *testing.T) {
	for _, tc := range divideTestCases {
		t.Run(tc.description, func(t *testing.T) {
			n1 := Number{tc.n1.a, tc.n1.b}
			n2 := Number{tc.n2.a, tc.n2.b}
			if got := n1.Divide(n2); !floatingPointEquals(got.Real(), tc.want.a) || !floatingPointEquals(got.Imaginary(), tc.want.b) {
				t.Errorf("Number%+v.Divide%+v\n got: %+v\nwant: %+v", tc.n1, tc.n2, got, tc.want)
			}
		})
	}
}

func TestNumber_Abs(t *testing.T) {
	for _, tc := range absTestCases {
		t.Run(tc.description, func(t *testing.T) {
			n := Number{tc.in.a, tc.in.b}
			if got := n.Abs(); !floatingPointEquals(got, tc.want) {
				t.Errorf("Number.Abs%+v = %v, want %v", tc.in, got, tc.want)
			}
		})
	}
}

func TestNumber_Conjugate(t *testing.T) {
	for _, tc := range conjugateTestCases {
		t.Run(tc.description, func(t *testing.T) {
			n := Number{tc.in.a, tc.in.b}
			if got := n.Conjugate(); !floatingPointEquals(got.Real(), tc.want.a) || !floatingPointEquals(got.Imaginary(), tc.want.b) {
				t.Errorf("Number%+v.Conjugate()\n got: %+v\nwant: %+v", tc.in, got, tc.want)
			}
		})
	}
}

func TestNumber_Exp(t *testing.T) {
	for _, tc := range expTestCases {
		t.Run(tc.description, func(t *testing.T) {
			n := Number{tc.in.a, tc.in.b}
			if got := n.Exp(); !floatingPointEquals(got.Real(), tc.want.a) || !floatingPointEquals(got.Imaginary(), tc.want.b) {
				t.Errorf("Number%+v.Exp()\n got: %+v\nwant: %+v", tc.in, got, tc.want)
			}
		})
	}
}

func BenchmarkNumber_Real(b *testing.B) {
	for range b.N {
		for _, tc := range realTestCases {
			Number{tc.in.a, tc.in.b}.Real()
		}
	}
}

func BenchmarkNumber_Imaginary(b *testing.B) {
	for range b.N {
		for _, tc := range imaginaryTestCases {
			Number{tc.in.a, tc.in.b}.Imaginary()
		}
	}
}

func BenchmarkNumber_Add(b *testing.B) {
	for range b.N {
		for _, tc := range addTestCases {
			n1 := Number{tc.n1.a, tc.n1.b}
			n2 := Number{tc.n2.a, tc.n2.b}
			n1.Add(n2)
		}
	}
}

func BenchmarkNumber_Subtract(b *testing.B) {
	for range b.N {
		for _, tc := range subtractTestCases {
			n1 := Number{tc.n1.a, tc.n1.b}
			n2 := Number{tc.n2.a, tc.n2.b}
			n1.Subtract(n2)
		}
	}
}

func BenchmarkNumber_Multiply(b *testing.B) {
	for range b.N {
		for _, tc := range multiplyTestCases {
			if tc.n2 == nil {
				b.Skip("skipping tests with factor used withNumber.Times()")
			}
			n1 := Number{tc.n1.a, tc.n1.b}
			n2 := Number{tc.n2.a, tc.n2.b}
			n1.Multiply(n2)
		}
	}
}

func BenchmarkNumber_Times(b *testing.B) {
	for range b.N {
		for _, tc := range multiplyTestCases {
			if tc.n2 != nil {
				b.Skip("skipping tests with complex multiplier used withNumber.Multiply()")
			}
			Number{tc.n1.a, tc.n1.b}.Times(tc.factor)
		}
	}
}

func BenchmarkNumber_Divide(b *testing.B) {
	for range b.N {
		for _, tc := range divideTestCases {
			n1 := Number{tc.n1.a, tc.n1.b}
			n2 := Number{tc.n2.a, tc.n2.b}
			n1.Divide(n2)
		}
	}
}

func BenchmarkNumber_Abs(b *testing.B) {
	for range b.N {
		for _, tc := range absTestCases {
			Number{tc.in.a, tc.in.b}.Abs()
		}
	}
}

func BenchmarkNumber_Conjugate(b *testing.B) {
	for range b.N {
		for _, tc := range conjugateTestCases {
			Number{tc.in.a, tc.in.b}.Conjugate()
		}
	}
}

func BenchmarkNumber_Exp(b *testing.B) {
	for range b.N {
		for _, tc := range expTestCases {
			Number{tc.in.a, tc.in.b}.Exp()
		}
	}
}
