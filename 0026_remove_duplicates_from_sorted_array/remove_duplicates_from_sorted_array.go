package leetcode0026

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for _, num := range nums {
		if num != nums[i] {
			i++
			nums[i] = num
		}
	}
	return i + 1
}
