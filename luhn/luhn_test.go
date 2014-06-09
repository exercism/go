package luhn

import "testing"

var validTests = []struct {
	n  string
	ok bool
}{
	{"738", false},
	{"8739567", true},
	{"1111", false},
	{"8763", true},
	{"    ", false},
	{"", false},
	{"2323 2005 7766 3554", true},
}

var addTests = []struct{ raw, luhn string }{
	{"123", "1230"},
	{"873956", "8739567"},
	{"837263756", "8372637564"},
	{"2323 2005 7766 355", "2323 2005 7766 3554"},
	// bonus Unicode cases
	// {"2323·2005·7766·355", "2323·2005·7766·3554"},
	// {"１２３", "１２３０"},
}

func TestValid(t *testing.T) {
	for _, test := range validTests {
		if ok := Valid(test.n); ok != test.ok {
			t.Fatalf("Valid(%s) = %t, want %t.", test.n, ok, test.ok)
		}
	}
}

func TestAddCheck(t *testing.T) {
	for _, test := range addTests {
		if luhn := AddCheck(test.raw); luhn != test.luhn {
			t.Fatalf("AddCheck(%s) = %s, want %s.", test.raw, luhn, test.luhn)
		}
	}
}

func BenchmarkValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Valid("2323 2005 7766 3554")
	}
}

func BenchmarkAddCheck(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddCheck("2323 2005 7766 355")
	}
}
