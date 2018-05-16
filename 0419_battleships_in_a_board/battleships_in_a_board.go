package leetcode0419

func countBattleships(board [][]byte) int {
	m := len(board)
	if m == 0 {
		return 0
	}
	n := len(board[0])
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'X' &&
				(i == 0 || board[i-1][j] == '.') &&
				(j == 0 || board[i][j-1] == '.') {
				count++
			}
		}
	}
	return count
}
