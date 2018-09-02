package leetcode0861

func matrixScore(A [][]int) int {
	m, n := len(A), len(A[0])
	res := (1 << uint(n-1)) * m
	for j := 1; j < n; j++ {
		cur := 0
		for i := 0; i < m; i++ {
			if A[i][j] == A[i][0] {
				cur++
			}
		}
		res += max(cur, m-cur) * (1 << uint(n-j-1))
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
