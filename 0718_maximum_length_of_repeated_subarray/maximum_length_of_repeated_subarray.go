package leetcode0718

func findLength(A []int, B []int) int {
	m, n := len(A), len(B)
	res := 0
	for k := 0; k < n; k++ {
		tmp := 0
		for i, j := 0, k; i < m && j < n; i, j = i+1, j+1 {
			if A[i] == B[j] {
				tmp++
				res = max(res, tmp)
			} else {
				tmp = 0
			}
		}
	}
	for k := 1; k < m; k++ {
		tmp := 0
		for i, j := k, 0; i < m && j < m; i, j = i+1, j+1 {
			if A[i] == B[j] {
				tmp++
				res = max(res, tmp)
			} else {
				tmp = 0
			}
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
