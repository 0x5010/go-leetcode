package leetcode0581

func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	l, r := 0, -1

	min, max := nums[n-1], nums[0]

	for i := 1; i < n; i++ {
		if max <= nums[i] {
			max = nums[i]
		} else {
			r = i
		}

		j := n - i - 1
		if min >= nums[j] {
			min = nums[j]
		} else {
			l = j
		}
	}
	return r - l + 1
}
