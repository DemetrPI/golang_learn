package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool;
// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(board Chessboard, file string) int {
	count := 0
	value, ok := board[file]
	if ok {
		for _, v := range value {
			if v {
				count++
			}
		}
	}
	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(board Chessboard, rank int) int {
	count := 0
	if rank < 1 || rank > 8 {
		return 0
	}
	for _, v := range board {
		if v[rank-1] {
			count++
		}
	}
	return count
}
// CountAll should count how many squares are present in the chessboard.
func CountAll(board Chessboard) int {
	count := 0
	for _, value := range board {
		count += len(value)
	}
	return count
}
// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(board Chessboard) int {
	count := 0
	for k := range board {
		count += CountInFile(board, k)

	}
	return count
}

