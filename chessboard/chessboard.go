package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File = [8]bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard = map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	cbFile, hasCbFile := cb[file]
	if !hasCbFile {
		return 0
	}

	count := 0
	for _, hasFile := range cbFile {
		if !hasFile {
			continue
		}
		count += 1
	}
	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 0 || rank > 8 {
		return 0
	}

	count := 0
	for _, file := range cb {
		if !file[rank-1] {
			continue
		}
		count += 1
	}

	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	// This seems to be what the question wants, but would not recommend
	// count := 0
	// for range cb {
	// 	for range cb {
	// 		count += 1
	// 	}
	// }
	// return count
	return 8 * 8
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	panic("Please implement CountOccupied()")
}
