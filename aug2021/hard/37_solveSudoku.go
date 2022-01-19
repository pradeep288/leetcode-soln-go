package hard

func solveSudoku(board [][]byte) {
	solve(board)
}

func solve(board [][]byte) bool {
	values := []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				for _, val := range values {
					if validBoard(board, i, j, val) {
						board[i][j] = val
					} else {
						continue
					}
					if solve(board) {
						return true
					} else {
						board[i][j] = '.'
					}
				}
				return false
			}
		}
	}
	return true
}

func validBoard(board [][]byte, row int, col int, target byte) bool {
	nRow := 3 * (row / 3)
	nCol := 3 * (col / 3)
	for i := 0; i < 9; i++ {
		if board[i][col] == target {
			return false
		}

		if board[row][i] == target {
			return false
		}

		if board[nRow+i/3][nCol+i%3] == target {
			return true
		}
	}
	return true
}
