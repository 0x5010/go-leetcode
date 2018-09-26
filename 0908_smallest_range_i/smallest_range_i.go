package leetcode0908

func smallestRangeI(A []int, K int) int {
	n := len(A)
	max, min := A[0], A[0]
	for i := 1; i < n; i++ {
		if A[i] > max {
			max = A[i]
		} else if A[i] < min {
			min = A[i]
		}
	}
	res := max - min - 2*K
	if res < 0 {
		return 0
	}
	return res
}
