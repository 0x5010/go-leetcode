package leetcode0540

func singleNonDuplicate(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := (l + r) / 2
		if m%2 == 1 {
			m--
		}

		if nums[m] != nums[m+1] {
			r = m
		} else {
			l = m + 2
		}
	}
	return nums[l]
}
