package cipher

import (
	"fmt"
	"testing"
)

// type for testing implementations of the Cipher interface
type cipherTest struct {
	source string // source text, any UTF-8
	cipher string // cipher text, result of Encode(st)
	plain  string // decoded plain text, result of Decode(ct)
}

var caesarPrepped = []cipherTest{
	{"iamapandabear", "ldpdsdqgdehdu", "iamapandabear"},
	{"programmingisawesome", "surjudpplqjlvdzhvrph", "programmingisawesome"},
	{"todayisholiday", "wrgdblvkrolgdb", "todayisholiday"},
	{"venividivici", "yhqlylglylfl", "venividivici"},
}

var caesarTests = []cipherTest{
	{"Go, go, gophers", "jrjrjrskhuv", "gogogophers"},
	{"I am a panda bear.", "ldpdsdqgdehdu", "iamapandabear"},
	{"Programming is AWESOME!", "surjudpplqjlvdzhvrph", "programmingisawesome"},
	{"today is holiday", "wrgdblvkrolgdb", "todayisholiday"},
	{"Twas the night before Christmas", "wzdvwkhqljkwehiruhfkulvwpdv", "twasthenightbeforechristmas"},
	{" -- @#!", "", ""},
	{"", "", ""},
}

func TestCaesar(t *testing.T) {
	c := NewCaesar()
	t.Run("no extra symbols", func(t *testing.T) {
		testCipher(c, caesarPrepped, t)
	})
	t.Run("with extra symbols", func(t *testing.T) {
		testCipher(c, caesarTests, t)
	})
}

func testCipher(c Cipher, tests []cipherTest, t *testing.T) {
	for _, test := range tests {
		t.Run(fmt.Sprintf("Encode(%s)", test.source), func(tt *testing.T) {
			if enc := c.Encode(test.source); enc != test.cipher {
				tt.Fatalf("Encode(%s): got %q, want %q.", test.source, enc, test.cipher)
			}
		})
		t.Run(fmt.Sprintf("Decode(%s)", test.cipher), func(tt *testing.T) {
			if dec := c.Decode(test.cipher); dec != test.plain {
				tt.Fatalf("Decode(%s): got %q, want %q.", test.cipher, dec, test.plain)
			}
		})
	}
}

var NSATests = []cipherTest{
	{"THE ENEMY IS NEAR", "qebbkbjvfpkbxo", "theenemyisnear"},
	{"SPIES SEND SECRET MESSAGES",
		"pmfbppbkapbzobqjbppxdbp",
		"spiessendsecretmessages"},
	{"THOMAS JEFFERSON DESIGNED A SUBSTITUTION CIPHER",
		"qeljxpgbccboplkabpfdkbaxprypqfqrqflkzfmebo",
		"thomasjeffersondesignedasubstitutioncipher"},
	{"the quick brown fox jumps over the lazy dog",
		"qebnrfzhyoltkclugrjmplsboqebixwvald",
		"thequickbrownfoxjumpsoverthelazydog"},
}

func TestShift(t *testing.T) {
	// test shift(3) against Caesar cases.
	t.Run(fmt.Sprintf("key=%d", 3), func(t *testing.T) {
		c := NewShift(3)
		if c == nil {
			t.Fatal("NewShift(3): got nil, want non-nil Cipher")
		}
		testCipher(c, caesarTests, t)
	})

	// NSA and WP say Caesar uses shift of -3
	t.Run(fmt.Sprintf("key=%d", -3), func(t *testing.T) {
		c := NewShift(-3)
		if c == nil {
			t.Fatal("NewShift(-3): got nil, want non-nil Cipher")
		}
		testCipher(c, NSATests, t)
	})

}

func TestWrongShiftKey(t *testing.T) {
	for _, s := range []int{-27, -26, 0, 26, 27} {
		if NewShift(s) != nil {
			t.Errorf("NewShift(%d): got non-nil, want nil", s)
		}
	}
}

var vtests = []struct {
	key   string
	tests []cipherTest
}{
	{"lemon", []cipherTest{{"ATTACKATDAWN", "lxfopvefrnhr", "attackatdawn"}}},
	{"abcdefghij", []cipherTest{
		{"aaaaaaaaaa", "abcdefghij", "aaaaaaaaaa"},
		{"zzzzzzzzzz", "zabcdefghi", "zzzzzzzzzz"},
	}},
	{"iamapandabear", []cipherTest{
		{"I am a panda bear.", "qayaeaagaciai", "iamapandabear"},
	}},
	{"duxrceqyaimciuucnelkeoxjhdyduu", []cipherTest{
		{"Diffie Hellman", "gccwkixcltycv", "diffiehellman"},
	}},
	{"qgbvno", []cipherTest{
		{"cof-FEE, 123!", "sugars", "coffee"},
	}},
}

func TestVigenere(t *testing.T) {
	for _, test := range vtests {
		t.Run(fmt.Sprintf("key=%s", test.key), func(t *testing.T) {
			v := NewVigenere(test.key)
			if v == nil {
				t.Fatalf("NewVigenere(%q): got nil, want non-nil Cipher",
					test.key)
			}
			testCipher(v, test.tests, t)
		})
	}
}

func TestVigenereWrongKey(t *testing.T) {
	for _, k := range []string{"", "a", "aa", "no way", "CAT", "3", "and,"} {
		if NewVigenere(k) != nil {
			t.Errorf("NewVigenere(%q): got non-nil, want nil", k)
		}
	}
}

// Benchmark combined time to run all tests.
// Note other ciphers test different data; times cannot be compared.
func BenchmarkEncodeCaesar(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	c := NewCaesar()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, test := range caesarTests {
			c.Encode(test.source)
		}
	}
}

func BenchmarkDecodeCaesar(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	c := NewCaesar()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, test := range caesarTests {
			c.Decode(test.cipher)
		}
	}
}

func BenchmarkNewShift(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for s := -27; s <= 27; s++ {
			NewShift(s)
		}
	}
}

func BenchmarkEncodeShift(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	s := NewShift(5)
	all := caesarTests
	all = append(all, NSATests...)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, test := range all {
			s.Encode(test.source)
		}
	}
}

func BenchmarkDecodeShift(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	s := NewShift(5)
	all := caesarTests
	all = append(all, NSATests...)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, test := range all {
			s.Decode(test.cipher)
		}
	}
}

func BenchmarkNewVigenere(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range vtests {
			NewVigenere(test.key)
		}
	}
}

func BenchmarkEncVigenere(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	v := make([]Cipher, len(vtests))
	for i, test := range vtests {
		v[i] = NewVigenere(test.key)
		if v[i] == nil {
			b.Skip("Benchmark requires valid Vigenere test cases")
		}
	}
	b.ResetTimer()
	for j := 0; j < b.N; j++ {
		for i, test := range vtests {
			vi := v[i]
			for _, test := range test.tests {
				vi.Encode(test.source)
			}
		}
	}
}

func BenchmarkDecVigenere(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	v := make([]Cipher, len(vtests))
	for i, test := range vtests {
		v[i] = NewVigenere(test.key)
		if v[i] == nil {
			b.Skip("Benchmark requires valid Vigenere test cases")
		}
	}
	b.ResetTimer()
	for j := 0; j < b.N; j++ {
		for i, test := range vtests {
			vi := v[i]
			for _, test := range test.tests {
				vi.Decode(test.cipher)
			}
		}
	}
}
