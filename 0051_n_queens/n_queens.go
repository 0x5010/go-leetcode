package leetcode0051

func solveNQueens(n int) [][]string {
	res := [][]string{}
	nQueens := make([]string, n)
	s := ""
	for i := 0; i < n; i++ {
		s += "."
	}
	for i := 0; i < n; i++ {
		nQueens[i] = s
	}
	flag := make([]bool, 5*n-2)
	var dfs func(row int)
	dfs = func(row int) {
		if row == n {
			tmp := make([]string, n)
			copy(tmp, nQueens)
			res = append(res, tmp)
			return
		}
		for col := 0; col != n; col++ {
			if !flag[col] && !flag[n+row+col] && !flag[4*n-2+col-row] {
				flag[col], flag[n+row+col], flag[4*n-2+col-row] = true, true, true
				nQueens[row] = nQueens[row][:col] + "Q" + nQueens[row][col+1:]
				dfs(row + 1)
				nQueens[row] = nQueens[row][:col] + "." + nQueens[row][col+1:]
				flag[col], flag[n+row+col], flag[4*n-2+col-row] = false, false, false
			}
		}
	}
	dfs(0)
	return res
}
