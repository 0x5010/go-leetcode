package leetcode0852

func peakIndexInMountainArray(A []int) int {
	for l, r := 0, len(A)-1; ; {
		m := (l + r) / 2
		if A[m] < A[m+1] {
			l = m
		} else if A[m-1] > A[m] {
			r = m
		} else {
			return m
		}
	}
}
