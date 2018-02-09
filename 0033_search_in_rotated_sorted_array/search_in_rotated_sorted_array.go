package leetcode0033

func search(nums []int, target int) int {
	index, maxIndex, numsLen := 0, 0, len(nums)
	if numsLen == 0 {
		return -1
	}

	for maxIndex+1 < numsLen && nums[maxIndex] < nums[maxIndex+1] {
		maxIndex++
	}

	for l, r := 0, numsLen; l < r; {
		m := (l + r) / 2
		index = m + maxIndex + 1
		if index >= numsLen {
			index -= numsLen
		}

		if nums[index] > target {
			r = m
		} else if nums[index] < target {
			l = m + 1
		} else {
			return index
		}
	}
	return -1
}
