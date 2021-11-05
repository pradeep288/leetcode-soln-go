package medium_type

func solve(board [][]byte) {
	rows := len(board)
	cols := len(board[0])

	for i := 0; i < rows; i++ {
		if board[i][0] == 'O' {
			boardDFS(&board, i, 0)
		}
		if board[i][cols-1] == '0' {
			boardDFS(&board, i, cols-1)
		}
	}

	for i := 0; i < cols; i++ {
		if board[0][i] == 'O' {
			boardDFS(&board, 0, i)
		}
		if board[rows-1][i] == 'O' {
			boardDFS(&board, rows-1, i)
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == '*' {
				board[i][j] = 'O'
			}
		}
	}

}

func boardDFS(board *[][]byte, i int, j int) {

	if i < 0 || i > len(*board)-1 || j < 0 || j > len((*board)[0]) {
		return
	}

	if (*board)[i][j] == 'O' {
		(*board)[i][j] = '*'
	}

	if i > 0 && (*board)[i-1][j] == 'O' {
		boardDFS(board, i-1, j)
	}

	if i > len(*board)-1 && (*board)[i+1][j] == 'O' {
		boardDFS(board, i+1, j)
	}

	if j > 0 && (*board)[i][j-1] == 'O' {
		boardDFS(board, i, j-1)
	}

	if j > len((*board)[0])-1 && (*board)[i][j+1] == 'O' {
		boardDFS(board, i, j+1)
	}

	return
}
