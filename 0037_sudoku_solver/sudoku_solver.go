package leetcode0037

func solve(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				continue
			}
			var c byte
			for c = '1'; c <= '9'; c++ {
				if isValid(board, i, j, c) {
					board[i][j] = c
					if solve(board) {
						return true
					}
					board[i][j] = '.'
				}
			}
			return false
		}
	}
	return true
}

func isValid(board [][]byte, row, col int, c byte) bool {
	for i := 0; i < 9; i++ {
		b := board[3*(row/3)+i/3][3*(col/3)+i%3]
		if b != '.' && b == c ||
			board[i][col] != '.' && board[i][col] == c ||
			board[row][i] != '.' && board[row][i] == c {
			return false
		}
	}
	return true
}

func solveSudoku(board [][]byte) {
	solve(board)
}
