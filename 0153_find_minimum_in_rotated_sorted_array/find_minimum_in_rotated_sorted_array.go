package leetcode0153

func findMin(nums []int) int {
	i, j := 0, len(nums)-1
	for i < j {
		if nums[i] < nums[j] {
			return nums[i]
		}
		mid := (i + j) / 2
		if nums[mid] >= nums[i] {
			i = mid + 1
		} else {
			j = mid
		}
	}
	return nums[i]
}
