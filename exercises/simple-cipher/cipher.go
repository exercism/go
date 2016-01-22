package cipher

type Cipher interface {
	Encode(string) string
	Decode(string) string
}
