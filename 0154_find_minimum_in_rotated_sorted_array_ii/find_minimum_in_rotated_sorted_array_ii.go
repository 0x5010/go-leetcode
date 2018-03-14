package leetcode0154

func findMin(nums []int) int {
	i, j := 0, len(nums)-1
	for i < j {
		mid := (i + j) / 2
		if nums[mid] == nums[j] {
			j--
		} else if nums[mid] > nums[j] {
			i = mid + 1
		} else {
			j = mid
		}
	}
	return nums[i]
}
