package leetcode0036

func isValidSudoku(board [][]byte) bool {
	row := make([]int, 9)
	col := make([]int, 9)
	block := make([]int, 9)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			index := 1 << (board[i][j] - '0')
			if row[i]&index != 0 || col[j]&index != 0 || block[i/3*3+j/3]&index != 0 {
				return false
			}
			row[i] |= index
			col[j] |= index
			block[i/3*3+j/3] |= index
		}
	}
	return true
}
