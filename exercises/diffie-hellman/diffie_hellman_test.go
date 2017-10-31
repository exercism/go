// Diffie-Hellman-Merkle key exchange
//
// Step 1:   PrivateKey(p *big.Int) *big.Int
// Step 2:   PublicKey(private, p *big.Int, g int64) *big.Int
// Step 2.1: NewPair(p *big.Int, g int64) (private, public *big.Int)
// Step 3:   SecretKey(private1, public2, p *big.Int) *big.Int
//
// Private keys should be generated randomly.

package diffiehellman

import (
	"math/big"
	"testing"
)

type testCase struct {
	g    int64    // prime, generator
	p    *big.Int // prime, modulus
	a, b *big.Int // private keys
	A, B *big.Int // public keys
	s    *big.Int // secret key
}

// WP example
var smallTest = testCase{
	5,
	big.NewInt(23),
	big.NewInt(6), big.NewInt(15),
	big.NewInt(8), big.NewInt(19),
	big.NewInt(2),
}

// 1024 bit example modulus from cryptopp.com wiki, random private keys
var biggerTest = testCase{
	2,
	mph("ab359aa76a6773ed7a93b214db0c25d0160817b8a893c001c761e198a3694509" +
		"ebe87a5313e0349d95083e5412c9fc815bfd61f95ddece43376550fdc624e92f" +
		"f38a415783b97261204e05d65731bba1ccff0e84c8cd2097b75feca1029261ae" +
		"19a389a2e15d2939314b184aef707b82eb94412065181d23e04bf065f4ac413f"),

	mph("2f6afe91cb53ecfa463d45cd800c948f7cb83bb9ddc62a07b5b3fd302d0cdf52" +
		"18ae53ad015a632d137001a3b34239d54715606a292b6cf895b09d7dcf1bdf7a"),

	mph("3651007bfa8a8b1cbaed2ae9326327599249c3bb6e9d8744b7407f3d4732cb8a" +
		"0708d95c382747bad640d444f135e7e599618d11b15b9ef32e7ac7194e547f4b"),

	mph("57d5489e3858cbd8fae75120907d1521f8e935cce2206d285b11762847e2a4c4" +
		"a341a4eea18bdd8b40036c8d0004ffc323d5ae19da55176b08ff6f2d0ac97c65" +
		"357c1f16756a6901ff926c8544c8af0a90ed2705966851f79a115cb77aac66be" +
		"ceff569aadd7f02859539c28d55c08c62a03e45613bc52d205ace0704ea7c610"),

	mph("6b189a36db5ca3ff83b66fb2c226078294c323f4c7cef35c506c237b0db7906d" +
		"64cceb05af15a3603a30fd49834d3a6971d917f520d9a577c159d3b7d2bd8813" +
		"5d19db47a9618340e4a51ec8845dbf5d50a4c6f14d6161def1eeaacecee8018f" +
		"8816113a294959399989b759f4618e35dbffc570ab2a5a74ac59fccef35f684c"),

	mph("64f74babc466f8e56d9b77ce2cc65d65fe1603b931c018b98a2a615d66172590" +
		"803a94ac230db02de4b8ae567497c1844a6f7bd5bed69b95d3137acc1a6462d5" +
		"aeba5b2b83a0e6b8ed4c072e5135a57c87b654ebe04cf128bbff49ee06df33b7" +
		"8615e0067fdc9df22f7673b1e0501fb57598c7bff9ff173ddff47270fbd6f98f"),
}

// must parse hex, short name contrived just to make test data line up with
// tab width 4.
func mph(h string) *big.Int {
	p, ok := new(big.Int).SetString(h, 16)
	if !ok {
		panic("invalid hex: " + h)
	}
	return p
}

var tests = []testCase{
	smallTest,
	biggerTest,
}

var _one = big.NewInt(1)

// test that PrivateKey returns numbers in range, returns different numbers.
func TestPrivateKey(t *testing.T) {
	priv := func(p *big.Int) *big.Int {
		a := PrivateKey(p)
		if a.Cmp(_one) <= 0 || a.Cmp(p) >= 0 {
			t.Fatalf("PrivateKey(%d) = %d, out of range (1, %d)", p, a, p)
		}
		return a
	}
	ms := map[string]bool{}
	mb := map[string]bool{}
	for i := 0; i < 100; i++ {
		ms[priv(smallTest.p).String()] = true
		mb[priv(biggerTest.p).String()] = true
	}
	if len(ms) == 1 {
		t.Fatalf("For prime %d same key generated every time.  "+
			"Want random keys.", smallTest.p)
	}
	if len(mb) < 100 {
		t.Fatal("For prime %d duplicate keys generated.  "+
			"Want unique keys.", biggerTest.p)
	}
}

// test that PublicKey returns known results.
func TestPublicKey(t *testing.T) {
	tp := func(a, A, p *big.Int, g int64) {
		if k := PublicKey(a, p, g); k.Cmp(A) != 0 {
			t.Fatalf("PublicKey(%x,\n%x,\n%d)\n= %x,\nwant %x.",
				a, p, g, k, A)
		}
	}
	for _, test := range tests {
		tp(test.a, test.A, test.p, test.g)
		tp(test.b, test.B, test.p, test.g)
	}
}

// test that SecretKey returns known results.
func TestSecretKeys(t *testing.T) {
	tp := func(a, B, p, s *big.Int) {
		if k := SecretKey(a, B, p); k.Cmp(s) != 0 {
			t.Fatalf("SecretKey(%x,\n%x,\n%x)\n= %x,\nwant %x.",
				a, B, p, k, s)
		}
	}
	for _, test := range tests {
		tp(test.a, test.B, test.p, test.s)
		tp(test.b, test.A, test.p, test.s)
	}
}

// test that NewPair produces working keys
func TestNewPair(t *testing.T) {
	p, g := biggerTest.p, biggerTest.g
	test := func(a, A *big.Int) {
		if a.Cmp(_one) <= 0 || a.Cmp(p) >= 0 {
			t.Fatalf("NewPair(%d, %d) private key = %d, out of range (1, %d)",
				p, g, a, p)
		}
		if A.Cmp(_one) <= 0 || A.Cmp(p) >= 0 {
			t.Fatalf("NewPair(%d, %d) public key = %d, out of range (1, %d)",
				p, g, A, p)
		}
	}
	a, A := NewPair(p, g)
	test(a, A)
	for i := 0; i < 20; i++ {
		b, B := NewPair(p, g)
		test(b, B)
		sa := SecretKey(a, B, p)
		sb := SecretKey(b, A, p)
		if sa.Cmp(sb) != 0 {
			t.Fatalf("NewPair() produced non-working keys.")
		}
		a, A = b, B
	}
}

func BenchmarkPrivateKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrivateKey(biggerTest.p)
	}
}

func BenchmarkPublicKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PublicKey(biggerTest.a, biggerTest.p, biggerTest.g)
	}
}

func BenchmarkNewPair(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewPair(biggerTest.p, biggerTest.g)
	}
}

func BenchmarkSecretKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SecretKey(biggerTest.a, biggerTest.B, biggerTest.p)
	}
}
