package leetcode0284

func findDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[slow]
	for slow != fast {
		slow, fast = nums[slow], nums[nums[fast]]
	}

	slow = 0
	for slow != fast {
		slow, fast = nums[slow], nums[fast]
	}
	return slow
}
