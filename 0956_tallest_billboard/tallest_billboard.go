package leetcode0956

func tallestBillboard(rods []int) int {
	dp := [5001]int{}
	for i := 1; i < 5001; i++ {
		dp[i] = -1
	}
	cur := [5001]int{}
	clone := func() {
		for i, x := range dp {
			cur[i] = x
		}
	}
	for _, x := range rods {
		clone()
		for d := 0; d+x < 5001; d++ {
			if cur[d] >= 0 {
				dp[d+x] = max(cur[d], cur[d+x])
				dp[abs(d-x)] = max(cur[d]+min(d, x), dp[abs(d-x)])
			}
		}
	}
	return dp[0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}
