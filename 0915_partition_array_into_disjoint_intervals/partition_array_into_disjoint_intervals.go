package leetcode0915

func partitionDisjoint(A []int) int {
	n, index, max, cur := len(A), 0, A[0], A[0]
	for i := 1; i < n; i++ {
		if cur > A[i] {
			cur = max
			index = i
		} else if A[i] > max {
			max = A[i]
		}
	}
	return index + 1
}
