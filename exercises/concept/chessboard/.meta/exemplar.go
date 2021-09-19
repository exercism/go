package chessboard

// Rank stores if a square is occupied by a piece
type Rank []bool

// Chessboard contains eight Ranks, accessed with values from '0' to '7'
type Chessboard map[int]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func (cb Chessboard) CountInRank(rank int) (ret int) {
	for _, r := range cb[rank] {
		if r {
			ret++
		}
	}
	return ret
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func (cb Chessboard) CountInFile(file int) (ret int) {
	if file < 1 || file > 8 {
		return
	}
	for _, f := range cb {
		if f[file-1] {
			ret++
		}
	}
	return ret
}

// CountAll should count how many squares are present in the chessboard
func (cb Chessboard) CountAll() (ret int) {
	for _, rank := range cb {
		for range rank {
			ret++
		}
	}
	return ret
}

// CountOccupied returns how many squares are occupied in the chessboard
func (cb Chessboard) CountOccupied() (ret int) {
	for rank := range cb {
		ret += cb.CountInRank(rank)
	}
	return ret
}
