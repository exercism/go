package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools

// Declare a type named Chessboard contains a map of eight Ranks, accessed with values from 1 to 8

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank int) (ret int) {
	panic("Please implement CountInRank()")
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) (ret int) {
	panic("Please implement CountInFile()")
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) (ret int) {
	panic("Please implement CountAll()")
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) (ret int) {
	panic("Please implement CountOccupied()")
}
