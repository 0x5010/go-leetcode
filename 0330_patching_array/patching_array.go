package leetcode0330

func minPatches(nums []int, n int) int {
	res, i, max, m := 0, 0, 1, len(nums)

	for max <= n {
		if i < m && nums[i] <= max {
			max += nums[i]
			i++
		} else {
			max <<= 1
			res++
		}
	}
	return res
}
