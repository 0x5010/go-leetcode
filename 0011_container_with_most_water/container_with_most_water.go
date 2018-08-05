package leetcode0011

func maxArea(height []int) int {
	maxarea, l, r := 0, 0, len(height)-1
	for l < r {
		maxarea = max(maxarea, min(height[l], height[r])*(r-l))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return maxarea
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
