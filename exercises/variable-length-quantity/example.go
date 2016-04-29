package variablelengthquantity

// EncodeVarint returns the varint encoding of x.
func EncodeVarint(x uint32) []byte {
	if x>>7 == 0 {
		return []byte{
			byte(x),
		}
	}

	if x>>14 == 0 {
		return []byte{
			byte(0x80 | x>>7),
			byte(127 & x),
		}
	}

	if x>>21 == 0 {
		return []byte{
			byte(0x80 | x>>14),
			byte(0x80 | x>>7),
			byte(127 & x),
		}
	}

	return []byte{
		byte(0x80 | x>>21),
		byte(0x80 | x>>14),
		byte(0x80 | x>>7),
		byte(127 & x),
	}
}

// DecodeVarint reads a varint-encoded integer from the slice.
// It returns the integer and the number of bytes consumed, or
// zero if there is not enough.
func DecodeVarint(buf []byte) (x uint32, n int) {
	if len(buf) < 1 {
		return 0, 0
	}

	if buf[0] <= 0x80 {
		return uint32(buf[0]), 1
	}

	var b byte
	for n, b = range buf {
		x = x << 7
		x |= uint32(b) & 0x7F
		if (b & 0x80) == 0 {
			return x, n
		}
	}

	return x, n
}
