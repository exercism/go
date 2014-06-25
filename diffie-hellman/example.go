package diffiehellman

import (
	"math/big"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().Unix()))
var two = big.NewInt(2)

func PrivateKey(p *big.Int) *big.Int {
	private := new(big.Int).Sub(p, two)
	return private.Add(two, private.Rand(r, private))
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	return new(big.Int).Exp(big.NewInt(g), private, p)
}

func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	private = PrivateKey(p)
	return private, PublicKey(private, p, g)
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	return new(big.Int).Exp(public2, private1, p)
}
