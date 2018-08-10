package leetcode0768

func maxChunksToSorted(arr []int) int {
	n := len(arr)
	maxL := make([]int, n)
	minR := make([]int, n)

	maxL[0] = arr[0]
	for i := 1; i < n; i++ {
		maxL[i] = max(maxL[i-1], arr[i])
	}

	minR[n-1] = arr[n-1]
	for i := n - 2; i >= 0; i-- {
		minR[i] = min(minR[i+1], arr[i])
	}

	res := 1
	for i := 0; i < n-1; i++ {
		if maxL[i] <= minR[i+1] {
			res++
		}
	}
	return res
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
