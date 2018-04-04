package leetcode0283

func moveZeroes(nums []int) {
	index := 0

	for i, num := range nums {
		if num != 0 {
			nums[i] = 0
			nums[index] = num
			index++
		}
	}
}
