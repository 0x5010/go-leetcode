package leetcode0275

func hIndex(citations []int) int {
	n := len(citations)
	l, r := 0, n-1
	for l <= r {
		m := (l + r) / 2
		if citations[m] >= n-m {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return n - l
}
