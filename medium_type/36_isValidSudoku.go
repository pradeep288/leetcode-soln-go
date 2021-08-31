package medium_type
// https://www.youtube.com/watch?v=TjFXEUCMqI8&ab_channel=NeetCode

func isValidSudoku(board [][]byte) bool {
	cols := make(map[int][]string)
	rows := make(map[int][]string)
	subbox := make(map[int][]string)

	for i, _ := range board {
		cols[i] = make([]string, 9)
		rows[i] = make([]string, 9)
		subbox[i] = make([]string, 9)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if string(board[i][j]) != "." {
				val := string(board[i][j])
				if valExistsInSudokoBoard(cols[j], val) ||
					valExistsInSudokoBoard(rows[i], val) ||
					valExistsInSudokoBoard(subbox[(i/3)*3+j/3], val) {
					return false
				}
				rows[i] = append(rows[i], val)
				cols[j] = append(cols[j], val)
				subbox[(i/3)*3+j/3] = append(subbox[(i/3)*3+j/3], val)
			}
		}
	}
	return true
}

func valExistsInSudokoBoard(elements []string, target string) bool {
	for _, v := range elements {
		if v == target {
			return true
		}
	}
	return false
}
