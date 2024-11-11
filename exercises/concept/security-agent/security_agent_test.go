package security

import (
	"fmt"
	"testing"
)

func TestImportantNumber(t *testing.T) {
	tests := []struct {
		num, power, want int
	}{
		{2, 3, 16},
		{1, 5, 32},
		{3, 4, 48},
		{0, 3, 0},
		{5, 0, 5},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("num=%d, power=%d", tt.num, tt.power), func(t *testing.T) {
			if got := ImportantNumber(tt.num, tt.power); got != tt.want {
				t.Errorf("ImportantNumber(%d, %d) = %d; want %d", tt.num, tt.power, got, tt.want)
			}
		})
	}
}

func TestCheckIntegrity(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{"AllEvenCount", []int{1, 2, 3, 1, 2, 3}, true},
		{"AllPairs", []int{7, 7, 8, 8, 9, 9, 10, 10}, true},
		{"OneSingleElement", []int{1, 1, 2, 2, 3}, false},
		{"EvenNumsAndWithoutPairs", []int{7, 4, 3, 7, 4, 8}, false},
		{"OneNumForAllPairs", []int{5, 5, 5, 5, 5, 5, 5, 5}, true},
		{"OneNumOnceWithoutAPair", []int{3, 3, 3, 3, 3}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckIntegrity(tt.nums); got != tt.want {
				t.Errorf("CheckIntegrity(%v) = %v; want %v", tt.nums, got, tt.want)
			}
		})
	}
}

func TestSetFlags(t *testing.T) {
	tests := []struct {
		name string
		num  uint8
		want uint8
	}{
		{"AllBitsSet", 0b1111_1111, 0b1111_0101},
		{"NoBitsSet", 0b0000_0000, 0b0000_0000},
		{"Clear2nd4thBits", 0b1110_1010, 0b1110_0000},
		{"MixedBits", 0b1010_1010, 0b1010_0000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetFlags(tt.num); got != tt.want {
				t.Errorf("SetFlags(%08b) = %08b; want %08b", tt.num, got, tt.want)
			}
		})
	}
}

func TestGenerateKey(t *testing.T) {
	tests := []struct{
		name string
		num uint8
		want uint32
	}{
		{"SomeNum1", 7, 7_877_048},
		{"SomeNum2", 83, 113_526_020},
		{"SomeNum3", 197, 255_222_646},
		{"Zero", 0, 7_405_745},
		{"MaxUint8", 255, 326_171_056},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateHash(tt.num); got != tt.want {
				t.Errorf("GenerateHash(%d) = %d; want %d", tt.num, got, tt.want)
			}
		})
	}
}
