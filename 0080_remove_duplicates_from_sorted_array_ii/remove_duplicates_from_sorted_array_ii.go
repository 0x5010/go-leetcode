package leetcode0080

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}

	res := 2
	for i := 2; i < n; i++ {
		if nums[i] != nums[res-2] {
			nums[res] = nums[i]
			res++
		}
	}
	return res
}
