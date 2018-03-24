package leetcode0213

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	}

	return max(robbing(nums[1:]), robbing(nums[:n-1]))
}

func robbing(nums []int) int {
	a, b := 0, 0
	for i, num := range nums {
		if i%2 == 0 {
			a = max(a+num, b)
		} else {
			b = max(a, b+num)
		}
	}
	return max(a, b)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
