package leetcode0440

func findKthNumber(n int, k int) int {
	res := 1
	k--
	for k > 0 {
		count := 0
		start, end := res, res+1
		for start <= n {
			count += min(n+1, end) - start
			start *= 10
			end *= 10
		}
		if count <= k {
			k -= count
			res++
		} else {
			k--
			res *= 10
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
