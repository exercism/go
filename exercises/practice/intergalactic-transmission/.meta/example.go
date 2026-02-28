package transmission

import (
	"errors"
	"math/bits"
)

func Transmit(message []byte) []byte {
	encoded := make([]byte, 0)
	var dataCarry byte = 0x00
	var dataMask byte = 0xFE

	for _, dataByte := range message {
		if dataMask == 0 {
			encoded = append(encoded, addParity(dataCarry))
			dataCarry = 0
			dataMask = 0xFE
		}

		dataOffset := bits.TrailingZeros8(dataMask)
		currentByte := (dataCarry << (8 - dataOffset)) | (dataByte >> dataOffset)
		encoded = append(encoded, addParity(currentByte))

		dataCarry = dataByte & (^dataMask)
		dataMask = dataMask << 1
	}

	if dataMask != 0xFE {
		lastByte := dataCarry << byte(bits.OnesCount8(dataMask))
		encoded = append(encoded, addParity(lastByte))
	}
	return encoded
}

func addParity(dataByte byte) byte {
	if bits.OnesCount8(dataByte)%2 == 0 {
		return dataByte << 1
	}
	return (dataByte << 1) | 0x01
}

func Decode(message []byte) ([]byte, error) {
	if len(message) == 0 {
		return []byte{}, nil
	}

	decoded := make([]byte, 0)
	var carryBits byte = 0x00
	var dataMask byte = 0xFF

	for _, byteToDecode := range message {
		if bits.OnesCount8(byteToDecode)%2 != 0 {
			return []byte{}, errors.New("wrong parity")
		}

		if dataMask == 0xFF {
			// This means we are at the first byte or
			// some multiple where there will be only
			//  7 bits of data.
			carryBits = byteToDecode & 0xFE
			dataMask = 0x80
		} else {
			currentByteData := byteToDecode & 0xFE
			contribution := currentByteData >> bits.TrailingZeros8(dataMask)
			decoded = append(decoded, carryBits|contribution)

			carryBits = (byteToDecode & ^(dataMask | 0x01)) << bits.OnesCount8(dataMask)
			dataMask = (dataMask >> 1) | 0x80
		}
	}
	return decoded, nil
}
