package cipher

import "testing"

// type for testing implementations of the Cipher interface
type cipherTest struct {
	description string
	source      string // source text, any UTF-8
	cipher      string // cipher text, result of Encode(st)
	plain       string // decoded plain text, result of Decode(ct)
}

var caesarTests = []cipherTest{
	{
		description: "Caeser iamapandabear",
		source:      "iamapandabear",
		cipher:      "ldpdsdqgdehdu",
		plain:       "iamapandabear",
	},
	{
		description: "Caeser programmingisawesome",
		source:      "programmingisawesome",
		cipher:      "surjudpplqjlvdzhvrph",
		plain:       "programmingisawesome",
	},
	{
		description: "Caeser todayisholiday",
		source:      "todayisholiday",
		cipher:      "wrgdblvkrolgdb",
		plain:       "todayisholiday",
	},
	{
		description: "Caeser venividivici",
		source:      "venividivici",
		cipher:      "yhqlylglylfl",
		plain:       "venividivici",
	},
	{
		description: "Caeser Go, go, gophers",
		source:      "Go, go, gophers",
		cipher:      "jrjrjrskhuv",
		plain:       "gogogophers",
	},
	{
		description: "Caeser I am a panda bear.",
		source:      "I am a panda bear.",
		cipher:      "ldpdsdqgdehdu",
		plain:       "iamapandabear",
	},
	{
		description: "Caeser Programming is AWESOME!",
		source:      "Programming is AWESOME!",
		cipher:      "surjudpplqjlvdzhvrph",
		plain:       "programmingisawesome",
	},
	{
		description: "Caeser today is holiday",
		source:      "today is holiday",
		cipher:      "wrgdblvkrolgdb",
		plain:       "todayisholiday",
	},
	{
		description: "Caeser Twas the night before Christmas",
		source:      "Twas the night before Christmas",
		cipher:      "wzdvwkhqljkwehiruhfkulvwpdv",
		plain:       "twasthenightbeforechristmas",
	},
	{
		description: "Caeser  -- @#!",
		source:      " -- @#!",
		cipher:      "",
		plain:       "",
	},
	{
		description: "Caeser empty",
		source:      "",
		cipher:      "",
		plain:       "",
	},
}

func TestCaesar(t *testing.T) {
	for _, tc := range caesarTests {
		t.Run(tc.description, func(tt *testing.T) {
			c := NewCaesar()
			if got := c.Encode(tc.source); got != tc.cipher {
				tt.Fatalf("NewCaesar().Encode(%q) = %q, want %q.", tc.source, got, tc.cipher)
			}
			if got := c.Decode(tc.cipher); got != tc.plain {
				tt.Fatalf("NewCaesar().Decode(%q) = %q, want %q.", tc.cipher, got, tc.plain)
			}
		})
	}
}

var nsaTests = []cipherTest{
	{
		description: "Shift() THE ENEMY IS NEAR",
		source:      "THE ENEMY IS NEAR",
		cipher:      "qebbkbjvfpkbxo",
		plain:       "theenemyisnear",
	},
	{
		description: "Shift() SPIES SEND SECRET MESSAGES",
		source:      "SPIES SEND SECRET MESSAGES",
		cipher:      "pmfbppbkapbzobqjbppxdbp",
		plain:       "spiessendsecretmessages",
	},
	{
		description: "Shift() THOMAS JEFFERSON DESIGNED A SUBSTITUTION CIPHER",
		source:      "THOMAS JEFFERSON DESIGNED A SUBSTITUTION CIPHER",
		cipher:      "qeljxpgbccboplkabpfdkbaxprypqfqrqflkzfmebo",
		plain:       "thomasjeffersondesignedasubstitutioncipher",
	},
	{
		description: "Shift() the quick brown fox jumps over the lazy dog",
		source:      "the quick brown fox jumps over the lazy dog",
		cipher:      "qebnrfzhyoltkclugrjmplsboqebixwvald",
		plain:       "thequickbrownfoxjumpsoverthelazydog",
	},
}

func TestShift(t *testing.T) {
	for _, tc := range nsaTests {
		t.Run(tc.description, func(tt *testing.T) {
			c := NewShift(-3)
			if got := c.Encode(tc.source); got != tc.cipher {
				tt.Fatalf("NewShift(-3).Encode(%q) = %q, want %q.", tc.source, got, tc.cipher)
			}
			if got := c.Decode(tc.cipher); got != tc.plain {
				tt.Fatalf("NewShift(-3).Decode(%q) = %q, want %q.", tc.cipher, got, tc.plain)
			}
		})
	}
}

type invalidShiftTestCase struct {
	description string
	shift       int
}

var invalidShiftTestCases = []invalidShiftTestCase{
	{description: "shift by -27", shift: -27},
	{description: "shift by -26", shift: -26},
	{description: "shift by 0", shift: 0},
	{description: "shift by 26", shift: 26},
	{description: "shift by 27", shift: 27},
}

func TestWrongShiftKey(t *testing.T) {
	for _, tc := range invalidShiftTestCases {
		t.Run(tc.description, func(tt *testing.T) {
			if NewShift(tc.shift) != nil {
				t.Errorf("NewShift(%d): got non-nil, want nil", tc.shift)
			}
		})
	}
}

