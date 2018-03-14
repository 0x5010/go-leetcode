package leetcode0162

func findPeakElement(nums []int) int {
	i, j := 0, len(nums)-1
	for i < j {
		mid := (i + j) / 2
		if nums[mid] < nums[mid+1] {
			i = mid + 1
		} else {
			j = mid
		}
	}
	return i
}
