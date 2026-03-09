package rationalnumbers

import "testing"

func TestReduce(t *testing.T) {
	for _, tc := range testCasesReduce {
		t.Run(tc.description, func(t *testing.T) {
			if actual := tc.input.Reduce(); actual != tc.expected {
				t.Fatalf("Reduce(%#v) = %#v, want: %#v", tc.input, actual, tc.expected)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	for _, tc := range testCasesAdd {
		t.Run(tc.description, func(t *testing.T) {
			if actual := tc.input[0].Add(tc.input[1]); actual != tc.expected {
				t.Fatalf("Add(%#v, %#v) = %#v, want: %#v", tc.input[0], tc.input[1], actual, tc.expected)
			}
		})
	}
}

func TestSub(t *testing.T) {
	for _, tc := range testCasesSub {
		t.Run(tc.description, func(t *testing.T) {
			if actual := tc.input[0].Sub(tc.input[1]); actual != tc.expected {
				t.Fatalf("Sub(%#v, %#v) = %#v, want: %#v", tc.input[0], tc.input[1], actual, tc.expected)
			}
		})
	}
}

func TestMul(t *testing.T) {
	for _, tc := range testCasesMul {
		t.Run(tc.description, func(t *testing.T) {
			if actual := tc.input[0].Mul(tc.input[1]); actual != tc.expected {
				t.Fatalf("Mul(%#v, %#v) = %#v, want: %#v", tc.input[0], tc.input[1], actual, tc.expected)
			}
		})
	}
}

func TestDiv(t *testing.T) {
	for _, tc := range testCasesDiv {
		t.Run(tc.description, func(t *testing.T) {
			if actual := tc.input[0].Div(tc.input[1]); actual != tc.expected {
				t.Fatalf("Div(%#v, %#v) = %#v, want: %#v", tc.input[0], tc.input[1], actual, tc.expected)
			}
		})
	}
}

func TestAbs(t *testing.T) {
	for _, tc := range testCasesAbs {
		t.Run(tc.description, func(t *testing.T) {
			if actual := tc.input.Abs(); actual != tc.expected {
				t.Fatalf("Abs(%#v) = %#v, want: %#v", tc.input, actual, tc.expected)
			}
		})
	}
}

func TestExprational(t *testing.T) {
	for _, tc := range testCasesExprational {
		t.Run(tc.description, func(t *testing.T) {
			if actual := tc.inputR.Exprational(tc.inputInt); actual != tc.expected {
				t.Fatalf("Exprational(%#v, %#v) = %#v, want: %#v", tc.inputR, tc.inputInt, actual, tc.expected)
			}
		})
	}
}

func TestExpreal(t *testing.T) {
	for _, tc := range testCasesExpreal {
		t.Run(tc.description, func(t *testing.T) {
			actual := tc.inputR.Expreal(tc.inputInt)
			diff := (actual - tc.expected)
			if diff < 0 {
				diff = -diff
			}
			if diff > 0.001 {
				t.Fatalf("Expreal(%#v, %#v) = %#v, want: %#v", tc.inputInt, tc.inputR, actual, tc.expected)
			}
		})
	}
}

func BenchmarkAbs(b *testing.B) {
	for range b.N {
		for _, tc := range testCasesAbs {
			tc.input.Abs()
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	for range b.N {
		for _, tc := range testCasesAdd {
			tc.input[0].Add(tc.input[1])
		}
	}
}

func BenchmarkSub(b *testing.B) {
	for range b.N {
		for _, tc := range testCasesSub {
			tc.input[0].Sub(tc.input[1])
		}
	}
}

func BenchmarkMul(b *testing.B) {
	for range b.N {
		for _, tc := range testCasesMul {
			tc.input[0].Mul(tc.input[1])
		}
	}
}

func BenchmarkDiv(b *testing.B) {
	for range b.N {
		for _, tc := range testCasesDiv {
			tc.input[0].Div(tc.input[1])
		}
	}
}

func BenchmarkExprational(b *testing.B) {
	for range b.N {
		for _, tc := range testCasesExprational {
			tc.inputR.Exprational(tc.inputInt)
		}
	}
}

func BenchmarkExpreal(b *testing.B) {
	for range b.N {
		for _, tc := range testCasesExpreal {
			tc.inputR.Expreal(tc.inputInt)
		}
	}
}

func BenchmarkReduce(b *testing.B) {
	for range b.N {
		for _, tc := range testCasesReduce {
			tc.input.Reduce()
		}
	}
}
