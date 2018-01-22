package railfence

import (
	"sort"
)

type pair struct {
	index, value int
	letter       byte
}

func pattern(rails, size int) []pair {
	var pairs []pair
	var v, direction = 0, 1
	for i := 0; i < size; i++ {
		pairs = append(pairs, pair{index: i, value: v})
		v += direction
		if v == 0 {
			direction = 1
		} else if v == rails-1 {
			direction = -1
		}
	}
	return pairs
}

func sortedPattern(rails, size int) []pair {
	fence := pattern(rails, size)
	sort.Slice(fence, func(i, j int) bool {
		f1, f2 := fence[i], fence[j]
		return f1.value < f2.value || (f1.value == f2.value && f1.index < f2.index)
	})
	return fence
}

// Encode encodes given message with given rails
func Encode(msg string, rails int) string {
	fence := sortedPattern(rails, len(msg))
	var encoded []byte
	for _, p := range fence {
		encoded = append(encoded, msg[p.index])
	}
	return string(encoded)
}

// Decode decodes given message with given rails
func Decode(msg string, rails int) string {
	fence := sortedPattern(rails, len(msg))
	for i := range msg {
		fence[i].letter = msg[i]
	}
	sort.Slice(fence, func(i, j int) bool { return fence[i].index < fence[j].index })
	var decoded []byte
	for _, p := range fence {
		decoded = append(decoded, p.letter)
	}
	return string(decoded)
}
