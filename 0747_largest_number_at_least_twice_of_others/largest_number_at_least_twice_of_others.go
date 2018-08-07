package leetcode0747

func dominantIndex(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}

	i1, i2 := 0, 1
	if nums[i1] < nums[i2] {
		i1, i2 = i2, i1
	}

	for i := 2; i < n; i++ {
		if nums[i1] < nums[i] {
			i2, i1 = i1, i
		} else if nums[i2] < nums[i] {
			i2 = i
		}
	}

	if nums[i2] == 0 || nums[i1]/nums[i2] >= 2 {
		return i1
	}
	return -1
}
