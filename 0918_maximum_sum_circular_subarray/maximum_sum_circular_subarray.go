package leetcode0918

func maxSubarraySumCircular(A []int) int {
	total, maxSum, minSum, curMax, curMin := 0, -30000, 30000, 0, 0
	for _, a := range A {
		curMax = max(curMax+a, a)
		maxSum = max(maxSum, curMax)
		curMin = min(curMin+a, a)
		minSum = min(minSum, curMin)
		total += a
	}
	if maxSum < 0 {
		return maxSum
	}
	return max(maxSum, total-minSum)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
