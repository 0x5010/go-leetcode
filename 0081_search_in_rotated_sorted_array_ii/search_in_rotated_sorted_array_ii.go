package leetcode0081

func search(nums []int, target int) bool {
	n := 0
	if n == 0 {
		return false
	}

	k := 0
	for k < n-1 && nums[k] <= nums[k+1] {
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