type vigenereTestCase struct {
	description string
	key         string
	source      string // source text, any UTF-8
	cipher      string // cipher text, result of Encode(st)
	plain       string // decoded plain text, result of Decode(ct)
}

var vtests = []vigenereTestCase{
	{
		description: "lemon dawn",
		key:         "lemon",
		source:      "ATTACKATDAWN",
		cipher:      "lxfopvefrnhr",
		plain:       "attackatdawn",
	},
	{
		description: "alpha a",
		key:         "abcdefghij",
		source:      "aaaaaaaaaa",
		cipher:      "abcdefghij",
		plain:       "aaaaaaaaaa",
	},
	{
		description: "alpha z",
		key:         "abcdefghij",
		source:      "zzzzzzzzzz",
		cipher:      "zabcdefghi",
		plain:       "zzzzzzzzzz",
	},
	{
		description: "panda",
		key:         "iamapandabear",
		source:      "I am a panda bear.",
		cipher:      "qayaeaagaciai",
		plain:       "iamapandabear",
	},
	{
		description: "crypto",
		key:         "duxrceqyaimciuucnelkeoxjhdyduu",
		source:      "Diffie Hellman",
		cipher:      "gccwkixcltycv",
		plain:       "diffiehellman",
	},
	{
		description: "coffee",
		key:         "qgbvno",
		source:      "cof-FEE, 123!",
		cipher:      "sugars",
		plain:       "coffee",
	},
}

func TestVigenere(t *testing.T) {
	for _, tc := range vtests {
		t.Run(tc.description, func(t *testing.T) {
			c := NewVigenere(tc.key)
			if c == nil {
				t.Fatalf("NewVigenere(%q): got nil, want non-nil Cipher", tc.key)
			}
			if got := c.Encode(tc.source); got != tc.cipher {
				t.Fatalf("NewVigenere(%q).Encode(%q) = %q, want %q.", tc.key, tc.source, got, tc.cipher)
			}
			if got := c.Decode(tc.cipher); got != tc.plain {
				t.Fatalf("NewVigenere(%q).Decode(%q) = %q, want %q.", tc.key, tc.cipher, got, tc.plain)
			}
		})
	}
}

type vigenereInvalidTestCase struct {
	description string
	key         string
}

var vigenereInvalidTestCases = []vigenereInvalidTestCase{
	{description: "Vigenere empty", key: ""},
	{description: "Vigenere a", key: "a"},
	{description: "Vigenere aa", key: "aa"},
	{description: "Vigenere no way", key: "no way"},
	{description: "Vigenere CAT", key: "CAT"},
	{description: "Vigenere 3", key: "3"},
	{description: "Vigenere and", key: "and,"},
}

func TestVigenereWrongKey(t *testing.T) {
	for _, tc := range vigenereInvalidTestCases {
		t.Run(tc.description, func(t *testing.T) {
			if NewVigenere(tc.key) != nil {
				t.Errorf("NewVigenere(%q): got non-nil, want nil", tc.key)
			}
		})
	}
}

// Benchmark combined time to run all tests.
// Note other ciphers test different data; times cannot be compared.
func BenchmarkEncodeCaesar(b *testing.B) {
	c := NewCaesar()
	b.ResetTimer()
	for range b.N {
		for _, tc := range caesarTests {
			c.Encode(tc.source)
		}
	}
}

func BenchmarkDecodeCaesar(b *testing.B) {
	c := NewCaesar()
	b.ResetTimer()
	for range b.N {
		for _, tc := range caesarTests {
			c.Decode(tc.cipher)
		}
	}
}

func BenchmarkNewShift(b *testing.B) {
	for range b.N {
		for s := -27; s <= 27; s++ {
			NewShift(s)
		}
	}
}

func BenchmarkEncodeShift(b *testing.B) {
	s := NewShift(5)
	all := caesarTests
	all = append(all, nsaTests...)
	b.ResetTimer()
	for range b.N {
		for _, tc := range all {
			s.Encode(tc.source)
		}
	}
}

func BenchmarkDecodeShift(b *testing.B) {
	s := NewShift(5)
	all := caesarTests
	all = append(all, nsaTests...)
	b.ResetTimer()
	for range b.N {
		for _, tc := range all {
			s.Decode(tc.cipher)
		}
	}
}

func BenchmarkNewVigenere(b *testing.B) {
	for range b.N {
		for _, tc := range vtests {
			NewVigenere(tc.key)
		}
	}
}

func BenchmarkEncVigenere(b *testing.B) {
	v := make([]Cipher, len(vtests))
	for i, tc := range vtests {
		v[i] = NewVigenere(tc.key)
		if v[i] == nil {
			b.Skip("Benchmark requires valid Vigenere tc cases")
		}
	}
	b.ResetTimer()
	for range b.N {
		for i, tc := range vtests {
			v[i].Encode(tc.source)
		}
	}
}

func BenchmarkDecVigenere(b *testing.B) {
	v := make([]Cipher, len(vtests))
	for i, tc := range vtests {
		v[i] = NewVigenere(tc.key)
		if v[i] == nil {
			b.Skip("Benchmark requires valid Vigenere tc cases")
		}
	}
	b.ResetTimer()
	for range b.N {
		for i, tc := range vtests {
			v[i].Decode(tc.cipher)
		}
	}
}
