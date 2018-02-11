package leetcode0035

func searchInsert(nums []int, target int) int {
	i := 0
	for i < len(nums) && nums[i] <= target {
		if nums[i] == target {
			return i
		}
		i++
	}
	return i
}
