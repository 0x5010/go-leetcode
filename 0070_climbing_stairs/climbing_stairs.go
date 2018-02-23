package leetcode0070

func climbStairs(n int) int {
	if n < 2 {
		return 1
	}

	res := make([]int, n)
	res[0], res[1] = 1, 2
	for i := 2; i < n; i++ {
		res[i] = res[i-1] + res[i-2]
	}
	return res[n-1]
}
