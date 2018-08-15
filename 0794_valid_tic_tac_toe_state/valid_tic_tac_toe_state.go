package leetcode0794

func validTicTacToe(board []string) bool {
	xs, os := 0, 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == 'X' {
				xs++
			} else if board[i][j] == 'O' {
				os++
			}
		}
	}
	win := func(b byte) bool {
		for i := 0; i < 3; i++ {
			if board[i][0] == b && board[i][1] == b && board[i][2] == b ||
				board[0][i] == b && board[1][i] == b && board[2][i] == b {
				return true
			}
		}
		return board[1][1] == b &&
			(board[0][0] == b && board[2][2] == b ||
				board[0][2] == b && board[2][0] == b)
	}
	if win('X') {
		return xs == os+1 && !win('O')
	}
	if win('O') {
		return xs == os && !win('X')
	}
	return xs == os+1 || xs == os
}
