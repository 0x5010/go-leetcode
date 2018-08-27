package leetcode0835

func largestOverlap(A [][]int, B [][]int) int {
	n := len(A)
	res := 0
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			a, b := 0, 0
			for i := 0; i < n-r; i++ {
				for j := 0; j < n-c; j++ {
					a += A[i][j] & B[i+r][j+c]
					b += B[i][j] & A[i+r][j+c]
				}
			}
			res = max(res, max(a, b))
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
