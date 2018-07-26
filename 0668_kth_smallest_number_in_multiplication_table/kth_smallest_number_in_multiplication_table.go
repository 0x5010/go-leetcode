package leetcode0668

func findKthNumber(m int, n int, k int) int {
	if m > n {
		m, n = n, m
	}
	l, r := 1, m*n
	for l < r {
		mid := (l + r) / 2
		c := count(mid, m, n)
		if c >= k {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

func count(v, m, n int) int {
	c := 0
	for i := 1; i < m+1; i++ {
		c += min(v/i, n)
	}
	return c
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
