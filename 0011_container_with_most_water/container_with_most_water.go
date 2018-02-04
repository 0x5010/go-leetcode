package leetcode0011

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

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
