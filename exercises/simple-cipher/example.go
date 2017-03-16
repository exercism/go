package cipher

import "strings"

const testVersion = 1

type shift int

func NewCaesar() Cipher { return NewShift(3) }

func NewShift(s int) Cipher {
	c := shift(s)
	switch {
	case c > 0 && c < 26:
		return c
	case c > -26 && c < 0:
		c += 26
		return c
	}
	return nil
}

func (c shift) Encode(s string) string {
	return strings.Map(
		func(r rune) rune { return enc(r, int(c)) },
		s)
}

func (c shift) Decode(ct string) string {
	return strings.Map(
		func(r rune) rune { return dec(r, int(c)) },
		ct)
}

func enc(r rune, sa int) rune {
	switch {
	case r >= 'a' && r <= 'z':
		return (r-'a'+rune(sa))%26 + 'a'
	case r >= 'A' && r <= 'Z':
		return (r-'A'+rune(sa))%26 + 'a'
	}
	return -1
}

func dec(r rune, sa int) rune {
	if r >= 'a' && r <= 'z' {
		return (r-'a'+rune(26-sa))%26 + 'a'
	}
	return -1
}

type vigenere string

func NewVigenere(key string) Cipher {
	ok := false
	for _, r := range key {
		switch {
		case r < 'a' || r > 'z':
			return nil
		case r > 'a':
			ok = true
		}
	}
	if !ok {
		return nil
	}
	return vigenere(key)
}

func (v vigenere) Encode(s string) string {
	x := 0
	return strings.Map(
		func(r rune) rune {
			if r = enc(r, int(v[x]-'a')); r >= 0 {
				x = (x + 1) % len(v)
			}
			return r
		},
		s)
}

func (v vigenere) Decode(ct string) string {
	x := 0
	return strings.Map(
		func(r rune) rune {
			if r = dec(r, int(v[x]-'a')); r >= 0 {
				x = (x + 1) % len(v)
			}
			return r
		},
		ct)
}
