package leetcode0931

func minFallingPathSum(A [][]int) int {
	n := len(A)
	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			tmp := A[i-1][j]
			if j > 0 {
				tmp = min(tmp, A[i-1][j-1])
			}
			if j < n-1 {
				tmp = min(tmp, A[i-1][j+1])
			}
			A[i][j] += tmp
		}
	}
	res := A[n-1][0]
	for i := 1; i < n; i++ {
		res = min(res, A[n-1][i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
