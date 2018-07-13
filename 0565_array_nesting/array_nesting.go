package leetcode0565

func arrayNesting(nums []int) int {
	res := 0
	for i := range nums {
		tmp := 1
		for nums[i] != i {
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
			tmp++
		}

		if tmp > res {
			res = tmp
		}
	}
	return res
}
