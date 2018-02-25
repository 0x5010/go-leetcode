package leetcode0081

func search(nums []int, target int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}

	k := 1
	for k < n && nums[k-1] <= nums[k] {
		k++
	}

	for i, j := 0, n-1; i <= j; {
		m := (i + j) / 2
		middle := (m + k) % n

		if nums[middle] == target {
			return true
		} else if nums[middle] < target {
			i = m + 1
		} else {
			j = m - 1
		}
	}
	return false
}
