package leetcode0704

func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	l, r := 0, n-1
	for l <= r {
		m := (l + r) / 2
		if nums[m] == target {
			return m
		} else if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return -1
}
