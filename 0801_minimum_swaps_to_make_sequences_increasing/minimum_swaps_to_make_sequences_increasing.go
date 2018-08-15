package leetcode0801

func minSwap(A []int, B []int) int {
	n := len(A)
	swap, fix := 1, 0
	for i := 1; i < n; i++ {
		if A[i-1] >= B[i] || B[i-1] >= A[i] {
			swap++
		} else if A[i-1] >= A[i] || B[i-1] >= B[i] {
			swap, fix = fix+1, swap
		} else {
			tmp := min(swap, fix)
			swap, fix = tmp+1, tmp
		}
	}
	return min(swap, fix)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
