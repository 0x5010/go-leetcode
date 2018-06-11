package leetcode0448

func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	res := []int{}
	for i := 0; i < n; i++ {
		k := abs(nums[i]) - 1
		if nums[k] > 0 {
			nums[k] = -nums[k]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			res = append(res, i+1)
		}
	}
	return res
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}
