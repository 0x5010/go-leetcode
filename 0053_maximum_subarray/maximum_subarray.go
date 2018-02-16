package leetcode0053

func maxSubArray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	tmp := nums[0]
	max := tmp
	for i := 1; i < n; i++ {
		if tmp < 0 {
			tmp = nums[i]
		} else {
			tmp += nums[i]
		}
		if max < tmp {
			max = tmp
		}
	}
	return max
}
