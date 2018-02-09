package leetcode0034

func extremeInsertionIndex(nums []int, target int, left bool) int {
	l, r := 0, len(nums)

	for l < r {
		m := (l + r) / 2
		if nums[m] > target || (left && target == nums[m]) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func searchRange(nums []int, target int) []int {
	targetRange := []int{-1, -1}
	leftIndex := extremeInsertionIndex(nums, target, true)

	if leftIndex == len(nums) || nums[leftIndex] != target {
		return targetRange
	}

	targetRange[0] = leftIndex
	targetRange[1] = extremeInsertionIndex(nums, target, false) - 1
	return targetRange
}
