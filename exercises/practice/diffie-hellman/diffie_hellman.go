package diffiehellman

import "math/big"

func PrivateKey(p *big.Int) *big.Int {
	panic("Please implement the PrivateKey function")
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	panic("Please implement the PublicKey function")
}

func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	panic("Please implement the NewPair function")
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	panic("Please implement the SecretKey function")
}
