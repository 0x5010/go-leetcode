package leetcode0084

func largestRectangleArea(heights []int) int {
	n := len(heights)
	if heights == nil || n == 0 {
		return 0
	}

	lessFromLeft := make([]int, n)
	lessFromRight := make([]int, n)
	lessFromLeft[0] = -1
	lessFromRight[n-1] = n

	for i := 1; i < n; i++ {
		p := i - 1
		for p >= 0 && heights[p] >= heights[i] {
			p = lessFromLeft[p]
		}
		lessFromLeft[i] = p
	}

	for i := n - 2; i >= 0; i-- {
		p := i + 1
		for p < n && heights[p] >= heights[i] {
			p = lessFromRight[p]
		}
		lessFromRight[i] = p
	}

	max := 0
	for i := 0; i < n; i++ {
		tmp := heights[i] * (lessFromRight[i] - lessFromLeft[i] - 1)
		if tmp > max {
			max = tmp
		}
	}
	return max
}
