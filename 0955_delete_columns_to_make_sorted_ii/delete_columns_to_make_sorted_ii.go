package leetcode0955

func minDeletionSize(A []string) int {
	m, n := len(A), len(A[0])
	sorted := make([]bool, m-1)
	res := 0

L:
	for j := 0; j < n; j++ {
		i := 0
		for ; i < m-1; i++ {
			if !sorted[i] && A[i][j] > A[i+1][j] {
				res++
				continue L
			}
		}
		for i = 0; i < m-1; i++ {
			if A[i][j] < A[i+1][j] {
				sorted[i] = true
			}
		}
	}
	return res
}
