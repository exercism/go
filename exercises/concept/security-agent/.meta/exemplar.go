package security

// ImportantNumber create an important number
func ImportantNumber(num, power int) int {
	return num << power
}

// CheckIntegrity check the slice nums for integrity
func CheckIntegrity(nums []int) bool {
	var res int

	for _, num := range nums {
		res ^= num
	}

	return res == 0
}

// SetFlags sets the 2nd and 4th bits to 0
func SetFlags(num uint8) uint8 {
	var mask uint8 = 0b1111_0101
	return num & mask
}

// GenerateKey generate key by rules
func GenerateHash(num uint8) uint32 {
	fourthByte := uint32(num - 79)
	thirdByte := uint32(num * 7) << 8
	secondByte := uint32(num + 113) << 16
	firstByte := uint32(num / 13) << 24

	return firstByte + secondByte + thirdByte + fourthByte
}
