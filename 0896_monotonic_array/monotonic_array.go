package leetcode0896

func isMonotonic(A []int) bool {
	n := len(A)
	color := 0
	for i := 1; i < n; i++ {
		if A[i-1] == A[i] {
			continue
		} else if A[i-1] < A[i] {
			if color > 0 {
				return false
			}
			color = -1
		} else {
			if color < 0 {
				return false
			}
			color = 1
		}
	}
	return true
}
