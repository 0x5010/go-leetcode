package leetcode0674

func findLengthOfLCIS(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	res := 1
	count := 1
	for i := 1; i < n; i++ {
		if nums[i]-nums[i-1] > 0 {
			count++
			if count > res {
				res = count
			}
		} else {
			count = 1
		}
	}
	return res
}
