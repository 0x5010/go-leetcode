package leetcode0035

func searchInsert(nums []int, target int) int {
	i := 0
	for i < len(nums) && nums[i] <= target {
		i++
	}
	return i
}
