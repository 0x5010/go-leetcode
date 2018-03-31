package leetcode0264

func nthUglyNumber(n int) int {
	res := make([]int, n)
	res[0] = 1
	i2, i3, i5 := 0, 0, 0
	for i := 1; i < n; i++ {
		next := min(2*res[i2], min(3*res[i3], 5*res[i5]))
		if next == 2*res[i2] {
			i2++
		}
		if next == 3*res[i3] {
			i3++
		}
		if next == 5*res[i5] {
			i5++
		}
		res[i] = next
	}
	return res[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
