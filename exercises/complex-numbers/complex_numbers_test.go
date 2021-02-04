package complexnumbers

import "testing"

func TestAddition(t *testing.T) {
	expected := 3 + 0i
	if observed := Addition(1+0i, 2+0i); observed != expected {
		t.Fatalf("Reasult = %v, want %v", observed, expected)
	}

	expected = 0 + 3i
	if observed := Addition(0+1i, 0+2i); observed != expected {
		t.Fatalf("Reasult = %v, want %v", observed, expected)
	}

	expected = 4 + 6i
	if observed := Addition(1+2i, 3+4i); observed != expected {
		t.Fatalf("Reasult = %v, want %v", observed, expected)
	}
}

func TestSubtraction(t *testing.T) {
	expected := -1 + 0i
	if observed := Subtraction(1+0i, 2+0i); observed != expected {
		t.Fatalf("Reasult = %v, want %v", observed, expected)
	}

	expected = 0 - 1i
	if observed := Subtraction(0+1i, 0+2i); observed != expected {
		t.Fatalf("Reasult = %v, want %v", observed, expected)
	}

	expected = -2 - 2i
	if observed := Subtraction(1+2i, 3+4i); observed != expected {
		t.Fatalf("Reasult = %v, want %v", observed, expected)
	}
}

func TestMultiply(t *testing.T) {
	expected := 2 + 0i
	if observed := Multiply(1+0i, 2+0i); observed != expected {
		t.Fatalf("Reasult = %v, want %v", observed, expected)
	}

	expected = -2 + 0i
	if observed := Multiply(0+1i, 0+2i); observed != expected {
		t.Fatalf("Reasult = %v, want %v", observed, expected)
	}

	expected = -5 + 10i
	if observed := Multiply(1+2i, 3+4i); observed != expected {
		t.Fatalf("Reasult = %v, want %v", observed, expected)
	}
}

func TestDivision(t *testing.T) {
	expected := 0.5 + 0i
	if observed := Division(1+0i, 2+0i); observed != expected {
		t.Fatalf("Reasult = %v, want %v", observed, expected)
	}

	expected = 0.5 + 0i
	if observed := Division(0+1i, 0+2i); observed != expected {
		t.Fatalf("Reasult = %v, want %v", observed, expected)
	}

	expected = 0.44 + 0.08i
	if observed := Division(1+2i, 3+4i); observed != expected {
		t.Fatalf("Reasult = %v, want %v", observed, expected)
	}
}

func TestConjugate(t *testing.T) {
	expected := 1 - 1i
	if observed := Conjugate(1 + 1i); observed != expected {
		t.Fatalf("Reasult = %v, want %v", observed, expected)
	}
	expected = 5 - 0i
	if observed := Conjugate(5 + 0i); observed != expected {
		t.Fatalf("Reasult = %v, want %v", observed, expected)
	}
	expected = 0 - 5i
	if observed := Conjugate(0 + 5i); observed != expected {
		t.Fatalf("Reasult = %v, want %v", observed, expected)
	}
}

func TestAbsolute(t *testing.T) {
	expected := 5.0
	if observed := Absolute(3 + 4i); observed != expected {
		t.Fatalf("Reasult = %v, want %v", observed, expected)
	}
	expected = 5.0
	if observed := Absolute(5 + 0i); observed != expected {
		t.Fatalf("Reasult = %v, want %v", observed, expected)
	}
	expected = 5.0
	if observed := Absolute(-5 + 0i); observed != expected {
		t.Fatalf("Reasult = %v, want %v", observed, expected)
	}
	expected = 5.0
	if observed := Absolute(0 + 5i); observed != expected {
		t.Fatalf("Reasult = %v, want %v", observed, expected)
	}
	expected = 5.0
	if observed := Absolute(0 - 5i); observed != expected {
		t.Fatalf("Reasult = %v, want %v", observed, expected)
	}
}

func TestExponent(t *testing.T) {
	expected := -0.9999987317275396 + 0.0015926529164868282i
	if observed := Exponent(0 + 3.14i); observed != expected {
		t.Fatalf("Reasult = %v, want %v", observed, expected)
	}
	expected = 1 + 0i
	if observed := Exponent(0 + 0i); observed != expected {
		t.Fatalf("Reasult = %v, want %v", observed, expected)
	}
	expected = 2.718281828459045 + 0i
	if observed := Exponent(1 + 0i); observed != expected {
		t.Fatalf("Reasult = %v, want %v", observed, expected)
	}
	expected = -1.9937130046685776 + 0.003175296858664687i
	if observed := Exponent(0.69 + 3.14i); observed != expected {
		t.Fatalf("Reasult = %v, want %v", observed, expected)
	}
}
