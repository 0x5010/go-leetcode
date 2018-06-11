package leetcode0453

func minMoves(nums []int) int {
	sum, min := 0, nums[0]
	for _, n := range nums {
		sum += n
		if n < min {
			min = n
		}
	}
	return sum - min*len(nums)
}
