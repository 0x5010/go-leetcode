package leetcode0673

func findNumberOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([][][2]int, n)
	count := 0
	for _, num := range nums {
		l, r := 0, count
		for l < r {
			mid := (l + r) / 2
			if dp[mid][len(dp[mid])-1][0] < num {
				l = mid + 1
			} else {
				r = mid
			}
		}
		pi := l
		paths := 1
		if pi > 0 {
			prev := dp[pi-1]
			l, r = 0, len(prev)
			for l < r {
				mid := (l + r) / 2
				if prev[mid][0] < num {
					r = mid
				} else {
					l = mid + 1
				}
			}
			paths = prev[len(prev)-1][1]
			if l > 0 {
				paths -= prev[l-1][1]
			}
		}
		if len(dp[pi]) > 0 {
			paths += dp[pi][len(dp[pi])-1][1]
		}
		dp[pi] = append(dp[pi], [2]int{num, paths})
		if pi == count {
			count++
		}
	}
	return dp[count-1][len(dp[count-1])-1][1]

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
